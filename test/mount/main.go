package main

import (
	"fmt"
	"io/fs"
	"ketexon/tg/vos"

	"github.com/spf13/afero"
)

func main() {
	f, err := vos.MountTar("./vfs/machine1.tar")
	if err != nil {
		panic(err)
	}

	afero.Walk(f, "/", func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		fmt.Printf("%s (%s)\n", path, info.Mode())
		if !info.IsDir() {
			content, err := afero.ReadFile(f, path)
			if err != nil {
				return err
			}
			fmt.Printf("Content: \n```\n%s```\n", string(content))
		}
		return nil
	})
}
