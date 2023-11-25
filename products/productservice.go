package products

import (
	"fmt"
	"os"
	"time"

	"github.com/blevesearch/bleve/v2"
)

const indexName = "example.bleve"

// IndexProducts indexes a list of products
func IndexProducts(products []ProductData) error {
	startTime := time.Now()

	// Remove existing index if it exists
	if err := os.RemoveAll(indexName); err != nil {
		fmt.Println("Error removing existing index:", err)
		return err
	}

	// Open a new index
	mapping := bleve.NewIndexMapping()
	index, err := bleve.New(indexName, mapping)
	if err != nil {
		return err
	}

	// Index each product in the list
	for _, product := range products {
		if err := index.Index(product.ID, product); err != nil {
			return err
		}
	}

	// Close the index to ensure all data is flushed and resources are released
	if err := index.Close(); err != nil {
		return err
	}

	indexingTime := time.Since(startTime)
	fmt.Printf("Indexing time: %v ms\n", indexingTime.Milliseconds())

	return nil
}

// SearchProducts searches for products by a given query
func SearchProducts(query string) (*bleve.SearchResult, error) {
	startTime := time.Now()

	// Open the existing index
	index, err := bleve.Open(indexName)
	if err != nil {
		return nil, err
	}

	// Use a String query but we will append wildcards
	query = "*" + query + "*"
	fuzzySearch := bleve.NewQueryStringQuery(query)
	searchRequest := bleve.NewSearchRequest(fuzzySearch)
	searchResults, err := index.Search(searchRequest)
	if err != nil {
		return nil, err
	}

	searchTime := time.Since(startTime)
	fmt.Printf("Search time: %v ms\n", searchTime.Milliseconds())

	return searchResults, nil
}
