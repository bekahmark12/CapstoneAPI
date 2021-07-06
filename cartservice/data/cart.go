package data

import (
	"encoding/json"
	"fmt"

	"github.com/go-redis/redis"
)

type (
	Cart struct {
		Products []*CartItem `json:"items_in_cart"`
	}

	CartItem struct {
		Item         *Item `json:"item"`
		ItemQuantity int32 `json:"item_quantity"`
	}

	Item struct {
		ID          uint    `json:"id"`
		Title       string  `json:"title"`
		Description string  `json:"description"`
		Price       float64 `json:"price"`
	}

	CartRepo struct {
		cache *redis.Client
	}
)

func NewCartRepo(cache *redis.Client) *CartRepo {
	cache.Set("cart", nil, 0)
	return &CartRepo{
		cache: cache,
	}
}

func (cr *CartRepo) AddItem(i *Item, qty int32) error {
	cart, err := cr.retrieveCart()
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

func (cr *CartRepo) RemoveItem(id uint) error {
	cart, err := cr.retrieveCart()
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

func (cr *CartRepo) GetCart() (*Cart, error) {
	cart, err := cr.retrieveCart()
	if err != nil {
		return &Cart{}, err
	}

	return cart, nil
}

func (cr *CartRepo) UpdateItemQuantity(id uint, qty int32) error {
	cart, err := cr.retrieveCart()
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

func (cr *CartRepo) ClearCart() error {
	return cr.cache.Set("cart", nil, 0).Err()
}

func (c *CartRepo) retrieveCart() (*Cart, error) {
	val, err := c.cache.Get("cart").Result()
	if err != nil {
		return nil, err
	}
	if val == "" {
		return &Cart{[]*CartItem{}}, nil
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
	return cr.cache.Set("cart", json, 0).Err()
}
