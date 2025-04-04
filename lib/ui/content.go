package ui

import (
	"baolhq/branded/lib/data"
	"baolhq/branded/lib/meta"

	"github.com/charmbracelet/bubbles/viewport"
)

// InitContent initializes the main game content viewport
func InitContent(content string) viewport.Model {
	vp := viewport.New(meta.MapWidth, meta.MapHeight)
	vp.SetContent(content)
	return vp
}

// UpdateContent updates the viewport with new game content
func UpdateContent(vp *viewport.Model, cX, cY int, c data.Chapter) {
	chapter := RenderChapter(cX, cY, c)
	vp.SetContent(chapter)
}

// Preview unit movements
func PreviewMovement(vp *viewport.Model, cX, cY int, c data.Chapter, u data.Unit) {

}
