package keymaps

import "github.com/charmbracelet/bubbles/key"

type fileTreeKeyMap struct {
	Up     key.Binding
	Down   key.Binding
	Left   key.Binding
	Right  key.Binding
	Esc    key.Binding
	Enter  key.Binding
	Add    key.Binding
	Toggle key.Binding
	Delete key.Binding
	Help   key.Binding
	Quit   key.Binding
}

func NewFileTreeKeyMap() fileTreeKeyMap {
	return fileTreeKeyMap{
		Up: key.NewBinding(
			key.WithKeys("k", "up"),
			key.WithHelp("k/↑", "move up"),
		),
		Down: key.NewBinding(
			key.WithKeys("j", "down"),
			key.WithHelp("j/↓", "move down"),
		),
		Left: key.NewBinding(
			key.WithKeys("h", "left", "backspace"),
			key.WithHelp("h/←", "move left"),
		),
		Right: key.NewBinding(
			key.WithKeys("l", "right"),
			key.WithHelp("l/→", "move right"),
		),
		Esc: key.NewBinding(
			key.WithKeys("esc"),
			key.WithHelp("esc", "escape current active window"),
		),
		Enter: key.NewBinding(
			key.WithKeys("enter"),
			key.WithHelp("enter", "read selected manga"),
		),
		Add: key.NewBinding(
			key.WithKeys("a"),
			key.WithHelp("a", "add manga collection"),
		),
		Toggle: key.NewBinding(
			key.WithKeys("ctrl+t"),
			// TODO: change the help msg if theyre in the file picker state to toggle hidden files
			key.WithHelp("ctrl+t", "show hidden directories/toggle file tree"),
		),
		Delete: key.NewBinding(
			key.WithKeys("d"),
			key.WithKeys("d", "delete manga collection"),
		),
		Help: key.NewBinding(
			key.WithKeys("?"),
			key.WithHelp("?", "show help menu"),
		),
		Quit: key.NewBinding(
			key.WithKeys("q", "ctrl+c"),
			key.WithHelp("q/ctrl+c", "quit mangareadertui"),
		),
	}
}

// TODO: key mappings for when reading manga
// func newImageViewerKeyMap() imageViewerKeyMap {}
// type genericKeyMaps struct{}
