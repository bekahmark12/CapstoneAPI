package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	"github.com/bekahmark12/CapstoneAPI/tree/main/userservice/data"
	"github.com/bekahmark12/CapstoneAPI/tree/main/userservice/handlers"
	"github.com/bekahmark12/CapstoneAPI/tree/main/userservice/register"
	"github.com/bekahmark12/CapstoneAPI/tree/main/userservice/routes"
)

func main() {
	sm := mux.NewRouter()
	logger := log.New(os.Stdout, "users-service", log.LstdFlags)
	consulClient := register.NewConsulClient("users-service")
	consulClient.RegisterService()
	userHandler := handlers.NewUserHandler(data.NewUserRepo(os.Getenv("DSN")), os.Getenv("SECRET_KEY"), logger, consulClient)

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
	signal.Notify(c, syscall.SIGTERM)
	sig := <-c
	if err := consulClient.DeregisterService(); err != nil {
		logger.Println(err)
	}
	logger.Println("Got Signal:", sig)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	server.Shutdown(ctx)
}
