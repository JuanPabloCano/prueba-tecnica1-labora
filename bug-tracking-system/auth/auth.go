package auth

import "prueba-tecnica1-labora/bug-tracking-system/models"

func GetUserByUsernameAndPassword(username string, password string) models.User {
	for idx, user := range models.Users {
		if user.Username == username && user.Password == password {
			return models.Users[idx]
		}
	}
	return models.User{}
}
