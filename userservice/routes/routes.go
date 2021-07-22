package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/yhung-mea7/SEN300-micro/tree/main/userservice/handlers"
)

func SetUpRoutes(sm *mux.Router, userHandler *handlers.UserHandler) {

	loginRouter := sm.Methods(http.MethodPost).Subrouter()
	loginRouter.HandleFunc("/", userHandler.Login())
	loginRouter.Use(userHandler.MiddlewareValidateLogin)

	signUpRouter := sm.Methods(http.MethodPost).Subrouter()
	signUpRouter.HandleFunc("/sign-up", userHandler.CreateUser())
	signUpRouter.Use(userHandler.MiddlewareValidateUser)

	checkUser := sm.Methods(http.MethodGet).Subrouter()
	checkUser.HandleFunc("/", userHandler.GetLoggedInUser())
	checkUser.Use(userHandler.Auth)

	healthHandler := sm.Methods(http.MethodGet).Subrouter()
	healthHandler.HandleFunc("/healthcheck", userHandler.HealthCheck())

}
