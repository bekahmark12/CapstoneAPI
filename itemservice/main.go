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
	"github.com/yhung-mea7/sen300-ex-1/handlers"
	"github.com/yhung-mea7/sen300-ex-1/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	sm := mux.NewRouter()
	logger := log.New(os.Stdout, "item-service", log.LstdFlags)
	ch := gohandlers.CORS(gohandlers.AllowedOrigins([]string{"*"}))

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  os.Getenv("DSN"),
		PreferSimpleProtocol: true,
	}), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	itemHandler := handlers.NewItemHandler(logger, models.NewItemRepo(db))

	postRouter := sm.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/", itemHandler.PostItem())
	postRouter.Use(itemHandler.MiddlewareValidateItem)

	putRouter := sm.Methods(http.MethodPut).Subrouter()
	putRouter.HandleFunc("/{id:[0-9]+}", itemHandler.UpdateItem())
	putRouter.Use(itemHandler.MiddlewareValidateItem)

	getRouter := sm.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/", itemHandler.GetAllItems())
	getRouter.HandleFunc("/{id:[0-9]+}", itemHandler.GetItemById())

	deleteRouter := sm.Methods(http.MethodDelete).Subrouter()
	deleteRouter.HandleFunc("/{id:[0-9]+}", itemHandler.DeleteItem())

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
