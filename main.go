package main

import (
	"fmt"
	"net/http"

	"main/config"
	"main/controller"

	"github.com/gorilla/mux"
)

func main() {
	//Init Router
	r := mux.NewRouter()

	controller.UserController(r)

	configuration := config.GetConfiguration()
	http.ListenAndServe(fmt.Sprintf(":%s", configuration.Port), r)

}
