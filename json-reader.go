package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

// readJSON returns a channel to read the
// products info from the JSON file item
// by item and a channel to read errors
// emerged during the process.
func readJSON(file *os.File, bufferSize int) (<-chan Product, <-chan error, error) {
	products := make(chan Product, bufferSize)
	errs := make(chan error)

	go func(file *os.File, products chan<- Product, errs chan<- error) {
		defer close(products)
		defer close(errs)

		scanner := bufio.NewScanner(file)
		hadData := false

		for scanner.Scan() {
			// Clear the raw JSON text data.
			line := scanner.Text()
			line = strings.TrimSpace(line)
			line = strings.TrimSuffix(line, ",")

			if line == "[" || line == "]" {
				continue
			}

			// Retrieve the product info.
			var product Product
			err := json.Unmarshal(unsafeGetBytes(
				line), &product)

			if err != nil {
				errs <- err
				return
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
