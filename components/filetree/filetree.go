// TODO: filetree comp functionality:
// ADD more then 1 library path
// DELETE a path from filetree
// allow navigation to select series and manga
// toggle and show expanded series files
// allow an alt window to configure and set your library collection to be viewed in the main filetree comp
package filetree

import (
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

func clearErrorAfter(t time.Duration) tea.Cmd {
	return tea.Tick(t, func(_ time.Time) tea.Msg {
		return clearErrMsg{}
	})
}

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
	height          int
	width           int
	err             error
}

// init initial model
func NewFileTreeModel(cfg *config.TuiConfig) FileTreeModel {
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
	return ftm, nil
}

func (m FileTreeModel) View() string {
	return ""
}

func (ftm FileTreeModel) Init() tea.Cmd {
	return ftm.filePickerModel.Init()
}
