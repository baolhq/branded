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
	Name         string
	Symbol       string
	Level        int
	Exp          int
	Gender       Gender
	Role         Role
	Movement     int
	Faction      int    // Initial faction of this unit toward the player's party
	Rso          []int  // Unit Resonance graph
	RecuitBy     []Unit // <Talk> to recuit this hostile unit
	Support      []Unit // <Support> to increase relationship
	SupportLevel []int  // Correspond to the <Support> array

	// Base attribute bonuses
	STR, DEX, CON int
	INT, WIS, CHA int

	// Level up bonuses, e.g: USTR 60 => 60% chance to increase STR on level up
	USTR, UDEX, UCON int
	UINT, UWIS, UCHA int

	PosX, PosY int // Unit position on the map
}
