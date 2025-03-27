package data

const (
	SwordProf int = iota
	AxeProf
	PolearmProf
	BowProf
	CrossbowProf
	AnimaProf
	LightProf
	DarkProf
)

const (
	Longsword int = iota
	Greatsword
	Shortsword
	Dagger
	RuneSword
	Rapier
	KillerSword

	Axe
	Greataxe
	Hammer
	Mace
	HandAxe
	KillerAxe

	Lance
	Pike
	Halberd
	Javelin
	KillerLance

	Longbow
	Shortbow
	Greatbow

	Crossbow
	HandCrossbow
	HeavyCrossbow
)

type Weapon struct {
	Id           int
	ProfId       int
	Name         string
	Desc         string
	Price        int
	MinRange     int
	MaxRange     int
	Weight       int  // More weight needs more constitution to avoid penalties
	IsTwoHanded  bool // Wheather this weapon MUST be wielded two-hand
	CanTwoHanded bool // Wheather this weapon CAN be wielded two-hand
	HitBonus     int  // e.g: +1d4 chance to hit
	DmgBonus     int  // e.g: +1d4 damage on hit
}
