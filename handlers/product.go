package handlers

import (
	"causeurgnocchi/store/models"
	"database/sql"
	"html/template"
	"log"
	"net/http"
)

type ProductHandler struct{
	Db *sql.DB
}

type ProductPageData struct{
	Product models.Product
}

func (h ProductHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	rows, err := h.Db.Query("select name, description, price from products where replace(lower(name), \" \", \"-\") = ?", r.PathValue("product"))
	if err != nil {
		log.Fatal(err)
	}
	
	var product models.Product
	if rows.Next() {
		rows.Scan(&product.Name, &product.Description, &product.Price)
	}

	tmpl := template.Must(template.ParseFiles("assets/html/product.html"))
	data := ProductPageData{
		Product: product, 
	}

	tmpl.Execute(w, data)
}