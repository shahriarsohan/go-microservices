package handlers

import (
	"net/http"

	"github.com/shahriarsohan/go-microservice-nic/data"
)

// swagger:route GET /products products listProducts
// Return a list of products
// responses:
//	200 : productResponse

func (p *Products) GetProducts(w http.ResponseWriter, r *http.Request) {
	lp := data.GetProducts()

	err := lp.ToJSON(w)

	if err != nil {
		http.Error(w, "Unable to marshal json", http.StatusInternalServerError)
	}
}
