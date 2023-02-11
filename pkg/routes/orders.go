package routes

import (
	"tracka/pkg/controllers"

	"github.com/gorilla/mux"
)

func registerOrders(api *mux.Router) {
	router := api.PathPrefix("/orders").Subrouter()

	router.HandleFunc("", controllers.OrdersList).Methods("GET")
	router.HandleFunc("", controllers.OrdersCreate).Methods("POST")
	router.HandleFunc("/{tracker}", controllers.OrderGet).Methods("GET")
	router.HandleFunc("/{tracker}", controllers.OrderStatusUpdate).Methods("PATCH")

	router.Use(authRequired)
}
