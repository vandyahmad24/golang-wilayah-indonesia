package wilayah

import (
	"database/sql"
	"log"
)

func RunMigration(db *sql.DB) {
	schema := `
	CREATE TABLE IF NOT EXISTS provinsi (
		id SERIAL PRIMARY KEY,
		name VARCHAR(100) NOT NULL,
		code VARCHAR(10) NOT NULL
	);

	CREATE TABLE IF NOT EXISTS kota (
		id SERIAL PRIMARY KEY,
		type VARCHAR(50) NOT NULL,
		name VARCHAR(100) NOT NULL,
		code VARCHAR(10) NOT NULL,
		full_code VARCHAR(10) NOT NULL,
		province_id INT NOT NULL REFERENCES provinsi(id) ON DELETE CASCADE
	);

	CREATE TABLE IF NOT EXISTS kecamatan (
		id SERIAL PRIMARY KEY,
		name VARCHAR(100) NOT NULL,
		code VARCHAR(10) NOT NULL,
		full_code VARCHAR(10) NOT NULL,
		city_id INT NOT NULL REFERENCES kota(id) ON DELETE CASCADE
	);

	CREATE TABLE IF NOT EXISTS kelurahan (
		id SERIAL PRIMARY KEY,
		name VARCHAR(100) NOT NULL,
		code VARCHAR(10) NOT NULL,
		full_code VARCHAR(10) NOT NULL,
		pos_code VARCHAR(10) NOT NULL,
		kecamatan_id INT NOT NULL REFERENCES kecamatan(id) ON DELETE CASCADE
	);


	`
	_, err := db.Exec(schema)
	if err != nil {
		panic(err)
	}
	log.Println("Migration Success")
}

func Seed(db *sql.DB, dataPath string) {
	SeedWilayah(db, dataPath)
}
