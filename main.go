package main

import (
	"causeurgnocchi/store/handlers"
	"database/sql"
	"embed"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

const ConnectionString = "root:example@(localhost)/store?parseTime=true"

//go:embed assets
var assets embed.FS
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

    {
        fs := http.FileServer(http.Dir("assets"))
        http.Handle("/static/", http.StripPrefix("/static/", fs))
    }

	http.Handle("/", handlers.DepartmentHandler{
        Assets: assets,
        Db: db,
    })

	http.ListenAndServe(":8080", nil)
}