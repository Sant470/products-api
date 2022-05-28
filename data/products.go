package data

import (
	"encoding/json"
	"fmt"
	"io"
	"time"
)

// ErrNotFound ...
var ErrNotFound = fmt.Errorf("Not Found")

// Product ...
type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	SKU         string  `json:"sku"`
	CreatedOn   string  `json:"-"`
	UpdatedOn   string  `json:"-"`
}

// Products ...
type Products []Product

// FromJSON ...
func (p *Product) FromJSON(r io.Reader) error {
	de := json.NewDecoder(r)
	return de.Decode(p)
}

// AddProduct ...
func (p *Product) AddProduct(){
	p.ID = productList[len(productList)-1].ID + 1
	fmt.Println("ID: ", p.ID)
	productList = append(productList, *p)
}

// UpdateProduct ...
func (p *Product) UpdateProduct(id int) error {
	i, err := searchProduct(id)
	if err != nil {
		return err
	}
	p.ID = id 
	productList[i] = *p
	return nil 
}

func searchProduct(id int) (i int, err error) {
	for i, prod := range productList {
		if prod.ID == id {
			return i, nil
		}
	}
	return -1, ErrNotFound
} 

// ToJSON ...
func (ps *Products) ToJSON(w io.Writer) error {
	ne := json.NewEncoder(w)
	return ne.Encode(ps)
}

// GetProducts ...
func GetProducts() Products {
	return productList
}

var productList = []Product{
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
