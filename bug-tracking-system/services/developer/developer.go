package developer

import (
	"prueba-tecnica1-labora/bug-tracking-system/shared"
	"prueba-tecnica1-labora/bug-tracking-system/storage"
)

var (
	pending = 1
)

func UpdateTaskStatus(id string, choice int) {
	storage.ChangeTaskStatus(id, shared.Status[choice])
}

func GetAllPendingTasks() {
	storage.GetAllTasksByStatus(pending)
}
