package routes

import (
	"net/http"
	"tracka/pkg/controllers"

	"github.com/gorilla/mux"
)

func registerAuth(api *mux.Router) {
	router := api.PathPrefix("/auth").Subrouter()

	router.HandleFunc("/signin", controllers.AuthSignin).Methods("POST", "OPTIONS")
	router.HandleFunc("/reset-password", controllers.AuthResetPassword).Methods("POST", "OPTIONS")

	router.HandleFunc("/logout", authRequired(http.HandlerFunc(controllers.AuthLogout)).ServeHTTP).Methods("POST", "OPTIONS")
}
