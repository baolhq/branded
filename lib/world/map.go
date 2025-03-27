package world

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

type Tile struct {
	Terrain data.Terrain
	Object  data.Object
	Faction data.Faction
}

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

func LoadMap() [meta.MapWidth][meta.MapHeight]TileStack {
	var m [meta.MapWidth][meta.MapHeight]TileStack
	for x := range m {
		for y := range m[x] {
			m[x][y] = TileStack{Tiles: []Tile{{Terrain: data.Terrain{Id: data.Plain}}}}
		}
	}
	return m
}

func (c *Chapter) TopTile(x, y int) Tile {
	stack := c.Map[x][y].Tiles
	if len(stack) > 0 {
		return stack[len(stack)-1]
	}
	return Tile{Terrain: data.Terrain{Id: data.Plain}}
}

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

func (c *Chapter) AddUnit(unit data.Unit) {
	if unit.PositionX >= 0 && unit.PositionX < meta.MapWidth &&
		unit.PositionY >= 0 && unit.PositionY < meta.MapHeight {
		if len(c.Map[unit.PositionX][unit.PositionY].Tiles) == 0 {
			c.AddObject(unit.PositionX, unit.PositionY, &data.Terrain{Id: data.Plain}, nil)
		}
		c.Units = append(c.Units, unit)
	}
}

func GetChapter(cursorX, cursorY int) string {
	var sb strings.Builder
	sb.Grow(meta.MapWidth * meta.MapHeight * 2)
	c := Chapter{Map: LoadMap()}

	c.AddObject(10, 5, &data.Terrain{Id: data.Lava}, nil)
	c.AddObject(20, 8, &data.Terrain{Id: data.Lava}, nil)

	c.AddObject(12, 13, nil, &data.Object{Id: data.Wall})
	c.AddObject(12, 14, nil, &data.Object{Id: data.Door})

	fighter := data.Unit{
		Role:      data.FighterRole(),
		Level:     5,
		Faction:   data.Party,
		PositionX: 10,
		PositionY: 5,
	}

	knight := data.Unit{
		Role:      data.KnightRole(),
		Level:     12,
		Faction:   data.Enemy,
		PositionX: 20,
		PositionY: 8,
	}

	c.AddUnit(fighter)
	c.AddUnit(knight)

	for y := 0; y < meta.MapHeight; y++ {
		for x := 0; x < meta.MapWidth; x++ {
			var char string
			unitFound := false

			for _, unit := range c.Units {
				if unit.PositionX == x && unit.PositionY == y {
					unitFound = true
					roleId := unit.Role.Id

					// Render role character, the roleId order is important
					switch {
					case roleId <= data.General:
						char = "f"
					case roleId <= data.SeraphKnight:
						char = "c"
					case roleId <= data.Spellblade:
						char = "m"
					case roleId <= data.Juggernaut:
						char = "b"
					case roleId <= data.Artillerist:
						char = "a"
					case roleId <= data.Sharpshooter:
						char = "r"
					case roleId <= data.Cryomancer:
						char = "g"
					case roleId <= data.Necromancer:
						char = "d"
					case roleId <= data.Inquisitor:
						char = "e"
					case roleId <= data.GreatLord:
						char = "l"
					default:
						char = "?"
					}

					style := setFactionStyle(lipgloss.NewStyle(), unit.Faction)
					if x == cursorX && y == cursorY {
						style = style.Background(cursorStyle.GetBackground())
					}
					sb.WriteString(style.Render(char))
					break
				}
			}

			if !unitFound {
				style := styles[0] // Default to white
				tile := c.TopTile(x, y)

				if tile.Object.Id > 0 {
					switch tile.Object.Id {
					case data.Wall:
						char = "#"
					case data.Door:
						char = "+"
					}
				} else {
					switch tile.Terrain.Id {
					case data.Plain:
						char = "."
					case data.Lava:
						char = "{"
					}
					style = styles[tile.Terrain.Id]
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
