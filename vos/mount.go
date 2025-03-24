package vos

import (
	"archive/tar"
	"fmt"
	"io"
	"log/slog"
	"os"

	"github.com/spf13/afero"
)

func MountTar(tarPath string) afero.Fs {
	// try open tarPath
	f, err := os.Open(tarPath)
	if err != nil {
		slog.Error("failed to open tar file", "path", tarPath, "error", err)
		panic(err)
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

		fmt.Printf("Contents of %s\n", header.Name)
		if _, err := io.Copy(os.Stdout, tr); err != nil {
			slog.Error("failed to copy tar entry", "error", err)
		}
		fmt.Println()
	}
	return fs
}
