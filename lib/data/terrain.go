package data

const (
	Plain int = iota
	Water
	Lava
	Snow
	Forest
	Mountain
	Desert
)

type Terrain struct {
	Id           int
	Name         string
	Info         string
	Letter       string
	MovementCost int
	Walkable     bool
	Diggable     bool
	BonusAc      int
}
