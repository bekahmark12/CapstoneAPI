package data

import (
	"encoding/json"
	"fmt"

	"github.com/go-redis/redis"
)

type (
	Cart struct {
		UserEmail string      `json:"user"`
		Products  []*CartItem `json:"items_in_cart"`
	}

	CartItem struct {
		Item         *Item `json:"item"`
		ItemQuantity int32 `json:"item_quantity"`
	}

	Item struct {
		ID          uint    `json:"id"`
		ImageURL    string  `json:"url"`
		Title       string  `json:"title"`
		Description string  `json:"description"`
		Price       float64 `json:"price"`
	}

	CartRepo struct {
		cache *redis.Client
	}
)

func NewCartRepo(cache *redis.Client) *CartRepo {
	return &CartRepo{
		cache: cache,
	}
}

func (cr *CartRepo) AddItem(email string, i *Item, qty int32) error {
	cart, err := cr.retrieveCart(email)
	if err != nil {
		return nil
	}
	for _, it := range cart.Products {
		if it.Item.ID == i.ID {
			return fmt.Errorf("Item already contained in list")
		}
	}
	cart.Products = append(cart.Products, &CartItem{i, qty})
	return cr.saveCart(cart)

}

func (cr *CartRepo) RemoveItem(email string, id uint) error {
	cart, err := cr.retrieveCart(email)
	if err != nil {
		return err
	}

	for i, k := range cart.Products {
		if k.Item.ID == id {
			cart.Products = append(cart.Products[:i], cart.Products[i+1:]...)
		}
	}
	return cr.saveCart(cart)

}

func (cr *CartRepo) GetCart(email string) (*Cart, error) {
	cart, err := cr.retrieveCart(email)
	if err != nil {
		return nil, err
	}

	return cart, nil
}

func (cr *CartRepo) UpdateItemQuantity(email string, id uint, qty int32) error {
	cart, err := cr.retrieveCart(email)
	if err != nil {
		return err
	}
	for _, i := range cart.Products {
		if i.Item.ID == id {
			i.ItemQuantity = qty
		}
	}
	return cr.saveCart(cart)
}

func (cr *CartRepo) ClearCart(email string) error {
	return cr.cache.Set(email, Cart{email, []*CartItem{}}, 0).Err()
}

func (c *CartRepo) retrieveCart(email string) (*Cart, error) {
	val, err := c.cache.Get(email).Result()
	if err == redis.Nil || val == "" {
		newCart := Cart{email, []*CartItem{}}
		c.cache.Set(email, newCart, 0)
		return &newCart, nil
	}
	outCart := Cart{}
	err = json.Unmarshal([]byte(val), &outCart)
	return &outCart, err
}

func (cr *CartRepo) saveCart(cart *Cart) error {
	json, err := json.Marshal(cart)
	if err != nil {
		return err
	}
	return cr.cache.Set(cart.UserEmail, json, 0).Err()
}
