package models

import (
	"github.com/google/uuid"
	"github.com/leosantosw/go-api-rest/database"
)

type User struct {
	Id        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt string    `json:"created_at"`
}

func ListUsers() []User {
	users := []User{}
	database.DB.Find(&users)
	return users
}

func FindUser(id string) User {
	user := User{}
	database.DB.First(&user, id)
	return user
}

func CreateUser(user User) {
	database.DB.Create(&user)
}

func DeleteUser(id string) {
	user := User{}
	database.DB.Delete(&user, &id)
}
