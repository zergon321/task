package main

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// This file demonstrates the alternative
// version of the data processing pipeline
// without utilizing channels.
//
// I decided not to utilize this approach
// because it leads to a lot of code duplication.
// I could get rid of it by using iterator
// or stretegy pattern, but the channels
// usage is more idiomatioc for Golang.

// processCSV processes the CSV file
// line by line and returns the max
// products by price and rating.
func processCSV(file *os.File) (Product, Product, error) {
	var maxPriceProduct, maxRatingProduct Product
	csvReader := csv.NewReader(file)
	skipHeader := false
	hadData := false

	for {
		data, err := csvReader.Read()

		if err == io.EOF {
			break
		}

		if err != nil {
			return Product{}, Product{}, err
		}

		if !skipHeader {
			skipHeader = true
			continue
		}

		// Convert fields.
		productName := data[0]
		productPrice, err := strconv.ParseFloat(data[1], 64)

		if err != nil {
			return Product{}, Product{}, err
		}

		productRating, err := strconv.ParseFloat(data[2], 64)

		if err != nil {
			return Product{}, Product{}, err
		}

		product := Product{
			Name:   productName,
			Price:  productPrice,
			Rating: productRating,
		}

		// Select max products.
		if product.Price > maxPriceProduct.Price {
			maxPriceProduct = product
		}

		if product.Rating > maxRatingProduct.Rating {
			maxRatingProduct = product
		}

		hadData = true
	}

	if !hadData {
		return Product{}, Product{}, fmt.Errorf("no data")
	}

	return maxPriceProduct, maxRatingProduct, nil
}

// processJSON processes the JSON file
// line by line and returns the max
// products by price and rating.
func processJSON(file *os.File) (Product, Product, error) {
	var maxPriceProduct, maxRatingProduct Product
	scanner := bufio.NewScanner(file)
	hadData := false

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)

		if line == "[" {
			continue
		}

		// Retrieve the product info.
		var product Product
		err := json.Unmarshal(unsafeGetBytes(
			line), &product)

		if err != nil {
			return Product{}, Product{}, err
		}

		// Select max products.
		if product.Price > maxPriceProduct.Price {
			maxPriceProduct = product
		}

		if product.Rating > maxRatingProduct.Rating {
			maxRatingProduct = product
		}

		hadData = true
	}

	if !hadData {
		return Product{}, Product{}, fmt.Errorf("no data")
	}

	return maxPriceProduct, maxRatingProduct, nil
}
