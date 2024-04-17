package models

import "prueba-tecnica1-labora/bug-tracking-system/shared"

type User struct {
	Id       string
	Username string
	Password string
	Role     string
}

var (
	Users = [3]User{
		{
			Id:       shared.GenerateUUID(),
			Username: "admin",
			Password: "root",
			Role:     "MANAGER",
		},
		{
			Id:       shared.GenerateUUID(),
			Username: "dev",
			Password: "developer",
			Role:     "DEVELOPER",
		},
		{
			Id:       shared.GenerateUUID(),
			Username: "qa",
			Password: "tester",
			Role:     "TESTER",
		},
	}
)
