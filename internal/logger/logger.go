package logger

import (
	"io"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/davecgh/go-spew/spew"
)

type LoggerModel struct {
	dump io.Writer
}

func newLogger(dump *os.File) *LoggerModel {
	return &LoggerModel{
		dump: dump,
	}
}

func (lm LoggerModel) Init() tea.Cmd {
	return nil
}

func (lm LoggerModel) Update(msg tea.Msg) (LoggerModel, tea.Cmd) {
	if lm.dump != nil {
		spew.Fdump(lm.dump, msg)
	}
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC:
			return lm, tea.Quit
		}
	}
	return lm, nil
}

func (lm LoggerModel) View() string {
	return ""
}

func InitializeLogger() *LoggerModel {
	var dump *os.File
	if _, ok := os.LookupEnv("DEBUG"); ok {
		var err error
		dump, err = os.OpenFile("messages.log", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0o644)
		if err != nil {
			os.Exit(1)
		}
	}

	m := newLogger(dump)
	return m
}
