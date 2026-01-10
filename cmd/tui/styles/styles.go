package styles

import "github.com/charmbracelet/lipgloss"

const (
	MainColor      = "#7AA2F7"
	SecondaryColor = "#B4DEF5"
	HighlightColor = "#FDFBD0"
	DisabledColor  = "#666666"
)

var (
	styleDim      = lipgloss.NewStyle().Foreground(lipgloss.Color(DisabledColor))
	styleSelected = lipgloss.NewStyle().Background(lipgloss.Color("#444444")).Bold(true) // think of a better selected color

	dirStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color(MainColor)).MarginRight(1)
	fileStyle = lipgloss.NewStyle().Foreground(lipgloss.Color(HighlightColor))
	enumStyle = lipgloss.NewStyle().Foreground(lipgloss.Color(SecondaryColor))
)
