package util

import (
	"baolhq/branded/lib/data"
	"baolhq/branded/lib/meta"
	"database/sql"
	"encoding/gob"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
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

// EncodeData generate .dat files from SQLite datasource
func EncodeData() {
	db, err := OpenDB(meta.DbPath)
	CheckErr(err)
	defer CloseDB(db)

	type encodeTarget struct {
		path string
		data any
	}

	items := GetItems(db)
	brands := GetBrands(db)
	encodeList := []encodeTarget{
		{meta.BrandPath, brands},
		{meta.ItemPath, items},
		{meta.WeaponPath, GetWeapons(db, items)},
		{meta.TerrainPath, GetTerrains(db)},
		{meta.ObjectPath, GetObjects(db)},
		{meta.UnitPath, GetUnits(db, brands)},
	}

	for _, e := range encodeList {
		func(path string, data any) {
			file, err := os.Create(path)
			CheckErr(err)
			defer file.Close()

			encoder := gob.NewEncoder(file)
			err = encoder.Encode(data)
			CheckErr(err)
		}(e.path, e.data)
	}
}

// decodeFile decodes a .dat file into the provided output variable
func decodeFile[T any](path string, out *T) {
	file, err := os.Open(path)
	CheckErr(err)
	defer file.Close()

	decoder := gob.NewDecoder(file)
	err = decoder.Decode(out)
	CheckErr(err)
}

func DecodeBrands() (out []data.Brand)     { decodeFile(meta.BrandPath, &out); return }
func DecodeItems() (out []data.Item)       { decodeFile(meta.ItemPath, &out); return }
func DecodeWeapons() (out []data.Weapon)   { decodeFile(meta.WeaponPath, &out); return }
func DecodeTerrains() (out []data.Terrain) { decodeFile(meta.TerrainPath, &out); return }
func DecodeObjects() (out []data.Object)   { decodeFile(meta.ObjectPath, &out); return }
func DecodeUnits() (out []data.Unit)       { decodeFile(meta.UnitPath, &out); return }

// queryRows executes a query and scans the results into a slice of type T
func queryRows[T any](db *sql.DB, query string, scan func(*sql.Rows) (T, error)) []T {
	rows, err := db.Query(query)
	CheckErr(err)
	defer rows.Close()

	var result []T
	for rows.Next() {
		item, err := scan(rows)
		CheckErr(err)
		result = append(result, item)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	return result
}

func GetBrands(db *sql.DB) []data.Brand {
	return queryRows(db, `SELECT id, name, info, letter, bonus_ac, movement,
		max_str, max_dex, max_con, max_int, max_wis, max_cha,
		bonus_str, bonus_dex, bonus_con, bonus_int, bonus_wis, bonus_cha FROM brands`,
		func(rows *sql.Rows) (data.Brand, error) {
			var b data.Brand
			err := rows.Scan(&b.Id, &b.Name, &b.Info, &b.Letter, &b.BonusAc, &b.Movement,
				&b.MaxStr, &b.MaxDex, &b.MaxCon, &b.MaxInt, &b.MaxWis, &b.MaxCha,
				&b.BonusStr, &b.BonusDex, &b.BonusCon, &b.BonusInt, &b.BonusWis, &b.BonusCha)
			return b, err
		})
}

func GetItems(db *sql.DB) []data.Item {
	return queryRows(db, `SELECT id, name, info, uses, price, is_consumable, 
		effect_type, effect_value, effect_duration FROM items`,
		func(rows *sql.Rows) (data.Item, error) {
			var i data.Item
			err := rows.Scan(&i.Id, &i.Name, &i.Info, &i.Uses, &i.Price, &i.IsConsumable,
				&i.EffectType, &i.EffectValue, &i.EffectDuration)
			return i, err
		})
}

func GetTerrains(db *sql.DB) []data.Terrain {
	return queryRows(db, `SELECT name, info, letter, movement_cost, walkable, diggable, bonus_ac FROM terrains`,
		func(rows *sql.Rows) (data.Terrain, error) {
			var t data.Terrain
			err := rows.Scan(&t.Name, &t.Info, &t.Letter, &t.MovementCost, &t.Walkable, &t.Diggable, &t.BonusAc)
			return t, err
		})
}

func GetObjects(db *sql.DB) []data.Object {
	return queryRows(db, `SELECT name, info, letter, walkable, destroyable, is_building, is_locked, bonus_ac FROM objects`,
		func(rows *sql.Rows) (data.Object, error) {
			var o data.Object
			err := rows.Scan(&o.Name, &o.Info, &o.Letter, &o.Walkable, &o.Destroyable, &o.IsBuilding, &o.IsLocked, &o.BonusAc)
			return o, err
		})
}

func GetWeapons(db *sql.DB, items []data.Item) []data.Weapon {
	itemMap := make(map[int]data.Item)
	for _, i := range items {
		itemMap[i.Id] = i
	}

	return queryRows(db, `SELECT item_id, weapon_type, weapon_rank, min_range, max_range,
		can_one_handed, can_two_handed, to_hit, to_dam, bonus_ac FROM weapons`,
		func(rows *sql.Rows) (data.Weapon, error) {
			var w data.Weapon
			var itemId int
			err := rows.Scan(&itemId, &w.Type, &w.Rank, &w.MinRange, &w.MaxRange,
				&w.CanOneHanded, &w.CanTwoHanded, &w.ToHit, &w.ToDam, &w.BonusAc)
			if err != nil {
				return w, err
			}
			w.Item = itemMap[itemId]
			return w, nil
		})
}

func GetUnits(db *sql.DB, brands []data.Brand) []data.Unit {
	brandMap := make(map[int]data.Brand)
	for _, b := range brands {
		brandMap[b.Id] = b
	}

	return queryRows(db, `SELECT id, brand_id, name, level, max_hp, gender, base_ac,
		faction, str, dex, con, int, wis, cha, bonus_str, bonus_dex, bonus_con,
		bonus_int, bonus_wis, bonus_cha FROM units`,
		func(rows *sql.Rows) (data.Unit, error) {
			var u data.Unit
			var brandId int
			err := rows.Scan(&u.Id, &brandId, &u.Name, &u.Level, &u.MaxHp, &u.Gender, &u.BaseAc, &u.Faction,
				&u.Str, &u.Dex, &u.Con, &u.Int, &u.Wis, &u.Cha, &u.BonusStr, &u.BonusDex, &u.BonusCon,
				&u.BonusInt, &u.BonusWis, &u.BonusCha)
			if err != nil {
				return u, err
			}
			u.Brand = brandMap[brandId]
			u.Hp = u.MaxHp
			return u, nil
		})
}
