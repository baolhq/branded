package core

import "github.com/charmbracelet/bubbles/viewport"

type Controls struct {
	CursorX, CursorY int
}

// MoveCursor handles movement logic
func (c *Controls) MoveCursor(direction string, width, height int) {
	switch direction {
	case "left", "h", "4":
		if c.CursorX > 0 {
			c.CursorX--
		}
	case "right", "l", "6":
		if c.CursorX < width-1 {
			c.CursorX++
		}
	case "up", "k", "8":
		if c.CursorY > 0 {
			c.CursorY--
		}
	case "down", "j", "2":
		if c.CursorY < height-1 {
			c.CursorY++
		}
	case "y", "7":
		if c.CursorX > 0 && c.CursorY > 0 {
			c.CursorX--
			c.CursorY--
		}
	case "u", "9":
		if c.CursorX < width-1 && c.CursorY > 0 {
			c.CursorX++
			c.CursorY--
		}
	case "b", "1":
		if c.CursorX > 0 && c.CursorY < height-1 {
			c.CursorX--
			c.CursorY++
		}
	case "n", "3":
		if c.CursorX < width-1 && c.CursorY < height-1 {
			c.CursorX++
			c.CursorY++
		}
	}
}

// ScrollMessages handles scrolling in the message log
func ScrollMessages(vp *viewport.Model, direction string) {
	switch direction {
	case "o":
		vp.LineUp(1)
	case "p":
		vp.LineDown(1)
	case "O":
		vp.GotoTop()
	case "P":
		vp.GotoBottom()
	}
}
