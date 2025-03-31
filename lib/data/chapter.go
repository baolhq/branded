package data

import "baolhq/branded/lib/meta"

// Chapter objectives
const (
	Conquest  int = iota // Defeat all enemies
	Seize                // Take control of a target area
	Intercept            // Prevent enemy from reaching a target area
	Raid                 // Destroy enemy building(s)
	Survive              // Hold your line
	Rescue               // Free captured or defend ally unit(s)
	Dethrone             // Take down enemy commander(s)
)

type Chapter struct {
	Name      string
	Objective int
	Map       [][]TileStack
	Units     []Unit
}

type Tile struct {
	Terrain Terrain
	Object  Object
	Faction Faction
}

type TileStack struct {
	Tiles []Tile
}

func LoadMap() [][]TileStack {
	m := make([][]TileStack, meta.MapWidth)
	for x := range m {
		m[x] = make([]TileStack, meta.MapHeight)
		for y := range m[x] {
			m[x][y] = TileStack{Tiles: []Tile{{Terrain: Terrain{Id: Plain, Symbol: "."}}}}
		}
	}
	return m
}

func (c *Chapter) AddObject(x, y int, terrain *Terrain, object *Object) {
	if x >= 0 && x < meta.MapWidth && y >= 0 && y < meta.MapHeight {
		if len(c.Map[x][y].Tiles) == 0 {
			c.Map[x][y].Tiles = append(c.Map[x][y].Tiles, Tile{Terrain: Terrain{Id: Plain}})
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

func (c *Chapter) AddUnit(unit Unit) {
	if unit.PosX >= 0 && unit.PosX < meta.MapWidth &&
		unit.PosY >= 0 && unit.PosY < meta.MapHeight {
		if len(c.Map[unit.PosX][unit.PosY].Tiles) == 0 {
			c.AddObject(unit.PosX, unit.PosY, &Terrain{Id: Plain}, nil)
		}
		c.Units = append(c.Units, unit)
	}
}

func (c *Chapter) GetUnitAt(x, y int) *Unit {
	for i := range c.Units {
		if c.Units[i].PosX == x && c.Units[i].PosY == y {
			return &c.Units[i]
		}
	}
	return nil
}
