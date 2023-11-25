// main.go
package main

import (
	"bleve_test/products"
	"fmt"
	"log"
)

func main() {
	// Get sample products
	sampleProducts := products.GetSampleProducts()

	// Index the products
	if err := products.IndexProducts(sampleProducts); err != nil {
		log.Fatal("Error indexing products:", err)
	}

	// Search for products by category
	query := "back"
	fmt.Printf("Searching for products with query: %s\n", query)

	searchResults, err := products.SearchProducts(query)
	if err != nil {
		log.Fatal("Error searching products:", err)
	}

	fmt.Printf("Search Results: %v\n", searchResults)
}
