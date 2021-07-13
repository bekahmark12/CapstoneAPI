package handlers

import (
	"log"
	"net/http"

	"github.com/yhung-mea7/SEN300-micro/tree/main/userservice/auth"
	"github.com/yhung-mea7/SEN300-micro/tree/main/userservice/data"
	"gorm.io/gorm"
)

type (
	UserHandler struct {
		repo *data.UserRepo
		jwt  *auth.JwtWrapper
		log  *log.Logger
	}
	generalError struct {
		Message string `json:"message"`
	}
	validationError struct {
		Message map[string]string `json:"message"`
	}
	clientInformation struct {
		Email string `json:"email"`
	}
)

func NewUserHandler(repo *data.UserRepo, key string, log *log.Logger) *UserHandler {
	return &UserHandler{
		repo: repo,
		jwt: &auth.JwtWrapper{
			SecretKey:       key,
			Issuer:          "user-service",
			ExpirationHours: 24,
		},
		log: log,
	}
}

func (uh *UserHandler) Login() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		uh.log.Println("POST LOGIN")
		login := r.Context().Value(keyvalue{}).(data.Login)
		user, err := uh.repo.GetUser(login.Email)
		if err == gorm.ErrRecordNotFound {
			rw.WriteHeader(http.StatusUnauthorized)
			data.ToJSON(&generalError{"Invalid Login information"}, rw)
			return
		}
		if err := user.CheckPassword(login.Password); err != nil {
			rw.WriteHeader(http.StatusUnauthorized)
			data.ToJSON(&generalError{"Invalid Login information"}, rw)
			return
		}
		token, err := uh.jwt.CreateJwToken(user.Email)
		if err != nil {
			uh.log.Println(err.Error())
			rw.WriteHeader(http.StatusInternalServerError)
			data.ToJSON(&generalError{"Failed to sign token"}, rw)
			return
		}
		data.ToJSON(
			struct {
				Token string `json:"token"`
			}{Token: token}, rw)
	}
}

func (uh *UserHandler) CreateUser() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		uh.log.Println("POST CREATE USER")
		user := r.Context().Value(keyvalue{}).(data.User)
		if err := uh.repo.CreateUser(&user); err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			uh.log.Println(err)
			data.ToJSON(&generalError{err.Error()}, rw)
			return
		}
		rw.WriteHeader(http.StatusNoContent)

	}
}

// func (uh *UserHandler) Options() http.HandlerFunc {
// 	return func(rw http.ResponseWriter, r *http.Request) {
// 		rw.WriteHeader(http.StatusOK)
// 	}
// }

func (uh *UserHandler) GetLoggedInUser() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		client := r.Context().Value(keyvalue{}).(clientInformation)
		data.ToJSON(&client, rw)
	}
}
