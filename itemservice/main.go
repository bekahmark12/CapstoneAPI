package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	// gohandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/yhung-mea7/sen300-ex-1/handlers"
	"github.com/yhung-mea7/sen300-ex-1/models"
	"github.com/yhung-mea7/sen300-ex-1/routes"
)

func main() {
	sm := mux.NewRouter()
	logger := log.New(os.Stdout, "item-service", log.LstdFlags)
	// ch := gohandlers.CORS(gohandlers.AllowedOrigins([]string{"*"}))

	itemHandler := handlers.NewItemHandler(logger, models.NewItemRepo(os.Getenv("DSN")))

	routes.SetUpRoutes(sm, itemHandler)

	server := http.Server{
		Addr:         os.Getenv("PORT"),
		Handler:      sm,
		ErrorLog:     logger,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	go func() {
		logger.Printf("Starting server on port: %v \n", server.Addr)
		err := server.ListenAndServe()
		if err != nil {
			logger.Printf("Error starting server: %v \n", err)
			os.Exit(1)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)
	sig := <-c
	logger.Println("Got Signal:", sig)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	server.Shutdown(ctx)
}
