package handlers

import (
	"net/http"

	"github.com/shahriarsohan/go-microservice-nic/data"
)

// swagger:route POST /products products createProduct
// Create a new product
//
// responses:
//	200: productResponse
//  422: errorValidation
//  501: errorResponse

// Create handles POST requests to add new products
func (p *Products) AddProduct(w http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle POST")
	prod := r.Context().Value(KeyProduct{}).(data.Product)
	data.AddProducts(&prod)
}
