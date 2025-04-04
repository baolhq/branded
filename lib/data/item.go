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
	BullionS
	BullionM
	BullionL
)

type Item struct {
	Id             int
	Name           string
	Info           string
	Uses           int  // Usage left
	Price          int  // Base price from shop
	IsConsumable   bool // Can select `Use` and lost upon use
	EffectType     string
	EffectValue    int
	EffectDuration int // 0 = Permanent, 1+ = Short-term
}

func (i *Item) UsageLeft() string {
	if i.Uses < 0 {
		return "--"
	}
	return fmt.Sprintf("%d", i.Uses)
}
