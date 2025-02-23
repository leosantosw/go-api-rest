package controllers

import (
	"encoding/json"
	"net/http"

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
