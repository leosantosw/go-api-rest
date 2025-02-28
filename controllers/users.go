package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/leosantosw/go-api-rest/models"
)

func ListUsers(w http.ResponseWriter, r *http.Request) {
	users := models.ListUsers()
	json.NewEncoder(w).Encode(users)
}

func FindUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	user := models.FindUser(id)
	json.NewEncoder(w).Encode(user)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)
	user.Id = uuid.New()
	models.CreateUser(user)
	json.NewEncoder(w).Encode(user)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	models.DeleteUser(id)
}

func EditUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	userUpdated := models.EditUser(id, r.Body)
	json.NewEncoder(w).Encode(userUpdated)
}
