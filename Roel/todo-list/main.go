package main

import (
	"fmt"
)

func pi(nombre []int) {
	fmt.Println(nombre)
}
func p(nombre string) {
	fmt.Println(nombre)
}

type Enfermeros struct {
	Nombre string
	Edad   int
	Correo string
}

func main() {
	p("**************SLICE**************************")

	// var a [5]int
	// a[0] = 10
	// var b = [...]int{12, 23, 233, 322, 32}
	// for i := 0; i < len(b); i++ {
	// 	fmt.Println(b[i])
	// }
	enfermeros := []string{"0", "1", "2", "3", "4", "5", "6", "7"}
	fmt.Println("Long", len(enfermeros))
	enfermeros = append(enfermeros[:7], enfermeros[:8]...)
	fmt.Println(enfermeros)
	fmt.Println("Long", len(enfermeros))
	fmt.Println("Cap", cap(enfermeros))
	p("**************MAKE**************************")
	nombres := make([]string, 5, 10) //5 elementos, capacidad 10
	nombres[0] = "Oscar Angarita"
	fmt.Println(nombres)
	rebanada1 := []int{1, 2, 3, 4, 5}
	rebanada2 := make([]int, 5)
	pi(rebanada1)
	copy(rebanada2, rebanada1)
	pi(rebanada1)
	pi(rebanada2)
	p("**************MAPAS***************")
	colors := map[string]string{
		"rojo":  "#ff0000",
		"verde": "#00ff00",
		"azul":  "#0000ff",
	}
	fmt.Println(colors["rojo"])
	colors["negro"] = "#00000"
	fmt.Println(colors)
	if valor, ok := colors["negr"]; ok {
		fmt.Println(valor)
	} else {
		fmt.Println("No existe la clave")
	}
	for clave, valor := range colors {
		fmt.Printf("Clave : %s,valor : %s\n", clave, valor)
	}
	delete(colors, "negro")
	fmt.Println(colors)
	p("*******************ESTRUCTURAS**************************************")

	var e Enfermeros
	e.Nombre = "Oscar"
	e.Edad = 60
	e.Correo = "oaangarita@gmail.com"
	fmt.Println(e)
	e2 := Enfermeros{"Rocio Pamela", 25, "rocio@angarita"}
	fmt.Println(e2)
	fmt.Println(e2.Nombre)
	e2.Edad = 24
	fmt.Println(e2)
	p("**********************Punteros y Metodos******************************************")

	var x int = 10

	// var p *int = &x

	// fmt.Println("&x ", &x)
	// fmt.Println("p ", p)

	fmt.Println(x)
	editar(&x)
	fmt.Println(x)
	e.Hola()
	ariel := Enfermeros{"Ariel", 38, "gmail"}
	ariel.Hola()

	p("************************ FIN MAIN **********************************")
}
func editar(x *int) {
	*x = 20
}
func (e *Enfermeros) Hola() {
	fmt.Println("Hola mi nombre es", e.Nombre)
}
