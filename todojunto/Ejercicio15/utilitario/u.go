package u

import (
	"fmt"
	"html/template"
	"net/http"
)

type P struct {
	Id       int
	Nombre   string
	Apellido string
}

func Loft(w http.ResponseWriter, r *http.Request) {
	oP := P{
		Id: 1, Nombre: "Oscar Alberto",
		Apellido: "Angarita",
	}
	template, _ := template.ParseFiles("plantillas/loft.html")
	template.Execute(w, oP)
}
func Principal(w http.ResponseWriter, r *http.Request) {
	oP := P{
		Id: 1, Nombre: "Oscar",
		Apellido: "Angarita",
	}
	template, _ := template.ParseFiles("plantillas/index.html")
	template.Execute(w, oP)
}
func Productos(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Pagina Productos")
}
func Categorias(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/", 301)
}
func Noencontro(w http.ResponseWriter, r *http.Request) {
	http.NotFound(w, r)
}
func Error(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Error en el servidor", 500)
}
func (Persona P) Saludar(nom string) string {
	return "Bienvenido ..." + nom
}
