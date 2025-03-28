package gen

import (
	"strings"

	"baolhq/branded/lib/data"
	"baolhq/branded/lib/meta"
	"baolhq/branded/lib/util"

	"github.com/charmbracelet/lipgloss"
)

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

func SeedData(c *data.Chapter) {
	c.AddObject(10, 5, &data.Terrain{Id: data.Lava, Symbol: "{"}, nil)
	c.AddObject(20, 8, &data.Terrain{Id: data.Lava, Symbol: "{"}, nil)

	c.AddObject(12, 13, nil, &data.Object{Id: data.Wall, Symbol: "#"})
	c.AddObject(12, 14, nil, &data.Object{Id: data.Door, Symbol: "+"})

	for i := 0; i < 20; i++ {
		unit := data.Unit{}
		unit.Role = data.Role{
			Id: util.RandInt(0, data.Lord),
		}

		for j := 0; j < 10; j++ {
			unit.Rso = append(unit.Rso, util.RandInt(-5, 5))
		}

		switch unit.Role.Id {
		case data.Fighter, data.Knight, data.Champion, data.Hero, data.General:
			unit.Symbol = "f"
		case data.Cavalry, data.Nomad, data.WyvernRider, data.PegasusKnight, data.Paladin, data.DragonRider, data.SeraphKnight:
			unit.Symbol = "c"
		case data.Mercenary, data.Duelist, data.Vanguard, data.Battlemage, data.SwordSaint, data.Warlord, data.Spellblade:
			unit.Symbol = "m"
		case data.Barbarian, data.Revenger, data.Berserker, data.Juggernaut:
			unit.Symbol = "b"
		case data.Archer, data.Marksman, data.Ballistician, data.Sniper, data.Beastmaster, data.Artillerist:
			unit.Symbol = "a"
		case data.Rogue, data.Stalker, data.Trickshot, data.Assassin, data.Trapper, data.Sharpshooter:
			unit.Symbol = "r"
		case data.Mage, data.Elementalist, data.Archmage, data.Pyromancer, data.Cryomancer:
			unit.Symbol = "g"
		case data.DarkMage, data.Warlock, data.Lich, data.NightWarden, data.SpiritCaller, data.Archlich, data.Necromancer:
			unit.Symbol = "d"
		case data.Enchanter, data.Priest, data.Cleric, data.Templar, data.Bishop, data.Crusader, data.Inquisitor:
			unit.Symbol = "e"
		default:
			unit.Symbol = "?"
		}

		if i < 10 {
			unit.Faction = data.Party
		} else {
			unit.Faction = data.Enemy
		}

		unit.PosX = util.RandInt(1, meta.MapWidth-1)
		unit.PosY = util.RandInt(1, meta.MapHeight-1)

		c.AddUnit(unit)
	}

	pLord := data.Unit{
		Role: data.Role{
			Id: data.Lord,
		},
		Rso:     []int{4, 2, 0, 1, -2, -3, -3, 0, 4, 5},
		Symbol:  "l",
		Faction: data.Party,
		PosX:    3,
		PosY:    3,
	}

	eLord := data.Unit{
		Role: data.Role{
			Id: data.Lord,
		},
		Rso:     []int{-4, -2, 0, 1, -3, 4, 2, 4, -2, 0},
		Symbol:  "l",
		Faction: data.Enemy,
		PosX:    13,
		PosY:    10,
	}

	c.AddUnit(pLord)
	c.AddUnit(eLord)
}

func GetChapter(cX, cY int) data.Chapter {
	c := data.Chapter{Map: data.LoadMap()}
	return c
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
				char = unit.Symbol
			} else {
				style = styles[0] // Default to white
				tile := c.Map[x][y].Tiles[len(c.Map[x][y].Tiles)-1]

				if tile.Object.Id > 0 {
					char = tile.Object.Symbol
				} else {
					char = tile.Terrain.Symbol
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
