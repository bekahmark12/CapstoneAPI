package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/yhung-mea7/SEN300-micro/tree/main/checkoutservice/handlers"
)

func SetUpRoutes(sm *mux.Router, checkoutHandler *handlers.Checkout) {

	postHandler := sm.Methods(http.MethodPost).Subrouter()
	postHandler.Handle("/", checkoutHandler.PostCheckout())
	postHandler.Use(checkoutHandler.MiddlewareValidateCheckout)
	postHandler.Use(checkoutHandler.Auth)

	healthHandler := sm.Methods(http.MethodGet).Subrouter()
	healthHandler.Handle("/healthcheck", checkoutHandler.HealthCheck())

	getHandler := sm.Methods(http.MethodGet).Subrouter()
	getHandler.Handle("/", checkoutHandler.GetPastCheckouts())
	getHandler.Use(checkoutHandler.Auth)
}
