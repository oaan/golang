package main

import (
	"fmt"
)

func divide(div, divs int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()
 validateZero(divs)
 fmt.Println(div / divs)
}
func validateZero(divs int) {
	if divs == 0 {
		panic("No puedes dividir por cero")
	}

}
func main() {
	divide(100, 10)
	divide(100, 0)
	divide(100, 10)
	divide(100, 10)

	divide(100, 0)
	divide(100, 0)
}
