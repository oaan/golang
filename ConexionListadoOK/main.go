package main

import (
	"a002/capadatos"
	"a002/handlers"
	"fmt"
	"net/http"

	_ "github.com/lib/pq"
)

func main() {

	oLista := capadatos.ListaProducto()
	for _, v := range oLista {
		fmt.Println(v.Descripcion, v.Precio, v.Cantidad, v.Fecha, v.Dolar, v.Id)
	}
	http.HandleFunc("/", handlers.Principal)
	Oscar("Oscar Alberto", "Angarita Roberts", 60)
	http.ListenAndServe(":8000", nil)
}

func Oscar(nombre, apellido string, edad int) {

	fmt.Println("\n********************************************************\nhttp://localhost:8000/", "Se activo la funcion...", edad, nombre, apellido)

}
