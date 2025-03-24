package vshell

import (
	"ketexon/tg/vos"
	"log/slog"
	"reflect"
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	logger        *slog.Logger
	input         textinput.Model
	vos           vos.VOS
	buffer        *strings.Builder
	activeProcess vos.VProc
	history       []string
	historyIndex  int
}

func New(vos vos.VOS) *Model {
	return &Model{
		logger:       slog.Default().WithGroup("tui").WithGroup("os"),
		vos:          vos,
		buffer:       &strings.Builder{},
		history:      make([]string, 1),
		historyIndex: 0,
	}
}

func (m *Model) Init() tea.Cmd {
	if m.vos == nil {
		m.logger.Error("VOS is nil")
		return tea.Quit
	}
	m.vos.Init()
	m.input = textinput.New()
	m.input.Prompt = m.vos.Prompt()
	return tea.Sequence(
		m.input.Focus(),
		textinput.Blink,
	)
}

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	if m.activeProcess != nil {
		switch msg.(type) {
		case vos.ProcessExited:
			m.buffer.WriteString(m.activeProcess.View())
			m.activeProcess = nil
			m.updateInput()
			return m, textinput.Blink
		default:
			var newProc tea.Model
			newProc, cmd = m.activeProcess.Update(msg)
			m.activeProcess = newProc.(vos.VProc)
			slog.Info("active process msg", "msg", msg, "type", reflect.TypeOf(msg))
			return m, cmd
		}
	}

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter:
			return m, m.submit()
		case tea.KeyCtrlC:
			m.input.SetValue(m.input.Value() + "^C")
			m.resetInput(false)
		case tea.KeyCtrlP, tea.KeyUp:
			m.changeHistory(-1)
		case tea.KeyCtrlN, tea.KeyDown:
			m.changeHistory(1)
		}
	}
	m.input, cmd = m.input.Update(msg)
	m.updateHistory()
	return m, cmd
}

func (m *Model) View() string {
	builder := &strings.Builder{}
	builder.WriteString(m.buffer.String())
	if m.activeProcess != nil {
		builder.WriteString(m.activeProcess.View())
	} else {
		builder.WriteString(m.input.View())
	}
	return builder.String()
}

func (m *Model) changeHistory(delta int) {
	newIndex := m.historyIndex + delta
	if newIndex >= 0 && newIndex < len(m.history) {
		m.historyIndex = newIndex
		m.input.SetValue(m.history[m.historyIndex])
		m.input.CursorEnd()
	}
}

func (m *Model) updateHistory() {
	input := m.input.Value()
	m.history[m.historyIndex] = input
}

func parseCommand(cmd string) []string {
	args := make([]string, 0)
	inQuotes := false
	quoteDelim := '\''
	escaping := false
	var currentArg strings.Builder

	for _, r := range cmd {
		if escaping {
			currentArg.WriteRune(r)
			escaping = false
			continue
		}
		if r == '\\' {
			escaping = true
			continue
		}
		if r == '\'' || r == '"' {
			if inQuotes && r == quoteDelim {
				inQuotes = false
				args = append(args, currentArg.String())
				currentArg.Reset()
				continue
			}
			if !inQuotes {
				inQuotes = true
				quoteDelim = r
				continue
			}
		}
		if r == ' ' && !inQuotes {
			if currentArg.Len() > 0 {
				args = append(args, currentArg.String())
				currentArg.Reset()
			}
			continue
		}
		currentArg.WriteRune(r)
	}

	if currentArg.Len() > 0 {
		args = append(args, currentArg.String())
	}

	return args
}

func (m *Model) updateInput() {
	m.input.Prompt = m.vos.Prompt()
}

func (m *Model) resetInput(addHistory bool) {
	input := m.input.Value()
	m.buffer.WriteString(m.input.Prompt + input + "\n")
	m.historyIndex = len(m.history) - 1
	m.history[m.historyIndex] = input
	if addHistory && input != "" {
		m.history = append(m.history, "")
		m.historyIndex = len(m.history) - 1
	}
	m.input.Reset()
	m.updateHistory()
}

func (m *Model) submit() tea.Cmd {
	input := m.input.Value()
	args := parseCommand(input)

	m.resetInput(true)

	if len(args) == 0 {
		return nil
	}

	procName := args[0]

	m.activeProcess = m.vos.StartProcess(procName, args)

	if m.activeProcess == nil {
		m.buffer.WriteString("Unknown command: " + procName + "\n")
		return nil
	}

	cmd := m.activeProcess.Init()
	startProcessCmd := func() tea.Msg { return vos.ProcessStarted{} }
	if cmd != nil {
		return tea.Sequence(cmd, startProcessCmd)
	} else {
		return startProcessCmd
	}
}
