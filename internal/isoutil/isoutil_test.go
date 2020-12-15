package isoutil

import (
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"testing"

	diskfs "github.com/diskfs/go-diskfs"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestIsoUtil(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Iso Util")
}

var _ = Describe("Isoutil", func() {
	var (
		isoDir   string
		isoFile  string
		filesDir string
		volumeID = "Assisted123"
	)
	BeforeSuite(func() {
		filesDir, isoDir, isoFile = createIsoViaGenisoimage(volumeID)
	})

	AfterSuite(func() {
		os.RemoveAll(filesDir)
		os.RemoveAll(isoDir)
	})

	Context("Extract", func() {
		It("extracts the files from an iso", func() {
			dir, err := ioutil.TempDir("", "isotest")
			Expect(err).ToNot(HaveOccurred())
			defer os.RemoveAll(dir)

			h := NewHandler(isoFile, dir).(*installerHandler)
			Expect(h.Extract()).To(Succeed())

			validateFileContent(filepath.Join(dir, "test"), "testcontent\n")
			validateFileContent(filepath.Join(dir, "testdir/stuff"), "morecontent\n")
		})
	})

	Context("Create", func() {
		It("generates an iso with the content in the given directory", func() {
			dir, err := ioutil.TempDir("", "isotest")
			Expect(err).ToNot(HaveOccurred())
			defer os.RemoveAll(dir)

			tmpIsoPath := filepath.Join(dir, "test.iso")
			h := NewHandler(tmpIsoPath, filepath.Join(filesDir, "files")).(*installerHandler)
			Expect(h.Create(isoFileSize(isoFile), "my-vol")).To(Succeed())

			d, err := diskfs.OpenWithMode(tmpIsoPath, diskfs.ReadOnly)
			Expect(err).ToNot(HaveOccurred())
			fs, err := d.GetFilesystem(0)
			Expect(err).ToNot(HaveOccurred())

			f, err := fs.OpenFile("/test", os.O_RDONLY)
			Expect(err).ToNot(HaveOccurred())
			content, err := ioutil.ReadAll(f)
			Expect(err).ToNot(HaveOccurred())
			Expect(string(content)).To(Equal("testcontent\n"))

			f, err = fs.OpenFile("/testdir/stuff", os.O_RDONLY)
			Expect(err).ToNot(HaveOccurred())
			content, err = ioutil.ReadAll(f)
			Expect(err).ToNot(HaveOccurred())
			Expect(string(content)).To(Equal("morecontent\n"))
		})
	})

	Context("fileExists", func() {
		It("returns true when file exists", func() {
			h := NewHandler("", filepath.Join(filesDir, "files")).(*installerHandler)
			exists, err := h.fileExists("test")
			Expect(err).ToNot(HaveOccurred())
			Expect(exists).To(BeTrue())

			exists, err = h.fileExists("testdir/stuff")
			Expect(err).ToNot(HaveOccurred())
			Expect(exists).To(BeTrue())
		})

		It("returns false when file does not exist", func() {
			h := NewHandler("", filepath.Join(filesDir, "files")).(*installerHandler)
			exists, err := h.fileExists("asdf")
			Expect(err).ToNot(HaveOccurred())
			Expect(exists).To(BeFalse())

			exists, err = h.fileExists("missingdir/things")
			Expect(err).ToNot(HaveOccurred())
			Expect(exists).To(BeFalse())
		})
	})

	Context("haveBootFiles", func() {
		It("returns true when boot files are present", func() {
			h := NewHandler("", filepath.Join(filesDir, "boot_files")).(*installerHandler)
			haveBootFiles, err := h.haveBootFiles()
			Expect(err).ToNot(HaveOccurred())
			Expect(haveBootFiles).To(BeTrue())
		})

		It("returns false when boot files are not present", func() {
			h := NewHandler("", filepath.Join(filesDir, "files")).(*installerHandler)
			haveBootFiles, err := h.haveBootFiles()
			Expect(err).ToNot(HaveOccurred())
			Expect(haveBootFiles).To(BeFalse())
		})
	})

	Context("readFile", func() {
		It("read existing file from ISO", func() {
			h := NewHandler(isoFile, "").(*installerHandler)
			reader, err := h.ReadFile("testdir/stuff")
			Expect(err).ToNot(HaveOccurred())
			fileContent, err := ioutil.ReadAll(reader)
			Expect(err).ToNot(HaveOccurred())
			Expect(string(fileContent)).To(Equal("morecontent\n"))
		})

		It("read non-existant file from ISO", func() {
			h := NewHandler(isoFile, "").(*installerHandler)
			reader, err := h.ReadFile("testdir/noexist")
			Expect(err).To(HaveOccurred())
			Expect(reader).To(BeNil())
		})
	})

	Context("getVolumeID", func() {
		It("get volume ID from existing ISO", func() {
			h := NewHandler(isoFile, "").(*installerHandler)
			volumeID, err := h.GetVolumeID()
			Expect(err).ToNot(HaveOccurred())
			Expect(volumeID).To(Equal(volumeID))
		})
	})
})

func isoFileSize(isoFile string) int64 {
	info, err := os.Stat(isoFile)
	Expect(err).NotTo(HaveOccurred())
	return info.Size()
}

func createIsoViaGenisoimage(volumeID string) (string, string, string) {
	filesDir, err := ioutil.TempDir("", "isotest")
	Expect(err).ToNot(HaveOccurred())

	isoDir, err := ioutil.TempDir("", "isotest")
	Expect(err).ToNot(HaveOccurred())
	isoFile := filepath.Join(isoDir, "test.iso")

	err = os.Mkdir(filepath.Join(filesDir, "files"), 0755)
	Expect(err).ToNot(HaveOccurred())
	err = ioutil.WriteFile(filepath.Join(filesDir, "files", "test"), []byte("testcontent\n"), 0664)
	Expect(err).ToNot(HaveOccurred())
	err = os.Mkdir(filepath.Join(filesDir, "files", "testdir"), 0755)
	Expect(err).ToNot(HaveOccurred())
	err = ioutil.WriteFile(filepath.Join(filesDir, "files", "testdir", "stuff"), []byte("morecontent\n"), 0664)
	Expect(err).ToNot(HaveOccurred())
	err = os.Mkdir(filepath.Join(filesDir, "boot_files"), 0755)
	Expect(err).ToNot(HaveOccurred())
	err = os.Mkdir(filepath.Join(filesDir, "boot_files", "images"), 0755)
	Expect(err).ToNot(HaveOccurred())
	err = ioutil.WriteFile(filepath.Join(filesDir, "boot_files", "images", "efiboot.img"), []byte(""), 0664)
	Expect(err).ToNot(HaveOccurred())
	err = os.Mkdir(filepath.Join(filesDir, "boot_files", "isolinux"), 0755)
	Expect(err).ToNot(HaveOccurred())
	err = ioutil.WriteFile(filepath.Join(filesDir, "boot_files", "isolinux", "boot.cat"), []byte(""), 0664)
	Expect(err).ToNot(HaveOccurred())
	err = ioutil.WriteFile(filepath.Join(filesDir, "boot_files", "isolinux", "isolinux.bin"), []byte(""), 0664)
	Expect(err).ToNot(HaveOccurred())
	cmd := exec.Command("genisoimage", "-rational-rock", "-J", "-joliet-long", "-V", volumeID, "-o", isoFile, filepath.Join(filesDir, "files"))
	err = cmd.Run()
	Expect(err).ToNot(HaveOccurred())

	return filesDir, isoDir, isoFile
}

func validateFileContent(filename string, content string) {
	fileContent, err := ioutil.ReadFile(filename)
	Expect(err).NotTo(HaveOccurred())
	Expect(string(fileContent)).To(Equal(content))
}
