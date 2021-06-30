package data

import "time"

type Product struct {
	ID int
	Name string
	Description string
	Price float32
	SKU string
	CreatedOn string
	UpdatedOn string
	DeletedOn string
}

var productList = []*Product{
	&Product{
		ID: 1,
		Name: "Latte",
		Description: "milky coffee",
		Price: 2.45,
		SKU: "abc123",
		CreatedOn: time.Now().UTC().String(),
		UpdatedOn: time.Now().UTC().String(),
	},
	&Product{
		ID: 2,
		Name: "Espresso",
		Description: "dark strong shot of coffee",
		Price: 1.99,
		SKU: "asdf34",
		CreatedOn: time.Now().UTC().String(),
		UpdatedOn: time.Now().UTC().String(),
	},
}