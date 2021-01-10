package images

import (
	"bytes"
	"context"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"text/template"
	"time"

	ign_3_1 "github.com/coreos/ignition/v2/config/v3_1"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"github.com/openshift/assisted-service/client/images"
	"github.com/openshift/assisted-service/client/installer"
	"github.com/openshift/assisted-service/internal/cluster"
	"github.com/openshift/assisted-service/internal/cluster/validations"
	"github.com/openshift/assisted-service/internal/common"
	"github.com/openshift/assisted-service/internal/events"
	"github.com/openshift/assisted-service/internal/identity"
	"github.com/openshift/assisted-service/internal/ignition"
	"github.com/openshift/assisted-service/internal/versions"
	"github.com/openshift/assisted-service/models"
	"github.com/openshift/assisted-service/pkg/filemiddleware"
	"github.com/openshift/assisted-service/pkg/s3wrapper"
	"github.com/openshift/assisted-service/restapi"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/vincent-petithory/dataurl"
)

var _ restapi.ImagesAPI = &Images{}

func NewImagesAPI(
	db *gorm.DB,
	log logrus.FieldLogger,
	cfg Config,
	eventsHandler events.Handler,
	objectHandler s3wrapper.API,
	versionsHandler versions.Handler,
) *Images {
	return &Images{
		db:              db,
		log:             log,
		Config:          cfg,
		eventsHandler:   eventsHandler,
		objectHandler:   objectHandler,
		versionsHandler: versionsHandler,
	}
}

type Config struct {
	ImageExpirationTime time.Duration `envconfig:"IMAGE_EXPIRATION_TIME" default:"4h"`
}

type Images struct {
	Config
	db              *gorm.DB
	log             logrus.FieldLogger
	eventsHandler   events.Handler
	objectHandler   s3wrapper.API
	versionsHandler versions.Handler
}

func (i *Images) GenerateClusterISO(ctx context.Context, params images.GenerateClusterISOParams) middleware.Responder {
	log := logutil.FromContext(ctx, i.log)
	c, apierr := cluster.GetCluster(ctx, i.log, i.db, params.ClusterID.String())
	if apierr != nil {
		return common.GenerateErrorResponder(apierr)
	}

	id := strfmt.UUID(uuid.New().String())
	log.Infof("Creating image with ID %s for cluster %s", id, params.ClusterID)

	url := installer.GetResourceURL{ClusterID: params.ClusterID, ResourceID: id, ResourceType: "images"}

	_, report, err := ign_3_1.Parse([]byte(params.ImageCreateParams.DiscoveryOverrides))
	if err != nil {
		log.WithError(err).Errorf("Failed to parse ignition config patch %s", params.ImageCreateParams.DiscoveryOverrides)
		return installer.NewUpdateDiscoveryIgnitionBadRequest().WithPayload(common.GenerateError(http.StatusBadRequest, err))
	}
	if report.IsFatal() {
		err = errors.Errorf("Ignition config patch %s failed validation: %s", params.ImageCreateParams.DiscoveryOverrides, report.String())
		log.Error(err)
		return installer.NewUpdateDiscoveryIgnitionBadRequest().WithPayload(common.GenerateError(http.StatusBadRequest, err))
	}

	if !c.PullSecretSet {
		errMsg := "Can't generate cluster ISO without pull secret"
		log.Error(errMsg)
		return installer.NewGenerateClusterISOBadRequest().
			WithPayload(common.GenerateError(http.StatusBadRequest, errors.New(errMsg)))
	}

	staticIpsConfig := formatStaticIPs(params.ImageCreateParams.StaticIpsConfig)

	image := models.Image{
		ID:                 id,
		Href:               swag.String(url.String()),
		Kind:               swag.String(models.ImageKindImage),
		ClusterID:          *c.ID,
		State:              models.ImageStateCreating,
		SSHPublicKey:       params.ImageCreateParams.SSHPublicKey,
		SizeBytes:          0,
		DownloadURL:        "",
		StaticIpsConfig:    staticIpsConfig,
		DiscoveryOverrides: swag.String(params),
		CreatedAt:          strfmt.DateTime(now),
		ExpiresAt:          strfmt.DateTime(now.Add(i.Config.ImageExpirationTime)),
	}

	if err = i.createImageInDB(ctx, image); err != nil {
		err = errors.
	}

	if imageExists {
		if err := b.updateImageInfoPostUpload(ctx, &cluster, clusterProxyHash); err != nil {
			return installer.NewGenerateClusterISOInternalServerError().
				WithPayload(common.GenerateError(http.StatusInternalServerError, err))
		}

		log.Infof("Re-used existing cluster <%s> image", params.ClusterID)
		b.eventsHandler.AddEvent(ctx, params.ClusterID, nil, models.EventSeverityInfo, "Re-used existing image rather than generating a new one", time.Now())
		return installer.NewGenerateClusterISOCreated().WithPayload(&cluster.Cluster)
	}
	ignitionConfig, formatErr := b.formatIgnitionFile(&cluster, params, log, false)
	if formatErr != nil {
		log.WithError(formatErr).Errorf("failed to format ignition config file for cluster %s", cluster.ID)
		msg := "Failed to generate image: error formatting ignition file"
		b.eventsHandler.AddEvent(ctx, params.ClusterID, nil, models.EventSeverityError, msg, time.Now())
		return installer.NewGenerateClusterISOInternalServerError().
			WithPayload(common.GenerateError(http.StatusInternalServerError, formatErr))
	}

	if err := b.objectHandler.Upload(ctx, []byte(ignitionConfig), fmt.Sprintf("%s/discovery.ign", cluster.ID)); err != nil {
		log.WithError(err).Errorf("Upload discovery ignition failed for cluster %s", cluster.ID)
		return common.NewApiError(http.StatusInternalServerError, err)
	}

	if err := b.objectHandler.UploadISO(ctx, ignitionConfig, b.objectHandler.GetBaseIsoObject(cluster.OpenshiftVersion),
		fmt.Sprintf(s3wrapper.DiscoveryImageTemplate, cluster.ID.String())); err != nil {
		log.WithError(err).Errorf("Upload ISO failed for cluster %s", cluster.ID)
		b.eventsHandler.AddEvent(ctx, params.ClusterID, nil, models.EventSeverityError, "Failed to upload image", time.Now())
		return installer.NewGenerateClusterISOInternalServerError().WithPayload(common.GenerateError(http.StatusInternalServerError, err))
	}

	if err := b.updateImageInfoPostUpload(ctx, &cluster, clusterProxyHash); err != nil {
		return installer.NewGenerateClusterISOInternalServerError().
			WithPayload(common.GenerateError(http.StatusInternalServerError, err))
	}

	ignitionConfigForLogging, _ := b.formatIgnitionFile(&cluster, params, log, true)
	log.Infof("Generated cluster <%s> image with ignition config %s", params.ClusterID, ignitionConfigForLogging)

	msg := "Generated image"

	var msgExtras []string

	if cluster.HTTPProxy != "" {
		msgExtras = append(msgExtras, fmt.Sprintf(`proxy URL is "%s"`, cluster.HTTPProxy))
	}

	sshExtra := "SSH public key is not set"
	if params.ImageCreateParams.SSHPublicKey != "" {
		sshExtra = "SSH public key is set"
	}

	msgExtras = append(msgExtras, sshExtra)

	msg = fmt.Sprintf("%s (%s)", msg, strings.Join(msgExtras, ", "))

	b.eventsHandler.AddEvent(ctx, params.ClusterID, nil, models.EventSeverityInfo, msg, time.Now())
	return installer.NewGenerateClusterISOCreated().WithPayload(&cluster.Cluster)
}

func (i *Images) createImageInDB(ctx context.Context, image *models.Image) error {
	tx := r.db.Begin()
	success := false
	defer func() {
		if rec := recover(); rec != nil || !success {
			r.log.Error("create image failed")
			tx.Rollback()
		}
	}()
	if tx.Error != nil {
		r.log.WithError(tx.Error).Error("failed to start transaction")
	}

	if err := tx.Create(image).Error; err != nil {
		r.log.Errorf("Error registering image %s", image)
		return err
	}

	if err := tx.Commit().Error; err != nil {
		return err
	}

	success = true
	return nil
}

func formatStaticIPs(staticIpsConfig []*models.StaticIPConfig) string {
	lines := make([]string, len(staticIpsConfig))

	// construct static IPs config string
	for i, entry := range staticIpsConfig {
		lines[i] = fmt.Sprintf("%s;%s;%s;%s;%s", entry.Mac, entry.IP, entry.Mask, entry.DNS, entry.Gateway)
	}

	sort.Strings(lines)

	return strings.Join(lines, "\n")
}

func (b *bareMetalInventory) DownloadClusterISO(ctx context.Context, params installer.DownloadClusterISOParams) middleware.Responder {
	log := logutil.FromContext(ctx, b.log)
	var cluster common.Cluster

	if err := b.db.First(&cluster, "id = ?", params.ClusterID).Error; err != nil {
		log.WithError(err).Errorf("failed to get cluster %s", params.ClusterID)
		return common.NewApiError(http.StatusNotFound, err)
	}

	imgName := getImageName(*cluster.ID)
	exists, err := b.objectHandler.DoesObjectExist(ctx, imgName)
	if err != nil {
		log.WithError(err).Errorf("Failed to get ISO for cluster %s", cluster.ID.String())
		b.eventsHandler.AddEvent(ctx, params.ClusterID, nil, models.EventSeverityError,
			"Failed to download image: error fetching from storage backend", time.Now())
		return installer.NewDownloadClusterISOInternalServerError().
			WithPayload(common.GenerateError(http.StatusInternalServerError, err))
	}
	if !exists {
		b.eventsHandler.AddEvent(ctx, params.ClusterID, nil, models.EventSeverityError,
			"Failed to download image: the image was not found (perhaps it expired) - please generate the image and try again", time.Now())
		return installer.NewDownloadClusterISONotFound().
			WithPayload(common.GenerateError(http.StatusNotFound, errors.New("The image was not found "+
				"(perhaps it expired) - please generate the image and try again")))
	}
	reader, contentLength, err := b.objectHandler.Download(ctx, imgName)
	if err != nil {
		log.WithError(err).Errorf("Failed to get ISO for cluster %s", cluster.ID.String())
		b.eventsHandler.AddEvent(ctx, params.ClusterID, nil, models.EventSeverityError,
			"Failed to download image: error fetching from storage backend", time.Now())
		return installer.NewDownloadClusterISOInternalServerError().
			WithPayload(common.GenerateError(http.StatusInternalServerError, err))
	}
	b.eventsHandler.AddEvent(ctx, params.ClusterID, nil, models.EventSeverityInfo, "Started image download", time.Now())

	return filemiddleware.NewResponder(installer.NewDownloadClusterISOOK().WithPayload(reader),
		fmt.Sprintf("cluster-%s-discovery.iso", params.ClusterID.String()),
		contentLength)
}

func (b *bareMetalInventory) updateImageInfoPostUpload(ctx context.Context, cluster *common.Cluster, clusterProxyHash string) error {
	updates := map[string]interface{}{}
	imgName := getImageName(*cluster.ID)
	imgSize, err := b.objectHandler.GetObjectSizeBytes(ctx, imgName)
	if err != nil {
		return errors.New("Failed to generate image: error fetching size")
	}
	updates["image_size_bytes"] = imgSize
	cluster.ImageInfo.SizeBytes = &imgSize

	// Presigned URL only works with AWS S3 because Scality is not exposed
	if b.objectHandler.IsAwsS3() {
		signedURL, err := b.objectHandler.GeneratePresignedDownloadURL(ctx, imgName, imgName, b.Config.ImageExpirationTime)
		if err != nil {
			return errors.New("Failed to generate image: error generating URL")
		}
		updates["image_download_url"] = signedURL
		cluster.ImageInfo.DownloadURL = signedURL
	}

	if cluster.ProxyHash != clusterProxyHash {
		updates["proxy_hash"] = clusterProxyHash
		cluster.ProxyHash = clusterProxyHash
	}

	dbReply := b.db.Model(&common.Cluster{}).Where("id = ?", cluster.ID.String()).Updates(updates)
	if dbReply.Error != nil {
		return errors.New("Failed to generate image: error updating image record")
	}

	return nil
}

func getImageName(clusterID strfmt.UUID) string {
	return fmt.Sprintf("%s.iso", fmt.Sprintf(s3wrapper.DiscoveryImageTemplate, clusterID.String()))
}

func (b *bareMetalInventory) formatIgnitionFile(cluster *common.Cluster, params installer.GenerateClusterISOParams, logger logrus.FieldLogger, safeForLogs bool) (string, error) {
	creds, err := validations.ParsePullSecret(cluster.PullSecret)
	if err != nil {
		return "", err
	}
	pullSecretToken := ""
	if b.authHandler.EnableAuth {
		r, ok := creds["cloud.openshift.com"]
		if !ok {
			return "", errors.Errorf("Pull secret does not contain auth for cloud.openshift.com")
		}
		pullSecretToken = r.AuthRaw
	}

	proxySettings, err := proxySettingsForIgnition(cluster.HTTPProxy, cluster.HTTPSProxy, cluster.NoProxy)
	if err != nil {
		return "", err
	}
	rhCa := ""
	if b.Config.InstallRHCa {
		rhCa = url.PathEscape(redhatRootCA)
	}
	var ignitionParams = map[string]string{
		"userSshKey":           b.getUserSshKey(params),
		"AgentDockerImg":       b.AgentDockerImg,
		"ServiceBaseURL":       strings.TrimSpace(b.ServiceBaseURL),
		"clusterId":            cluster.ID.String(),
		"PullSecretToken":      pullSecretToken,
		"AGENT_MOTD":           url.PathEscape(agentMessageOfTheDay),
		"PULL_SECRET":          url.PathEscape(cluster.PullSecret),
		"RH_ROOT_CA":           rhCa,
		"PROXY_SETTINGS":       proxySettings,
		"HTTPProxy":            cluster.HTTPProxy,
		"HTTPSProxy":           cluster.HTTPSProxy,
		"NoProxy":              cluster.NoProxy,
		"SkipCertVerification": strconv.FormatBool(b.SkipCertVerification),
		"AgentTimeoutStartSec": strconv.FormatInt(int64(b.AgentTimeoutStart.Seconds()), 10),
		"SELINUX_POLICY":       base64.StdEncoding.EncodeToString([]byte(selinuxPolicy)),
	}
	if safeForLogs {
		for _, key := range []string{"userSshKey", "PullSecretToken", "PULL_SECRET", "RH_ROOT_CA"} {
			ignitionParams[key] = "*****"
		}
	}
	if b.ServiceCACertPath != "" {
		var caCertData []byte
		caCertData, err = ioutil.ReadFile(b.ServiceCACertPath)
		if err != nil {
			return "", err
		}
		ignitionParams["ServiceCACertData"] = dataurl.EncodeBytes(caCertData)
		ignitionParams["HostCACertPath"] = common.HostCACertPath
	}
	if b.Config.ServiceIPs != "" {
		ignitionParams["ServiceIPs"] = dataurl.EncodeBytes([]byte(ignition.GetServiceIPHostnames(b.Config.ServiceIPs)))
	}

	if cluster.ImageInfo.StaticIpsConfig != "" {
		ignitionParams["StaticIPsData"] = b64.StdEncoding.EncodeToString([]byte(cluster.ImageInfo.StaticIpsConfig))
		ignitionParams["StaticIPsConfigScript"] = b64.StdEncoding.EncodeToString([]byte(ignition.ConfigStaticIpsScript))
	}

	tmpl, err := template.New("ignitionConfig").Parse(ignitionConfigFormat)
	if err != nil {
		return "", err
	}
	buf := &bytes.Buffer{}
	if err = tmpl.Execute(buf, ignitionParams); err != nil {
		return "", err
	}

	res := buf.String()
	if cluster.IgnitionConfigOverrides != "" {
		res, err = ignition.MergeIgnitionConfig(buf.Bytes(), []byte(cluster.IgnitionConfigOverrides))
		if err != nil {
			return "", err
		}
		logger.Infof("Applying ignition overrides %s for cluster %s, resulting ignition: %s", cluster.IgnitionConfigOverrides, cluster.ID, res)
	}

	return res, nil
}

func (b *bareMetalInventory) getUserSshKey(params installer.GenerateClusterISOParams) string {
	sshKey := params.ImageCreateParams.SSHPublicKey
	if sshKey == "" {
		return ""
	}
	return fmt.Sprintf(`{
		"name": "core",
		"passwordHash": "!",
		"sshAuthorizedKeys": [
		"%s"],
		"groups": [ "sudo" ]}`, sshKey)
}

func (b *bareMetalInventory) GetDiscoveryIgnition(ctx context.Context, params installer.GetDiscoveryIgnitionParams) middleware.Responder {
	log := logutil.FromContext(ctx, b.log)

	c, err := b.getCluster(ctx, params.ClusterID.String())
	if err != nil {
		return common.GenerateErrorResponder(err)
	}
	isoParams := installer.GenerateClusterISOParams{ClusterID: params.ClusterID, ImageCreateParams: &models.ImageCreateParams{}}

	cfg, err := b.formatIgnitionFile(c, isoParams, log, false)
	if err != nil {
		log.WithError(err).Error("Failed to format ignition config")
		return common.GenerateErrorResponder(err)
	}

	configParams := models.DiscoveryIgnitionParams{Config: cfg}
	return installer.NewGetDiscoveryIgnitionOK().WithPayload(&configParams)
}

func (b *bareMetalInventory) UpdateDiscoveryIgnition(ctx context.Context, params installer.UpdateDiscoveryIgnitionParams) middleware.Responder {
	log := logutil.FromContext(ctx, b.log)

	c, err := b.getCluster(ctx, params.ClusterID.String())
	if err != nil {
		return common.GenerateErrorResponder(err)
	}

	_, report, err := ign_3_1.Parse([]byte(params.DiscoveryIgnitionParams.Config))
	if err != nil {
		log.WithError(err).Errorf("Failed to parse ignition config patch %s", params.DiscoveryIgnitionParams)
		return installer.NewUpdateDiscoveryIgnitionBadRequest().WithPayload(common.GenerateError(http.StatusBadRequest, err))
	}
	if report.IsFatal() {
		err = errors.Errorf("Ignition config patch %s failed validation: %s", params.DiscoveryIgnitionParams, report.String())
		log.Error(err)
		return installer.NewUpdateDiscoveryIgnitionBadRequest().WithPayload(common.GenerateError(http.StatusBadRequest, err))
	}

	err = b.db.Model(&common.Cluster{}).Where(identity.AddUserFilter(ctx, "id = ?"), params.ClusterID).Update("ignition_config_overrides", params.DiscoveryIgnitionParams.Config).Error
	if err != nil {
		return installer.NewUpdateDiscoveryIgnitionInternalServerError().WithPayload(common.GenerateError(http.StatusInternalServerError, err))
	}

	existed, err := b.objectHandler.DeleteObject(ctx, getImageName(*c.ID))
	if err != nil {
		return common.NewApiError(http.StatusInternalServerError, err)
	}
	if existed {
		b.eventsHandler.AddEvent(ctx, *c.ID, nil, models.EventSeverityInfo, "Deleted image from backend because its ignition was updated. The image may be regenerated at any time.", time.Now())
	}

	return installer.NewUpdateDiscoveryIgnitionCreated()
}
