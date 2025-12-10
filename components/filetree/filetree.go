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

type fileTreeState int

const (
	stateLibraryView    fileTreeState = iota // main component that will render along side the rest of the tui
	stateSelectPathView                      // alt window that will pop up for them to select their collection path
)

const cbzExt = ".cbz"

type FileTreeModel struct {
	// data model list of paths that the user selects for their manga library, and stores in a conf.json
	compState         fileTreeState
	config            config.TuiConfig
	mangaLibraryRoots []string // maybe it should be a map?
	cursor            int
	offset            int
	selectedSeries    string
	selectedManga     string // might not need this value since the series is being stored in a map[str][]str
	expandedPaths     map[string]bool

	// --- Component: Path Picker ---
	// This is the bubble used ONLY when state == stateSelectPathView.
	// You allow the user to navigate the OS here.
	// When they press "Select", you append the path to LibraryRoots and switch state back.
	filePickerModel filepicker.Model
	err             error
}

var ErrGettingHomeDir error = errors.New("Error getting home directory")

type clearErrMsg struct{}

func clearErrorAfter(t time.Duration) tea.Cmd {
	return tea.Tick(t, func(_ time.Time) tea.Msg {
		return clearErrMsg{}
	})
}

func (ftm FileTreeModel) Init() tea.Cmd {
	return ftm.filePickerModel.Init()
}

func (ftm FileTreeModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// TODO: instead of quiting the app have it Toggle the component?
	// switch msg := msg.(type) {
	// case clearErrMsg:
	// 	ftm.Err = nil
	// }

	var cmd tea.Cmd
	ftm.filePickerModel, cmd = ftm.filePickerModel.Update(msg)

	// Did the user select a file?
	if didSelect, path := ftm.filePickerModel.DidSelectFile(msg); didSelect {
		// Get the path of the selected file.
		ftm.selectedFile = path
	}

	// Did the user select a disabled file?
	// This is only necessary to display an error to the user.
	if didSelect, path := ftm.filePickerModel.DidSelectDisabledFile(msg); didSelect {
		// Let's clear the selectedFile and display an error.
		ftm.err = errors.New(path + " is not valid.")
		ftm.selectedFile = ""
		return ftm, tea.Batch(cmd, clearErrorAfter(2*time.Second))
	}

	return ftm, cmd
}

func (ftm FileTreeModel) View() string {
	// if ftm.Toggle {
	//
	// 	return ""
	// }

	var s strings.Builder
	s.WriteString("\n ")
	if ftm.err != nil {
		s.WriteString(ftm.filePickerModel.Styles.DisabledFile.Render(ftm.err.Error()))
	} else if ftm.selectedFile == "" {
		s.WriteString("pick a file: ")
	} else {
		s.WriteString("Manga selected, starting Viewer... " + ftm.filePickerModel.Styles.Selected.Render(ftm.selectedFile))
	}
	s.WriteString("\n\n" + ftm.filePickerModel.View() + "\n")
	return s.String()
}

func NewFileTreeModel() (FileTreeModel, error) {
	fp := filepicker.New()
	fp.AllowedTypes = nil
	fp.DirAllowed = true
	fp.FileAllowed = false
	fp.ShowHidden = true

	// NOTE: for now this will just be set to the home dir until i get the config side going
	// TODO: handle the error better by defaulting to their home path if the config isn't set
	defaultDir, err := os.UserHomeDir()
	if err != nil {
		return FileTreeModel{}, errors.Join(ErrGettingHomeDir, err)
	}

	fp.CurrentDirectory = defaultDir
	return FileTreeModel{filePickerModel: fp}, nil
}
