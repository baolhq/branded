package data

// Unit represent a single unit in game
type Unit struct {
	Name          string
	Level         int
	Hp            int
	MaxHp         int
	Exp           int
	Gender        string
	Brand         Brand
	Movement      int
	BaseAc        int
	Faction       string         // Initial faction of this unit toward the player's party
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
