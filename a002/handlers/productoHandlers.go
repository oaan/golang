package handlers

import (
	"a002/capadatos"
	"a002/models"
	"a002/utilitarios"
	"net/http"
)

type ListarProductoForm struct {
	ListaProducto []models.Producto
	Nombre        string
}

func Principal(w http.ResponseWriter, r *http.Request) {
	var nombreProducto string
	var oProducto []models.Producto
	if r.Method == "GET" {
		oProducto = capadatos.ListaProducto()
	} else {
		nombreProducto = r.FormValue("nombre")
		oProducto = capadatos.FiltraProducto(nombreProducto)
	}
	obj := ListarProductoForm{ListaProducto: oProducto, Nombre: nombreProducto}
	utilitarios.Request(w, "index", obj)
}
