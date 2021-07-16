package handlers

import (
	"context"
	"net/http"

	"github.com/yhung-mea7/SEN300-micro/tree/main/cartservice/data"
)

func (ch *CartHandler) Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Add("Content-Type", "application/json")
		token := r.Header.Get("Authorization")
		if token == "" {
			rw.WriteHeader(http.StatusForbidden)
			data.ToJSON(&generalError{"No token provided"}, rw)
			return
		}
		serr, err := ch.reg.LookUpService("users-service")
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			data.ToJSON(&generalError{err.Error()}, rw)
			return
		}
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
			rw.WriteHeader(http.StatusInternalServerError)
			data.ToJSON(&generalError{err.Error()}, rw)
			return
		}
		ctx := context.WithValue(r.Context(), keyValue{}, userInfo)
		r = r.WithContext(ctx)

		next.ServeHTTP(rw, r)
	})
}
