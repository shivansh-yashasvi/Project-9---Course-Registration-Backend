package DBConnection

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func ConDB() *sql.DB {

	db, err := sql.Open("mysql", "root:root@(127.0.0.1:3306)/VIT?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}

	return db
}
