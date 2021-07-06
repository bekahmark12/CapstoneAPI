package handlers

import (
	"log"
	"net/http"

	"github.com/yhung-mea7/SEN300-micro/tree/main/checkoutservice/data"
	"github.com/yhung-mea7/SEN300-micro/tree/main/checkoutservice/messaging"
)

type (
	Checkout struct {
		l      *log.Logger
		repo   *data.CheckoutRepo
		broker *messaging.Messager
	}

	generalError struct {
		Message string `json:"message"`
	}
	validationError struct {
		Message map[string]string
	}
)

func NewCheckOutHandler(l *log.Logger, cr *data.CheckoutRepo, broker *messaging.Messager) *Checkout {
	return &Checkout{l, cr, broker}
}

func (ch *Checkout) PostCheckout() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		ch.l.Println("POST CHECKOUT")
		checkout := r.Context().Value(keyvalue{}).(data.Checkout)
		resp, err := http.Get("http://cartapi:8080/")
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			data.ToJSON(&generalError{"Failed to establish connection to cart api"}, rw)
			return
		}
		defer resp.Body.Close()

		cart := data.Cart{}
		if err := data.FromJSON(&cart, resp.Body); err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			data.ToJSON(&generalError{err.Error()}, rw)
			return
		}

		if err := ch.repo.CheckoutOrder(&checkout, &cart); err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			data.ToJSON(&generalError{"Failed to submit order"}, rw)
			return
		}

		if err := ch.broker.SubmitToMessageBroker(&messaging.Message{checkout.Name, checkout.Email, "Your order has been submitted!!"}); err != nil {
			ch.l.Println("Failed to submit email")
		}
		rw.WriteHeader(http.StatusAccepted)
	}
}

func (ch *Checkout) GetPastCheckouts() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		ch.l.Panicln("GET CHECKOUT")
		orders, err := ch.repo.GetAllOrders()
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			data.ToJSON(&generalError{"Failed to receive orders"}, rw)
			return
		}
		data.ToJSON(orders, rw)
	}
}