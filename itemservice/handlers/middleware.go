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

// func (i *Item) Auth(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
// 		token := r.Header.Get("Authorization")
// 		if token == "" {
// 			rw.WriteHeader(http.StatusForbidden)
// 			models.ToJSON(&GeneralError{"No token provided"}, rw)
// 			return
// 		}
// 		req, err := http.NewRequest("GET", "http://userapi:8080/", nil)
// 		if err != nil {
// 			rw.WriteHeader(http.StatusInternalServerError)
// 			models.ToJSON(&GeneralError{err.Error()}, rw)
// 			return
// 		}
// 		req.Header.Add("Authorization", token)
// 		client := &http.Client{}
// 		resp, err := client.Do(req)
// 		if err != nil {
// 			rw.WriteHeader(http.StatusInternalServerError)
// 			models.ToJSON(&GeneralError{err.Error()}, rw)
// 			return
// 		}
// 		defer resp.Body.Close()
// 		i.logger.Println(resp.Body)
// 		if resp.StatusCode != http.StatusOK {
// 			rw.WriteHeader(http.StatusUnauthorized)
// 			models.ToJSON(&GeneralError{"You are not authorized to make this request"}, rw)
// 			return
// 		}
// 		next.ServeHTTP(rw, r)
// 	})
// }
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
