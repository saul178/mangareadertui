package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	filetree "github.com/saul178/mangareadertui/cmd/tui/components"
	"github.com/saul178/mangareadertui/internal/config"
)

// TODO: eventually main will stitch everything together here
// this will hold all the components child models used for the main parent app to use
// the appModel will choose which child component will be used

var failedToStartTui string = "error starting mangareadertui: %v\n"

const (
	minWidth  = 50
	minHeight = 100
)

type appComponents int

const (
	filetreeComp appComponents = iota
	imageViewerComp
	searchComp
	pageStatusBarComp
	helpComp
)

type mainAppModel struct {
	// NOTE: subcomponents
	conf     *config.TuiConfig
	filetree filetree.FileTreeModel
	// imageViewer imageView.Model
	// search search.Model
	// statusBar statusBar.Model

	// NOTE: Main application state
	activeComp         appComponents // tracks which component is focused on
	toggleFileTreeComp bool
	toggleHelpComp     bool
	width              int
	height             int
	err                error // report any errors encountered
}

// TODO: this will initialize all subcomponents
func (mam mainAppModel) Init() tea.Cmd {
	// batch up all the subcomp inits for the app
	// TODO: maybe this should be done sequentially?
	// return tea.Batch(
	// 	mam.filetree.Init(),
	// 	// am.ImageViewer etc
	// )
	return nil
}

// TODO: Main application Update: Delegates messages and handles custom messages.
// depending on which component is selected
func (mam mainAppModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		mam.width = msg.Width
		mam.height = msg.Height
		if mam.width < minWidth || mam.height < minHeight {
			// TODO: present an error screen that informs the user that the window size is too small
			// for the app
		}

		// Optionally forward to children if they need it
		// For now, filetree doesn't need it, but image viewer will!
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return mam, tea.Quit
			// NOTE: idea to cycle through the components
			// case "tab":
			// 	mam.activeComp = (mam.activeComp + 1) % 4
			// 	return mam, nil
		}
	}

	switch mam.activeComp {
	case filetreeComp:
		mam.filetree, cmd = mam.filetree.Update(msg)
		cmds = append(cmds, cmd)
		// case imageViewerComp:
		// case searchComp:
		// case statusBarComp:
	}

	// TODO: delegate the msgs to the current active component

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
	mainContentView := "\t\t\t\tthis is a tmp str"
	// }
	//
	// TODO: Combine them (using simple string concatenation for this example)
	// In a real TUI, you would use lipgloss to tile or stack these views.
	output := lipgloss.JoinHorizontal(
		lipgloss.Top,
		lipgloss.PlaceVertical(mam.height, lipgloss.Top, fileTreeView),
		lipgloss.PlaceVertical(mam.height, lipgloss.Top, mainContentView),
	)
	// add status bar at the bottom
	// return lipgloss.JoinVertical(lipgloss.Left, output, mam.statusBar.View())
	return output
}

func main() {
	// init conf on app start up
	cfg, err := config.LoadConfig()
	if err != nil {
		// TODO: log this error to some log file
		fmt.Printf("failed to load conf: %v\n", err)
	}

	ft := filetree.NewFileTreeModel(cfg)

	mangareadertui := mainAppModel{
		activeComp: filetreeComp, // set default active comp
		filetree:   ft,
		conf:       cfg,
	}

	p := tea.NewProgram(&mangareadertui, tea.WithAltScreen()) // TODO: look into other rendering options
	if _, err := p.Run(); err != nil {
		fmt.Printf(failedToStartTui, err)
		os.Exit(1)
	}
}
