package greetings
import "fmt"
//Hello devuelve un saludo
func Hello(name string)string{
  message:=fmt.Sprintf("Hola , %v! Bienvenido!",name)
  return message
}

