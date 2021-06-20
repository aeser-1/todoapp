package dbconn

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func Connection() *sql.DB {

	db, err := sql.Open("mysql", "root:pass123@tcp(localhost:3306)/todo?parseTime=true")
	if err != nil {
		log.Fatalf("Error: %+v\n", err)
	}
	return db
}
