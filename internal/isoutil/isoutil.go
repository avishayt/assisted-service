package isoutil

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"

	diskfs "github.com/diskfs/go-diskfs"
	"github.com/diskfs/go-diskfs/disk"
	"github.com/diskfs/go-diskfs/filesystem"
	"github.com/diskfs/go-diskfs/filesystem/iso9660"
)

type Handler interface {
	Extract(isoPath string) error
	Create(isoPath string, size int64, volumeLabel string) error
}

type installerHandler struct {
	workDir string
}

func NewHandler(workDir string) Handler {
	return &installerHandler{workDir: workDir}
}

func (h *installerHandler) Extract(isoPath string) error {
	d, err := diskfs.OpenWithMode(isoPath, diskfs.ReadOnly)
	if err != nil {
		return err
	}

	fs, err := d.GetFilesystem(0)
	if err != nil {
		return err
	}

	files, err := fs.ReadDir("/")
	if err != nil {
		return err
	}
	err = copyAll(fs, "/", files, h.workDir)
	if err != nil {
		return err
	}

	return nil
}

// recursive function for unpacking all files and directores from the given iso filesystem starting at fsDir
func copyAll(fs filesystem.FileSystem, fsDir string, infos []os.FileInfo, targetDir string) error {
	for _, info := range infos {
		osName := filepath.Join(targetDir, info.Name())
		fsName := filepath.Join(fsDir, info.Name())

		if info.IsDir() {
			if err := os.Mkdir(osName, info.Mode().Perm()); err != nil {
				return err
			}

			files, err := fs.ReadDir(fsName)
			if err != nil {
				return err
			}
			if err := copyAll(fs, fsName, files[:], osName); err != nil {
				return err
			}
		} else {
			fsFile, err := fs.OpenFile(fsName, os.O_RDONLY)
			if err != nil {
				return err
			}
			osFile, err := os.Create(osName)
			if err != nil {
				return err
			}

			_, err = io.Copy(osFile, fsFile)
			if err != nil {
				osFile.Close()
				return err
			}

			if err := osFile.Close(); err != nil {
				return err
			}
		}
	}
	return nil
}

func (h *installerHandler) Create(isoPath string, size int64, volumeLabel string) error {
	d, err := diskfs.Create(isoPath, size, diskfs.Raw)
	if err != nil {
		return err
	}

	d.LogicalBlocksize = 2048
	fspec := disk.FilesystemSpec{Partition: 0, FSType: filesystem.TypeISO9660, VolumeLabel: volumeLabel}
	fs, err := d.CreateFilesystem(fspec)
	if err != nil {
		return err
	}

	addFileToISO := func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		p, err := filepath.Rel(h.workDir, path)
		if err != nil {
			return err
		}

		if info.IsDir() {
			return fs.Mkdir(p)
		}

		content, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}

		rw, err := fs.OpenFile(p, os.O_CREATE|os.O_RDWR)
		if err != nil {
			return err
		}

		_, err = rw.Write(content)
		return err
	}
	if err := filepath.Walk(h.workDir, addFileToISO); err != nil {
		return err
	}

	iso, ok := fs.(*iso9660.FileSystem)
	if !ok {
		return fmt.Errorf("not an iso9660 filesystem")
	}

	options := iso9660.FinalizeOptions{
		RockRidge:        true,
		VolumeIdentifier: volumeLabel,
	}

	if haveFiles, err := h.haveBootFiles(); err != nil {
		return err
	} else if haveFiles {
		options.ElTorito = &iso9660.ElTorito{
			BootCatalog: "isolinux/boot.cat",
			Entries: []*iso9660.ElToritoEntry{
				{
					Platform:  iso9660.BIOS,
					Emulation: iso9660.NoEmulation,
					BootFile:  "isolinux/isolinux.bin",
					BootTable: true,
					LoadSize:  4,
				},
				{
					Platform:  iso9660.EFI,
					Emulation: iso9660.NoEmulation,
					BootFile:  "images/efiboot.img",
				},
			},
		}
	}

	return iso.Finalize(options)
}

func (h *installerHandler) haveBootFiles() (bool, error) {
	files := []string{"isolinux/boot.cat", "isolinux/isolinux.bin", "images/efiboot.img"}
	for _, f := range files {
		if exists, err := h.fileExists(f); err != nil {
			return false, err
		} else if !exists {
			return false, nil
		}
	}

	return true, nil
}

func (h *installerHandler) fileExists(relName string) (bool, error) {
	name := filepath.Join(h.workDir, relName)
	if _, err := os.Stat(name); os.IsNotExist(err) {
		return false, nil
	} else if err != nil {
		return false, err
	}
	return true, nil
}
