package s3wrapper

import (
	"context"
	"fmt"
	"net/url"
	"os"
	"strings"

	"github.com/kdomanski/iso9660"
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

func ExtractFilesFromISOAndUploadStream(ctx context.Context, log logrus.FieldLogger, isoFilePath, isoObjectName string, api API) error {
	destPrefix := strings.TrimSuffix(isoObjectName, ".iso")
	for _, fileType := range ISOFileTypes {
		objectName := destPrefix + "." + fileType
		exists, err := api.DoesObjectExist(ctx, objectName)
		if err != nil {
			return errors.Wrapf(err, "Failed searching for object %s", objectName)
		}
		if exists {
			log.Infof("Object %s already exists, skipping upload", objectName)
			continue
		}
		err = uploadFileFromISO(ctx, isoFilePath, fileType, objectName, api)
		if err != nil {
			log.WithError(err).Fatalf("Failed to extract and upload file %s from ISO %s", fileType, isoFilePath)
		}

		log.Infof("Successfully uploaded object %s", objectName)
	}
	return nil
}

func uploadFileFromISO(ctx context.Context, isoFilePath, fileType, objectName string, api API) error {
	filename, err := convertPXEFileTypeToFileName(fileType)
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

func convertPXEFileTypeToFileName(fileType string) (string, error) {
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
