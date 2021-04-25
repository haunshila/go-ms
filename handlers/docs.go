// Package classification of Product API
//
// Documentations for Product API
//
// Schemes: http
// BasePath: /
// Version: 1.0.0
//
// Consumes:
// - application/json
//
// Produces:
// - application/json
//
// swagger:meta
package handlers

import "github.com/haunshila/go-ms/data"

// A list of products
// swagger:response productsResponse
type productsResponseWrapper struct {
	// All current products
	// in: body
	Body []data.Product
}
