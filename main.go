package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/prakashsingha/go-webapp/controller"
	"github.com/prakashsingha/go-webapp/middleware"
	"github.com/prakashsingha/go-webapp/model"

	_ "github.com/lib/pq"

	_ "net/http/pprof"
)

func main() {
	templates := populateTemplates()
	db := connectToDatabase()
	defer db.Close()

	controller.Startup(templates)
	http.ListenAndServeTLS(":8000", "cert.pem", "key.pem", &middleware.TimeoutMiddleware{new(middleware.GzipMiddleware)})
}

func populateTemplates() map[string]*template.Template {
	result := make(map[string]*template.Template)
	const basePath = "templates"

	layout := template.Must(template.ParseFiles(basePath + "/_layout.html"))
	template.Must(layout.ParseFiles(basePath+"/_header.html", basePath+"/_footer.html"))

	dir, err := os.Open(basePath + "/content")
	if err != nil {
		panic("Unable to open folder")
	}

	files, err := dir.Readdir(-1)
	if err != nil {
		panic("Unable to read the contents of the directory")
	}

	for _, fi := range files {
		f, err := os.Open(basePath + "/content/" + fi.Name())
		if err != nil {
			panic("Failed to open template '" + fi.Name() + "'")
		}

		content, err := ioutil.ReadAll(f)
		if err != nil {
			panic("Failed to read the content from file '" + f.Name() + "'")
		}
		f.Close()

		tmpl := template.Must(layout.Clone())
		_, err = tmpl.Parse(string(content))
		if err != nil {
			panic("Failed to parse contents of '" + fi.Name() + "' as template")
		}
		result[fi.Name()] = tmpl
	}
	return result
}

func connectToDatabase() *sql.DB {
	db, err := sql.Open("postgres", "postgres://go-webapp:password@localhost/go-webapp?sslmode=disable")
	if err != nil {
		log.Fatalln(fmt.Errorf("Unable to connect to database: %v", err))
	}

	model.SetDatabase(db)
	return db
}
