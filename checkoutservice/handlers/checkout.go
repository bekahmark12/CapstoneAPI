package handlers

import (
	"log"
	"net/http"

	"github.com/yhung-mea7/SEN300-micro/tree/main/checkoutservice/data"
	"github.com/yhung-mea7/SEN300-micro/tree/main/checkoutservice/messaging"
	register "github.com/yhung-mea7/SEN300-micro/tree/main/checkoutservice/register"
)

type (
	Checkout struct {
		l      *log.Logger
		repo   *data.CheckoutRepo
		broker *messaging.Messager
		reg    *register.ConsulClient
	}

	generalError struct {
		Message string `json:"message"`
	}
	validationError struct {
		Message map[string]string
	}
	clientInformation struct {
		UserType int    `json:"user_type"`
		Email    string `json:"email"`
	}
	keyValue struct{}
)

func NewCheckOutHandler(l *log.Logger, cr *data.CheckoutRepo, broker *messaging.Messager, reg *register.ConsulClient) *Checkout {
	return &Checkout{l, cr, broker, reg}
}

func (ch *Checkout) PostCheckout() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		ch.l.Println("POST CHECKOUT")
		checkout := r.Context().Value(keyValue{}).(data.Checkout)
		service, err := ch.reg.LookUpService("cart-service")
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			data.ToJSON(&generalError{err.Error()}, rw)
			return
		}

		req, err := http.NewRequest("GET", service.GetHTTP(), nil)
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			data.ToJSON(&generalError{err.Error()}, rw)
			return
		}
		req.Header.Add("Authorization", r.Header.Get("Authorization"))
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			data.ToJSON(&generalError{err.Error()}, rw)
			return
		}
		defer resp.Body.Close()

		cart := data.Cart{}
		if err := data.FromJSON(&cart, resp.Body); err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			data.ToJSON(&generalError{err.Error()}, rw)
			return
		}
		usrSerr, _ := ch.reg.LookUpService("users-service")
		req, err = http.NewRequest("GET", usrSerr.GetHTTP(), nil)
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			data.ToJSON(&generalError{err.Error()}, rw)
			return
		}

		req.Header.Set("Authorization", r.Header.Get("Authorization"))
		resp, err = client.Do(req)
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			data.ToJSON(&generalError{err.Error()}, rw)
			return
		}

		defer resp.Body.Close()
		clientInfo := clientInformation{}
		if err := data.FromJSON(&clientInfo, resp.Body); err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			data.ToJSON(&generalError{err.Error()}, rw)
			return
		}
		if err := ch.repo.CheckoutOrder(clientInfo.Email, &checkout, &cart); err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			data.ToJSON(&generalError{"Failed to submit order"}, rw)
			return
		}
		req, err = http.NewRequest("DELETE", service.GetHTTP(), nil)
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			data.ToJSON(&generalError{err.Error()}, rw)
			return
		}
		req.Header.Add("Authorization", r.Header.Get("Authorization"))
		_, err = client.Do(req)
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			data.ToJSON(&generalError{err.Error()}, rw)
			return
		}
		if err := ch.broker.SubmitToMessageBroker(&messaging.Message{
			Name:    checkout.Name,
			Email:   clientInfo.Email,
			Content: "Your order has been submitted!!",
		}); err != nil {
			ch.l.Println("Failed to submit email")
		}
		rw.WriteHeader(http.StatusAccepted)
	}
}

func (ch *Checkout) HealthCheck() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		data.ToJSON(&generalError{"service good to go"}, rw)
	}
}

func (ch *Checkout) GetPastCheckouts() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		ch.l.Println("GET CHECKOUT")
		rw.Header().Set("Content-type", "application/json")
		// userInfo := r.Context().Value(keyValue{}).(clientInformation)
		userInfo := clientInformation{}
		usrSerr, err := ch.reg.LookUpService("users-service")
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			data.ToJSON(&generalError{err.Error()}, rw)
			return
		}
		req, _ := http.NewRequest("GET", usrSerr.GetHTTP(), nil)
		req.Header.Set("Authorization", r.Header.Get("Authorization"))
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			data.ToJSON(&generalError{err.Error()}, rw)
			return
		}
		defer resp.Body.Close()
		data.FromJSON(&userInfo, resp.Body)

		if userInfo.UserType != 1 {
			rw.WriteHeader(http.StatusForbidden)
			data.ToJSON(&generalError{"You are not authorized to view these orders"}, rw)
			return
		}
		orders, err := ch.repo.GetAllOrders()
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			data.ToJSON(&generalError{"Failed to receive orders"}, rw)
			return
		}
		data.ToJSON(orders, rw)
	}
}
