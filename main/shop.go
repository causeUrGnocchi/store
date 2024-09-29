package main

import (
	db "causeurgnocchi/shop"
	"fmt"
	"net/http"
)

func main() {
	db.Connect()

    http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Welcome to my website!")
    })

	http.HandleFunc("GET /product/{id}", func (w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Requested product of id %s", r.PathValue("id"))
	})

    fs := http.FileServer(http.Dir("static/"))
    http.Handle("/static/", http.StripPrefix("/static/", fs))

    http.ListenAndServe(":8080", nil)
}