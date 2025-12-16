// TODO: filetree comp functionality:
// ADD more then 1 library path
// DELETE a path from filetree
// allow navigation to select series and manga
// toggle and show expanded series files
// allow an alt window to configure and set your library collection to be viewed in the main filetree comp
package filetree

import (
	"errors"
	"os"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/filepicker"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/saul178/mangareadertui/internal/config"
)

type (
	pathSelectedMsg string
	// cancelPathSelectedMsg struct{}
	clearErrMsg struct{}
)

type fileTreeState int

const (
	stateLibraryView    fileTreeState = iota // standard mode that will render along side the rest of the tui
	stateSelectPathView                      // add mode alt window that will pop up for them to select their collection path
)

type FileTreeModel struct {
	// data model list of paths that the user selects for their manga library, and stores in a conf.json
	compState         fileTreeState
	config            *config.TuiConfig
	mangaLibraryRoots []string // maybe it should be a map?
	expandedPaths     map[string]bool
	cursor            int
	selectedSeries    string // keep track of what is selected
	selectedManga     string
	offset            int

	// --- Component: Path Picker ---
	// This is the bubble used ONLY when state == stateSelectPathView.
	// You allow the user to navigate the OS here.
	// When they press "Select", you append the path to LibraryRoots and switch state back.
	filePickerModel filepicker.Model
	err             error
}

var ErrGettingHomeDir error = errors.New("Error getting home directory")

func clearErrorAfter(t time.Duration) tea.Cmd {
	return tea.Tick(t, func(_ time.Time) tea.Msg {
		return clearErrMsg{}
	})
}

// -- file picker state  --
func NewFilePickerModel(cfg *config.TuiConfig) FileTreeModel {
	fp := filepicker.New()
	fp.AllowedTypes = nil
	fp.DirAllowed = true
	fp.FileAllowed = false
	fp.ShowHidden = true // TODO: have this be toggled by user

	return FileTreeModel{
		compState:         stateLibraryView,
		config:            cfg,
		mangaLibraryRoots: cfg.CollectionPaths,
		expandedPaths:     make(map[string]bool),
		filePickerModel:   fp,
	}
}

func (ftm FileTreeModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// TODO: instead of quiting the app have it Toggle the component?
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case clearErrMsg:
		ftm.err = nil

	case tea.KeyMsg:
		switch msg.String() {
		case "a":
			ftm.compState = stateSelectPathView
		case "enter":
			ftm.mangaLibraryRoots = append(ftm.mangaLibraryRoots, ftm.filePickerModel.CurrentDirectory)
			ftm.compState = stateLibraryView
		}
	case pathSelectedMsg:
		path := string(msg)
		ftm.config.CollectionPaths = append(ftm.config.CollectionPaths, path)
		ftm.mangaLibraryRoots = ftm.config.CollectionPaths

		if err := config.SaveConfig(ftm.config); err != nil {
			ftm.err = errors.New(err.Error())
			return ftm, clearErrorAfter(time.Second * 2)
		}
		ftm.compState = stateLibraryView
		return ftm, cmd

	}
	return ftm, cmd
}

func (ftm FileTreeModel) View() string {
	var s strings.Builder
	s.WriteString("\n ")
	if ftm.compState == stateSelectPathView {
		return ftm.filePickerModel.View()
	}

	if len(ftm.mangaLibraryRoots) == 0 {
		s.WriteString("No collection paths added yet, push some button to add")
		return s.String()
	}
	s.WriteString("testing initial view")
	return s.String()
}

func (ftm FileTreeModel) Init() tea.Cmd {
	return ftm.filePickerModel.Init()
}
