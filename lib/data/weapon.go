package data

type WeaponRank int

const (
	RankD WeaponRank = iota
	RankC
	RankB
	RankA
	RankS
)

func (w WeaponRank) String() string {
	return [...]string{"D", "C", "B", "A", "S"}[w]
}

type WeaponType int

const (
	Sword WeaponType = iota
	Axe
	Lance
	Bow
	Crossbow
	Arcane
	Divine
	Occult
)

func (w WeaponType) String() string {
	return [...]string{"Sword", "Axe", "Lance", "Bow", "Crossbow", "Arcane", "Divine", "Occult"}[w]
}

type Weapon struct {
	Item         Item
	Rank         WeaponRank
	Type         WeaponType
	MinRange     int
	MaxRange     int
	CanOneHanded bool // Wheather this weapon can be wielded one-hand
	CanTwoHanded bool // Wheather this weapon can be wielded two-hand
	ToHit        int  // e.g: +d4 chance to hit
	ToDam        int  // e.g: +d4 damage on hit
	BonusAc      int
}
