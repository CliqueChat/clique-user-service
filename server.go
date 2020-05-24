package main

import (
	"github.com/CliqueChat/clique-user-service/handlers"
	"github.com/CliqueChat/clique-user-service/helpers"
	"github.com/CliqueChat/clique-user-service/resources"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

var prop = resources.GetApplicationProfile()

func main() {

	// Read host and port from property file
	host, _ := prop.Get(helpers.HOST)
	port, _ := prop.Get(helpers.PORT)

	// Setting up routes
	r := setupRoutes()

	log.Println("APPLICATION STARTED ON " + host + ":" + port)
	log.Fatal(http.ListenAndServe(host+":"+port, r))
}

func setupRoutes() *mux.Router {
	r := mux.NewRouter()
	handlers.InitUserHandles(r)
	return r
}
