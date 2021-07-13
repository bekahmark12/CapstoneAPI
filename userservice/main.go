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
	"github.com/yhung-mea7/SEN300-micro/tree/main/userservice/data"
	"github.com/yhung-mea7/SEN300-micro/tree/main/userservice/handlers"
	"github.com/yhung-mea7/SEN300-micro/tree/main/userservice/routes"
)

func main() {
	sm := mux.NewRouter()
	logger := log.New(os.Stdout, "user-service", log.LstdFlags)
	// ch := gohandlers.CORS(gohandlers.AllowedOrigins([]string{"*"}))
	// ch := cors.New(cors.Options{
	// 	AllowedOrigins:   []string{"*"},
	// 	AllowCredentials: true,
	// 	AllowedHeaders:   []string{"*"},
	// 	AllowedMethods:   []string{"GET", "POST", "DELETE", "OPTIONS"},
	// })
	userHandler := handlers.NewUserHandler(data.NewUserRepo(os.Getenv("DSN")), os.Getenv("SECRET_KEY"), logger)

	routes.SetUpRoutes(sm, userHandler)

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
