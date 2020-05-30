package main

import (
	"github.com/CliqueChat/clique-user-service/handlers"
	"github.com/CliqueChat/clique-user-service/helpers"
	"github.com/CliqueChat/clique-user-service/repositories"
	"github.com/CliqueChat/clique-user-service/resources"
	"github.com/CliqueChat/clique-user-service/services"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

var prop, profile = resources.GetApplicationProfile()

func main() {

	// Read host and port from property file
	host, _ := prop.Get(helpers.HOST)
	port, _ := prop.Get(helpers.PORT)

	// Connecting MongoDb
	repositories.Connect()

	// Connecting Email Service
	services.ConnectEmailServer()

	// Setting up routes
	r := setupRoutes()

	log.Println("APPLICATION STARTING IN " + profile + " PROFILE")
	log.Println("APPLICATION STARTED ON " + host + ":" + port)

	log.Fatal(http.ListenAndServe(host+":"+port, r))
}

func setupRoutes() *mux.Router {
	r := mux.NewRouter()
	handlers.InitUserHandles(r)
	return r
}
