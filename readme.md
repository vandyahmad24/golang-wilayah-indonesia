
# 🇮🇩 golang-wilayah-indonesia

Package Go untuk **import data wilayah Indonesia** (Provinsi, Kota/Kabupaten, Kecamatan) dari file JSON ke database **MySQL** atau **PostgreSQL** secara otomatis.

---

## 🚀 Fitur Utama
✅ Auto generate table (migration)  
✅ Auto import data dari JSON  
✅ Support **PostgreSQL** dan **MySQL**  
✅ Simple dan reusable, tinggal panggil dari `main.go`

---

## 📦 Instalasi
Tambahkan package ini ke project kamu:
```bash
go get github.com/vandyahmad24/golang-wilayah-indonesia
```

---

## 📂 Struktur JSON Data
Contoh file JSON di folder `data/`:
```json
// province.json
[
    { "id": 1, "name": "Aceh (NAD)", "code": "11" }
]
```
```json
// city.json
[
    { "id": 1, "type": "Kabupaten", "name": "Aceh Barat", "code": "05", "full_code": "1105", "provinsi_id": 1 }
]
```
```json
// district.json
[
    { "id": 1, "name": "Air Majunto", "code": "13", "full_code": "170613", "kabupaten_id": 1 }
]
```

---

## 🛠 Cara Pakai

### 1️⃣ Import package
```go
import "github.com/vandyahmad24/golang-wilayah-indonesia/wilayah"
```

### 2️⃣ Siapkan koneksi database
Contoh koneksi PostgreSQL:
```go
db := wilayah.ConnectDB("postgres", "host=localhost user=postgres password=secret dbname=your_db port=5432 sslmode=disable")
defer db.Close()
```

Untuk MySQL:
```go
db := wilayah.ConnectDB("mysql", "root:password@tcp(localhost:3306)/your_db")
defer db.Close()
```

---

### 3️⃣ Jalankan Migration (Auto create table jika belum ada)
```go
wilayah.RunMigration(db)
```

---

### 4️⃣ Seed data dari JSON ke database
Pastikan folder `data/` berisi `province.json`, `city.json`, dan `district.json`
```go
wilayah.Seed(db, "data")
```

---

## ✅ Contoh Full `main.go`
```go
package main

import (
	"github.com/username/golang-wilayah-indonesia/wilayah"
)

func main() {
	db := wilayah.ConnectDB("postgres", "host=localhost user=postgres password=secret dbname=your_db port=5432 sslmode=disable")
	defer db.Close()

	wilayah.RunMigration(db)
	wilayah.Seed(db, "data")
}
```

---

## 💾 Struktur Database yang Dibuat
- **provinces** (id, name, code)
- **cities** (id, type, name, code, full_code, province_id)
- **districts** (id, name, code, full_code, city_id)

Relasi terjaga dan data akan terisi otomatis sesuai JSON.

---

## 📖 Cara Kerja
✅ Baca file JSON dari folder `data/`  
✅ Insert ke database (skip jika sudah ada - id conflict)  
✅ Logging error jika ada  
✅ Siap pakai di production / development

---

## 📄 License
MIT License

---

## 🤝 Kontribusi
Pull Request dan kontribusi sangat terbuka.  
Silakan open issue jika ada pertanyaan.

---

