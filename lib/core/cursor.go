package core

import (
	"baolhq/branded/lib/meta"
)

type Controls struct {
	CursorX, CursorY int
}

// Handles cursor movement logic
func (c *Controls) MoveCursor(direction string) {
	switch direction {
	case "left", "h", "4":
		if c.CursorX > 0 {
			c.CursorX--
		}
	case "right", "l", "6":
		if c.CursorX < meta.MapWidth-1 {
			c.CursorX++
		}
	case "up", "k", "8":
		if c.CursorY > 0 {
			c.CursorY--
		}
	case "down", "j", "2":
		if c.CursorY < meta.MapHeight-1 {
			c.CursorY++
		}
	case "y", "7":
		if c.CursorX > 0 && c.CursorY > 0 {
			c.CursorX--
			c.CursorY--
		}
	case "u", "9":
		if c.CursorX < meta.MapWidth-1 && c.CursorY > 0 {
			c.CursorX++
			c.CursorY--
		}
	case "b", "1":
		if c.CursorX > 0 && c.CursorY < meta.MapHeight-1 {
			c.CursorX--
			c.CursorY++
		}
	case "n", "3":
		if c.CursorX < meta.MapWidth-1 && c.CursorY < meta.MapHeight-1 {
			c.CursorX++
			c.CursorY++
		}
	}
}
