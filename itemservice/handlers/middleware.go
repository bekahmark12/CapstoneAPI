package handlers

import (
	"context"
	"net/http"
	"regexp"
	"strings"

	"github.com/yhung-mea7/sen300-ex-1/models"
)

func (i *Item) MiddlewareValidateItem(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Add("Content-Type", "application/json")
		item := models.Item{}
		err := models.FromJSON(&item, r.Body)
		if err != nil {
			i.logger.Println("[ERROR] deserializing item", err)
			rw.WriteHeader(http.StatusBadRequest)
			models.ToJSON(&GeneralError{Message: err.Error()}, rw)
			return
		}
		i.logger.Println(item)
		err = item.Validate()
		if err != nil {
			i.logger.Println("[ERROR] validating item", err)
			rw.WriteHeader(http.StatusBadRequest)
			models.ToJSON(&ValidationError{Message: formatValidationError(err.Error())}, rw)
			return
		}
		ctx := context.WithValue(r.Context(), KeyValue{}, item)
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
