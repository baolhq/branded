package ui

import (
	"baolhq/branded/lib/meta"

	"github.com/charmbracelet/lipgloss"
)

var (
	chapterStyle = lipgloss.NewStyle().Width(meta.StatusWidth/2 - 4).PaddingLeft(1).Bold(true)
	statsStyle   = lipgloss.NewStyle().Width(meta.StatusWidth/2 + 4).PaddingRight(1).Align(lipgloss.Right)
)

// UpdateStatus generates the status bar display
func UpdateStatus() string {
	chapterStyle := chapterStyle.Render("Chapter XVIII: Oath of Steel")
	stats := statsStyle.Render("CONQUEST - TURN: 99 READY: 99 ENEMIES: 99")

	return lipgloss.JoinHorizontal(lipgloss.Top, chapterStyle, stats)
}
