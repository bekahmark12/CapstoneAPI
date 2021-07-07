package handlers_test

import (
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/yhung-mea7/SEN300-micro/tree/main/userservice/handlers"
)

func TestValidateUser(t *testing.T) {
	nextHttpHandler := http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {

	})
	logger := log.New(os.Stdout, "test-service", log.LstdFlags)

	testHandler := handlers.NewUserHandler(nil, "", logger)
	hand := testHandler.MiddlewareValidateUser(nextHttpHandler)
	req := httptest.NewRequest("POST", "http://testing", nil)
	hand.ServeHTTP(httptest.NewRecorder(), req)
}
