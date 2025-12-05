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
type appComponents int

const (
	filetreeComp appComponents = iota
	imageViewerComp
	searchComp
	statusBarComp
)

type mainAppModel struct {
	// NOTE: subcomponents
	filetree filetree.FileTreeModel
	// imageViewer imageView.Model
	// search search.Model
	// statusBar statusBar.Model

	// NOTE: Main application state
	// activeComp appComponents // tracks which component is focused on
	// width  int
	// height int
	// err error // report any errors encountered
}

// TODO: this will initialize all subcomponents
func (am mainAppModel) Init() tea.Cmd {
	ft, err := filetree.NewFileTreeModel()
	if err != nil {
		fmt.Printf("failed to initialize file tree component: %v\n", err)
		return nil
	}

	// assign the models to the mainAppModel to be initialized
	am.filetree = ft

	// set the initial focus
	// am.activeComp = filetreeComp

	// batch up all the subcomp inits for the app
	return tea.Batch(
		am.filetree.Init(),
	)
}

// TODO: Main application Update: Delegates messages and handles custom messages.
// depending on which component is selected
func (fm mainAppModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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
func (am mainAppModel) View() string {
	return am.filetree.View()
}

func main() {
	ft, err := filetree.NewFileTreeModel()
	if err != nil {
		fmt.Printf("failed to to start filetree component: %v\n", err)
		os.Exit(1)
	}
	p := tea.NewProgram(&ft)
	if _, err := p.Run(); err != nil {
		fmt.Println("Alas, there's been an error: ", err)
		os.Exit(1)
	}
}
