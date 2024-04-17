package manager

import (
	"prueba-tecnica1-labora/bug-tracking-system/models"
	"prueba-tecnica1-labora/bug-tracking-system/shared"
	"prueba-tecnica1-labora/bug-tracking-system/storage"
)

func CreateTask(name string) {
	newTask := models.Task{
		Id:     shared.GenerateUUID(),
		Name:   name,
		Status: "PENDING",
	}
	storage.CreateTask(newTask)
}

// TODO: create file service
