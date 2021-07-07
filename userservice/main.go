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
	"github.com/yhung-mea7/SEN300-micro/tree/main/userservice/data"
	"github.com/yhung-mea7/SEN300-micro/tree/main/userservice/handlers"
)

func main() {
	sm := mux.NewRouter()
	logger := log.New(os.Stdout, "user-service", log.LstdFlags)
	ch := gohandlers.CORS(gohandlers.AllowedOrigins([]string{"*"}))
	userHandler := handlers.NewUserHandler(data.NewUserRepo(os.Getenv("DSN")), os.Getenv("SECRET_KEY"), logger)

	loginRouter := sm.Methods(http.MethodPost).Subrouter()
	loginRouter.HandleFunc("/login", userHandler.Login())
	loginRouter.Use(userHandler.MiddlewareValidateLogin)

	signUpRouter := sm.Methods(http.MethodPost).Subrouter()
	signUpRouter.HandleFunc("/sign-up", userHandler.CreateUser())
	signUpRouter.Use(userHandler.MiddlewareValidateUser)

	checkUser := sm.Methods(http.MethodGet).Subrouter()
	checkUser.HandleFunc("/", userHandler.GetLoggedInUser())
	checkUser.Use(userHandler.Auth)

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
