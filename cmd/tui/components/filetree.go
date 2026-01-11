// TODO: filetree comp functionality:
// ADD more then 1 library path
// DELETE a path from filetree
// allow navigation to select series and manga
// toggle and show expanded series files
// allow an alt window to configure and set your library collection to be viewed in the main filetree comp
package components

import (
	"os"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/filepicker"
	tea "github.com/charmbracelet/bubbletea"

	// "github.com/charmbracelet/lipgloss"
	"github.com/saul178/mangareadertui/internal/config"
)

type (
	pathSelectedMsg  string
	savedToConfigMsg string

	cancelPathSelectedMsg struct{}
	clearErrMsg           struct{}
)

func clearErrorAfter(t time.Duration) tea.Cmd {
	return tea.Tick(t, func(_ time.Time) tea.Msg {
		return clearErrMsg{}
	})
}

type fileTreeState int

const (
	stateFileTree   fileTreeState = iota // standard mode that will render along side the rest of the tui
	stateFilePicker                      // add mode alt window that will pop up for them to select their collection path
)

// TODO: reorganize the struct
type FileTreeModel struct {
	// data model list of paths that the user selects for their manga library, and stores in a conf.json
	compState         fileTreeState
	filePickerModel   filepicker.Model
	config            *config.TuiConfig
	selectedField     string   // keep track of what is selected
	mangaLibraryRoots []string // i dont think i need this?
	expandedPaths     map[string]struct{}
	cursor            int
	offset            int
	height            int
	width             int
	err               error
}

// init initial model
func NewFileTreeModel(cfg *config.TuiConfig) FileTreeModel {
	fp := filepicker.New()
	fp.CurrentDirectory, _ = os.UserHomeDir() // TODO: dont ignore the error on prod and think of a better way to do this
	fp.AllowedTypes = nil
	fp.DirAllowed = true
	fp.FileAllowed = false
	fp.ShowHidden = false

	return FileTreeModel{
		compState:         stateFileTree,
		config:            cfg,
		mangaLibraryRoots: cfg.CollectionPaths,
		expandedPaths:     make(map[string]struct{}),
		filePickerModel:   fp,
	}
}

func (ftm FileTreeModel) Update(msg tea.Msg) (FileTreeModel, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		ftm.width = msg.Width
		ftm.height = msg.Height

	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return ftm, tea.Quit
		case "esc":
			if ftm.compState == stateFilePicker {
				ftm.compState = stateFileTree
			} else {
				return ftm, nil
			}
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
			ftm.compState = stateFilePicker
			ftm.filePickerModel.SetHeight(ftm.height - 10)
			cmds = append(cmds, ftm.filePickerModel.Init())
			return ftm, tea.Batch(cmds...)
		case "backspace": //, "h", "left"
			// TODO: allow to navigate back up one dir up to the root of the manga library only
		case "ctrl+a":
			// toggle hidden files when in filepicker mode
			if ftm.compState == stateFilePicker {
				ftm.filePickerModel.ShowHidden = !ftm.filePickerModel.ShowHidden
				cmds = append(cmds, ftm.filePickerModel.Init())
			} else {
				return ftm, nil
			}
		}

	}

	// update which model to be used depending on state
	switch ftm.compState {
	case stateFilePicker:
		if didSelect, path := ftm.filePickerModel.DidSelectFile(msg); didSelect {
			ftm.config.CollectionPaths = append(ftm.config.CollectionPaths, path)
			config.SaveConfig(ftm.config)
			// TODO: display a success msg or fail msg if saving succeeded
			ftm.compState = stateFileTree
			return ftm, nil
		}
		ftm.filePickerModel, cmd = ftm.filePickerModel.Update(msg)
		cmds = append(cmds, cmd)
	case stateFileTree:
	}
	return ftm, tea.Batch(cmds...)
}

func (ftm FileTreeModel) View() string {
	var s strings.Builder
	switch ftm.compState {
	case stateFilePicker:
		// TODO: render this correctly with lipgloss and on it's own window
		s.WriteString("\n\n" + ftm.filePickerModel.View() + "\n")
		return s.String()
	default:
		return "this will soon be a filetree"
	}
}

func (ftm FileTreeModel) Init() tea.Cmd {
	return ftm.filePickerModel.Init()
}
