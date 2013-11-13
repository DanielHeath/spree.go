package spree

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Data struct {
	Products []Product `json:"products"`
}

type Product struct {
	Name        string `json:"name"`
	AvailableOn string `json:"available_on"`
}

func MustProductsList() []Product {
	result, err := ProductsList()
	if err != nil {
		panic(err)
	}
	return result
}

func ProductsList() ([]Product, error) {
	body, err := getProductJson()
	if err != nil {
		return nil, err
	}
	return ParseProductJson(body)
}

func ParseProductJson(body []byte) (products []Product, err error) {
	dat := Data{}
	err = json.Unmarshal(body, &dat)
	if err != nil {
		return
	}
	return dat.Products, nil
}

func getProductJson() (body []byte, err error) {
	r, err := http.Get("https://gist.github.com/radar/7440196/raw/eea66a32c4446eb04532f8f286f3c065cc73fec2/products.json")
	if err != nil {
		return
	}
	defer r.Body.Close()

	// read the body of the http message into a byte array
	return ioutil.ReadAll(r.Body)
}
