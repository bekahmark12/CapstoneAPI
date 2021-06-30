package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/yhung-mea7/sen300-ex-1/models"
)

type (
	Item struct {
		logger *log.Logger
		repo   *models.ItemRepo
	}
	GeneralError struct {
		Message string `json:"message"`
	}
	ValidationError struct {
		Message map[string]string `json:"message"`
	}
	KeyValue struct{}
)

func NewItemHandler(l *log.Logger, r *models.ItemRepo) *Item {
	return &Item{
		logger: l,
		repo:   r,
	}

}

func (i *Item) PostItem() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		i.logger.Println("POST Item")
		item := r.Context().Value(KeyValue{}).(models.Item)
		err := i.repo.CreateItem(&item)
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			models.ToJSON(&GeneralError{err.Error()}, rw)
			return
		}
		rw.WriteHeader(http.StatusNoContent)
	}
}

func (i *Item) UpdateItem() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		i.logger.Println("PUT Item")
		id := getItemId(r)
		item := r.Context().Value(KeyValue{}).(models.Item)
		err := i.repo.UpdateItem(uint(id), &item)
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			models.ToJSON(&GeneralError{err.Error()}, rw)
			return
		}

		if item.ID == 0 {
			rw.WriteHeader(http.StatusNotFound)
			models.ToJSON(&GeneralError{"Item not found"}, rw)
			return
		}

		item.ID = uint(id)
		rw.WriteHeader(http.StatusAccepted)
		models.ToJSON(item, rw)

	}
}

func (i *Item) GetItemById() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		i.logger.Println("GET Item BY ID")
		rw.Header().Add("Content-type", "application/json")
		id := getItemId(r)
		item := i.repo.GetItemById(uint(id))
		if item.ID == 0 {
			rw.WriteHeader(http.StatusNotFound)
			models.ToJSON(&GeneralError{"Item not found"}, rw)
			return
		}
		models.ToJSON(item, rw)
	}
}

func (i *Item) GetAllItems() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		i.logger.Println("GET ALL Items")
		rw.Header().Add("Content-type", "application/json")
		items := i.repo.GetAllItems()
		models.ToJSON(items, rw)
	}
}

func (i *Item) DeleteItem() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		i.logger.Println("DELETE Item")
		id := getItemId(r)
		err := i.repo.DeleteItem(uint(id))

		if err != nil {
			rw.WriteHeader(http.StatusNotFound)
			models.ToJSON(&GeneralError{err.Error()}, rw)
			return
		}
		rw.WriteHeader(http.StatusAccepted)
	}
}

func getItemId(r *http.Request) int {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		panic(err) //do something better than panic
	}
	return id
}
