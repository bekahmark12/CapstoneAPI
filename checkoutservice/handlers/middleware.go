package handlers

import (
	"context"
	"net/http"
	"regexp"
	"strings"

	"github.com/yhung-mea7/SEN300-micro/tree/main/checkoutservice/data"
)

func (c *Checkout) MiddlewareValidateCheckout(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Set("Content-Type", "application/json")
		checkout := data.Checkout{}
		if err := data.FromJSON(&checkout, r.Body); err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			data.ToJSON(&generalError{err.Error()}, rw)
			return
		}
		if err := checkout.Validate(); err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			data.ToJSON(&validationError{formatValidationError(err.Error())}, rw)
			return
		}
		ctx := context.WithValue(r.Context(), keyValue{}, checkout)
		r = r.WithContext(ctx)
		next.ServeHTTP(rw, r)
	})
}

func (ch *Checkout) Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if token == "" {
			rw.WriteHeader(http.StatusForbidden)
			data.ToJSON(&generalError{"No token provided"}, rw)
			return
		}
		serr, err := ch.reg.LookUpService("users-service")
		req, err := http.NewRequest("GET", serr.GetHTTP(), nil)
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			data.ToJSON(&generalError{err.Error()}, rw)
			return
		}
		req.Header.Add("Authorization", token)
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			data.ToJSON(&generalError{err.Error()}, rw)
			return
		}
		defer resp.Body.Close()
		if resp.StatusCode != http.StatusOK {
			rw.WriteHeader(http.StatusUnauthorized)
			data.ToJSON(&generalError{"You are not authorized to make this request"}, rw)
			return
		}
		userInfo := clientInformation{}
		if err := data.FromJSON(&userInfo, resp.Body); err != nil {
			rw.WriteHeader(http.StatusUnauthorized)
			data.ToJSON(&generalError{"You are not authorized to make this request"}, rw)
			return
		}

		// ctx := context.WithValue(r.Context(), keyValue{}, userInfo)
		// r = r.WithContext(ctx)
		next.ServeHTTP(rw, r)
	})
}

func formatValidationError(err string) map[string]string {
	messages := strings.Split(err, "\n")
	output := map[string]string{}
	rgx := regexp.MustCompile(`^Key: (.*) Error:(.*)$`)
	for _, line := range messages {
		m := rgx.FindStringSubmatch(line)
		output[m[1]] = m[2]

	}
	return output
}
