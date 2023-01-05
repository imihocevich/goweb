package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
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
	reg, err := os.ReadFile("products.json")
	if err != nil {
		fmt.Println(err)
	}

	err = json.Unmarshal(reg, &products)
	if err != nil {
		fmt.Println(err)
	}

	router := gin.Default()
	//Crear una ruta /ping que debe respondernos con un string que contenga pong con el status 200 OK.
	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	//Crear una ruta /products que nos devuelva la lista de todos los productos en la slice.
	router.GET("/products", func(c *gin.Context) {
		c.JSON(http.StatusOK, products)
	})

	//Crear una ruta /products/:id que nos devuelva un producto por su id.
	router.GET("/products/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		}
		var productid Product
		for _, v := range products {
			if v.Id == id {
				productid = v
				break
			}

		}
		c.JSON(http.StatusOK, productid)

	})

	//Crear una ruta /products/search que nos permita buscar por parámetro los productos cuyo precio sean mayor a un valor priceGt.
	router.GET("/products/search", func(c *gin.Context) {
		var query Query
		
	}
	})

	err = router.Run()
	if err != nil {
		fmt.Println(err)
	}

}
