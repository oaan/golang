package main

import (
  "fmt"
)
func main(){
  fmt.Println("Hola aca estoy!!!")
  var x int=10
  fmt.Println(x)
  editar(&x)
}
func editar(x *int){
  *x=20
}



