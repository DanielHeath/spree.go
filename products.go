package spree

import (
	"encoding/json"
	// "fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

type ProductClient struct{}

type Product struct {
	Name        string
	AvailableOn string `json:"available_on"`
}

func (self *ProductClient) List() []*Product {
	// define a wrapper function for the Product List, so that we can
	// cleanly parse the JSON
	type listProductResp struct{ Data []*Product }
	resp := listProductResp{}

	var reqBody io.Reader
	reqBody = strings.NewReader("")
	req, err := http.NewRequest("GET", "http://localhost:3000/api/products.json", reqBody)
	if err != nil {
		return err
	}
	// submit the http request
	r, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	// read the body of the http message into a byte array
	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		return err
	}
	products := json.Unmarshal(body, resp)

	return products
}

// func main() {
// 	var data Products
// 	err := json.Unmarshal(body, &data)
// 	if err != nil {
// 		fmt.Println("something went wrong!")
// 	}

// 	fmt.Println(data.Products[0].AvailableOn)
// }
