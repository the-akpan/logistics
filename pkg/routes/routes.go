package routes

import (
	"net/http"
	"tracka/pkg/config"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func Init(router *mux.Router) http.Handler {
	c := cors.New(cors.Options{
		AllowCredentials: true,
		AllowedHeaders:   []string{"Accept", "Accept-Encoding", "Content-Type", "X-Auth-Token", "X-CSRF-Token"},
		AllowedMethods:   []string{"GET", "OPTIONS", "POST", "PUT"},
		AllowedOrigins:   []string{config.Get().CorsOrigin},
	})

	api := router.PathPrefix("/api").Subrouter()
	api.NotFoundHandler = http.HandlerFunc(notFound)
	registerAuth(api)
	registerOrders(api)

	spa := spaHandler{
		staticPath: config.Get().Static.StaticPath,
		indexFile:  config.Get().Static.IndexPath,
	}
	frontend := router.PathPrefix("/").Subrouter()
	frontend.HandleFunc("", spa.ServeHTTP)
	frontend.NotFoundHandler = http.HandlerFunc(spa.ServeHTTP)

	return c.Handler(router)
}
