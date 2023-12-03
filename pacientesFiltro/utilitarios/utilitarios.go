package utilitarios

import (
	"html/template"
	"net/http"
)

var templateTodos = template.Must(template.New("T").ParseGlob("html/*.html"))
var templateError = template.Must(template.ParseFiles("html/error.html"))

func Request(w http.ResponseWriter, nombrePagina string, estructura interface{}) {
	w.Header().Set("Content-Type", "text/html")
	err := templateTodos.ExecuteTemplate(w, nombrePagina, estructura)
	if err != nil {
		w.WriteHeader(500)
		templateError.Execute(w, nil)
	}
}
