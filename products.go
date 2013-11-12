package spree

import (
	"encoding/json"
	// "fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

type Data struct {
	Products Products
}

type Products []Product

type Product struct {
	Name        string
	AvailableOn string `json:"available_on"`
}

func ProductsList() (products []*Product, err error) {
	var reqBody io.Reader

	reqBody = strings.NewReader("")
	req, err := http.NewRequest("GET", "https://gist.github.com/radar/7440196/raw/eea66a32c4446eb04532f8f286f3c065cc73fec2/products.json", reqBody)
	if err != nil {
		return
	}
	// submit the http request
	r, err := http.DefaultClient.Do(req)
	if err != nil {
		return
	}

	// read the body of the http message into a byte array
	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		return
	}
	json.Unmarshal(body, &products)
	return
}
