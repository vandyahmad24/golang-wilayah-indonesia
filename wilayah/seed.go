package wilayah

import (
	"database/sql"
	"encoding/json"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

func SeedWilayah(db *sql.DB, dataPath string) {
	seedProvince(db, filepath.Join(dataPath, "provinsi.json"))
	seedCity(db, filepath.Join(dataPath, "kota.json"))
	seedDistrict(db, filepath.Join(dataPath, "kecamatan.json"))
	seedKelurahan(db, filepath.Join(dataPath, "kelurahan.json"))
	log.Println("Seed Success")
}

func seedProvince(db *sql.DB, path string) {
	log.Println("Seed Province")
	// Akses internal data JSON di library
	_, b, _, _ := runtime.Caller(0)
	basePath := filepath.Dir(b)
	file, err := os.Open(filepath.Join(basePath, "data", "provinsi.json"))
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	var provinsi []Province
	if err := json.NewDecoder(file).Decode(&provinsi); err != nil {
		log.Fatal(err)
	}

	for _, p := range provinsi {
		_, err := db.Exec(`INSERT INTO provinsi (id, name, code) VALUES ($1, $2, $3) ON CONFLICT (id) DO NOTHING`, p.ID, p.Name, p.Code)
		if err != nil {
			log.Println("Province Error:", err)
		}
	}
}

func seedCity(db *sql.DB, path string) {
	log.Println("Seed City")
	_, b, _, _ := runtime.Caller(0)
	basePath := filepath.Dir(b)
	file, err := os.Open(filepath.Join(basePath, "data", "kota.json"))
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	var kota []City
	if err := json.NewDecoder(file).Decode(&kota); err != nil {
		log.Fatal(err)
	}

	for _, c := range kota {
		_, err := db.Exec(`INSERT INTO kota (id, type, name, code, full_code, province_id) 
			VALUES ($1, $2, $3, $4, $5, $6) 
			ON CONFLICT (id) DO NOTHING`, c.ID, c.Type, c.Name, c.Code, c.FullCode, c.ProvinceID)
		if err != nil {
			log.Println("City Error:", err)
		}
	}
}

func seedDistrict(db *sql.DB, path string) {
	log.Println("Seed District")
	_, b, _, _ := runtime.Caller(0)
	basePath := filepath.Dir(b)
	file, err := os.Open(filepath.Join(basePath, "data", "kecamatan.json"))
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	var kecamatan []District
	if err := json.NewDecoder(file).Decode(&kecamatan); err != nil {
		log.Fatal(err)
	}

	for _, d := range kecamatan {
		_, err := db.Exec(`INSERT INTO kecamatan (id, name, code, full_code, city_id) 
			VALUES ($1, $2, $3, $4, $5) 
			ON CONFLICT (id) DO NOTHING`, d.ID, d.Name, d.Code, d.FullCode, d.CityID)
		if err != nil {
			log.Println("District Error:", err)
		}
	}
}

func seedKelurahan(db *sql.DB, path string) {
	log.Println("Seed Village")
	_, b, _, _ := runtime.Caller(0)
	basePath := filepath.Dir(b)
	file, err := os.Open(filepath.Join(basePath, "data", "kelurahan.json"))
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	var kelurahan []Village
	if err := json.NewDecoder(file).Decode(&kelurahan); err != nil {
		log.Fatal(err)
	}

	for _, v := range kelurahan {
		_, err := db.Exec(`INSERT INTO kelurahan (id, name, code, full_code, pos_code, kecamatan_id) 
			VALUES ($1, $2, $3, $4, $5, $6) 
			ON CONFLICT (id) DO NOTHING`, v.ID, v.Name, v.Code, v.FullCode, v.PosCode, v.DistrictID)
		if err != nil {
			log.Println("Village Error:", err)
		}
	}
}
