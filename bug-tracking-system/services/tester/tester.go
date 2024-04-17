package tester

import (
	"prueba-tecnica1-labora/bug-tracking-system/models"
	"prueba-tecnica1-labora/bug-tracking-system/shared"
	"prueba-tecnica1-labora/bug-tracking-system/storage"
)

var (
	resolved = 2
)

func CreateBug(description string, taskId string) {
	newBug := models.Bug{
		Id:          shared.GenerateUUID(),
		Description: description,
		Status:      shared.Status[1],
		TaskId:      taskId,
	}
	storage.CreateBug(newBug)
}

func UpdateBugStatus(id string, choice int) {
	storage.ChangeBugStatus(id, shared.Status[choice])
}

func GetAllCompletedTasks() {
	storage.GetAllTasksByStatus(resolved)
}

func GetAllBugs() {
	storage.GetAllBugs()
}
