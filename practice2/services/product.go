package services

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/imihocevich/goweb/practice2/services/models"
)

var Products []models.Product

func ReadProduct() {
	reg, err := os.ReadFile("/Users/imihocevich/Desktop/goweb/products.json")
	if err != nil {
		fmt.Println(err)
	}
	err = json.Unmarshal(reg, &Products)
	if err != nil {
		fmt.Println(err)
	}

}
