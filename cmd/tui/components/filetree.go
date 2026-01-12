// TODO: filetree comp functionality:
// ADD more then 1 library path
// DELETE a path from filetree
// allow navigation to select series and manga
// toggle and show expanded series files
// allow an alt window to configure and set your library collection to be viewed in the main filetree comp
// create custom keymaps for filetree a = add d = delete etc
package components

import (
	"os"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/filepicker"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"

	// "github.com/charmbracelet/lipgloss"
	"github.com/saul178/mangareadertui/cmd/tui/keymaps"
	"github.com/saul178/mangareadertui/internal/config"
)

var keys = keymaps.NewFileTreeKeyMap()

type clearErrMsg struct{}

func clearErrorAfter(t time.Duration) tea.Cmd {
	return tea.Tick(t, func(_ time.Time) tea.Msg {
		return clearErrMsg{}
	})
}

type fileTreeState int

const (
	stateFileTree fileTreeState = iota
	stateFilePicker
)

type ConfigUpdatedMsg struct {
	PathSaved   bool
	PathDeleted bool
	Path        *string
}

type ErrMsg error

// TODO: reorganize the struct
type FileTreeModel struct {
	compState         fileTreeState
	filePickerModel   filepicker.Model
	config            *config.TuiConfig
	selectedField     string   // keep track of what is selected either current navigating directory or selected manga
	mangaLibraryRoots []string // i dont think i need this?
	expandedPaths     map[string]struct{}
	cursor            int // keeps track of the users position
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
		switch {
		case key.Matches(msg, keys.Quit):
			return ftm, tea.Quit
		case key.Matches(msg, keys.Esc):
			// TODO: depending on state this should just escape from the current active state up to the default
			// state which will be the file tree state
			if ftm.compState == stateFilePicker {
				ftm.compState = stateFileTree
			} else {
				return ftm, nil
			}
		case key.Matches(msg, keys.Down):
			// TODO: cursor logic: navigation is dependent if paths are expanded or not
		case key.Matches(msg, keys.Up):
			// TODO: cursor logic: navigation is dependent if paths are expanded or not
		case key.Matches(msg, keys.Left):
			// TODO: allow to navigate back up one dir up to the root of the manga library only
			// TODO: if its a directory then it should expand revealing the children dirs
			// if it's a normal file do nothing
		case key.Matches(msg, keys.Right):
			// TODO: move down one directory up to the end of it
			// if its a non empty directory then it should expand revealing the children dirs
			// if it's a normal file do nothing until they press enter to read their selectedManga
		case key.Matches(msg, keys.Enter): //
			// TODO: if its a valid comic file then we perform the operations for it to be read and load it
			// to the imageviewer
		case key.Matches(msg, keys.Delete):
			// TODO: this action should delete a path saved from the config file and save the changes
			// it must be the root directory or else im pretty sure weird behavior will happen if children
			// are deleted
		case key.Matches(msg, keys.Add):
			// TODO: here we change the state of the filetree component to filepicker component
			// it should open up a separate window where the user can navigate and select their path
			// the path selected should then be saved in the conf.json and recursively insert any sub dir
			// and valid cbz files
			ftm.compState = stateFilePicker
			ftm.filePickerModel.SetHeight(ftm.height - 10)
			cmds = append(cmds, ftm.filePickerModel.Init())
			return ftm, tea.Batch(cmds...)
		case key.Matches(msg, keys.Toggle):
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
		ftm.filePickerModel, cmd = ftm.filePickerModel.Update(msg)
		cmds = append(cmds, cmd)

		if didSelect, path := ftm.filePickerModel.DidSelectFile(msg); didSelect {
			ftm.config.CollectionPaths = append(ftm.config.CollectionPaths, path)
			err := config.SaveConfig(ftm.config)
			ftm.compState = stateFileTree
			return ftm, func() tea.Msg {
				if err != nil {
					return ErrMsg(err)
				}
				return &ConfigUpdatedMsg{
					PathSaved: didSelect,
					Path:      &path,
				}
			}
		}
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
