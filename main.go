package main

import (
	"fmt"
	"net/http"

	"github.com/leosantosw/go-api-rest/database"
	"github.com/leosantosw/go-api-rest/routers"
)

func HandleRequests() {
	routers.Routers()
	http.ListenAndServe(":8080", nil)
}

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to the home page!")
}

func main() {
	database.Connect()
	HandleRequests()
}
