package dbconn

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func Connection() *sql.DB {

	db, err := sql.Open("mysql", "root:pass123@tcp(localhost:3306)/todo?parseTime=true")
	if err != nil {
		panic(err.Error())
	}
	return db
}
