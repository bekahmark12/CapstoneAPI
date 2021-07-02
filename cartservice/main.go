package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/go-redis/redis"
	gohandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/yhung-mea7/SEN300-micro/tree/main/cartservice/data"
	"github.com/yhung-mea7/SEN300-micro/tree/main/cartservice/handlers"
)

func main() {
	redisCli := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT")),
		Password: "",
		DB:       0,
	})
	sm := mux.NewRouter()
	logger := log.New(os.Stdout, "cart-service", log.LstdFlags)
	ch := gohandlers.CORS(gohandlers.AllowedOrigins([]string{"*"}))
	cartHandler := handlers.NewCartHandler(logger, data.NewCartRepo(redisCli))

	postHandler := sm.Methods(http.MethodPost).Subrouter()
	postHandler.HandleFunc("/{id:[0-9]+}", cartHandler.PostItemToCart())

	getHandler := sm.Methods(http.MethodGet).Subrouter()
	getHandler.HandleFunc("/", cartHandler.GetItemCart())

	deleteHandler := sm.Methods(http.MethodDelete).Subrouter()
	deleteHandler.HandleFunc("/{id:[0-9]+}", cartHandler.DeleteItem())
	deleteHandler.HandleFunc("/", cartHandler.ClearCart())

	patchHandler := sm.Methods(http.MethodPatch).Subrouter()
	patchHandler.HandleFunc("/{id:[0-9]+}", cartHandler.PatchItemQuantity())

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
