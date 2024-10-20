package handlers

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"
	"strings"
	"causeurgnocchi/store/models"
)

type Department struct {
	Id       int
	Name     string
}

type DepartmentPageData struct {
	PageTitle string
	Department Department
    Products []models.Product
    Departments []Department
}

type DepartmentHandler struct {
	Db *sql.DB
}

func (h DepartmentHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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

	name := r.PathValue("name")
	products := h.productsByDepartment(name)

	var department Department
	for _, d := range(departments) {
		if strings.ToLower(d.Name) == name {
			department = d
			break
		}
	}

	tmpl := template.Must(template.New("department.tmpl").Funcs(template.FuncMap{
		"toLower": strings.ToLower,
		"kebabCase": func (s string) string { return strings.ReplaceAll(strings.ToLower(s), " ", "-") },
	}).ParseFiles("assets/html/base.html", "assets/html/department.html"))

	data := DepartmentPageData{
		PageTitle: department.Name + " - Store",
		Department: department,
		Products: products,
		Departments: departments,
	}

	tmpl.ExecuteTemplate(w, "base", data)
}


func (h DepartmentHandler) productsByDepartment(name string) []models.Product {
	rows, err := h.Db.Query("select p.id, p.name, p.description, p.price from products p inner join departments d on p.department_id = d.id and d.name = ?", name)
	if (err != nil) {
		log.Fatal(err)
	}
	defer rows.Close()

	var products []models.Product
	for rows.Next() {
		var p models.Product
		err := rows.Scan(&p.Id, &p.Name, &p.Description, &p.Price)
		if err != nil {
			log.Fatal(err)
		}
		products = append(products, p)
	}
	return products
}