package data

const (
	Herb int = iota + 1
	RareHerb
	MysticHerb
	Torch

	// Stat boosters
	DragonClaw
	HarpieFeather
	TitanEye
	MagicMirror
	WisemanJournal
	SirenRose
	GoddessIcon

	// Held items
	PhoenixDown
	BloodEmblem
	DarkOnesTalisman
	AegisPendant
	EarthmotherTotem
	RabbitFluff
	OracleCrystal
	DemonMask

	// Valuables
	CursedIdol
	FiveLeafClover
	SilkCloth
	SilverIngot
	MithrilOre
)

type Item struct {
	Id         int
	Price      int // Base price from shop
	Desc       string
	Consumable bool // Can select `Use` and lost upon use
}
