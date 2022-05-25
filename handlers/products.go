package handlers

import (
	"golang-microservice/data"
	"log"
	"net/http"
)

type Products struct {
	lgr *log.Logger
}

func NewProducts(lgr *log.Logger) *Products{
	return &Products{lgr}
}

func (prd *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		prd.getProducts(rw, r)
		return 
	}
	if r.Method == http.MethodPost {
		prd.addProducts(rw, r)
		return 
	}
	rw.WriteHeader(http.StatusMethodNotAllowed)
}

func (prd *Products) getProducts(rw http.ResponseWriter, r *http.Request) {
	products := data.GetProducts()
	err := products.ToJSON(rw)
	if err != nil {
		http.Error(rw, "error encoding into json", http.StatusInternalServerError)
	}
}

func (prd *Products) addProducts(rw http.ResponseWriter, r *http.Request) {
	prod := data.Product{}
	err := prod.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "Could not Decode Body!", http.StatusBadRequest)
	}
	prd.lgr.Printf("prod: %#v", prod)
	prod.AddProduct()
}

func (prd *Products) updateProduct(rw http.ResponseWriter, r *http.Request) {
	
}