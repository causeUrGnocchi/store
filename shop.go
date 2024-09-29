package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

const ConnectionString = "root:example@(localhost)/shop?parseTime=true"

type Product struct {
	id          int
	name        string
	description string
	price       float32
}

func main() {
	db, err := sql.Open("mysql", ConnectionString)
	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		rows, err := db.Query("select * from products")
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()

		var products []Product
		for rows.Next() {
			var p Product

			err := rows.Scan(&p.id, &p.name, &p.description, &p.price)
			if err != nil {
				log.Fatal(err)
			}

			products = append(products, p)
		}

		for _, p := range products {
			fmt.Fprintf(w, "Id: %d, Name: %s, Description: %s, Price: %f\n", p.id, p.name, p.description, p.price)
		}
	})

	{
		fs := http.FileServer(http.Dir("static/"))
		http.Handle("/static/", http.StripPrefix("/static/", fs))
	}

	http.ListenAndServe(":8080", nil)
}
