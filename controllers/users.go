package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/leosantosw/go-api-rest/models"
)

func Users(w http.ResponseWriter, r *http.Request) {
	users := models.ListUsers()
	json.NewEncoder(w).Encode(users)
}
