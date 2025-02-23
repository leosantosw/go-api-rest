package main

import (
	"net/http"

	"github.com/leosantosw/go-api-rest/database"
	"github.com/leosantosw/go-api-rest/routers"
)

func HandleRequests() {
	routers.Routers()
	http.ListenAndServe(":8080", nil)
}

func main() {
	database.Connect()
	HandleRequests()
}
