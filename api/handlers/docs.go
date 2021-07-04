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

import "github.com/sssaang/go-cafe/api/data"

// A list of products returns in the response
// swagger:response productsResponse
type productsResponseWrapper struct {
	// products in the system
	// in: body
	Body []data.Product
}

// Data structure representing a single product
// swagger:response productResponse
type productResponseWrapper struct {
	// Newly created product
	// in: body
	Body data.Product
}

// A response with no content
// swagger:response noContentResponse
type noContentResponse struct {}

// swagger:parameters deleteProduct updateProduct
type productIdParamWrapper struct {
	// id of the product to be deleted or modified from the database
	// in:path
	// required: true
	ID int `json:"id"`
}

// swagger:parameters updateProduct createProduct
type productParamsWrapper struct {
	// Product data structure to Update or Create.
	// Note: the id field is ignored by update and create operations
	// in: body
	// required: true
	Body data.Product
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