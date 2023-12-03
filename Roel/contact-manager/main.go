package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

type Contact struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

// guarda contactos en un archivo json
func saveContactsToFile(contacts []Contact) error {
	file, err := os.Create("contacts.json")
	if err != nil {
		return err
	}

	defer file.Close()
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&contacts)
	if err != nil {
		return err
	}
	return nil
}

func loadContactsFromFile(contacts *[]Contact) error {
	file, err := os.Open("contacts.json")
	if err != nil {
		return err
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&contacts)
	if err != nil {
		return err
	}
	return nil
}
func main() {
	//slice de contactos
	var contacts []Contact
	//cargar contactos existentes desde el archivo

	err := loadContactsFromFile(&contacts)
	if err != nil {
		fmt.Println("Error al cargar los contactos: ", err)
	}
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(`====Gestor de Contactos =====
            1.Agregar un contacto
            2.Mostrar todos los contactos
            3.Salir
            Elige una opcion: `)
		//leer la opcion del usuario
		var option int
		_, err = fmt.Scanln(&option)
		if err != nil {
			fmt.Println("Error al leer la opcion: ", err)
			return
		}
		switch option {
		case 1:
			// ingresar y crear contacto
			var c Contact
			fmt.Print("Nombre: ")
			c.Name, _ = reader.ReadString('\n')

			fmt.Print("Email: ")
			c.Email, _ = reader.ReadString('\n')

			fmt.Print("Telefono: ")
			c.Phone, _ = reader.ReadString('\n')
			// agregar un contacto a slice
			contacts = append(contacts, c)
			// guardar en un archivo json
			if err := saveContactsToFile(contacts); err != nil {
				fmt.Println("Error al guardar el contacto: ", err)
			}
		case 2:
			fmt.Println("======================================================")
			for index, contact := range contacts {
				fmt.Printf("%d. Nombre: %s Email: %s Telefono: %s\n",
					index+1, contact.Name, contact.Email, contact.Phone)
			}
			fmt.Println("======================================================")
		case 3:
			//salir del programa
			return
		default:
		}
	}
}
