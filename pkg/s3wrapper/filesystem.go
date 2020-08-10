package s3wrapper

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"

	logutil "github.com/openshift/assisted-service/pkg/log"

	"github.com/moby/moby/pkg/ioutils"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type FSClient struct {
	log     *logrus.Logger
	basedir string
}

func NewFSClient(basedir string, logger *logrus.Logger) *FSClient {
	return &FSClient{log: logger, basedir: basedir}
}

func (f *FSClient) CreateBucket() error {
	return nil
}

func (f *FSClient) Upload(ctx context.Context, data []byte, objectName string) error {
	log := logutil.FromContext(ctx, f.log)
	filePath := filepath.Join(f.basedir, objectName)
	if err := ioutil.WriteFile(filePath, data, 0600); err != nil {
		err = errors.Wrapf(err, "Unable to write data to file %s", filePath)
		log.Error(err)
		return err
	}
	return nil
}

func (f *FSClient) Download(ctx context.Context, objectName string) (io.ReadCloser, int64, error) {
	log := logutil.FromContext(ctx, f.log)
	filePath := filepath.Join(f.basedir, objectName)
	fp, err := os.Open(filePath)
	if err != nil {
		err = errors.Wrapf(err, "Unable to open file %s", filePath)
		log.Error(err)
		return nil, 0, err
	}
	info, err := fp.Stat()
	if err != nil {
		fp.Close()
		err = errors.Wrapf(err, "Unable to stat file %s", filePath)
		log.Error(err)
		return nil, 0, err
	}
	contentLength := info.Size()
	return ioutils.NewReadCloserWrapper(fp, fp.Close), contentLength, nil
}

func (f *FSClient) DownloadFromPublicURL(ctx context.Context, location string) (io.ReadCloser, int64, error) {
	log := logutil.FromContext(ctx, f.log)
	log.Infof("Fetching from file: %s", filepath.Join(f.basedir, location))
	return f.Download(ctx, location)
}

func (f *FSClient) DoesObjectExist(ctx context.Context, objectName string) (bool, error) {
	filePath := filepath.Join(f.basedir, objectName)
	info, err := os.Stat(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, errors.Wrap(err, fmt.Sprintf("failed to get file %s", filePath))
	}
	return !info.IsDir(), nil
}

func (f *FSClient) DeleteObject(ctx context.Context, objectName string) error {
	log := logutil.FromContext(ctx, f.log)
	filePath := filepath.Join(f.basedir, objectName)
	err := os.Remove(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return errors.Wrap(err, fmt.Sprintf("Failed to delete file %s", filePath))
	}

	log.Infof("Deleted file %s", filePath)
	return nil
}

func (f *FSClient) GetObjectSizeBytes(ctx context.Context, objectName string) (int64, error) {
	filePath := filepath.Join(f.basedir, objectName)
	info, err := os.Stat(filePath)
	if err != nil {
		return 0, errors.Wrap(err, fmt.Sprintf("failed to get file %s", filePath))
	}
	return info.Size(), nil
}

func (f *FSClient) GeneratePresignedDownloadURL(ctx context.Context, objectName string, duration time.Duration) (string, error) {
	return "", nil
}

func (f *FSClient) UpdateObjectTimestamp(ctx context.Context, objectName string) (bool, error) {
	log := logutil.FromContext(ctx, f.log)
	filePath := filepath.Join(f.basedir, objectName)
	log.Infof("Updating timestamp of file %s", filePath)
	now := time.Now()
	if err := os.Chtimes(filePath, now, now); err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, errors.Wrap(err, fmt.Sprintf("Failed to update timestamp for file %s", filePath))
	}
	return true, nil
}

func (f *FSClient) ExpireObjects(ctx context.Context, prefix string, deleteTime time.Duration, callback func(ctx context.Context, objectName string)) {
	log := logutil.FromContext(ctx, f.log)
	now := time.Now()

	log.Info("Checking for expired files...")
	err := filepath.Walk(f.basedir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if strings.HasPrefix(path, prefix) {
			f.handleFile(ctx, log, path, info, now, deleteTime, callback)
		}
		return nil
	})
	if err != nil {
		log.WithError(err).Error("Error listing files")
		return
	}
}

func (f *FSClient) handleFile(ctx context.Context, log logrus.FieldLogger, filePath string, fileInfo os.FileInfo, now time.Time,
	deleteTime time.Duration, callback func(ctx context.Context, objectName string)) {
	if now.Before(fileInfo.ModTime().Add(deleteTime)) {
		return
	}
	err := os.Remove(filePath)
	if err != nil {
		if !os.IsNotExist(err) {
			log.WithError(err).Errorf("Failed to delete file %s", filePath)
		}
		return
	}
	log.Infof("Deleted expired file %s", filePath)
	callback(ctx, filePath)
}
