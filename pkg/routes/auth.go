package routes

import (
	"log"
	"net/http"
	"tracka/pkg/controllers"

	"github.com/gorilla/mux"
)

func authRequired(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		log.Println("Authentication Required")
		next.ServeHTTP(res, req)
	})
}

func registerAuth(api *mux.Router) {
	router := api.PathPrefix("/auth").Subrouter()
	router.StrictSlash(false)

	router.HandleFunc("/signin", controllers.AuthSignin).Methods("POST")
	router.HandleFunc("/reset-password", controllers.AuthResetPassword).Methods("POST")

	router.HandleFunc("/logout", authRequired(controllers.AuthLogout)).Methods("POST")
}
