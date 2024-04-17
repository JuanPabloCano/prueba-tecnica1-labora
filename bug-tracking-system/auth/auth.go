package auth

import "github.com/google/uuid"

type User struct {
	Id       string
	Username string
	Password string
	Role     string
}

var (
	Users = [3]User{
		{
			Id:       generateUUID(),
			Username: "admin",
			Password: "root",
			Role:     "MANAGER",
		},
		{
			Id:       generateUUID(),
			Username: "dev",
			Password: "developer",
			Role:     "DEVELOPER",
		},
		{
			Id:       generateUUID(),
			Username: "qa",
			Password: "tester",
			Role:     "TESTER",
		},
	}
)

func generateUUID() string {
	return uuid.New().String()[:8]
}
