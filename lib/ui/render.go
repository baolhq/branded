package ui

import (
	"strings"

	"baolhq/branded/lib/data"
	"baolhq/branded/lib/meta"

	"github.com/charmbracelet/lipgloss"
)

var styles = map[int]lipgloss.Style{
	data.Plain: lipgloss.NewStyle().Foreground(meta.White),
	data.Lava:  lipgloss.NewStyle().Foreground(meta.Red),
}

var cursorStyle = lipgloss.NewStyle().Background(meta.Gray).Bold(true)

func setFactionStyle(style lipgloss.Style, faction data.Faction) lipgloss.Style {
	switch faction {
	case data.Other:
		return style.Foreground(meta.White)
	case data.Party:
		return style.Foreground(meta.Blue)
	case data.Enemy:
		return style.Foreground(meta.Red)
	case data.Ally:
		return style.Foreground(meta.Green)
	}
	return style
}

func RenderChapter(cX, cY int, c data.Chapter) string {
	var sb strings.Builder
	sb.Grow(meta.MapWidth * meta.MapHeight * 2)

	for y := 0; y < meta.MapHeight; y++ {
		for x := 0; x < meta.MapWidth; x++ {
			var style lipgloss.Style
			var char string

			if unit := c.GetUnitAt(x, y); unit != nil {
				style = setFactionStyle(lipgloss.NewStyle(), unit.Faction)
				char = unit.Brand.Letter
			} else {
				style = styles[0] // Default to white
				tile := c.Map[x][y].Tiles[len(c.Map[x][y].Tiles)-1]

				if tile.Object.Id > 0 {
					char = tile.Object.Letter
				} else {
					char = tile.Terrain.Letter
				}
			}

			if x == cX && y == cY {
				style = style.Background(cursorStyle.GetBackground())
			}
			sb.WriteString(style.Render(char))
		}
		sb.WriteByte('\n')
	}

	return sb.String()
}
