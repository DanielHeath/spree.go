package spree

import (
	"testing"
	// "time"
	"fmt"
)

func TestListProducts(t *testing.T) {
	products := Products.List()
	fmt.Println(products[0].AvailableOn)

}
