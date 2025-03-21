package wilayah

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"log"
	"path/filepath"
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
	data, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	var provinsi []Province
	json.Unmarshal(data, &provinsi)

	for _, p := range provinsi {
		_, err := db.Exec(`INSERT INTO provinsi (id, name, code) VALUES ($1, $2, $3) ON CONFLICT (id) DO NOTHING`, p.ID, p.Name, p.Code)
		if err != nil {
			log.Println("Province Error:", err)
		}
	}
}

func seedCity(db *sql.DB, path string) {
	log.Println("Seed City")
	data, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	var kota []City
	json.Unmarshal(data, &kota)

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
	data, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	var kecamatan []District
	json.Unmarshal(data, &kecamatan)

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
	data, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	var kecamatan []Village
	json.Unmarshal(data, &kecamatan)

	for _, v := range kecamatan {
		_, err := db.Exec(`INSERT INTO kelurahan (id, name, code, full_code, pos_code, kecamatan_id) 
			VALUES ($1, $2, $3, $4, $5, $6) 
			ON CONFLICT (id) DO NOTHING`, v.ID, v.Name, v.Code, v.FullCode, v.PosCode, v.DistrictID)
		if err != nil {
			log.Println("Village Error:", err)
		}
	}
}
