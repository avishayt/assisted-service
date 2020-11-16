package s3wrapper

import (
	"context"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"os/exec"
	"path/filepath"
	"testing"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/sirupsen/logrus"
)

func TestJob(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Util")
}

var _ = Describe("FixEndpointURL", func() {
	It("returns the original string with a valid http URL", func() {
		endpoint := "http://example.com/stuff"
		result, err := FixEndpointURL(endpoint)
		Expect(err).ToNot(HaveOccurred())
		Expect(result).To(Equal("http://example.com/stuff"))
	})

	It("returns the original string with a valid https URL", func() {
		endpoint := "https://example.com/stuff"
		result, err := FixEndpointURL(endpoint)
		Expect(err).ToNot(HaveOccurred())
		Expect(result).To(Equal("https://example.com/stuff"))
	})

	It("prefixes an invalid endpoint with http:// when S3_USE_SSL is not set", func() {
		endpoint := "example.com"
		result, err := FixEndpointURL(endpoint)
		Expect(err).ToNot(HaveOccurred())
		Expect(result).To(Equal("http://example.com"))
	})

	It("prefixes and invalid endpoint with https:// when S3_USE_SSL is \"true\"", func() {
		endpoint := "example.com"
		os.Setenv("S3_USE_SSL", "true")
		defer os.Unsetenv("S3_USE_SSL")
		result, err := FixEndpointURL(endpoint)
		Expect(err).ToNot(HaveOccurred())
		Expect(result).To(Equal("https://example.com"))
	})

	It("returns an error when a prefix does not produce a valid URL", func() {
		endpoint := ":example.com"
		result, err := FixEndpointURL(endpoint)
		Expect(result).To(Equal(""))
		Expect(err).To(HaveOccurred())
	})
})

var _ = Describe("UploadBootFiles", func() {
	var (
		ctx          = context.Background()
		log          logrus.FieldLogger
		ctrl         *gomock.Controller
		mockS3Client *MockAPI
	)

	BeforeEach(func() {
		log = logrus.New()
		ctrl = gomock.NewController(GinkgoT())
		mockS3Client = NewMockAPI(ctrl)
	})

	AfterEach(func() {
		ctrl.Finish()
	})

	It("all files already uploaded", func() {
		mockS3Client.EXPECT().DoAllBootFilesExist(ctx, RHCOSBaseObjectName).Return(true, nil)
		err := UploadBootFiles(ctx, log, RHCOSBaseObjectName, RHCOSBaseURL, mockS3Client)
		Expect(err).ToNot(HaveOccurred())
	})

	It("missing iso and rootfs", func() {
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			filesDir, err := ioutil.TempDir("", "isotest")
			Expect(err).ToNot(HaveOccurred())
			err = os.MkdirAll(filepath.Join(filesDir, "files/images/pxeboot"), 0755)
			Expect(err).ToNot(HaveOccurred())
			err = ioutil.WriteFile(filepath.Join(filesDir, "files/images/pxeboot/rootfs.img"), []byte("this is rootfs"), 0664)
			Expect(err).ToNot(HaveOccurred())
			isoPath := filepath.Join(filesDir, "file.iso")
			cmd := exec.Command("genisoimage", "-rational-rock", "-J", "-joliet-long", "-V", "volumeID", "-o", isoPath, filepath.Join(filesDir, "files"))
			err = cmd.Run()
			Expect(err).ToNot(HaveOccurred())
			file, err := os.Open(isoPath)
			Expect(err).ToNot(HaveOccurred())
			defer file.Close()
			_, err = io.Copy(w, file)
			Expect(err).ToNot(HaveOccurred())
		}))
		defer ts.Close()

		mockS3Client.EXPECT().DoAllBootFilesExist(ctx, RHCOSBaseObjectName).Return(false, nil)
		mockS3Client.EXPECT().DoesObjectExist(ctx, RHCOSBaseObjectName).Return(false, nil)
		mockS3Client.EXPECT().DoesObjectExist(ctx, BootFileTypeToObjectName(RHCOSBaseObjectName, "initrd.img")).Return(true, nil)
		mockS3Client.EXPECT().DoesObjectExist(ctx, BootFileTypeToObjectName(RHCOSBaseObjectName, "rootfs.img")).Return(false, nil)
		mockS3Client.EXPECT().DoesObjectExist(ctx, BootFileTypeToObjectName(RHCOSBaseObjectName, "vmlinuz")).Return(true, nil)

		mockS3Client.EXPECT().UploadFileToPublicBucket(ctx, gomock.Any(), RHCOSBaseObjectName)
		mockS3Client.EXPECT().UploadStreamToPublicBucket(ctx, gomock.Any(), BootFileTypeToObjectName(RHCOSBaseObjectName, "rootfs.img"))

		err := UploadBootFiles(ctx, log, RHCOSBaseObjectName, ts.URL, mockS3Client)
		Expect(err).ToNot(HaveOccurred())
	})
})
