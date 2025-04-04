package util

import (
	"baolhq/branded/lib/data"
	"database/sql"
	"encoding/gob"
	"fmt"
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

// GetBrands retrieves all brands from the database and converts them into a slice of Brand objects
func GetBrands(db *sql.DB) ([]data.Brand, error) {
	rows, err := db.Query(`SELECT name, info, letter, bonus_ac,
		max_str, max_dex, max_con, max_int, max_wis, max_cha,
		bonus_str, bonus_dex, bonus_con, bonus_int, bonus_wis, bonus_cha FROM brands`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var brands []data.Brand
	for rows.Next() {
		var b data.Brand
		err := rows.Scan(&b.Name, &b.Info, &b.Letter, &b.BonusAc,
			&b.MaxStr, &b.MaxDex, &b.MaxCon, &b.MaxInt, &b.MaxWis, &b.MaxCha,
			&b.BonusStr, &b.BonusDex, &b.BonusCon, &b.BonusInt, &b.BonusWis, &b.BonusCha)
		if err != nil {
			return nil, err
		}
		brands = append(brands, b)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return brands, nil
}

// GenerateData take the input db path to generate .dat file at the output path
func GenerateData(input, output string) {
	db, err := OpenDB(input)
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}
	defer CloseDB(db)

	brands, err := GetBrands(db)
	if err != nil {
		log.Fatalf("Error getting brands: %v", err)
	}

	file, err := os.Create(output)
	if err != nil {
		log.Fatalf("Error creating .dat file: %v", err)
	}
	defer file.Close()

	// Encode the brands slice to the file
	encoder := gob.NewEncoder(file)
	err = encoder.Encode(brands)
	if err != nil {
		log.Fatalf("Error encoding data: %v", err)
	}

	fmt.Println("Brands have been successfully encoded into brands.dat")
}

// Load data from SQLite file using connection string
func LoadSql(path string) ([]data.Brand, error) {
	db, err := OpenDB(path)
	if err != nil {
		return nil, err
	}
	defer CloseDB(db)

	brands, err := GetBrands(db)
	if err != nil {
		return nil, err
	}
	return brands, nil
}

// Load generated .dat file at specific path
func LoadData(path string) ([]data.Brand, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Decode the data into the brands slice
	var brands []data.Brand
	decoder := gob.NewDecoder(file)
	err = decoder.Decode(&brands)
	if err != nil {
		return nil, err
	}

	return brands, nil
}
