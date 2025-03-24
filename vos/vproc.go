package vos

import (
	tea "github.com/charmbracelet/bubbletea"
)

type VProc interface {
	tea.Model
	Execute(VOS, string, []string)
}

type ProcessStarted struct{}
type ProcessExited struct{}
