package main

import (
	"bufio"
	"fmt"
	"os"
)

type Tareas struct {
	nombre     string
	desc       string
	completado bool
}
type ListaTareas struct {
	tareas []Tareas
}

// metodo que recibe (receptor) dato tipo tarea,instancia l puntero a ListaTareas, t instancia Tarea
// slice l.tareas llama la instancia y le agrega t
func (l *ListaTareas) agregarTarea(t Tareas) {
	l.tareas = append(l.tareas, t)
}

// metodo que recibe un index y lo uso para buscar
func (l *ListaTareas) marcarCompletado(index int) {
	l.tareas[index].completado = true
}
func (l *ListaTareas) editarTarea(index int, t Tareas) {
	l.tareas[index] = t
}
func (l *ListaTareas) eliminarTarea(index int) {
	l.tareas = append(l.tareas[:index], l.tareas[index+1:]...)
}

func main() {
	//instancia de ListaTareas
	lista := ListaTareas{}
	leer := bufio.NewReader(os.Stdin)
	for {
		var opcion int
		fmt.Println(`Seleccione una opcion:
          1.Agregar tarea
          2.Marcar tarea como completada
          3.Editar tarea
          4.Eliminar tarea
          5.Salir
          
              Ingrese la opcion:`)
		fmt.Scanln(&opcion)
		switch opcion {
		case 1:
			var t Tareas
			fmt.Print("Ingrese el nombre de la tarea..")
			t.nombre, _ = leer.ReadString('\n')
			fmt.Print("Ingrese la descripcion de la tarea..")
			t.desc, _ = leer.ReadString('\n')
			lista.agregarTarea(t)
			fmt.Println("Tarea agregada correctamente")

		case 2:
			var index int
			fmt.Println("Ingrese el indice de la tarea que desea marcar como completada:")
			fmt.Scanln(&index)
			lista.marcarCompletado(index)
			fmt.Println("Tarea marcada como completada correctamente.")
		case 3:
			var index int
			var t Tareas
			fmt.Print("Ingrese el indice de la tarea que desea actualizar:")
			fmt.Scanln(&index)
			fmt.Println("Ingrese el nombre de la tarea: ")
			t.nombre, _ = leer.ReadString('\n')
			fmt.Println("Ingrese la descripcion de la tarea: ")
			t.desc, _ = leer.ReadString('\n')
			lista.editarTarea(index, t)
			fmt.Println("Tarea actualizada correctamente.")

		case 4:
			var index int
			fmt.Print("Ingrese el indice de la tarea que desea eliminar: ")
			fmt.Scanln(&index)
			lista.eliminarTarea(index)
			fmt.Println("Tarea eliminada correctamente")
		case 5:
			fmt.Println("Saliendo del programa....")
			return
		default:
			fmt.Println("Opcion invalida.")

			//listar todas las tareas
			fmt.Println("Lista de Tareas:")
			fmt.Println("===================================================================")
			for i, t := range lista.tareas {
				fmt.Printf("%d. %s - %s - Completado: %t\n", i, t.nombre, t.desc, t.completado)
			}
			fmt.Println("===================================================================")
		}
	}
}
