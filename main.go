package main

import (
	"fmt"
	"net/http"
)

func HandleRequests() {
	http.HandleFunc("/", Home)
	http.ListenAndServe(":8080", nil)
}

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to the home page!")
}

func main() {
	fmt.Println("Hello, World!")
	HandleRequests()
}
