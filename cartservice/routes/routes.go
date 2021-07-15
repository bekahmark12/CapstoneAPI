package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/yhung-mea7/SEN300-micro/tree/main/cartservice/handlers"
)

func SetUpRoutes(sm *mux.Router, cartHandler *handlers.CartHandler) {

	postHandler := sm.Methods(http.MethodPost).Subrouter()
	postHandler.HandleFunc("/{id:[0-9]+}", cartHandler.PostItemToCart())
	postHandler.Use(cartHandler.Auth)

	getHandler := sm.Methods(http.MethodGet).Subrouter()
	getHandler.HandleFunc("/", cartHandler.GetItemCart())
	getHandler.Use(cartHandler.Auth)

	deleteHandler := sm.Methods(http.MethodDelete).Subrouter()
	deleteHandler.HandleFunc("/{id:[0-9]+}", cartHandler.DeleteItem())
	deleteHandler.HandleFunc("/", cartHandler.ClearCart())
	deleteHandler.Use(cartHandler.Auth)

	patchHandler := sm.Methods(http.MethodPatch).Subrouter()
	patchHandler.HandleFunc("/{id:[0-9]+}", cartHandler.PatchItemQuantity())
	patchHandler.Use(cartHandler.Auth)

	healthHandler := sm.Methods(http.MethodGet).Subrouter()
	healthHandler.HandleFunc("/healthcheck", cartHandler.HealthCheck())

}
