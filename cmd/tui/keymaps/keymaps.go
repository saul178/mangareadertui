package keymaps

import "github.com/charmbracelet/bubbles/key"

type GenericKeyMap struct {
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

func NewFileTreeKeyMap() GenericKeyMap {
	return GenericKeyMap{
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
