package vos

import (
	"fmt"

	"github.com/spf13/afero"
)

var linuxProcesses = map[string](func() VProc){
	"echo": func() VProc { return NewEcho() },
	"ls":   func() VProc { return NewLs() },
	"cd":   func() VProc { return NewCd() },
	"cat":  func() VProc { return NewCat() },
}

type fsLoader func() afero.Fs

type LinuxVOS struct {
	afero.Fs
	Cwd string
	fsLoader
}

func (v *LinuxVOS) Init() {
	v.Fs = v.fsLoader()
}

func (v *LinuxVOS) Prompt() string {
	return fmt.Sprintf("%s> ", v.Cwd)
}

func (v *LinuxVOS) StartProcess(cmd string, args []string) VProc {
	if procFunc, ok := linuxProcesses[cmd]; ok {
		proc := procFunc()
		proc.Execute(v, args[0], args)
		return proc
	}
	return nil
}

func (v *LinuxVOS) GetFs() afero.Fs {
	return v.Fs
}

func (v *LinuxVOS) GetCwd() string {
	return v.Cwd
}

func (v *LinuxVOS) SetCwd(cwd string) {
	v.Cwd = cwd
}
