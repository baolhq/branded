package data

// Gender enum
type Gender int

const (
	Male Gender = iota
	Female
	Unknown
)

// Faction enum
type Faction int

const (
	Other int = iota
	Party
	Enemy
	Ally
)

// Unit represent a single unit in game
type Unit struct {
	Name     string
	Level    int
	Exp      int
	Gender   Gender
	Role     Role
	Faction  int    // Default faction of this unit
	RecuitBy []Unit // <Talk> to recuit this hostile unit
	Support  []Unit // <Support> to increase relationship

	// Base attribute bonuses
	STR, DEX, CON int
	INT, WIS, CHA int

	// Level up bonuses, e.g: USTR 60 => 60% chance to increase STR on level up
	USTR, UDEX, UCON int
	UINT, UWIS, UCHA int

	// Unit position on the map
	PositionX, PositionY int
}
