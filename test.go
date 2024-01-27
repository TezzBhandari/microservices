package main

import "fmt"

func Remove() {
	// Sample slice of products
	products := []string{"product1", "product2", "product3", "product4", "product5"}

	// Index of the element you want to remove
	indexToRemove := 2

	// Check if the index is valid
	if indexToRemove >= 0 && indexToRemove < len(products) {
		// Remove the element by creating a new slice without the specified index
		products = append(products[:indexToRemove], products[indexToRemove+1:]...)

		// Display the modified slice
		fmt.Println(products)
	} else {
		fmt.Println("Invalid index to remove.")
	}
}
