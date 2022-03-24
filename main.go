package main

import (
	"fmt"
	"io"
	"os"

	"github.com/gabriel-vasile/mimetype"
)

const (
	channelBufferSize = 32
)

func main() {
	if len(os.Args) < 2 {
		panic("input file not specified")
	}

	filename := os.Args[1]

	inputFile, err := os.Open(filename)
	handleError(err)

	mimetype.SetLimit(1024 * 1024)
	mtype, err := mimetype.DetectReader(inputFile)
	handleError(err)
	_, err = inputFile.Seek(0, io.SeekStart)
	handleError(err)

	var products <-chan Product
	var errs <-chan error

	switch mtype.String() {
	case "text/csv":
		products, errs, err = readCSV(
			inputFile, channelBufferSize)

	case "application/json":
		products, errs, err = readJSON(
			inputFile, channelBufferSize)
	}

	maxPriceProduct, maxRatingProduct,
		err := receive(products, errs)
	handleError(err)

	fmt.Println("The max price product:", maxPriceProduct.String())
	fmt.Println("The max rating product:", maxRatingProduct.String())
}

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}
