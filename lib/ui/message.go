package ui

import (
	"baolhq/branded/lib/meta"

	"github.com/charmbracelet/bubbles/viewport"
	"github.com/charmbracelet/lipgloss"
)

var (
	oldMessageStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#888888"))
	newMessageStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#fff"))
)

// InitMessage initializes the header viewport
func InitMessage(width int) viewport.Model {
	vp := viewport.New(meta.MsgWidth, meta.MsgHeight)
	vp.SetContent(renderMessage())
	vp.GotoBottom()
	return vp
}

// renderMessage returns the formatted header content
func renderMessage() string {
	msg := oldMessageStyle.Render("This is an old message")
	temp := msg
	for i := 1; i < 20; i++ {
		temp = lipgloss.JoinVertical(lipgloss.Top, temp, msg)
	}

	return lipgloss.JoinVertical(lipgloss.Top,
		temp,
		oldMessageStyle.Render("Even more logs..."),
		newMessageStyle.Render("Too many logs! Scroll up or down..."),
	)
}
