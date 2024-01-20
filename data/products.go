package data

import (
	"encoding/json"
	"fmt"
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
	DeletedOn   string  `json:"-"`
}

type Products []*Product

func (products *Products) ToJSON(rw io.Writer) error {
	fmt.Println(products)
	e := json.NewEncoder(rw)
	return e.Encode(products)
}

func (product *Product) FromJSON(r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(product)
}

var productList = Products{
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

func AddProduct(product *Product) {
	newId := getNextId()
	product.ID = newId
	productList = append(productList, product)
}

func getNextId() int {
	return len(productList) + 1
}

func GetProducts() Products {
	return productList
}

func UpdateProduct(id int, update *Product) error {
	product, pos, err := findProductById(id)
	if err != nil {
		return err
	}

	prevId := product.ID
	productList[pos] = update
	productList[pos].ID = prevId
	return nil
}

var ErrorProductNotFound = fmt.Errorf("product not found")

func findProductById(id int) (*Product, int, error) {
	for index, product := range productList {
		if product.ID == id {
			return product, index, nil
		}
	}
	return nil, -1, ErrorProductNotFound
}
