package models

type Amigos []string
type Producto struct {
	Descripcion string
	Precio      float64
	Cantidad    int
	Fecha       string
	Dolar       int
	Id          int
	Amigos      []string
}
type Lista []Producto
