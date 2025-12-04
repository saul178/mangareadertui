package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/saul178/mangareadertui/components/filetree"
)

// TODO: eventually main will stitch everything together here
// this will hold all the components child models used for the main parent app to use
// the appModel will choose which child component will be used
type AppModel struct {
	filetree filetree.FileTreeModel
	// ImageViewer component
	// Search/filter component
	// status bar component
	// other states that the main app will need
	// width  int
	// height int
}

// TODO: this will initialize all subcomponents
func (am AppModel) Init() tea.Cmd {
	am.filetree = filetree.FileTreeModel{}

	return am.filetree.Init()
}

// TODO: Main application Update: Delegates messages and handles custom messages.
// depending on which component is selected
func (fm AppModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q", "esc":
			return fm, tea.Quit
		}
	}

	return fm, cmd
}

// TODO: this will define the whole view of the tui by rendering the sub components togather
func (am AppModel) View() string {
	return am.filetree.View()
}

func main() {
	p := tea.NewProgram(AppModel{})
	if _, err := p.Run(); err != nil {
		fmt.Println("Alas, there's been an error: ", err)
		os.Exit(1)
	}
}
