package main

import (
	"context"
	"github.com/NewChakrit/golang_web_development/config"
	"github.com/NewChakrit/golang_web_development/db"
	"github.com/NewChakrit/golang_web_development/routes"
	"log"
	"net/http"
)

func main() {
	handler := routes.MounteRoutes()

	db.InitDB()

	server := &http.Server{
		Addr:    config.Config.AppPort,
		Handler: handler,
	}

	config.Config.LoadConfig()
	println("Server running at", config.Config.AppPort)
	defer db.DB.Close(context.Background())
	err := server.ListenAndServe()
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
