package presenters

import (
	"fmt"
	"prueba-tecnica1-labora/bug-tracking-system/services/manager"
	"prueba-tecnica1-labora/bug-tracking-system/shared"
)

func InitManagerMenu() {
	for {
		fmt.Println("Bienvenido manager")
		fmt.Println("1. Crear tarea")
		fmt.Println("2. Crear reporte")
		fmt.Println("3. Cerrar sesión")
		fmt.Println("Seleccione una opción: ")
		choice := shared.GetScanner()

		if choice != -1 {
			switch choice {
			case 1:
				fmt.Println("Ingrese un nombre para la tarea: ")
				name := shared.GetScannerString()
				manager.CreateTask(name)
			case 2:
				manager.GenerateBugReport()
			case 3:
				fmt.Println("Hasta pronto")
				return
			default:
				fmt.Println("Opción inválida, intentelo nuevamente")
			}
		}
	}
}
