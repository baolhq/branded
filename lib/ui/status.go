package ui

import (
	"baolhq/branded/lib/meta"

	"github.com/charmbracelet/lipgloss"
)

var (
	chapterStyle = lipgloss.NewStyle().Width(meta.TermWidth/2 - 7).PaddingLeft(1).Bold(true)
	statsStyle   = lipgloss.NewStyle().Width(meta.TermWidth/2 + 5).Align(lipgloss.Right).PaddingRight(1)
)

// RenderStatus generates the status bar display
func RenderStatus() string {
	chapterStyle := chapterStyle.Render("[Kros the Champion             ]")
	stats := statsStyle.Render("Heavy Crossbow AC: 96 MV: 10 TN: 101")

	return lipgloss.NewStyle().
		Width(meta.StatusWidth).
		Border(lipgloss.NormalBorder(), true, false, false, false).
		Render(lipgloss.JoinHorizontal(lipgloss.Top, chapterStyle, stats))
}
