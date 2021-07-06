package data

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

type Order struct {
	Email         string `json:"user_email"`
	Name          string `json:"user_name"`
	StreetAddress string `json:"street_address"`
	Items         *Cart  `json:"items_ordered"`
}

func (cr *CheckoutRepo) CheckoutOrder(c *Checkout, cart *Cart) error {
	order := Order{
		Email:         c.Email,
		Name:          c.Name,
		StreetAddress: c.StreetAddress,
		Items:         cart,
	}

	collection := cr.client.Database(cr.dbname).Collection("orders")
	_, err := collection.InsertOne(context.TODO(), order)
	return err
}

func (cr *CheckoutRepo) GetAllOrders() ([]*Order, error) {
	orders := []*Order{}

	collection := cr.client.Database(cr.dbname).Collection("orders")
	cursor, err := collection.Find(context.TODO(), bson.D{{}})
	if err != nil {
		return orders, err
	}

	for cursor.Next(context.TODO()) {
		current := Order{}
		if err := cursor.Decode(&current); err != nil {
			return orders, err
		}
		orders = append(orders, &current)
	}
	cursor.Close(context.TODO())
	return orders, nil
}