package main

import (
	"causeurgnocchi/store/handlers"
	"database/sql"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

const ConnectionString = "root:example@(localhost)/store?parseTime=true"

var db *sql.DB

func main() {
    {
        var err error
        db, err = sql.Open("mysql", ConnectionString)
        if err != nil {
            log.Fatal(err)
        }
        if err := db.Ping(); err != nil {
            log.Fatal(err)
        }
    }

    assets := http.FileServer(http.Dir("assets"))
    http.Handle("/assets/", http.StripPrefix("/assets/", assets))

	http.Handle("/", http.RedirectHandler("/departments/1", 302))

	http.Handle("/departments/{id}", handlers.DepartmentHandler{
		Db: db,
	})

	http.ListenAndServe(":8080", nil)
}