package util

import (
	"baolhq/branded/lib/data"
	"baolhq/branded/lib/meta"
	"fmt"
	"log"
)

func SeedData(c *data.Chapter) {
	c.AddObject(10, 5, &data.Terrain{Id: data.Lava, Letter: "{"}, nil)
	c.AddObject(20, 8, &data.Terrain{Id: data.Lava, Letter: "{"}, nil)

	c.AddObject(12, 13, nil, &data.Object{Id: data.Wall, Letter: "#"})
	c.AddObject(12, 14, nil, &data.Object{Id: data.Door, Letter: "+"})

	brands, err := LoadData("./res/db/brands.dat")
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < len(brands)-2; i++ {
		unit := data.Unit{}
		unit.Name = "Soldier"
		unit.Level = i
		unit.Exp = ClampInt(i*5, 0, 99)
		unit.Hp = ClampInt(i*5-10, 0, 99)
		unit.MaxHp = ClampInt(i*5-1, 0, 99)
		unit.Brand = brands[i]
		unit.Str = ClampInt(i*5-1, 0, 99)
		unit.Dex = ClampInt(i*5-2, 0, 99)
		unit.Con = ClampInt(i*5-3, 0, 99)
		unit.Int = ClampInt(i*5-4, 0, 99)
		unit.Wis = ClampInt(i*5-5, 0, 99)
		unit.Cha = ClampInt(i*5-6, 0, 99)

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
		Str: 99, Dex: 99, Con: 99,
		Int: 99, Wis: 99, Cha: 99,
		Brand: data.Brand{
			Letter: "@",
			Name:   "Lord",
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
		Faction: data.Party,
		PosX:    3,
		PosY:    2,
	}

	eLord := data.Unit{
		Name:  "Morgan",
		Level: 20,
		Exp:   99,
		Hp:    99, MaxHp: 99,
		Str: 99, Dex: 99, Con: 99,
		Int: 99, Wis: 99, Cha: 99,
		Brand: data.Brand{
			Letter: "@",
			Name:   "Lord",
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
		Faction: data.Enemy,
		PosX:    40,
		PosY:    19,
	}

	c.AddUnit(pLord)
	c.AddUnit(eLord)
}
