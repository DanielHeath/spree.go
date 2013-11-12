package spree

import (
	"encoding/json"
	// "fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

type Products []Product

type Product struct {
	Name        string
	AvailableOn string `json:"available_on"`
}

func (products Products) List() (Products, error) {
	var reqBody io.Reader

	reqBody = strings.NewReader("")
	req, err := http.NewRequest("GET", "http://localhost:3000/api/products.json", reqBody)
	if err != nil {
		return products, err
	}
	// submit the http request
	r, err := http.DefaultClient.Do(req)
	if err != nil {
		return products, err
	}

	// read the body of the http message into a byte array
	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		return products, err
	}
	json.Unmarshal(body, &products)
	return products, err
}
