package data

import (
	"encoding/json"
	"io"
	"time"
)

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	SKU         string  `json:"sku"`
	CreatedOn   string  `json:"-"`
	UpdatedOn   string  `json:"-"`
}

type Products []Product

func (p *Product) FromJSON(r io.Reader) error {
	de := json.NewDecoder(r)
	return de.Decode(p)
}

func (p *Product) AddProduct(){
	p.ID = productList[len(productList)-1].ID + 1
	productList = append(productList, *p)
}

func (ps *Products) ToJSON(w io.Writer) error {
	ne := json.NewEncoder(w)
	return ne.Encode(ps)
}


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
