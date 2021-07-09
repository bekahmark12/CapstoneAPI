package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/yhung-mea7/SEN300-micro/tree/main/cartservice/handlers"
)

func SetUpRoutes(sm *mux.Router, cartHandler *handlers.CartHandler) {
	sm.Use(cartHandler.Auth)

	postHandler := sm.Methods(http.MethodPost).Subrouter()
	postHandler.HandleFunc("/{id:[0-9]+}", cartHandler.PostItemToCart())

	getHandler := sm.Methods(http.MethodGet).Subrouter()
	getHandler.HandleFunc("/", cartHandler.GetItemCart())

	deleteHandler := sm.Methods(http.MethodDelete).Subrouter()
	deleteHandler.HandleFunc("/{id:[0-9]+}", cartHandler.DeleteItem())
	deleteHandler.HandleFunc("/", cartHandler.ClearCart())

	patchHandler := sm.Methods(http.MethodPatch).Subrouter()
	patchHandler.HandleFunc("/{id:[0-9]+}", cartHandler.PatchItemQuantity())
}
