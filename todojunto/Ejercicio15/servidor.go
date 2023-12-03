package main

import (
	// "Ejercicio15/db"
	"Ejercicio15/utilitario"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", u.Principal)
	http.HandleFunc("/loft", u.Loft)
	http.HandleFunc("/productos", u.Productos)
	http.HandleFunc("/categorias", u.Categorias)
	http.HandleFunc("/noencontro", u.Noencontro)
	http.HandleFunc("/error", u.Error)
	fmt.Println("http://localhost:8000/")
	http.ListenAndServe(":8000", nil)

}
