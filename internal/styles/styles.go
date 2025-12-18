package styles

import "github.com/charmbracelet/lipgloss"

var (
	styleDim      = lipgloss.NewStyle().Foreground(lipgloss.Color("#666666"))
	styleSelected = lipgloss.NewStyle().Background(lipgloss.Color("#444444")).Bold(true)

	dirStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("#7AA2F7")).MarginRight(1)
	fileStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#FDFBD0"))
	enumStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#B4DEF5"))
)
