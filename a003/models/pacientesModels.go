package models

type Lista []Producto

type Producto struct {
	Descripcion string
	Precio      float64
	Cantidad    int
	Fecha       string
	Dolar       int
	Id          int
	Amigos      []string
}

type ListaPacientes []Pacientes
type Pacientes struct {
	HCL int
	AP  string
	NOM string
}
