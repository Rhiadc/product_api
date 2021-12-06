package handlers

import (
	"net/http"
	"productApi/data"
	"strconv"

	"github.com/gorilla/mux"
)

// swagger:route DELETE /products/{id} products deleteProduct
// Update a products details
//
// responses:
//	201: noContentResponse
//  404: errorResponse
//  501: errorResponse

// Delete handles DELETE requests and removes items from the database

func (p *Products) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle DELETE Products")
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Unable to convert id", http.StatusBadRequest)
		return
	}

	err = data.DeleteProduct(id)

	if err != nil {
		http.Error(w, "Unable to found given id", http.StatusNotFound)
		return
	}
}
