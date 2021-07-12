package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/yhung-mea7/SEN300-micro/tree/main/checkoutservice/handlers"
)

func SetUpRoutes(sm *mux.Router, checkoutHandler *handlers.Checkout) {
	sm.Use(checkoutHandler.Auth)
	postHandler := sm.Methods(http.MethodPost).Subrouter()
	postHandler.Handle("/", checkoutHandler.PostCheckout())
	postHandler.Use(checkoutHandler.MiddlewareValidateCheckout)
}
