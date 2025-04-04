package util

import "baolhq/branded/lib/data"

// Get party lord position, useful for setup initial cursor position
// and/or new position after every turn
func GetLordPosition(c *data.Chapter) (int, int) {
	lord := GetPartyLord(c)
	return lord.PosX, lord.PosY
}

// Given the chapter, return the party lord
func GetPartyLord(c *data.Chapter) *data.Unit {
	for x := range c.Map {
		for y := range c.Map[x] {
			unit := c.GetUnitAt(x, y)
			if unit != nil &&
				unit.Brand.Name == "Lord" &&
				unit.Faction == "Party" {
				return unit
			}
		}
	}

	return nil
}
