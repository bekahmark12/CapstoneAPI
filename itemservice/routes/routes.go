package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/yhung-mea7/sen300-ex-1/handlers"
)

func SetUpRoutes(sm *mux.Router, itemHandler *handlers.ItemHandler) {
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
}
