package s3wrapper

import (
	"context"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/sirupsen/logrus"

	"github.com/openshift/assisted-service/internal/isoutil"
	"github.com/pkg/errors"
	"github.com/prometheus/common/log"
)

func FixEndpointURL(endpoint string) (string, error) {
	_, err := url.ParseRequestURI(endpoint)
	if err == nil {
		return endpoint, nil
	}

	prefix := "http://"
	if os.Getenv("S3_USE_SSL") == "true" {
		prefix = "https://"
	}

	new_url := prefix + endpoint
	_, err = url.ParseRequestURI(new_url)
	if err != nil {
		return "", err
	}
	return new_url, nil
}

func UploadBootFiles(ctx context.Context, log logrus.FieldLogger, isoObjectName, isoURL string, api API) error {
	exist, err := api.DoAllBootFilesExist(ctx, isoObjectName)
	if err != nil {
		return err
	}
	if exist {
		return nil
	}

	log.Infof("Starting Base ISO download for %s", isoObjectName)
	tmpfile, err := downloadToTemporaryFile(isoURL)
	if err != nil {
		log.Error(err)
		return err
	}
	defer os.Remove(tmpfile)

	exists, err := api.DoesObjectExist(ctx, isoObjectName)
	if err != nil {
		return err
	}
	if !exists {
		err = api.UploadFileToPublicBucket(ctx, tmpfile, isoObjectName)
		if err != nil {
			return err
		}
		log.Infof("Successfully uploaded object %s", isoObjectName)
	}

	isoHandler := isoutil.NewHandler(tmpfile, "")

	for fileType := range ISOFileTypes {
		objectName := BootFileTypeToObjectName(isoObjectName, fileType)
		exists, err := api.DoesObjectExist(ctx, objectName)
		if err != nil {
			return errors.Wrapf(err, "Failed searching for object %s", objectName)
		}
		if exists {
			log.Infof("Object %s already exists, skipping upload", objectName)
			continue
		}
		log.Infof("Starting to upload %s from Base ISO %s", fileType, isoObjectName)
		err = uploadFileFromISO(ctx, isoHandler, fileType, objectName, api)
		if err != nil {
			log.WithError(err).Fatalf("Failed to extract and upload file %s from ISO", fileType)
		}

		log.Infof("Successfully uploaded object %s", objectName)
	}
	return nil
}

func downloadToTemporaryFile(url string) (string, error) {
	tmpfile, err := ioutil.TempFile("", "isodownload")
	if err != nil {
		return "", errors.Wrap(err, "Error creating temporary file")
	}

	resp, err := http.Get(url)
	if err != nil {
		return "", errors.Wrapf(err, "Failed fetching from URL %s", url)
	}

	_, err = io.Copy(tmpfile, resp.Body)
	if err != nil {
		return "", errors.Wrapf(err, "Failed downloading file from %s to %s", url, tmpfile.Name())
	}

	tmpfile.Close()

	return tmpfile.Name(), nil
}

func uploadFileFromISO(ctx context.Context, isoHandler isoutil.Handler, fileType, objectName string, api API) error {
	filename := ISOFileTypes[fileType]
	reader, err := isoHandler.ReadFile(filename)
	if err != nil {
		return errors.Wrapf(err, "Failed to read file %s from ISO", filename)
	}

	err = api.UploadStreamToPublicBucket(ctx, reader, objectName)
	if err != nil {
		return err
	}
	return nil
}

func BootFileTypeToObjectName(isoObjectName, fileType string) string {
	return strings.TrimSuffix(isoObjectName, ".iso") + "." + fileType
}

func DoAllBootFilesExist(ctx context.Context, isoObjectName string, api API) (bool, error) {
	for _, fileType := range BootFileExtensions {
		objectName := BootFileTypeToObjectName(isoObjectName, fileType)
		exists, err := api.DoesObjectExist(ctx, objectName)
		if err != nil {
			log.Error(err)
			return false, err
		}
		if !exists {
			return false, nil
		}
	}
	return true, nil
}
