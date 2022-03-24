package main

import "fmt"

// Product is a data structure
// to save the single product info.
type Product struct {
	Name   string  `json:"product"`
	Price  float64 `json:"price"`
	Rating float64 `json:"rating"`
}

// String returns the string
// representation of a product.
func (product *Product) String() string {
	return fmt.Sprintf(
		"the %s product with the rating of %.0f and the price of %.0f",
		product.Name, product.Rating, product.Price)
}
