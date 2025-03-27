package ui

import "github.com/charmbracelet/lipgloss"

const Width = 80

var (
	nameStyle  = lipgloss.NewStyle().Width(Width/2 - 7).PaddingLeft(1).Bold(true)
	statsStyle = lipgloss.NewStyle().Width(Width/2 + 5).Align(lipgloss.Right).PaddingRight(1)
)

// RenderStatus generates the status bar display
func RenderStatus() string {
	name := nameStyle.Render("[Kros the Champion             ]")
	stats := statsStyle.Render("Heavy Crossbow AC: 96 MV: 10 TN: 101")

	return lipgloss.NewStyle().
		Width(Width - 2).
		Border(lipgloss.RoundedBorder()).
		Render(lipgloss.JoinHorizontal(lipgloss.Top, name, stats))
}
