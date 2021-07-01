package data

import (
	"encoding/json"

	"github.com/go-redis/redis"
)

type (
	Cart struct {
		Products []*CartItem `json:"items_in_cart"`
	}

	CartItem struct {
		ItemName     *Item `json:"item"`
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
	cache.Set("cart", Cart{}, 0)
	return &CartRepo{
		cache: cache,
	}
}

func (cr *CartRepo) AddItem(i *Item, qty int32) error {
	cart := Cart{}
	err := cr.retrieveCart(&cart)
	if err != redis.Nil {
		return err
	}
	if cart.Products == nil {
		cart.Products = make([]*CartItem, 0)
	}
	cart.Products = append(cart.Products, &CartItem{i, qty})
	err = cr.saveCart(&cart)
	return err
}

func (cr *CartRepo) RemoveItem(id uint) error {
	var cart Cart
	err := cr.retrieveCart(&cart)
	if err != nil {
		return err
	}

	for i, k := range cart.Products {
		if k.ItemName.ID == id {
			cart.Products = append(cart.Products[:i], cart.Products[i+1:]...)
		}
	}
	err = cr.saveCart(&cart)
	return err
}

func (cr *CartRepo) GetCart() ([]*CartItem, error) {
	var cart Cart
	err := cr.retrieveCart(&cart)

	if err != nil {
		return nil, err
	}
	return cart.Products, nil
}

func (cr *CartRepo) ClearCart() error {
	return cr.cache.Set("cart", nil, 0).Err()
}

func (c *CartRepo) retrieveCart(cart *Cart) error {
	val, err := c.cache.Get("cart").Result()
	if err != nil {
		return err
	}
	err = json.Unmarshal([]byte(val), &cart)
	return err
}

func (cr *CartRepo) saveCart(cart *Cart) error {
	json, err := json.Marshal(cart)
	if err != nil {
		return err
	}
	_, err = cr.cache.Set("cart", json, 0).Result()
	return err
}
