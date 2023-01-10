package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/imihocevich/goweb/practice2/cmd/handlers"
	"github.com/imihocevich/goweb/practice2/services"
)

func main() {
	services.ReadProduct()

	r := gin.Default()

	products := r.Group("/products")
	products.GET("", handlers.GetAll)
	products.POST("", handlers.AddProduct)

	if err := r.Run(); err != nil {
		log.Fatal(err)
	}
}
