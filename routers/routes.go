package routers

import (
	"net/http"

	"github.com/leosantosw/go-api-rest/controllers"
)

func Routers() {
	http.HandleFunc("/users", controllers.Users)
}
