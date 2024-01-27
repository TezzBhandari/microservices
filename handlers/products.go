package handlers

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strconv"

	"github.com/TezzBhandari/lecture-03/data"
)

type Products struct {
	logger *log.Logger
}

func NewProduct(logger *log.Logger) *Products {
	return &Products{logger}
}

func (ph *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		ph.getProducts(rw, r)
		return
	}

	if r.Method == http.MethodPost {
		ph.addProduct(rw, r)
		return
	}

	if r.Method == http.MethodDelete {
		ph.logger.Println("handling delete request")
		// expect an id in the url path
		// regex expression.
		// used to get id from the url path
		reg := regexp.MustCompile("/([0-9]+)")
		g := reg.FindAllStringSubmatch(r.URL.Path, -1)
		// ph.updateProduct(rw, r)
		ph.logger.Println(g[0])

		ph.logger.Println(len(g[0]))
		if len(g) != 1 {
			http.Error(rw, "invalid url path", http.StatusBadRequest)
			return
		}

		if len(g[0]) != 2 {
			http.Error(rw, "invalid url path", http.StatusBadRequest)
			return
		}

		idString := g[0][1]
		// converting idstring into integer
		id, err := strconv.Atoi(idString)
		if err != nil {
			http.Error(rw, "invalid id", http.StatusBadRequest)
			return
		}

		// update the product with corresponding id
		ph.deleteProduct(id, rw, r)
		return
	}

	if r.Method == http.MethodPut {
		ph.logger.Println("handling PUT reqeust")
		// expect an id in the url path
		// regex expression.
		// used to get id from the url path
		reg := regexp.MustCompile("/([0-9]+)")
		g := reg.FindAllStringSubmatch(r.URL.Path, -1)
		// ph.updateProduct(rw, r)
		ph.logger.Println(g[0])

		ph.logger.Println(len(g[0]))
		if len(g) != 1 {
			http.Error(rw, "invalid url path", http.StatusBadRequest)
			return
		}

		if len(g[0]) != 2 {
			http.Error(rw, "invalid url path", http.StatusBadRequest)
			return
		}

		idString := g[0][1]
		// converting idstring into integer
		id, err := strconv.Atoi(idString)
		if err != nil {
			http.Error(rw, "invalid id", http.StatusBadRequest)
			return
		}

		// update the product with corresponding id
		ph.updateProduct(id, rw, r)
		return
	}
	rw.WriteHeader(http.StatusMethodNotAllowed)
	rw.Write([]byte("Method not implemented"))
}

func (ph *Products) getProducts(rw http.ResponseWriter, r *http.Request) {
	ph.logger.Println("handles GET request")
	// get the product from model
	productList := data.GetProducts()

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

func (ph *Products) addProduct(rw http.ResponseWriter, r *http.Request) {
	ph.logger.Println("handles POST request")
	product := &data.Product{}

	err := product.FromJSON(r.Body)

	// _, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(rw, "unable to read body", http.StatusBadRequest)
		return
	}

	ph.logger.Printf("product: %#v", product)
	data.AddProduct(product)
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)
	fmt.Fprintf(rw, product.Name)
}

// updates product
func (ph *Products) updateProduct(id int, rw http.ResponseWriter, r *http.Request) {
	product := &data.Product{}
	err := product.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "Unable to read body", http.StatusBadRequest)
		return
	}

	ph.logger.Printf("updates: %#v", product)
	err = data.UpdateProduct(id, product)

	// update product returns error if failed to update product
	// if the product is not found in the list it returns product not found error
	if err == data.ErrorProductNotFound {
		http.Error(rw, "product not found", http.StatusNotFound)
	}

	// handling errors other than product not found
	if err != nil {
		http.Error(rw, "failed to update product", http.StatusInternalServerError)
		return
	}
	rw.Header().Add("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)
	fmt.Fprintf(rw, "successfully updated product")
}

func (ph *Products) deleteProduct(id int, rw http.ResponseWriter, r *http.Request) {

	err := data.DeleteProduct(id)

	if err == data.ErrorProductNotFound {
		http.Error(rw, "product not found", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(rw, "failed to delete product", http.StatusInternalServerError)
		return
	}
	rw.Header().Add("Content-Type", "text/plain")
	rw.WriteHeader(http.StatusOK)
	fmt.Fprintf(rw, "successfully deleted product")
}
