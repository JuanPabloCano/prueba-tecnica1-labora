package presenters

import (
	"fmt"
	"prueba-tecnica1-labora/bug-tracking-system/shared"
)

func InitMenu() {
	for {
		fmt.Println(":: EJERCICIO ENTREGABLE ::")
		fmt.Println("Bienvenido")
		fmt.Println("1. Iniciar sesión")
		fmt.Println("0. Volver al menú principal")
		fmt.Print("Seleccione una opción: ")
		choice := shared.GetScanner()

		switch choice {
		case 1:
		case 0:
			fmt.Println("Hasta luego")
			return
		default:
			fmt.Println("Opción incorrecta, intentelo de nuevo")
		}
	}
}
