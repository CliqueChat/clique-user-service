package handlers

import (
	"encoding/json"
	"github.com/CliqueChat/clique-user-service/services"
	"github.com/CliqueChat/clique-user-service/structs"
	"github.com/gorilla/mux"
	"net/http"
)

func InitUserHandles(r *mux.Router) {
	r.HandleFunc("/user/profile", getUserProfile).Methods(http.MethodGet)
	r.HandleFunc("/user/create", createUser).Methods(http.MethodPost)

}

func createUser(w http.ResponseWriter, r *http.Request) {

	var user structs.User

	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = services.CreateANewUser(user)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//TODO Return success status

}

func getUserProfile(w http.ResponseWriter, r *http.Request) {

}
