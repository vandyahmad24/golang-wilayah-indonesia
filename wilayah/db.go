package wilayah

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

func ConnectDB(driver, dsn string) *sql.DB {
	db, err := sql.Open(driver, dsn)
	if err != nil {
		log.Fatal("Failed connect:", err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal("Failed ping:", err)
	}
	return db
}
