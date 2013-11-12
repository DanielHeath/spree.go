package spree

import (
	"testing"
	// "time"
	"fmt"
)

func TestListProducts(t *testing.T) {
	products, _ := ProductsList()
	fmt.Println(products)
}
