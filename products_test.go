package spree

import (
	"testing"
	// "time"
)

var (
	Products = new(ProductClient)
)

func TestListProducts(t *testing.T) {
	products = Products.List()

}
