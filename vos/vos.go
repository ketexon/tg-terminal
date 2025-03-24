package vos

import (
	"path"

	"github.com/spf13/afero"
)

var VOSMap = map[string]VOS{
	"user1": &LinuxVOS{
		Cwd: "/home/user1",
		fsLoader: func() afero.Fs {
			return afero.NewBasePathFs(afero.NewOsFs(), "./vfs/machine1")
		},
	},
}

type VOS interface {
	Init()
	Prompt() string
	StartProcess(string, []string) VProc
	GetFs() afero.Fs
	GetCwd() string
	SetCwd(string)
}

func resolvePath(vos VOS, pathStr string) string {
	if path.IsAbs(pathStr) {
		return path.Clean(pathStr)
	}
	return path.Clean(path.Join(vos.GetCwd(), pathStr))
}
