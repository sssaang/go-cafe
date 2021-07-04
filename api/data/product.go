package data

import (
	"encoding/json"
	"fmt"
	"io"
	"regexp"
	"time"

	"github.com/go-playground/validator"
)

// Product defines the structure for an API product
// swagger:model
type Product struct {
	// the id for the user
	//
	// required: true
	// min: 1
	ID int `json:"id"`

	// the name for the product
	//
	// required: true
	// max length: 255
	Name string `json:"name" validate:"required"`

	// the description for the product
	//
	// required: false
	// max length: 10000
	Description string `json:"description"`

	// the price for the product
	//
	// required: true
	// min: 0.01
	Price float32 `json:"price" validate:"required,gt=0"`

	// the SKU for the product
	//
	// required: true
	// pattern: [a-z]+-[a-z]+-[a-z]+
	SKU string `json:"sku" validate:"sku"`
	CreatedOn string `json:"-"`
	UpdatedOn string `json:"-"`
	DeletedOn string `json:"-"`
}


func (p *Product) Validate() error {
	validate := validator.New()
	validate.RegisterValidation("sku", valdiateSKU)
	return validate.Struct(p)
}

func valdiateSKU(fl validator.FieldLevel) bool {
	regex := regexp.MustCompile(`[a-z]+-[a-z]+-[a-z]+`)
	matches := regex.FindAllString(fl.Field().String(), -1)

	return len(matches) == 1
}

func (p *Product) FromJSON(r io.Reader) error {
	decoder := json.NewDecoder(r)

	return decoder.Decode(p)
}

type Products []*Product

func (p *Products) ToJSON(w io.Writer) error {
	encoder := json.NewEncoder(w)

	// encoding itself and if returns error if there's an error
	return encoder.Encode(p)
}

var productList = Products{
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

func GetProducts() Products {
	return productList
}

func AddProduct(p *Product) {
	p.ID = getNextID()
	productList = append(productList, p)
}

func getNextID() int {
	lp := productList[len(productList) - 1]
	return lp.ID + 1
}

func UpdateProduct(id int, prod *Product) error {
	_, pos, err := findProduct(id)
	if err != nil {
		return err
	}

	prod.ID = id
	productList[pos] = prod
	return nil
}

func DeleteProduct(id int) error {
	_, pos, err := findProduct(id)
	if err != nil {
		return err
	}

	productList = append(productList[:pos], productList[pos+1:]...)
	return nil
}

var ErrProductNotFound = fmt.Errorf("Product not found")

func findProduct(id int) (*Product, int, error)  {
	for i, p := range productList {
		if p.ID == id {
			return p, i, nil
		}
	}
	return nil, -1, ErrProductNotFound
}