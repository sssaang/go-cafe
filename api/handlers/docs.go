package handlers

import "github.com/sssaang/go-cafe/api/data"

// A list of products returns in the response
// swagger:response productsResponse
type productsResponseWrapper struct {
	// All products in the system
	// in: body
	Body []data.Product
}

// swagger:response noContentResponse
type noContentResponse struct {}

// swagger:parameters deleteProduct
type productDeleteParamWrapper struct {
	// id of the product to be deleted from the database
	// in:path
	// required: true
	ID int `json:"id"`
}

// swagger:parameters updateProduct
type updateProduct struct {
	// id of the product to be modified from the database
	// in:path
	// required: true
	ID int `json:"id"`
}

// Generic error message returned as a string
// swagger:response errorResponse
type errorResponseWrapper struct {
	// Description of the error
	// in: body
	Body GenericError
}

// Validation errors defined as an array of strings
// swagger:response errorValidation
type errorValidationWrapper struct {
	// Collection of the errors
	// in: body
	Body ValidationError
}

type GenericError struct {
	Message string `json:"message"`
}

type ValidationError struct {
	Message string `json:"message"`
}