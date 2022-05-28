package handlers

import (
	"golang-microservice/data"
	"log"
	"net/http"
	"regexp"
	"strconv"
)

// Products ...
type Products struct {
	lgr *log.Logger
}

// NewProducts ...
func NewProducts(lgr *log.Logger) *Products {
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
	if r.Method == http.MethodPut {
		reg := regexp.MustCompile(`/([0-9]+)`)
		matches := reg.FindAllStringSubmatch(r.URL.Path, -1)
		prd.lgr.Println("matches: ", matches)
		if len(matches) != 1 {
			http.Error(rw, "Not a valid request!", http.StatusBadRequest)
		}
		if len(matches[0]) != 2 {
			http.Error(rw, "Not a valid regex group!", http.StatusInternalServerError)
		}
		id, err := strconv.Atoi(matches[0][1])
		if err != nil {
			http.Error(rw, "Not a valid id", http.StatusInternalServerError)
		}
		prd.updateProduct(id, rw, r)
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
	prod.AddProduct()
	prd.lgr.Printf("prod: %#v", prod)
}

func (prd *Products) updateProduct(id int, rw http.ResponseWriter, r *http.Request) {
	prod := data.Product{}
	err := prod.FromJSON(r.Body)
	if err != nil {
		prd.lgr.Printf("could not decode request :%#v", err)
	}
	err = prod.UpdateProduct(id)
	if err == data.ErrNotFound {
		http.Error(rw, err.Error(), http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return 
	}
	// rw.Write([]byte("successfully updated document!"))
}
