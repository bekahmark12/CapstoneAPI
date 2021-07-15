package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-redis/redis"
	"github.com/gorilla/mux"
	"github.com/yhung-mea7/SEN300-micro/tree/main/cartservice/data"
	"github.com/yhung-mea7/SEN300-micro/tree/main/cartservice/handlers"
	register "github.com/yhung-mea7/SEN300-micro/tree/main/cartservice/registry"
	"github.com/yhung-mea7/SEN300-micro/tree/main/cartservice/routes"
)

func main() {
	redisCli := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT")),
		Password: "",
		DB:       0,
	})
	sm := mux.NewRouter()
	logger := log.New(os.Stdout, "cart-service", log.LstdFlags)

	consulClient := register.NewConsulClient("cart-service")
	consulClient.RegisterService()

	cartHandler := handlers.NewCartHandler(logger, data.NewCartRepo(redisCli), consulClient)
	routes.SetUpRoutes(sm, cartHandler)

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
