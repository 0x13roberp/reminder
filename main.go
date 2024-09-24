package main

import (
	"goapi/dbconnect"
	"goapi/routes"
	"log"
	"net/http"

	"html/template"
)

var tmpl *template.Template

// funcion para carchar los archivos estaticos
func init() {
	var err error
	tmpl, err = template.ParseGlob("./views/*.html")

	if err != nil {
		log.Fatal("error rendering templates: ", err)
	}

}

func handler(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{
		"Title":  "My music",
		"Author": "Massano"}
	err := tmpl.ExecuteTemplate(w, "index.html", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	db, err := dbconnect.Connect()
	if err != nil {
		log.Fatalf("could not connect to database: %v", err)
	}

	r := routes.InitRoutes(db)

	r.HandleFunc("/", handler)

	log.Println("running on port 3000")
	log.Fatal(http.ListenAndServe(":3000", r))
}
