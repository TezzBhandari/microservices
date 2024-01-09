package handlers

import (
	"io"
	"log"
	"net/http"

	products "github.com/TezzBhandari/lecture-03/data"
)

type Products struct {
	logger *log.Logger
}

func NewProduct(logger *log.Logger) *Products {
	return &Products{logger}
}

func (ph *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		getProducts(rw, r)
		return
	}

	if r.Method == http.MethodPost {
		addProduct(rw, r)
		return
	}
	rw.WriteHeader(http.StatusMethodNotAllowed)
	rw.Write([]byte("Method not implemented"))
}

func getProducts(rw http.ResponseWriter, r *http.Request) {
	// get the product from model
	productList := products.GetProducts()

	// encode the product into json
	// productByte, err := json.Marshal(productList)

	// using the encoder with ToJSON
	rw.Header().Set("Content-Type", "applicaton/json")
	rw.WriteHeader(http.StatusOK)
	err := productList.ToJSON(rw)

	// handles any error during json encoding
	if err != nil {
		// rw.Header().Set("Content-Type", "text/plain; charset=utf-8")
		// rw.Header().Set("X-Content-Type-Options", "nosniff")
		// rw.WriteHeader(http.StatusBadRequest)
		// rw.Write([]byte("Unable to marshal the json"))
		// return
		http.Error(rw, "unable to marshal the json", http.StatusInternalServerError)
		return
	}

	// returns the http response with proper content type headers
	// rw.Header().Add("Content-Type", "application/json")
	// rw.WriteHeader(http.StatusOK)
	// rw.Write(productByte)
}

func addProduct(rw http.ResponseWriter, r *http.Request) {
	_, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(rw, "unable to read body", http.StatusBadRequest)
		return
	}

}
