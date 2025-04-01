package core

import (
	"baolhq/branded/lib/meta"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

type Controls struct {
	CursorX, CursorY int
}

// Handles cursor movement logic
func (c *Controls) MoveCursor(msg tea.KeyMsg) {
	switch {
	case key.Matches(msg, Keys.Left):
		if c.CursorX > 0 {
			c.CursorX--
		}
	case key.Matches(msg, Keys.Right):
		if c.CursorX < meta.MapWidth-1 {
			c.CursorX++
		}
	case key.Matches(msg, Keys.Up):
		if c.CursorY > 0 {
			c.CursorY--
		}
	case key.Matches(msg, Keys.Down):
		if c.CursorY < meta.MapHeight-1 {
			c.CursorY++
		}
	case key.Matches(msg, Keys.UpLeft):
		if c.CursorX > 0 && c.CursorY > 0 {
			c.CursorX--
			c.CursorY--
		}
	case key.Matches(msg, Keys.UpRight):
		if c.CursorX < meta.MapWidth-1 && c.CursorY > 0 {
			c.CursorX++
			c.CursorY--
		}
	case key.Matches(msg, Keys.DownLeft):
		if c.CursorX > 0 && c.CursorY < meta.MapHeight-1 {
			c.CursorX--
			c.CursorY++
		}
	case key.Matches(msg, Keys.DownRight):
		if c.CursorX < meta.MapWidth-1 && c.CursorY < meta.MapHeight-1 {
			c.CursorX++
			c.CursorY++
		}
	}
}
