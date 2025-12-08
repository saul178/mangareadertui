// TODO: clean this up
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

const cbzExt = ".cbz"

type cbzSelectedMsg string

var errGettingHomeDir error = errors.New("Error getting home directory")

type FileTreeModel struct {
	filepicker         filepicker.Model
	config             config.TuiConfig
	collectionRootPath string
	selectedPath       string
	expanded           bool
	currentEntries     []os.DirEntry
	err                error
}

type clearErrMsg struct{}

func clearErrorAfter(t time.Duration) tea.Cmd {
	return tea.Tick(t, func(_ time.Time) tea.Msg {
		return clearErrMsg{}
	})
}

func (ftm FileTreeModel) Init() tea.Cmd {
	return ftm.filepicker.Init()
}

func (ftm FileTreeModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// TODO: instead of quiting the app have it Toggle the component?
	// switch msg := msg.(type) {
	// case clearErrMsg:
	// 	ftm.Err = nil
	// }

	var cmd tea.Cmd
	ftm.filepicker, cmd = ftm.filepicker.Update(msg)

	// Did the user select a file?
	if didSelect, path := ftm.filepicker.DidSelectFile(msg); didSelect {
		// Get the path of the selected file.
		ftm.selectedFile = path
	}

	// Did the user select a disabled file?
	// This is only necessary to display an error to the user.
	if didSelect, path := ftm.filepicker.DidSelectDisabledFile(msg); didSelect {
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
		s.WriteString(ftm.filepicker.Styles.DisabledFile.Render(ftm.err.Error()))
	} else if ftm.selectedFile == "" {
		s.WriteString("pick a file: ")
	} else {
		s.WriteString("Manga selected, starting Viewer... " + ftm.filepicker.Styles.Selected.Render(ftm.selectedFile))
	}
	s.WriteString("\n\n" + ftm.filepicker.View() + "\n")
	return s.String()
}

func NewFileTreeModel() (FileTreeModel, error) {
	fp := filepicker.New()
	fp.AllowedTypes = []string{cbzExt}
	fp.ShowHidden = true

	// NOTE: for now this will just be set to the home dir until i get the config side going
	// TODO: handle the error better by defaulting to their home path if the config isn't set
	defaultDir, err := os.UserHomeDir()
	if err != nil {
		return FileTreeModel{}, errors.Join(errGettingHomeDir, err)
	}

	fp.CurrentDirectory = defaultDir
	return FileTreeModel{filepicker: fp}, nil
}
