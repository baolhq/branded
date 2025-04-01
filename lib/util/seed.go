package util

import (
	"baolhq/branded/lib/data"
	"baolhq/branded/lib/meta"
	"fmt"
)

func SeedData(c *data.Chapter) {
	c.AddObject(10, 5, &data.Terrain{Id: data.Lava, Symbol: "{"}, nil)
	c.AddObject(20, 8, &data.Terrain{Id: data.Lava, Symbol: "{"}, nil)

	c.AddObject(12, 13, nil, &data.Object{Id: data.Wall, Symbol: "#"})
	c.AddObject(12, 14, nil, &data.Object{Id: data.Door, Symbol: "+"})

	for i := 0; i < 20; i++ {
		unit := data.Unit{}
		unit.Name = "Soldier"
		unit.Level = i
		unit.Exp = i*5 - 1
		unit.Hp = i*5 - 10
		unit.MaxHp = i*5 - 1
		unit.Brand = data.Brand{
			Type: data.BrandType(RandInt(0, 46)),
		}
		unit.Brand.Name = unit.Brand.Type.String()
		unit.STR = i*5 - 1
		unit.DEX = i*5 - 2
		unit.CON = i*5 - 3
		unit.INT = i*5 - 4
		unit.WIS = i*5 - 5
		unit.CHA = i*5 - 6

		unit.Items = []*data.Item{
			{
				Name: "Bronze Sword (E)",
				Uses: 99,
			},
			{
				Name: "Iron Lance",
				Uses: 10,
			},
			{
				Name: fmt.Sprintf("Vulnararies %d", i),
				Uses: -1,
			},
		}

		switch unit.Brand.Type {
		case data.Fighter:
			unit.Symbol = "f"
		case data.Champion, data.Hero:
			unit.Symbol = "F"
		case data.Knight:
			unit.Symbol = "k"
		case data.GreatKnight, data.General:
			unit.Symbol = "K"
		case data.Cavalier:
			unit.Symbol = "c"
		case data.Paladin, data.DreadKnight:
			unit.Symbol = "C"
		case data.Nomad:
			unit.Symbol = "n"
		case data.BowKnight, data.Strider:
			unit.Symbol = "N"
		case data.PegasusKnight:
			unit.Symbol = "p"
		case data.FalconKnight, data.SeraphKnight:
			unit.Symbol = "P"
		case data.WyvernRider:
			unit.Symbol = "w"
		case data.WyvernLord, data.DragonRider:
			unit.Symbol = "W"
		case data.Mercenary:
			unit.Symbol = "m"
		case data.Duelist, data.Vanguard:
			unit.Symbol = "M"
		case data.Spellblade:
			unit.Symbol = "s"
		case data.Hexblade, data.Battlemage:
			unit.Symbol = "S"
		case data.Barbarian:
			unit.Symbol = "b"
		case data.Revenger, data.Berserker:
			unit.Symbol = "B"
		case data.Archer:
			unit.Symbol = "a"
		case data.Marksman, data.Sniper:
			unit.Symbol = "A"
		case data.Rogue:
			unit.Symbol = "r"
		case data.Assassin, data.Sharpshooter:
			unit.Symbol = "R"
		case data.Mage:
			unit.Symbol = "g"
		case data.Archmage, data.Pyromancer:
			unit.Symbol = "G"
		case data.DarkMage:
			unit.Symbol = "d"
		case data.Warlock, data.NightWarden:
			unit.Symbol = "D"
		case data.Lich:
			unit.Symbol = "l"
		case data.Archlich, data.Necromancer:
			unit.Symbol = "L"
		case data.Priest:
			unit.Symbol = "i"
		case data.Cleric, data.Bishop:
			unit.Symbol = "I"
		case data.Templar:
			unit.Symbol = "t"
		case data.Crusader, data.Inquisitor:
			unit.Symbol = "T"
		default:
			unit.Symbol = "?"
		}

		if RandInt(0, 2) < 1 {
			unit.Faction = data.Party
		} else {
			unit.Faction = data.Enemy
		}

		unit.PosX = RandInt(0, meta.MapWidth)
		unit.PosY = RandInt(0, meta.MapHeight)

		c.AddUnit(unit)
	}

	pLord := data.Unit{
		Name:  "Kros",
		Level: 20,
		Exp:   99,
		Hp:    99, MaxHp: 99,
		STR: 99, DEX: 99, CON: 99,
		INT: 99, WIS: 99, CHA: 99,
		Brand: data.Brand{
			Type: data.Lord,
			Name: "Lord",
		},
		Items: []*data.Item{
			{
				Name: "Falchion (E)",
				Uses: 99,
			},
			{
				Name: "Iron Lance",
				Uses: 10,
			},
		},
		Symbol:  "@",
		Faction: data.Party,
		PosX:    21,
		PosY:    11,
	}

	eLord := data.Unit{
		Name:  "Morgan",
		Level: 20,
		Exp:   99,
		Hp:    99, MaxHp: 99,
		STR: 99, DEX: 99, CON: 99,
		INT: 99, WIS: 99, CHA: 99,
		Brand: data.Brand{
			Type: data.Lord,
			Name: "Lord",
		},
		Items: []*data.Item{
			{
				Name: "Steel Lance (E)",
				Uses: 99,
			},
			{
				Name: "Iron Lance",
				Uses: 10,
			},
		},
		Symbol:  "@",
		Faction: data.Enemy,
		PosX:    43,
		PosY:    21,
	}

	c.AddUnit(pLord)
	c.AddUnit(eLord)
}
