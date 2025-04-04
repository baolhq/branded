package util

import (
	"baolhq/branded/lib/data"
	"baolhq/branded/lib/meta"
	"database/sql"
	"encoding/gob"
	"log"
	"os"
)

func OpenDB(filename string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", filename)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func CloseDB(db *sql.DB) {
	if err := db.Close(); err != nil {
		log.Printf("error closing database: %v", err)
	}
}

// EncodeData generate .dat files from SQLite database
func EncodeData() {
	db, err := OpenDB(meta.DbPath)
	CheckErr(err)
	defer CloseDB(db)

	// Encoding brands
	file, err := os.Create(meta.BrandPath)
	CheckErr(err)
	defer file.Close()
	brands := GetBrands(db)
	encoder := gob.NewEncoder(file)
	err = encoder.Encode(brands)
	CheckErr(err)

	// Encoding items
	file, err = os.Create(meta.ItemPath)
	CheckErr(err)
	defer file.Close()
	items := GetItems(db)
	encoder = gob.NewEncoder(file)
	err = encoder.Encode(items)
	CheckErr(err)

	// Encoding weapons
	file, err = os.Create(meta.WeaponPath)
	CheckErr(err)
	defer file.Close()
	weapons := GetWeapons(db, items)
	encoder = gob.NewEncoder(file)
	err = encoder.Encode(weapons)
	CheckErr(err)

	// Encoding terrains
	file, err = os.Create(meta.TerrainPath)
	CheckErr(err)
	defer file.Close()
	terrains := GetTerrains(db)
	encoder = gob.NewEncoder(file)
	err = encoder.Encode(terrains)
	CheckErr(err)

	// Encoding objects
	file, err = os.Create(meta.ObjectPath)
	CheckErr(err)
	defer file.Close()
	objects := GetObjects(db)
	encoder = gob.NewEncoder(file)
	err = encoder.Encode(objects)
	CheckErr(err)
}

func DecodeBrands() []data.Brand {
	file, err := os.Open(meta.BrandPath)
	CheckErr(err)
	defer file.Close()

	var brands []data.Brand
	decoder := gob.NewDecoder(file)
	err = decoder.Decode(&brands)
	CheckErr(err)

	return brands
}

func DecodeItems() []data.Item {
	file, err := os.Open(meta.ItemPath)
	CheckErr(err)
	defer file.Close()

	var items []data.Item
	decoder := gob.NewDecoder(file)
	err = decoder.Decode(&items)
	CheckErr(err)

	return items
}

func DecodeWeapons() []data.Weapon {
	file, err := os.Open(meta.WeaponPath)
	CheckErr(err)
	defer file.Close()

	var weapons []data.Weapon
	decoder := gob.NewDecoder(file)
	err = decoder.Decode(&weapons)
	CheckErr(err)

	return weapons
}

func DecodeTerrains() []data.Terrain {
	file, err := os.Open(meta.TerrainPath)
	CheckErr(err)
	defer file.Close()

	var terrains []data.Terrain
	decoder := gob.NewDecoder(file)
	err = decoder.Decode(&terrains)
	CheckErr(err)

	return terrains
}

func DecodeObjects() []data.Object {
	file, err := os.Open(meta.ObjectPath)
	CheckErr(err)
	defer file.Close()

	var objects []data.Object
	decoder := gob.NewDecoder(file)
	err = decoder.Decode(&objects)
	CheckErr(err)

	return objects
}

func GetBrands(db *sql.DB) []data.Brand {
	rows, err := db.Query(`SELECT name, info, letter, bonus_ac,
		max_str, max_dex, max_con, max_int, max_wis, max_cha,
		bonus_str, bonus_dex, bonus_con, bonus_int, bonus_wis, bonus_cha FROM brands`)
	CheckErr(err)
	defer rows.Close()

	var brands []data.Brand
	for rows.Next() {
		var b data.Brand
		err := rows.Scan(&b.Name, &b.Info, &b.Letter, &b.BonusAc,
			&b.MaxStr, &b.MaxDex, &b.MaxCon, &b.MaxInt, &b.MaxWis, &b.MaxCha,
			&b.BonusStr, &b.BonusDex, &b.BonusCon, &b.BonusInt, &b.BonusWis, &b.BonusCha)
		CheckErr(err)
		brands = append(brands, b)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	return brands
}

func GetItems(db *sql.DB) []data.Item {
	rows, err := db.Query(`SELECT id, name, info, uses, price, is_consumable, 
		effect_type, effect_value, effect_duration FROM items`)
	CheckErr(err)
	defer rows.Close()

	var items []data.Item
	for rows.Next() {
		var i data.Item
		err := rows.Scan(&i.Id, &i.Name, &i.Info, &i.Uses, &i.Price, &i.IsConsumable,
			&i.EffectType, &i.EffectValue, &i.EffectDuration)
		CheckErr(err)
		items = append(items, i)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	return items
}

func GetWeapons(db *sql.DB, items []data.Item) []data.Weapon {
	rows, err := db.Query(`SELECT item_id, weapon_type, weapon_rank, min_range, max_range,
		can_one_handed, can_two_handed, to_hit, to_dam, bonus_ac FROM weapons`)
	CheckErr(err)
	defer rows.Close()

	var weapons []data.Weapon
	for rows.Next() {
		var w data.Weapon
		var itemId int
		err := rows.Scan(&itemId, &w.Type, &w.Rank, &w.MinRange, &w.MaxRange,
			&w.CanOneHanded, &w.CanTwoHanded, &w.ToHit, &w.ToDam, &w.BonusAc)
		CheckErr(err)

		for _, item := range items {
			if item.Id == itemId {
				w.Item = item
				break
			}
		}

		weapons = append(weapons, w)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	return weapons
}

func GetTerrains(db *sql.DB) []data.Terrain {
	rows, err := db.Query(`SELECT name, info, letter, movement_cost, walkable, diggable, bonus_ac FROM terrains`)
	CheckErr(err)
	defer rows.Close()

	var terrains []data.Terrain
	for rows.Next() {
		var t data.Terrain
		err := rows.Scan(&t.Name, &t.Info, &t.Letter, &t.MovementCost, &t.Walkable,
			&t.Diggable, &t.BonusAc)
		CheckErr(err)
		terrains = append(terrains, t)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	return terrains
}

func GetObjects(db *sql.DB) []data.Object {
	rows, err := db.Query(`SELECT name, info, letter, walkable, destroyable,
		is_building, is_locked, bonus_ac FROM objects`)
	CheckErr(err)
	defer rows.Close()

	var objects []data.Object
	for rows.Next() {
		var o data.Object
		err := rows.Scan(&o.Name, &o.Info, &o.Letter, &o.Walkable, &o.Destroyable,
			&o.IsBuilding, &o.IsLocked, &o.BonusAc)
		CheckErr(err)
		objects = append(objects, o)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	return objects
}
