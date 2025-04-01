package data

type WeaponProf int

const (
	RankD WeaponProf = iota
	RankC
	RankB
	RankA
	RankS
)

func (w WeaponProf) String() string {
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
	Prof         WeaponProf
	Type         WeaponType
	MinRange     int
	MaxRange     int
	Weight       int  // More weight needs more constitution to avoid penalties
	CanOneHanded bool // Wheather this weapon can be wielded one-hand
	CanTwoHanded bool // Wheather this weapon can be wielded two-hand
	ToHit        int  // e.g: +d4 chance to hit
	ToDam        int  // e.g: +d4 damage on hit
}
