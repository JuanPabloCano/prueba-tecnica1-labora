package presenters

import (
	"fmt"
	"prueba-tecnica1-labora/bug-tracking-system/services/tester"
	"prueba-tecnica1-labora/bug-tracking-system/shared"
)

func InitTesterMenu() {
	for {
		fmt.Println("Bienvenido tester")
		fmt.Println("Lista de tareas completadas: ")
		tester.GetAllCompletedTasks()
		fmt.Println("1. Crear bug")
		fmt.Println("2. Actualizar estado de bug")
		fmt.Println("3. Ver listado de bugs")
		fmt.Println("4. Cerrar sesión")
		fmt.Println("Seleccione una opción: ")
		choice := shared.GetScanner()

		switch choice {
		case 1:
			fmt.Println("Ingrese una descripción del bug: ")
			description := shared.GetScannerString()
			fmt.Println("Ingrese el Id de la tarea: ")
			taskId := shared.GetScannerString()
			tester.CreateBug(description, taskId)
		case 2:
			fmt.Println("Lista de Bugs: ")
			tester.GetAllBugs()
			fmt.Println("Ingrese el ID del bug: ")
			id := shared.GetScannerString()
			fmt.Println("Ingrese uno de los siguientes estados: ")
			fmt.Println("1. PENDING - 2. RESOLVED")
			status := shared.GetScanner()
			tester.UpdateBugStatus(id, status)
		case 3:
			tester.GetAllBugs()
		case 4:
			fmt.Println("Hasta pronto")
			return
		default:
			fmt.Println("Opción inválida, intentelo nuevamente")
		}
	}
}
