package routes

import (
	"net/http"
	"tracka/pkg/config"

	"github.com/gorilla/mux"
)

func Init(router *mux.Router) {
	api := router.PathPrefix("/api").Subrouter()
	api.NotFoundHandler = http.HandlerFunc(notFound)
	registerAuth(api)
	registerOrders(api)

	spa := spaHandler{
		staticPath: config.Get().Static.StaticPath,
		indexFile:  config.Get().Static.IndexPath,
	}
	frontend := router.PathPrefix("/").Subrouter()
	frontend.HandleFunc("/", spa.ServeHTTP)
	frontend.NotFoundHandler = http.HandlerFunc(spa.ServeHTTP)
}
