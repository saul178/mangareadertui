package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/saul178/mangareadertui/components/filetree"
	"github.com/saul178/mangareadertui/internal/config"
)

// TODO: eventually main will stitch everything together here
// this will hold all the components child models used for the main parent app to use
// the appModel will choose which child component will be used
type appComponents int

var failedToStartTui string = "error starting mangareadertui: %v\n"

const (
	filetreeComp appComponents = iota
	imageViewerComp
	searchComp
	statusBarComp
)

type mainAppModel struct {
	// NOTE: subcomponents
	filetree filetree.FileTreeModel
	conf     *config.TuiConfig
	// imageViewer imageView.Model
	// search search.Model
	// statusBar statusBar.Model

	// NOTE: Main application state
	activeComp         appComponents // tracks which component is focused on
	toggleFileTreeComp bool
	width              int
	height             int
	err                error // report any errors encountered
}

// TODO: this will initialize all subcomponents
func (mam mainAppModel) Init() tea.Cmd {
	// batch up all the subcomp inits for the app
	return tea.Batch(
		mam.filetree.Init(),
		// am.ImageViewer etc
	)
}

// TODO: Main application Update: Delegates messages and handles custom messages.
// depending on which component is selected
func (mam mainAppModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q", "esc":
			return mam, tea.Quit
		// NOTE: idea to cycle through the components
		case "tab":
			mam.activeComp = (mam.activeComp + 1) % 4
			return mam, nil
		}
		// NOTE:
		// this is to resize the window of the main app but might need to send the msg down to the subcomponents too
		// case tea.WindowSizeMsg:
		// 	am.width = msg.Width
		// 	am.height = msg.Height
	}

	// TODO: delegate the msgs to the current active component
	switch mam.activeComp {
	case filetreeComp:
		newModel, newCmd := mam.filetree.Update(msg)
		mam.filetree = newModel.(filetree.FileTreeModel) // need to type assert here
		cmds = append(cmds, newCmd)
		// case imageViewerComp:
		// case searchComp:
		// case statusBarComp:
	}

	return mam, tea.Batch(cmds...)
}

// TODO: this will define the whole view of the tui by rendering the sub components togather
// using lipgloss for style and composition
func (mam mainAppModel) View() string {
	// TODO: use lipgloss for layout
	// example: Define a style for the focused component to highlight it.

	// Render file tree view
	fileTreeView := mam.filetree.View()

	// TODO: render the main content
	// mainContentView := "No main content selected."
	// if mam.activeComp == imageViewerComp {
	// 	mainContentView = mam.imageViewer.View()
	// } else if mam.activeComp == searchComp {
	// 	mainContentView = mam.search.View()
	// } else {
	mainContentView := "this is a tmp str"
	// }
	//
	// TODO: Combine them (using simple string concatenation for this example)
	// In a real TUI, you would use lipgloss to tile or stack these views.
	output := lipgloss.JoinHorizontal(
		lipgloss.Top,
		// File tree gets 30% width
		lipgloss.PlaceVertical(mam.height, lipgloss.Top, fileTreeView),
		// Main content gets the rest
		lipgloss.PlaceVertical(mam.height, lipgloss.Top, mainContentView),
	)
	// add status bar at the bottom
	// return lipgloss.JoinVertical(lipgloss.Left, output, mam.statusBar.View())
	return output
}

func main() {
	ft, err := filetree.NewFileTreeModel()
	if err != nil {
		fmt.Printf("failed to initialize file tree component: %w\n", err)
	}

	// init conf on app start up
	cfg, err := config.LoadConfig()
	if err != nil {
		fmt.Printf("failed to load conf: %w\n", err)
	}
	mangareadertui := mainAppModel{
		activeComp: filetreeComp, // set default active comp
		filetree:   ft,
		conf:       cfg,
	}

	p := tea.NewProgram(&mangareadertui) // TODO: look into other rendering options
	if _, err := p.Run(); err != nil {
		fmt.Printf(failedToStartTui, err)
		os.Exit(1)
	}
}
