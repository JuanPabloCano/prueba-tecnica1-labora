package manager

import (
	"fmt"
	"os"
	"prueba-tecnica1-labora/bug-tracking-system/models"
	"prueba-tecnica1-labora/bug-tracking-system/shared"
	"prueba-tecnica1-labora/bug-tracking-system/storage"
	"time"
)

func CreateTask(name string) {
	newTask := models.Task{
		Id:     shared.GenerateUUID(),
		Name:   name,
		Status: "PENDING",
	}
	storage.CreateTask(newTask)
}

func GenerateBugReport() {
	fileName := fmt.Sprintf("bug-tracking-system/reports/registro_cantidad_teclas_%s.txt", time.Now().Format("20060102_150405"))
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Printf("Error al crear el archivo de registro: %v\n", err)
		return
	}
	defer file.Close()

	keystrokes := shared.CountKeyStrokes()
	_, err = file.WriteString(fmt.Sprintf("Cantidad de teclas pulsadas: %v", keystrokes))
	if err != nil {
		fmt.Printf("Error al escribir en el archivo de registro: %v\n", err)
		return
	}
	fmt.Println("Archivo de registro de cantidad de teclas pulsadas %s creado exitosamente.\n", fileName)
}
