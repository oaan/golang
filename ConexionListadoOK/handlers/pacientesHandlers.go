package handlers

import (
	"a002/capadatos"
	"a002/models"
	"a002/utilitarios"
	"net/http"
)

func Principal(w http.ResponseWriter, r *http.Request) {
	var oProducto []models.Producto
	if r.Method == "GET" {
		oProducto = capadatos.ListaProducto()
	} else {
		nombre := r.FormValue("nombre")
		oProducto = capadatos.FiltraProducto(nombre)
	}
	utilitarios.Request(w, "index", oProducto)
}
