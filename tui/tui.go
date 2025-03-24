package tui

import (
	"ketexon/tg/message"
	"log/slog"
	"reflect"

	tea "github.com/charmbracelet/bubbletea"
)

type SwitchModel struct {
	NextModel tea.Model
}

type TUI struct {
	logger   *slog.Logger
	server   *message.Server
	curModel tea.Model
}

func NewTUI(server *message.Server, startModel tea.Model) *TUI {
	return &TUI{
		logger:   slog.Default().WithGroup("tui"),
		server:   server,
		curModel: startModel,
	}
}

func (t *TUI) Init() tea.Cmd {
	if t.curModel != nil {
		return t.curModel.Init()
	}
	return nil
}

func (t *TUI) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// Uncomment the following line to enable debug logging
	// for messages
	t.logger.Debug("Received message", "msg", msg, "type", reflect.TypeOf(msg))

	switch msg := msg.(type) {
	case SwitchModel:
		t.curModel = msg.NextModel
		return t, t.curModel.Init()
	case tea.KeyMsg:
		if msg.Type == tea.KeyCtrlD {
			t.logger.Info("Received Ctrl+D, quitting")
			return t, tea.Quit
		}
	}
	var cmd tea.Cmd
	if t.curModel != nil {
		t.curModel, cmd = t.curModel.Update(msg)
		return t, cmd
	}
	return t, nil
}

func (t *TUI) View() string {
	if t.curModel != nil {
		return t.curModel.View()
	}
	return "<no model>"
}
