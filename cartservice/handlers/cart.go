package handlers

import (
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
	"github.com/yhung-mea7/SEN300-micro/tree/main/cartservice/data"
)

type (
	CartHandler struct {
		logger *log.Logger
		repo   *data.CartRepo
	}

	generalError struct {
		Message string `json:"message"`
	}
	clientInformation struct {
		Email string `json:"email"`
	}
	keyValue struct{}
)

func NewCartHandler(l *log.Logger, r *data.CartRepo) *CartHandler {
	return &CartHandler{
		logger: l,
		repo:   r,
	}
}

func (ch *CartHandler) PostItemToCart() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		itemId := getItemId(r)
		resp, err := http.Get(strings.Join([]string{"http://itemapi:8080/", strconv.Itoa(itemId)}, ""))
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			data.ToJSON(&generalError{"Failed to establish connection to Item api"}, rw)
			return
		}
		defer resp.Body.Close()
		item := data.Item{}
		data.FromJSON(&item, resp.Body)
		qty := getURLqty(r)
		userInfo := r.Context().Value(keyValue{}).(clientInformation)

		if err = ch.repo.AddItem(userInfo.Email, &item, qty); err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			data.ToJSON(&generalError{err.Error()}, rw)
			return
		}
		rw.WriteHeader(http.StatusNoContent)

	}
}

func (ch *CartHandler) GetItemCart() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Set("Content-Type", "application/json")
		userInfo := r.Context().Value(keyValue{}).(clientInformation)
		d, err := ch.repo.GetCart(userInfo.Email)
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			data.ToJSON(&generalError{"Error receiving cart"}, rw)
			return
		}
		data.ToJSON(d, rw)
	}
}

func (ch *CartHandler) DeleteItem() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		id := getItemId(r)
		userInfo := r.Context().Value(keyValue{}).(clientInformation)

		if err := ch.repo.RemoveItem(userInfo.Email, uint(id)); err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			data.ToJSON(&generalError{err.Error()}, rw)
			return
		}
		rw.WriteHeader(http.StatusAccepted)
	}
}

func (ch *CartHandler) PatchItemQuantity() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		id := getItemId(r)
		qty := getURLqty(r)
		userInfo := r.Context().Value(keyValue{}).(clientInformation)

		if err := ch.repo.UpdateItemQuantity(userInfo.Email, uint(id), qty); err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			data.ToJSON(&generalError{err.Error()}, rw)
			return
		}
		rw.WriteHeader(http.StatusAccepted)
	}
}

func (ch *CartHandler) ClearCart() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		userInfo := r.Context().Value(keyValue{}).(clientInformation)

		if err := ch.repo.ClearCart(userInfo.Email); err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			data.ToJSON(&generalError{"Failed to clear shopping cart"}, rw)
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

func getURLqty(r *http.Request) int32 {
	qty, ok := r.URL.Query()["qty"]

	if !ok || len(qty[0]) < 1 {
		return 0
	}

	q, err := strconv.Atoi(qty[0])
	if err != nil {
		return 0
	}
	return int32(q)

}
