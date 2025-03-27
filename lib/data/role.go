package data

// WeaponMastery represents weapon proficiency
type WeaponMastery int

// Character role IDs, the order is important
const (
	// Fighter line (f)
	Fighter int = iota
	Knight
	Champion
	Hero
	General

	// Cavalry line (c)
	Cavalry
	Nomad
	WyvernRider
	PegasusKnight
	Paladin
	DragonRider
	SeraphKnight

	// Mercenary line (m)
	Mercenary
	Duelist
	Vanguard
	Battlemage
	SwordSaint
	Warlord
	Spellblade

	// Barbarian line (b)
	Barbarian
	Revenger
	Berserker
	Juggernaut

	// Archer line (a)
	Archer
	Marksman
	Ballistician
	Sniper
	Beastmaster
	Artillerist

	// Rogue line (r)
	Rogue
	Stalker
	Trickshot
	Assassin
	Trapper
	Sharpshooter

	// Mage line (g)
	Mage
	Elementalist
	Archmage
	Pyromancer
	Cryomancer

	// Dark mage line (d)
	DarkMage
	Warlock
	Lich
	NightWarden
	SpiritCaller
	Archlich
	Necromancer

	// Enchanter line (e)
	Enchanter
	Priest
	Cleric
	Templar
	Bishop
	Crusader
	Inquisitor

	// Lord line (l)
	Lord
	KnightLord
	GreatLord
)

// Prerequisite to promote to a role
type Prerequisite struct {
	RequiredClass  string // e.g., "Fighter"
	RequiredLevel  int    // e.g., Level 10
	RequiredGender Gender // e.g., Female only
}

// Role represents a character archetype
type Role struct {
	Id           int
	Desc         string
	Prerequisite Prerequisite
	WeaponProf   []int

	// Maximum attribute
	MSTR, MDEX, MCON int
	MINT, MWIS, MCHA int

	// Promotion attribute bonuses
	STR, DEX, CON int
	INT, WIS, CHA int
}

// FighterArchetype returns a base Fighter archetype
func FighterRole() Role {
	return Role{
		Id:   Fighter,
		Desc: "Basic melee combatant",
		Prerequisite: Prerequisite{
			RequiredClass: "None", // No prior class needed
		},
		WeaponProf: []int{SwordProf},
	}
}

// KnightArchetype returns a promoted Knight archetype
func KnightRole() Role {
	return Role{
		Id:   Knight,
		Desc: "Heavy armored unit",
		Prerequisite: Prerequisite{
			RequiredClass: "Fighter",
			RequiredLevel: 10,
		},
		STR: 2, DEX: 0, CON: 4, INT: 2, WIS: 0, CHA: 3,
		WeaponProf: []int{SwordProf, AxeProf},
	}
}

func (w WeaponMastery) String() string {
	return [...]string{"Sword", "Axe", "Polearm", "Bow", "Crossbow", "Anima", "Light", "Dark"}[w]
}

func (g Gender) String() string {
	return [...]string{"Any", "Male", "Female"}[g]
}
