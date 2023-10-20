package main

import (
	"customerapp/controller"
	"customerapp/dbrepos"
	"customerapp/router"
	"go.uber.org/zap"
	"log"
	"net/http"
)

func main() {
	logger, _ := zap.NewProduction()             // Create Uber's Zap logger
	defer logger.Sync()                          // flushes buffer, if any
	repo, err := dbrepos.NewInmemoryRepository() // With in-memory database
	// repo, err := dbrepos.NewMongoDBRepository() // With MongoDB database
	if err != nil {
		log.Fatal("Error:", err)
	}
	h := &controller.CustomerController{
		Repository: repo, // Injecting dependency
		Logger:     logger,
	}

	server := &http.Server{
		Addr:    ":8080",
		Handler: router.InitializeRoutes(h), // configure routes
	}
	log.Println("Listening...")
	if err := server.ListenAndServe(); err != nil { // Run the http server
		log.Fatal("Error:", err)
	}
}
