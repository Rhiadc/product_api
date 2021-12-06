package data

import (
	"fmt"
	"time"
)

// Product defines the structure for an API product
// swagger:model
type Product struct {
	// the id for this user
	//
	// required: true
	// min: 1
	//
	ID          int     `json:"id"`
	Name        string  `json:"name" validate:"required"`
	Description string  `json:"description"`
	Price       float32 `json:"price" validate:"gt=0"`
	SKU         string  `json:"sku" validate:"required,sku"`
	CreatedOn   string  `json:"-"`
	UpdatedOn   string  `json:"-"`
	DeletedOn   string  `json:"-"`
}

// Products is a collection of Product
type Products []*Product

// GetProducts returns a list of products
func GetProducts() Products {
	return productList
}

func AddProduct(p *Product) {
	p.ID = getNexID()
	productList = append(productList, p)
}

func DeleteProduct(id int) error {
	_, _, err := findProduct(id)
	if err != nil {
		return ErrProductNotFound
	}

	new := []*Product{}
	for _, p := range productList {
		if p.ID != id {
			new = append(new, p)
		}
	}
	productList = new
	return nil
}

func UpdateProduct(id int, p *Product) error {
	_, pos, err := findProduct(id)
	if err != nil {
		return err
	}

	p.ID = id
	productList[pos] = p

	return nil
}

var ErrProductNotFound = fmt.Errorf("Product not found")

func GetById(id int) (*Product, error) {
	for _, p := range productList {
		if p.ID == id {
			return p, nil
		}
	}
	return nil, ErrProductNotFound
}

func findProduct(id int) (*Product, int, error) {
	for i, p := range productList {
		if p.ID == id {
			return p, i, nil
		}
	}
	return nil, -1, ErrProductNotFound
}

func getNexID() int {
	lp := productList[len(productList)-1]
	return lp.ID + 1
}

// productList is a hard coded list of products for this
// example data source
var productList = []*Product{
	{
		ID:          1,
		Name:        "Latte",
		Description: "Frothy milky coffee",
		Price:       2.45,
		SKU:         "abc323",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	{
		ID:          2,
		Name:        "Espresso",
		Description: "Short and strong coffee without milk",
		Price:       1.99,
		SKU:         "fjd34",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
}
