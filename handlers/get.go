package handlers

import (
	"net/http"
	"productApi/data"
	"strconv"

	"github.com/gorilla/mux"
)

// swagger:route GET /products products listProducts
// Return a list of products from the database
// responses:
//	200: productsResponse

// ListAll handles GET requests and returns all current products
func (p *Products) GetProducts(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle GET Products")

	// fetch the products from the datastore
	lp := data.GetProducts()

	// serialize the list to JSON
	err := data.ToJSON(lp, rw)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}

func (p *Products) GetProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Unable to convert id", http.StatusBadRequest)
		return
	}

	prod, err := data.GetById(id)
	if err != nil {
		http.Error(w, "Product not found", http.StatusNotFound)
	}

	err = data.ToJSON(prod, w)

}
