package data

import "time"

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	SKU         string  `json:"sku"`
	CreatedOn   string  `json:"-"`
	UpdatedOn   string  `json:"-"`
	DeletedOn   string  `json:"-"`
}

func GetProducts() []*Product {
	return prodcutList
}

var prodcutList = []*Product{
	&Product{
		ID:          1,
		Name:        "Laptop",
		Description: "Super super fast laptop",
		Price:       310.1,
		SKU:         "D1234",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          2,
		Name:        "Mobile",
		Description: "Super compact mobile",
		Price:       210.1,
		SKU:         "R1234",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
}