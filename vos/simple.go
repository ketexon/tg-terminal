package vos

import (
	"path"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/afero"
)

type SimpleVProc struct {
	fn   func(vos VOS, cmd string, args []string) string
	text string
}

var _ VProc = (*SimpleVProc)(nil)

func (s *SimpleVProc) Execute(vos VOS, cmd string, args []string) {
	s.text = s.fn(vos, cmd, args)
}

func (s *SimpleVProc) Init() tea.Cmd {
	return nil
}

func (s *SimpleVProc) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return s, func() tea.Msg {
		return ProcessExited{}
	}
}

func (s *SimpleVProc) View() string {
	return s.text
}

func NewSimple(fn func(vos VOS, cmd string, args []string) string) *SimpleVProc {
	return &SimpleVProc{
		fn: fn,
	}
}

func NewEcho() VProc {
	return NewSimple(func(vos VOS, cmd string, args []string) string {
		return strings.Join(args[1:], " ")
	})
}

func NewLs() VProc {
	return NewSimple(func(vos VOS, cmd string, args []string) string {
		if len(args) > 2 {
			return "Usage: ls [directory]\n"
		}
		fs := vos.GetFs()
		var dirPath string
		if len(args) > 1 {
			dirPath = resolvePath(vos, args[1])
		} else {
			dirPath = path.Clean(vos.GetCwd())
		}

		dir, err := fs.Open(dirPath)
		if err != nil {
			return "Error opening directory: " + err.Error() + "\n"
		}
		defer dir.Close()
		files, err := dir.Readdirnames(-1)
		if err != nil {
			return "Error reading directory: " + err.Error() + "\n"
		}
		return strings.Join(files, "\n") + "\n"
	})
}

func NewCd() *SimpleVProc {
	return NewSimple(func(vos VOS, cmd string, args []string) string {
		if len(args) < 2 {
			return "Usage: cd <directory>\n"
		}
		if len(args) > 2 {
			return "Too many arguments\n"
		}
		var dirPath string
		if len(args) > 1 {
			dirPath = resolvePath(vos, args[1])
		} else {
			dirPath = path.Clean(vos.GetCwd())
		}

		fs := vos.GetFs()
		exists, err := afero.DirExists(fs, dirPath)
		if err != nil {
			return "Error checking directory: " + err.Error() + "\n"
		}
		if !exists {
			return "No such directory: " + dirPath + "\n"
		}
		vos.SetCwd(dirPath)
		return ""
	})
}

func NewCat() VProc {
	return NewSimple(func(vos VOS, cmd string, args []string) string {
		if len(args) < 2 {
			return "Usage: cat <file>\n"
		}
		sb := strings.Builder{}
		for _, arg := range args[1:] {
			var filePath string = resolvePath(vos, arg)
			fs := vos.GetFs()
			exists, err := afero.Exists(fs, filePath)
			if err != nil {
				return "Error checking file: " + err.Error() + "\n"
			}
			if !exists {
				return "No such file: " + filePath + "\n"
			}

			data, err := afero.ReadFile(fs, filePath)
			if err != nil {
				return "Error reading file: " + err.Error() + "\n"
			}
			sb.Write(data)
		}
		return sb.String()
	})
}
