package data

import (
	"context"
	"time"

	"github.com/go-playground/validator"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type (
	Checkout struct {
		Name          string `json:"name" validate:"required"`
		StreetAddress string `json:"street_address" validate:"required"`
		Card          *card  `json:"card" validate:"required"`
	}

	CheckoutRepo struct {
		client *mongo.Client
		dbname string
	}
)

func NewCheckoutRepo(connString string, dbname string) *CheckoutRepo {
	client, err := mongo.NewClient(options.Client().ApplyURI(connString))
	if err != nil {
		panic(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err = client.Connect(ctx); err != nil {
		panic(err)
	}
	return &CheckoutRepo{client, dbname}
}

func (ch *Checkout) Validate() error {
	validate := validator.New()
	validate.RegisterValidation("card", validateCard)
	return validate.Struct(ch)
}
