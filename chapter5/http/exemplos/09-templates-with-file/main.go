package main

import (
	"html/template"
	"net/http"
	"strings"
)

type Curso struct {
	Nome string
	CargaHoraria int
}

type Cursos []Curso

func main() {

	var templates []string = []string{
		"content.html",
		"header.html",
		"footer.html",
	}
	
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){

		t := template.New("content.html")
		t.Funcs(template.FuncMap{"ToUpper": strings.ToUpper})

		tmp := template.Must(t.ParseFiles(templates...))
		error := tmp.Execute(w, Cursos{
			{"Go", 40},
			{"Java", 90},
			{"Javascript", 60},
		})
	
		if error != nil {
			panic(error)
		}
	})

	http.ListenAndServe(":8080", nil)
}