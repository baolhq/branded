package data

import "fmt"

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
