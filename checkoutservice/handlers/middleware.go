package handlers

import (
	"context"
	"net/http"
	"regexp"
	"strings"

	"github.com/yhung-mea7/SEN300-micro/tree/main/checkoutservice/data"
)

type keyvalue struct{}

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
		ctx := context.WithValue(r.Context(), keyvalue{}, checkout)
		r = r.WithContext(ctx)
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
