package data

type BrandType int

const (
	// Fighter line (f)
	Fighter BrandType = iota
	Champion
	Hero

	// Knight line (k)
	Knight
	GreatKnight
	General

	// Cavalier line (c)
	Cavalier
	Paladin
	DreadKnight

	// Nomad line (n)
	Nomad
	BowKnight
	Strider

	// Pegasus knight line (p)
	PegasusKnight
	FalconKnight
	SeraphKnight

	// Wyvern rider line(w)
	WyvernRider
	WyvernLord
	DragonRider

	// Mercenary line (m)
	Mercenary
	Duelist
	Vanguard

	// Spellblade line (s)
	Spellblade
	Hexblade
	Battlemage

	// Barbarian line (b)
	Barbarian
	Revenger
	Berserker

	// Archer line (a)
	Archer
	Marksman
	Sniper

	// Rogue line (r)
	Rogue
	Assassin
	Sharpshooter

	// Mage line (g)
	Mage
	Archmage
	Pyromancer

	// Dark mage line (d)
	DarkMage
	Warlock
	NightWarden

	// Lich line (l)
	Lich
	Archlich
	Necromancer

	// Priest line (i)
	Priest
	Cleric
	Bishop

	// Templar line (t)
	Templar
	Crusader
	Inquisitor

	// Lord line (@)
	Lord
	GreatLord
)

func (b BrandType) String() string {
	switch b {
	case Fighter:
		return "Fighter"
	case Champion:
		return "Champion"
	case Hero:
		return "Hero"
	case Knight:
		return "Knight"
	case GreatKnight:
		return "Great Knight"
	case General:
		return "General"
	case Cavalier:
		return "Cavalier"
	case Paladin:
		return "Paladin"
	case DreadKnight:
		return "Dread Knight"
	case Nomad:
		return "Nomad"
	case BowKnight:
		return "Bow Knight"
	case Strider:
		return "Strider"
	case PegasusKnight:
		return "Pegasus Knight"
	case FalconKnight:
		return "Falcon Knight"
	case SeraphKnight:
		return "Seraph Knight"
	case WyvernRider:
		return "Wyvern Rider"
	case WyvernLord:
		return "Wyvern Lord"
	case DragonRider:
		return "Dragon Rider"
	case Mercenary:
		return "Mercenary"
	case Duelist:
		return "Duelist"
	case Vanguard:
		return "Vanguard"
	case Spellblade:
		return "Spellblade"
	case Hexblade:
		return "Hexblade"
	case Battlemage:
		return "Battlemage"
	case Barbarian:
		return "Barbarian"
	case Revenger:
		return "Revenger"
	case Berserker:
		return "Berserker"
	case Archer:
		return "Archer"
	case Marksman:
		return "Marksman"
	case Sniper:
		return "Sniper"
	case Rogue:
		return "Rogue"
	case Assassin:
		return "Assassin"
	case Sharpshooter:
		return "Sharpshooter"
	case Mage:
		return "Mage"
	case Archmage:
		return "Archmage"
	case Pyromancer:
		return "Pyromancer"
	case DarkMage:
		return "Dark Mage"
	case Warlock:
		return "Warlock"
	case NightWarden:
		return "Night Warden"
	case Lich:
		return "Lich"
	case Archlich:
		return "Archlich"
	case Necromancer:
		return "Necromancer"
	case Priest:
		return "Priest"
	case Cleric:
		return "Cleric"
	case Bishop:
		return "Bishop"
	case Templar:
		return "Templar"
	case Crusader:
		return "Crusader"
	case Inquisitor:
		return "Inquisitor"
	case Lord:
		return "Lord"
	case GreatLord:
		return "Great Lord"
	default:
		return "Unknown"
	}
}

// Prerequisite to promote to an advanced brand
type Prerequisite struct {
	RequiredClass  string // e.g., "Fighter"
	RequiredLevel  int    // e.g., Level 10
	RequiredGender int    // e.g., Female only
}

// Brand represents a character archetype
type Brand struct {
	Type         BrandType
	Name         string
	Desc         string
	Prerequisite Prerequisite
	WeaponProf   []WeaponProf

	// Maximum attribute
	MSTR, MDEX, MCON int
	MINT, MWIS, MCHA int

	// Promotion attribute bonuses
	STR, DEX, CON int
	INT, WIS, CHA int
}
