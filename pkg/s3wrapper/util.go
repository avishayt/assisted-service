package s3wrapper

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/kdomanski/iso9660"
	"github.com/openshift/assisted-service/pkg/leader"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
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

func uploadBootFilesWithLeader(ctx context.Context, api API, log logrus.FieldLogger, uploadLeader leader.ElectorInterface, isoFilePath string) error {
	var objectName string
	var haveISOFile bool
	fileTypes := []string{"iso", "initrd.img", "rootfs.img", "vmlinuz"}

	if isoFilePath != "" {
		// Use this ISO file in place of a URL
		volumeID, err := getVolumeIDFromISO(isoFilePath)
		if err != nil {
			log.Error(err)
			return err
		}
		baseISOObjectName = strings.TrimSpace(volumeID) + ".iso"
		haveISOFile = true
	} else {
		baseISOObjectName = RHCOSBaseISOObjectName
		haveISOFile = false
	}

	for _, fileType := range fileTypes {
		objectName = getBootFileName(baseISOObjectName, fileType)
		exists, err := api.DoesObjectExist(ctx, objectName)
		if err != nil {
			err = errors.Wrapf(err, "Failed to check existence of %s", objectName)
			log.Error(err)
			return err
		}
		if exists {
			continue
		}
		if !haveISOFile {
			// We need to upload a boot file but don't have it locally
			tmpFile, fileErr := ioutil.TempFile(os.TempDir(), "tmpiso-")
			if fileErr != nil {
				fileErr = errors.Wrapf(fileErr, "Failed to create temp file in directory %s", os.TempDir())
				log.Error(fileErr)
				return fileErr
			}
			defer os.Remove(tmpFile.Name())

			fileErr = downloadToFile(RHCOSBaseISOURL, tmpFile)
			if fileErr != nil {
				log.Error(fileErr)
				return fileErr
			}

			isoFilePath = tmpFile.Name()
			haveISOFile = true

		}
		if fileType == "iso" {
			err = uploadLeader.RunWithLeader(ctx, func() error {
				log.Infof("Starting upload of ISO %s", objectName)
				return api.UploadFile(ctx, isoFilePath, objectName)
			})
		} else {
			err = uploadLeader.RunWithLeader(ctx, func() error {
				log.Infof("Starting upload of file %s", objectName)
				return uploadFileFromISO(ctx, isoFilePath, fileType, objectName, api)
			})
		}
		if err != nil {
			err = errors.Wrapf(err, "Failed to upload boot file %s", objectName)
			log.Error(err)
			return err
		}
	}

	return nil
}

func downloadToFile(url string, destFile *os.File) error {
	resp, err := http.Get(url)
	if err != nil {
		return errors.Wrapf(err, "Failed to access URL %s", url)
	}
	defer resp.Body.Close()
	defer destFile.Close()

	if _, err = io.Copy(destFile, resp.Body); err != nil {
		return errors.Wrapf(err, "Failed to download URL to %s", destFile.Name())
	}
	return nil
}

func getVolumeIDFromISO(isoFilePath string) (string, error) {
	iso, err := os.Open(isoFilePath)
	if err != nil {
		return "", errors.Wrapf(err, "Failed to open provided base image for inspection: %s", isoFilePath)
	}
	defer iso.Close()

	// Need a method to identify the ISO provided
	// Based on the ISO 9660 standard we should consistently be able to
	// grab the Volume ID here.
	volumeId := make([]byte, 32)
	_, err = iso.ReadAt(volumeId, 32808)
	if err != nil {
		return "", errors.Wrapf(err, "Failed to read volume ID from provided base image")
	}

	// TODO: Should we verify that the volume id is of the form `rhcos-<version>`?
	return strings.TrimSpace(string(volumeId)), nil
}

func getBootFileName(isoObjectName, fileType string) string {
	return strings.TrimSuffix(isoObjectName, ".iso") + "." + fileType
}

func uploadFileFromISO(ctx context.Context, isoFilePath, fileType, objectName string, api API) error {
	filename, err := convertBootFileTypeToFileName(fileType)
	if err != nil {
		return err
	}
	f, err := os.Open(isoFilePath)
	if err != nil {
		return errors.Wrapf(err, "Failed to open ISO file: %s", isoFilePath)
	}
	defer f.Close()
	image, err := iso9660.OpenImage(f)
	if err != nil {
		return err
	}
	isoFile, err := image.RootDir()
	if err != nil {
		return err
	}
	pathElements := strings.Split(filename, "/")
	for _, pathElement := range pathElements {
		element := strings.TrimSpace(pathElement)
		if element == "" {
			continue
		}
		if !isoFile.IsDir() {
			return errors.New(fmt.Sprintf("Expected directory in ISO while searching in %s for %s", isoFile.Name(), element))
		}
		isoFile, err = findFileInISODirectory(isoFile, element)
		if err != nil {
			return err
		}
	}
	err = api.UploadStream(ctx, isoFile.Reader(), objectName)
	if err != nil {
		return err
	}
	return nil
}

func convertBootFileTypeToFileName(fileType string) (string, error) {
	var filename string
	switch fileType {
	case "initrd.img":
		filename = "/IMAGES/PXEBOOT/INITRD.IMG"
	case "rootfs.img":
		filename = "/IMAGES/PXEBOOT/ROOTFS.IMG"
	case "vmlinuz":
		filename = "/IMAGES/PXEBOOT/VMLINUZ"
	}
	if filename == "" {
		return "", errors.New(fmt.Sprintf("Invalid file type specified: %s", fileType))
	}
	return filename, nil
}

func findFileInISODirectory(directory *iso9660.File, pathElement string) (*iso9660.File, error) {
	children, err := directory.GetChildren()
	if err != nil {
		return nil, err
	}
	for _, child := range children {
		if child.Name() == pathElement {
			return child, nil
		}
	}
	return nil, errors.New(fmt.Sprintf("Failed to find path element %s in ISO directory %s", pathElement, directory.Name()))
}
