// Documentation for Product API
//
//	Schemes: http
//	BasePath: /
//	Version: 1.0.0
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
// swagger:meta
package handlers

import (
	"context"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/sssaang/go-cafe/api/data"
)

// A list of products returns in the response
// swagger:response productsResponse
type productsResponseWrapper struct {
	// All products in the system
	// in: body
	Body []data.Product
}

// swagger:response noContent
type noContent struct {}

// swagger:parameters deleteProduct
type productDeleteParamWrapper struct {
	// id of the product to be deleted from the database
	// in:path
	// required: true
	ID int `json:"id"`
}

type Products struct {
	l *log.Logger
}

func NewProduct(l *log.Logger) *Products {
	return &Products{l}
}

// swagger:route GET /products products listProducts
// Returns a list of all products
// responses:
//  200: productsResponse

// Retrieves all products from the database
func (p *Products) GetProducts(rw http.ResponseWriter, r *http.Request) {
	lp := data.GetProducts()
	err := lp.ToJSON(rw)

	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}

// Inserts a product to the database
func (p *Products) AddProduct(rw http.ResponseWriter, r *http.Request) {
	prod := r.Context().Value(KeyProduct{}).(*data.Product)

	p.l.Printf("Prod: %#v", prod)
	data.AddProduct(prod)
}

// Updates a product from the database
func (p *Products) UpdateProduct(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(rw, "Invalid Id", http.StatusBadRequest)
	}

	prod := r.Context().Value(KeyProduct{}).(*data.Product)

	err = data.UpdateProduct(id, prod)

	if err == data.ErrProductNotFound {
		http.Error(rw, "Product not found!", http.StatusInternalServerError)
	}
}

// swagger:route DELETE /products/{id} products deleteProduct
// Deletes a product
// responses:
//  201: noContent

// Deletes a product from the database
func (p *Products) DeleteProduct(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(rw, "Invalid Id", http.StatusBadRequest)
	}
	err = data.DeleteProduct(id)
	if err == data.ErrProductNotFound {
		http.Error(rw, "Product not found!", http.StatusInternalServerError)
	}
}

type KeyProduct struct {}

func (p *Products) MiddlewareProductValidation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		prod := &data.Product{}

		err := prod.FromJSON(r.Body)
		if err != nil {
			http.Error(rw, "Unable to unmarshal data", http.StatusBadRequest)
			return
		}

		err = prod.Validate()
		if err != nil {
			http.Error(rw, "Invalid data format", http.StatusBadRequest)
			return
		}

		ctx := context.WithValue(r.Context(), KeyProduct{}, prod)
		req := r.WithContext(ctx)

		next.ServeHTTP(rw, req)
	})
}
