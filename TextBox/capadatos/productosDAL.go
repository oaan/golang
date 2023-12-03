package capadatos

import (
	"a002/db"
	"a002/models"
)

func ListaProducto() models.Lista {
	oLista := models.Lista{}
	sqlQuery := `select * from uspListarProducto()`
	db.Open()
	rows, _ := db.Query(sqlQuery)
	for rows.Next() {
		oProducto := models.Producto{}
		rows.Scan(&oProducto.Descripcion, &oProducto.Precio, &oProducto.Cantidad,
			&oProducto.Fecha, &oProducto.Dolar, &oProducto.Id)
		oLista = append(oLista, oProducto)
	}
	return oLista
}
func FiltraProducto(nombreProducto string) models.Lista {
	oLista := models.Lista{}
	sqlQuery := `select * from uspFiltrarProducto($1)`
	db.Open()
	rows, _ := db.Query(sqlQuery, nombreProducto)
	for rows.Next() {
		oProducto := models.Producto{}
		rows.Scan(&oProducto.Descripcion, &oProducto.Precio, &oProducto.Cantidad,
			&oProducto.Fecha, &oProducto.Dolar, &oProducto.Id)
		oLista = append(oLista, oProducto)
	}
	return oLista
}
