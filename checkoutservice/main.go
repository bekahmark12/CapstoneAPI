package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	gohandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/yhung-mea7/SEN300-micro/tree/main/checkoutservice/data"
	"github.com/yhung-mea7/SEN300-micro/tree/main/checkoutservice/handlers"
	"github.com/yhung-mea7/SEN300-micro/tree/main/checkoutservice/messaging"
)

func main() {
	sm := mux.NewRouter()
	logger := log.New(os.Stdout, "cart-service", log.LstdFlags)
	ch := gohandlers.CORS(gohandlers.AllowedOrigins([]string{"*"}))
	checkoutHandler := handlers.NewCheckOutHandler(logger, data.NewCheckoutRepo(os.Getenv("MONGO_URI"), os.Getenv("MONGO_DB")), messaging.NewMessager(os.Getenv("RABBIT_CONN")))

	postHandler := sm.Methods(http.MethodPost).Subrouter()
	postHandler.Handle("/", checkoutHandler.PostCheckout())
	postHandler.Use(checkoutHandler.MiddlewareValidateCheckout)

	server := http.Server{
		Addr:         os.Getenv("PORT"),
		Handler:      ch(sm),
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
