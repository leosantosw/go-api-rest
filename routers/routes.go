package routers

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/leosantosw/go-api-rest/controllers"
)

func Routers() {
	r := mux.NewRouter()
	r.HandleFunc("/users", controllers.CreateUser).Methods("POST")
	r.HandleFunc("/users", controllers.ListUsers).Methods("GET")
	r.HandleFunc("/users/{id}", controllers.FindUser).Methods("GET")
	r.HandleFunc("/users/{id}", controllers.DeleteUser).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", r))
}
