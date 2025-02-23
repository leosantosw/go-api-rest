package models

import "github.com/leosantosw/go-api-rest/database"

type User struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Password  int    `json:"password"`
	CreatedAt string `json:"created_at"`
}

func ListUsers() []User {
	users := []User{}
	database.DB.Find(&users)
	return users
}
