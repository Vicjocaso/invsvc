package types

import "time"

type Product struct {
	Id          int        `json:"product_id,omitempty"`
	ProductId   string     `json:"product_id_id,omitempty"`
	ProductName string     `json:"product_name,omitempty"`
	Description string     `json:"product_description,omitempty"`
	Tag         string     `json:"product_tag,omitempty"`
	Image       string     `json:"product_image,omitempty"`
	Price       string     `json:"product_price,omitempty"`
	InStock     string     `json:"product_in_stock,omitempty"`
	Ingredients string     `json:"product_ingredients,omitempty"`
	CreateAt    *time.Time `json:"product_create_at,omitempty"`
	UpdateAt    *time.Time `json:"product_update_at,omitempty"`
}
