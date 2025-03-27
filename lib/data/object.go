package data

// One-indexed objectId enum
const (
	Wall int = iota + 1
	Fork
	Ruins
	House
	Statue
	Chest
	Door
	Tree
	Sink
	Staircase
	Throne
)

type Object struct {
	Id          int
	Walkable    bool
	Destroyable bool
	IsBuilding  bool
}
