package wilayah

import (
	"database/sql"
	"log"
)

func RunMigration(db *sql.DB) {
	schema := `
	CREATE TABLE IF NOT EXISTS provinces (
		id SERIAL PRIMARY KEY,
		name VARCHAR(100) NOT NULL,
		code VARCHAR(10) NOT NULL
	);

	CREATE TABLE IF NOT EXISTS cities (
		id SERIAL PRIMARY KEY,
		type VARCHAR(50) NOT NULL,
		name VARCHAR(100) NOT NULL,
		code VARCHAR(10) NOT NULL,
		full_code VARCHAR(10) NOT NULL,
		province_id INT NOT NULL REFERENCES provinces(id) ON DELETE CASCADE
	);

	CREATE TABLE IF NOT EXISTS districts (
		id SERIAL PRIMARY KEY,
		name VARCHAR(100) NOT NULL,
		code VARCHAR(10) NOT NULL,
		full_code VARCHAR(10) NOT NULL,
		city_id INT NOT NULL REFERENCES cities(id) ON DELETE CASCADE
	);

	CREATE TABLE IF NOT EXISTS villages (
		id SERIAL PRIMARY KEY,
		name VARCHAR(100) NOT NULL,
		code VARCHAR(10) NOT NULL,
		full_code VARCHAR(10) NOT NULL,
		pos_code VARCHAR(10) NOT NULL,
		district_id INT NOT NULL REFERENCES districts(id) ON DELETE CASCADE
	);
	`
	if _, err := db.Exec(schema); err != nil {
		log.Fatal("Migration Failed:", err)
	}
	log.Println("Migration Success")
}

func Seed(db *sql.DB, dataPath string) {
	SeedWilayah(db, dataPath)
}
