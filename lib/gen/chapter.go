package gen

import (
	"strings"

	"baolhq/branded/lib/data"
	"baolhq/branded/lib/meta"

	"github.com/charmbracelet/lipgloss"
)

type Chapter struct {
	Name  string
	Map   [meta.MapWidth][meta.MapHeight]TileStack
	Units []data.Unit
}

// A tile represent a character, terrain or objects
// Characters will have their faction, others don't
type Tile struct {
	Terrain data.Terrain
	Object  data.Object
	Faction data.Faction
}

// If multiple objects is on the same place, store their stack
type TileStack struct {
	Tiles []Tile
}

var styles = map[int]lipgloss.Style{
	data.Plain: lipgloss.NewStyle().Foreground(lipgloss.Color("#FFFFFF")),
	data.Lava:  lipgloss.NewStyle().Foreground(lipgloss.Color("#F95454")),
}

var cursorStyle = lipgloss.NewStyle().Background(lipgloss.Color("#A1A1A1")).Bold(true)

func setFactionStyle(style lipgloss.Style, faction int) lipgloss.Style {
	switch faction {
	case data.Other:
		return style.Foreground(lipgloss.Color("#FFFFFF"))
	case data.Party:
		return style.Foreground(lipgloss.Color("#0D92F4"))
	case data.Enemy:
		return style.Foreground(lipgloss.Color("#F95454"))
	case data.Ally:
		return style.Foreground(lipgloss.Color("#A0C878"))
	}
	return style
}

// Load default map with only plains terrain
func LoadMap() [meta.MapWidth][meta.MapHeight]TileStack {
	var m [meta.MapWidth][meta.MapHeight]TileStack
	for x := range m {
		for y := range m[x] {
			m[x][y] = TileStack{Tiles: []Tile{{Terrain: data.Terrain{Id: data.Plain, Symbol: "."}}}}
		}
	}
	return m
}

// Get the top-most object to render
func (c *Chapter) TopTile(x, y int) Tile {
	stack := c.Map[x][y].Tiles
	if len(stack) > 0 {
		return stack[len(stack)-1]
	}
	return Tile{Terrain: data.Terrain{Id: data.Plain}}
}

// Adding object and/or terrain into a stack
func (c *Chapter) AddObject(x, y int, terrain *data.Terrain, object *data.Object) {
	if x >= 0 && x < meta.MapWidth && y >= 0 && y < meta.MapHeight {
		// Ensure there's at least a default terrain tile
		if len(c.Map[x][y].Tiles) == 0 {
			c.Map[x][y].Tiles = append(c.Map[x][y].Tiles, Tile{Terrain: data.Terrain{Id: data.Plain}})
		}

		newTile := Tile{}
		if terrain != nil {
			newTile.Terrain = *terrain
		}
		if object != nil {
			newTile.Object = *object
		}

		c.Map[x][y].Tiles = append(c.Map[x][y].Tiles, newTile)
	}
}

// Adding an unit into a stack, store their location as well
func (c *Chapter) AddUnit(unit data.Unit) {
	if unit.PosX >= 0 && unit.PosX < meta.MapWidth &&
		unit.PosY >= 0 && unit.PosY < meta.MapHeight {
		if len(c.Map[unit.PosX][unit.PosY].Tiles) == 0 {
			c.AddObject(unit.PosX, unit.PosY, &data.Terrain{Id: data.Plain}, nil)
		}
		c.Units = append(c.Units, unit)
	}
}

func GetChapter(cursorX, cursorY int) string {
	var sb strings.Builder
	sb.Grow(meta.MapWidth * meta.MapHeight * 2)
	c := Chapter{Map: LoadMap()}

	c.AddObject(10, 5, &data.Terrain{Id: data.Lava, Symbol: "{"}, nil)
	c.AddObject(20, 8, &data.Terrain{Id: data.Lava, Symbol: "{"}, nil)

	c.AddObject(12, 13, nil, &data.Object{Id: data.Wall, Symbol: "#"})
	c.AddObject(12, 14, nil, &data.Object{Id: data.Door, Symbol: "+"})

	fighter := data.Unit{
		Role:    data.FighterRole(),
		Symbol:  "f",
		Level:   5,
		Faction: data.Party,
		PosX:    10,
		PosY:    5,
	}

	knight := data.Unit{
		Role:    data.KnightRole(),
		Symbol:  "f",
		Level:   12,
		Faction: data.Enemy,
		PosX:    20,
		PosY:    8,
	}

	c.AddUnit(fighter)
	c.AddUnit(knight)

	for y := 0; y < meta.MapHeight; y++ {
		for x := 0; x < meta.MapWidth; x++ {
			unitFound := false

			for _, unit := range c.Units {
				if unit.PosX == x && unit.PosY == y {
					unitFound = true
					style := setFactionStyle(lipgloss.NewStyle(), unit.Faction)

					if x == cursorX && y == cursorY {
						style = style.Background(cursorStyle.GetBackground())
					}

					sb.WriteString(style.Render(unit.Symbol))
					break
				}
			}

			if !unitFound {
				var char string
				style := styles[0] // Default to white
				tile := c.TopTile(x, y)

				if tile.Object.Id > 0 {
					char = tile.Object.Symbol
				} else {
					char = tile.Terrain.Symbol
				}

				if x == cursorX && y == cursorY {
					style = style.Background(cursorStyle.GetBackground())
				}

				sb.WriteString(style.Render(char))
			}
		}
		sb.WriteByte('\n')
	}

	return sb.String()
}
