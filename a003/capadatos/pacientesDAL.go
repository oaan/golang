package capadatos

import (
	"a002/db"
	"a002/models"
)

func ListaPacientes() models.ListaPacientes {
	oLista := models.ListaPacientes{}
	sqlQuery := `select "HCL","AP","NOM" FROM "PARTOS"`
	db.Open()
	rows, _ := db.Query(sqlQuery)
	for rows.Next() {
		oPaciente := models.Pacientes{}
		rows.Scan(&oPaciente.HCL, &oPaciente.AP, &oPaciente.NOM)
		oLista = append(oLista, oPaciente)
	}
	return oLista
}
func FiltraPaciente(AP string) models.ListaPacientes {
	oLista := models.ListaPacientes{}
	sqlQuery := `select * from uspFiltrarProducto($1)`
	db.Open()
	rows, _ := db.Query(sqlQuery, AP)
	for rows.Next() {
		oPaciente := models.Pacientes{}
		rows.Scan(&oPaciente.HCL, &oPaciente.AP, &oPaciente.NOM)
		oLista = append(oLista, oPaciente)
	}
	return oLista
}
func FiltraProductoDolar(dolar int) models.Lista {
	oLista := models.Lista{}
	sqlQuery := `select * from uspFiltrarProductoDolar($1)`
	db.Open()
	rows, _ := db.Query(sqlQuery, dolar)
	for rows.Next() {
		oProducto := models.Producto{}
		rows.Scan(&oProducto.Descripcion, &oProducto.Precio, &oProducto.Cantidad,
			&oProducto.Fecha, &oProducto.Dolar, &oProducto.Id)
		oLista = append(oLista, oProducto)
	}
	return oLista
}
