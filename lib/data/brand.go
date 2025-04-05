package data

// Prerequisite to promote to an advanced brand
type Prerequisite struct {
	RequiredClass  string // e.g., "Fighter"
	RequiredLevel  int    // e.g., Level 10
	RequiredGender int    // e.g., Female only
}

// Represents a unit class
type Brand struct {
	Id           int
	Name         string
	Info         string
	Letter       string
	BonusAc      int
	Movement     int
	Prerequisite Prerequisite
	WeaponProf   map[string]string // Proficiency according for each weapon type

	// Maximum attribute
	MaxStr, MaxDex, MaxCon int
	MaxInt, MaxWis, MaxCha int

	// Level up attribute bonuses
	BonusStr, BonusDex, BonusCon int
	BonusInt, BonusWis, BonusCha int
}
