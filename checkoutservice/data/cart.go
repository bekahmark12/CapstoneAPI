package data

type (
	Cart struct {
		Products []*cartItem `json:"items_in_cart"`
	}

	cartItem struct {
		Item         *item `json:"item"`
		ItemQuantity int32 `json:"item_quantity"`
	}
	item struct {
		ID          uint    `json:"id"`
		Title       string  `json:"title"`
		Description string  `json:"description"`
		Price       float64 `json:"price"`
	}
)
