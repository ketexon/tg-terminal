package vos

import (
	"archive/tar"
	"io"
	"log/slog"
	"os"
	"path"

	"github.com/spf13/afero"
)

func MountTar(tarPath string) (afero.Fs, error) {
	// try open tarPath
	f, err := os.Open(tarPath)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	tr := tar.NewReader(f)

	// create a new in-memory file system
	fs := afero.NewMemMapFs()

	// extract files from the tar archive
	for {
		header, err := tr.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			slog.Error("failed to read tar entry", "error", err)
		}

		name := path.Clean(path.Join("/", header.Name))
		mode := os.FileMode(header.Mode)

		switch header.Typeflag {
		case tar.TypeDir:
			if err := fs.MkdirAll(name, mode); err != nil {
				return nil, err
			}
		case tar.TypeReg:
			file, err := fs.OpenFile(name, os.O_CREATE|os.O_WRONLY, mode)
			if err != nil {
				return nil, err
			}
			defer file.Close()
			if _, err := io.Copy(file, tr); err != nil {
				return nil, err
			}
		}
	}
	return fs, nil
}
