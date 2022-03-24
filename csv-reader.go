package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
)

// readCSV returns a channel to read the
// products info from the CSV file item
// by item and a channel to read errors
// emerged during the process.
func readCSV(file *os.File, bufferSize int) (<-chan Product, <-chan error, error) {
	products := make(chan Product, bufferSize)
	errs := make(chan error)

	go func(file *os.File, products chan<- Product, errs chan<- error) {
		defer close(products)
		defer close(errs)

		csvReader := csv.NewReader(file)
		skipHeader := false
		hadData := false

		for {
			data, err := csvReader.Read()

			if err == io.EOF {
				break
			}

			if err != nil {
				errs <- err
				return
			}

			if !skipHeader {
				skipHeader = true
				continue
			}

			// Convert fields.
			productName := data[0]
			productPrice, err := strconv.ParseFloat(data[1], 64)

			if err != nil {
				errs <- err
				return
			}

			productRating, err := strconv.ParseFloat(data[2], 64)

			if err != nil {
				errs <- err
				return
			}

			product := Product{
				Name:   productName,
				Price:  productPrice,
				Rating: productRating,
			}

			products <- product
			hadData = true
		}

		if !hadData {
			errs <- fmt.Errorf("no data")
		}
	}(file, products, errs)

	return products, errs, nil
}
