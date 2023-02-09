package main

import (
	"log"
	"net/http"
	"time"
	"tracka/pkg/config"
	"tracka/pkg/models"
	"tracka/pkg/routes"

	"github.com/gorilla/mux"
)

func main() {
	// setup project config
	config.Init()

	// initialize models
	models.Init()

	router := mux.NewRouter()
	routes.Init(router)

	server := &http.Server{
		Handler:      router,
		Addr:         config.Get().PORT,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Println("Start server @", config.Get().PORT)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
