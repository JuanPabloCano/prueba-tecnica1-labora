package storage

import (
	"fmt"
	"github.com/olekukonko/tablewriter"
	"os"
	"prueba-tecnica1-labora/bug-tracking-system/models"
	"prueba-tecnica1-labora/bug-tracking-system/shared"
)

var (
	tasks = []models.Task{
		{
			Id:     "bd83a5d0",
			Name:   "Task One",
			Status: "PENDING",
		},
		{
			Id:     "bd65a5f7",
			Name:   "Task Two",
			Status: "PENDING",
		},
	}
	bugs = []models.Bug{
		{
			Id:          "sv54r5b1",
			Description: "lorem ipsum dolor sit amet",
			Status:      "PENDING",
			TaskId:      "bd83a5d0",
		},
		{
			Id:          "pm58q3b4",
			Description: "lorem ipsum dolor sit amet",
			Status:      "PENDING",
			TaskId:      "bd65a5f7",
		},
	}
)

func CreateTask(task models.Task) {
	tasks = append(tasks, task)
	fmt.Println("Tarea creada exitosamente")
}

func CreateBug(bug models.Bug) {
	bugs = append(bugs, bug)
	fmt.Println("Bug creado exitosamente")
}

func GetAllTasksByStatus(status int) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Id", "Name", "Status"})

	for _, task := range tasks {
		if task.Status == shared.Status[status] {
			taskData := []string{
				task.Id,
				task.Name,
				task.Status,
			}
			table.Append(taskData)
		}
	}
	table.Render()
}

func GetAllBugs() {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Id", "Description", "Status", "TaskId"})

	for _, bug := range bugs {
		bugsData := []string{
			bug.Id,
			bug.Description,
			bug.Status,
			bug.TaskId,
		}
		table.Append(bugsData)
	}
	table.Render()
}

func ChangeTaskStatus(id string, status string) {
	for idx, task := range tasks {
		if task.Id == id {
			newTask := task.UpdateStatus(status)
			tasks[idx] = newTask
			fmt.Println("Estado de tarea actualizado correctamente")
			return
		}
	}
	fmt.Println("Tarea no encontrada, intentelo de nuevo")
}

func ChangeBugStatus(id string, status string) {
	for idx, bug := range bugs {
		if bug.Id == id {
			newBug := bug.UpdateStatus(status)
			bugs[idx] = newBug
			fmt.Println("Estado de bug actualizado correctamente")
			return
		}
	}
	fmt.Println("Bug no encontrado, intentelo de nuevo")
}
