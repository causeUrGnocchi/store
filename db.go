package db

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func Connect() {
	db, err := sql.Open("mysql", "root:example@(localhost)/shop?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
}
