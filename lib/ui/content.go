package ui

import (
	"baolhq/branded/lib/world"

	"github.com/charmbracelet/bubbles/viewport"
)

// InitContent initializes the main game content viewport
func InitContent(width, height int, content string) viewport.Model {
	vp := viewport.New(width-4, height-4)
	vp.SetContent(content)
	return vp
}

// UpdateContent updates the viewport with new game content
func UpdateContent(vp *viewport.Model, x, y int) {
	vp.SetContent(world.GetChapter(x, y))
}
