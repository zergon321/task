package main

// receive receives the products data
// from the given input channel and
// also tracks the errors.
//
// It returns the max product by price
// and the max product by rating.
func receive(products <-chan Product, errs <-chan error) (Product, Product, error) {
	var maxPriceProduct, maxRatingProduct Product
	shouldStop := false

	for !shouldStop {
		select {
		case product, ok := <-products:
			// Select max products.
			if product.Price > maxPriceProduct.Price {
				maxPriceProduct = product
			}

			if product.Rating > maxRatingProduct.Rating {
				maxRatingProduct = product
			}

			if !ok {
				shouldStop = true
				break
			}

		case err := <-errs:
			if err != nil {
				return Product{}, Product{}, err
			}
		}
	}

	return maxPriceProduct, maxRatingProduct, nil
}
