package main

import (
	"fmt"
)

func main() {
	medicos := map[int]string{
		1: "Oscar",
		2: "Sebastian",
		3: "Hector",
		4: "Laura",
	}
	fmt.Println(medicos)

	fmt.Println(medicos[1])
	medicos[5] = "Moyano"
	fmt.Println(medicos[4])
	fmt.Println(medicos)
	if valor, ok := medicos[1]; !ok {
		fmt.Println("No existe la clave...")
	} else {
		fmt.Println("Ok", valor)

	}
	delete(medicos, 1)
	fmt.Println(medicos)
}
