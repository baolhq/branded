package data

type Weapon struct {
	Item         Item
	Rank         string // e.g: D, C, B, A, S
	Type         string // e.g: Sword, Axe, Bow, Staff, Tome, etc.
	MinRange     int
	MaxRange     int
	CanOneHanded bool // Wheather this weapon can be wielded one-hand
	CanTwoHanded bool // Wheather this weapon can be wielded two-hand
	ToHit        int  // e.g: +d4 chance to hit
	ToDam        int  // e.g: +d4 damage on hit
	BonusAc      int
}
