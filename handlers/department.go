package handlers

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

type Department struct {
	Id       int
	Name     string
}

type Product struct {
	Id           int
	Name         string
	Description  string
	Price        float32
}

type DepartmentPageData struct {
    Products []Product
    Departments []Department
	DepartmentId int
}

type DepartmentHandler struct {
	Db *sql.DB
}

func (h DepartmentHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	tmpl := template.Must(template.New("department.html").Funcs(template.FuncMap{
		"decrement": func(a int, b int) int {
			return a - b
		},
	}).ParseFiles("assets/html/department.html"))
		
	rows, err := h.Db.Query("select id, name from departments")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var departments []Department
	for rows.Next() {
		var d Department
		err := rows.Scan(&d.Id, &d.Name)
		if err != nil {
			log.Fatal(err)
		}
		departments = append([]Department{d}, departments...)
	}

	id, err := strconv.Atoi(req.PathValue("id"))
	if (err != nil) {
		id = 1
	}

	data := DepartmentPageData {
		Products: h.productsByDepartment(id),
		Departments: departments,
		DepartmentId: id,
	}
	tmpl.Execute(w, data)
}


func (h DepartmentHandler) productsByDepartment(department int) []Product {
	rows, err := h.Db.Query("select id, name, description, price from products where department_id = ?", department)
	if (err != nil) {
		log.Fatal(err)
	}
	defer rows.Close()

	var products []Product
	for rows.Next() {
		var p Product
		err := rows.Scan(&p.Id, &p.Name, &p.Description, &p.Price)
		if err != nil {
			log.Fatal(err)
		}
		products = append(products, p)
	}
	return products
}