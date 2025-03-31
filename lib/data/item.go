package data

import "fmt"

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
	Name       string
	Desc       string
	Uses       int  // Usage left
	Price      int  // Base price from shop
	Consumable bool // Can select `Use` and lost upon use
}

func (i *Item) UsageLeft() string {
	if i.Uses < 0 {
		return "--"
	}
	return fmt.Sprintf("%d", i.Uses)
}
