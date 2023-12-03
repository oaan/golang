package handlers

import (
	"a002/capadatos"
	"a002/models"
	"a002/utilitarios"
	"net/http"
)

type ListarPacienteForm struct {
	ListaPaciente []models.Paciente
	Nombre        string
}

type ListarPacienteComboForm struct {
	ListaPaciente []models.Paciente
	HCL           int
	OS            string
}

func Paciente(w http.ResponseWriter, r *http.Request) {
	var nombrePaciente string
	var oPaciente []models.Paciente
	if r.Method == "GET" {
		oPaciente = capadatos.ListaPaciente()
	} else {
		nombrePaciente = r.FormValue("nombre")
		oPaciente = capadatos.FiltraPaciente(nombrePaciente)
	}
	obj := ListarPacienteForm{ListaPaciente: oPaciente, Nombre: nombrePaciente}
	utilitarios.Request(w, "index", obj)
}
