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
func NewFileTreeModel(cfg *config.TuiConfig) *FileTreeModel {
	fp := filepicker.New()
	fp.AllowedTypes = nil
	fp.DirAllowed = true
	fp.FileAllowed = false
	fp.ShowHidden = true // TODO: have this be toggled by user

	return &FileTreeModel{
		compState:         stateLibraryView,
		config:            cfg,
		mangaLibraryRoots: cfg.CollectionPaths,
		expandedPaths:     make(map[string]bool),
		filePickerModel:   fp,
	}
}

func (ftm FileTreeModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "esc", "ctrl+c":
			return ftm, tea.Quit
		case "j", "down":
			// TODO: cursor logic: navigation is dependent if paths are expanded or not
		case "k", "up":
			// TODO: cursor logic: navigation is dependent if paths are expanded or not
		case "enter": //, "l", "right"
			// TODO: if its a directory then it should expand revealing the children dirs
			// if its a valid comic form file the we can do operations to extract the cbz to be viewed
		case "a":
			// TODO: here we change the state of the filetree component to filepicker component
			// it should open up a separate window where the user can navigate and select their path
			// the path selected should then be saved in the conf.json and recursively insert any sub dir
			// and valid cbz files
			ftm.compState = stateSelectPathView
		case "backspace": //, "h", "left"
			// TODO: allow to navigate back up one dir up to the root of the manga library only
		}
		// update which model to be used depending on state
		switch ftm.compState {
		case stateLibraryView:
		case stateSelectPathView:
			ftm.filePickerModel, cmd = ftm.filePickerModel.Update(msg)
		}
	}

	return ftm, nil
}

func (m FileTreeModel) View() string {
	return ""
}

func (ftm FileTreeModel) Init() tea.Cmd {
	return ftm.filePickerModel.Init()
}
