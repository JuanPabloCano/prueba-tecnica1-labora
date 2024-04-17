package presenters

import (
	"fmt"
	"prueba-tecnica1-labora/bug-tracking-system/services/developer"
	"prueba-tecnica1-labora/bug-tracking-system/shared"
)

func InitDeveloperMenu() {
	for {
		fmt.Println("Bienvenido developer")
		fmt.Println("Lista de tareas pendientes: ")
		developer.GetAllPendingTasks()
		fmt.Println("1. Actualizar estado de tarea")
		fmt.Println("2. Cerrar sesión")
		fmt.Println("Seleccione una opción: ")
		choice := shared.GetScanner()

		switch choice {
		case 1:
			fmt.Println("Ingrese el ID de la tarea: ")
			id := shared.GetScannerString()
			fmt.Println("Ingrese uno de los siguientes estados: ")
			fmt.Println("1. PENDING - 2. RESOLVED")
			status := shared.GetScanner()
			developer.UpdateTaskStatus(id, status)
		case 2:
			fmt.Println("Hasta pronto")
			return
		default:
			fmt.Println("Opción inválida, intentelo nuevamente")
		}
	}
}
