package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Product struct {
	Id           int     `json:"id"`
	Name         string  `json:"name"`
	Quantity     int     `json:"quantity"`
	Code_value   string  `json:"code_value"`
	Is_published bool    `json:"is_published"`
	Expiration   string  `json:"expiration"`
	Price        float64 `json:"price"`
}

var products []Product

func main() {
	reg, err := os.ReadFile("/Users/imihocevich/Desktop/goweb/products.json")
	if err != nil {
		fmt.Println(err)
	}
	err = json.Unmarshal(reg, &products)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(products)

}
