package capadatos

import (
	"a002/db"
	"a002/models"
)

func ListaPaciente() models.ListaPacientes {
	oLista := models.ListaPacientes{}
	sqlQuery := `select "HCL","NOM","AP","OS","IdHCL","AFILIADO" FROM "PARTOS"`
	db.Open()
	rows, _ := db.Query(sqlQuery)
	for rows.Next() {
		oPaciente := models.Paciente{}
		rows.Scan(&oPaciente.HCL, &oPaciente.NOM,
			&oPaciente.AP, &oPaciente.OS, &oPaciente.IdHCL, &oPaciente.AFILIADO)
		oLista = append(oLista, oPaciente)
	}
	return oLista
}

func FiltraPaciente(nombrePaciente string) models.ListaPacientes {
	oLista := models.ListaPacientes{}
	sqlQuery := `select * from uspFiltrarPaciente($1)`
	db.Open()
	rows, _ := db.Query(sqlQuery, nombrePaciente)
	for rows.Next() {
		oPaciente := models.Paciente{}
		rows.Scan(&oPaciente.HCL, &oPaciente.NOM, &oPaciente.AP, &oPaciente.OS)
		oLista = append(oLista, oPaciente)
	}
	return oLista
}
func FiltraPacienteXHCL(HCL int) models.ListaPacientes {
	oLista := models.ListaPacientes{}
	sqlQuery := `select * from uspFiltrarPacienteXHCL($1)`
	db.Open()
	rows, _ := db.Query(sqlQuery, HCL)
	for rows.Next() {
		oPaciente := models.Paciente{}
		rows.Scan(&oPaciente.HCL, &oPaciente.NOM, &oPaciente.AP, &oPaciente.OS)
		oLista = append(oLista, oPaciente)
	}
	return oLista
}
func FiltraPacienteXOS(OS string) models.ListaPacientes {
	oLista := models.ListaPacientes{}
	sqlQuery := `select * from uspFiltrarPacienteXOS($1)`
	db.Open()
	rows, _ := db.Query(sqlQuery, OS)
	for rows.Next() {
		oPaciente := models.Paciente{}
		rows.Scan(&oPaciente.HCL, &oPaciente.NOM, &oPaciente.AP, &oPaciente.OS)
		oLista = append(oLista, oPaciente)
	}
	return oLista
}
