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
	MovementCost int
	Walkable     bool
	Diggable     bool
}
