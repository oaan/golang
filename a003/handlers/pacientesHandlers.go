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
	var oPaciente []models.Pacientes
	if r.Method == "GET" {
		oPaciente = capadatos.ListaPacientes()
	} else {
		oPaciente = r.FormValue("nombre")
		
	}
	utilitarios.Request(w, "index", oPaciente)
}
