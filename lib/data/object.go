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
	Name        string
	Info        string
	Letter      string
	Walkable    bool
	Destroyable bool
	IsBuilding  bool // Can be `Visit`
	IsLocked    bool // Can be `Unlock`
	BonusAc     int
}
