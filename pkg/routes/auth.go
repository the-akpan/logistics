package routes

import (
	"tracka/pkg/controllers"

	"github.com/gorilla/mux"
)

func registerAuth(api *mux.Router) {
	router := api.PathPrefix("/auth").Subrouter()
	router.HandleFunc("/signin", controllers.AuthSignin).Methods("POST")
}
