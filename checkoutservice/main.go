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
	"github.com/yhung-mea7/SEN300-micro/tree/main/checkoutservice/data"
	"github.com/yhung-mea7/SEN300-micro/tree/main/checkoutservice/handlers"
	"github.com/yhung-mea7/SEN300-micro/tree/main/checkoutservice/messaging"
	"github.com/yhung-mea7/SEN300-micro/tree/main/checkoutservice/register"
	"github.com/yhung-mea7/SEN300-micro/tree/main/checkoutservice/routes"
)

func main() {
	sm := mux.NewRouter()
	logger := log.New(os.Stdout, "order-service", log.LstdFlags)

	consulClient := register.NewConsulClient("order-service")
	checkoutHandler := handlers.NewCheckOutHandler(logger, data.NewCheckoutRepo(os.Getenv("MONGO_URI"), os.Getenv("MONGO_DB")), messaging.NewMessager(os.Getenv("RABBIT_CONN")), consulClient)
	consulClient.RegisterService()
	routes.SetUpRoutes(sm, checkoutHandler)

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
	logger.Println("Got Signal:", sig)
	if err := consulClient.DeregisterService(); err != nil {
		logger.Println(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	server.Shutdown(ctx)
}
