package data

type Gender int

const (
	Unknown Gender = iota
	Male
	Female
)

func (g Gender) String() string {
	switch g {
	case Unknown:
		return "Unknown"
	case Male:
		return "Male"
	default:
		return "Female"
	}
}

type Faction int

const (
	Other Faction = iota
	Party
	Enemy
	Ally
)

func (f Faction) String() string {
	switch f {
	case Other:
		return "Other"
	case Party:
		return "Party"
	case Enemy:
		return "Enemy"
	default:
		return "Ally"
	}
}

// Unit represent a single unit in game
type Unit struct {
	Name          string
	Level         int
	Hp            int
	MaxHp         int
	Exp           int
	Gender        Gender
	Brand         Brand
	Movement      int
	BaseAc        int
	Faction       Faction        // Initial faction of this unit toward the player's party
	RecuitBy      []Unit         // <Talk> to recuit this hostile unit
	Items         []*Item        // Consumables, held items
	ActiveEffects map[string]int // Short-term effects

	// Base attribute bonuses
	Str, Dex, Con int
	Int, Wis, Cha int

	// Level up bonuses, e.g: BonusStr 60 => 60% chance to increase STR on level up
	BonusStr, BonusDex, BonusCon int
	BonusInt, BonusWis, BonusCha int

	PosX, PosY int // Unit position on the map
}
