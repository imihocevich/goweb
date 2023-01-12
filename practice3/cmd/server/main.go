package main

import (
	"encoding/json"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/imihocevich/goweb/practice3/cmd/server/handler"
	"github.com/imihocevich/goweb/practice3/internal/domain"
	"github.com/imihocevich/goweb/practice3/internal/product"
)

func main() {
	productList := []domain.Product{}
	loadFile("/Users/imihocevich/Desktop/goweb/products.json", &productList)

	repository := product.NewRepository(productList)
	service := product.NewService(repository)
	productHandler := handler.NewpHandler(service)

	r := gin.Default()
	pr := r.Group("/products")
	{
		pr.GET("", productHandler.GetAll())
		pr.GET(":id", productHandler.GetByID())
		pr.GET("/search", productHandler.Search())
		pr.POST("", productHandler.Post())
		pr.DELETE(":id", productHandler.Delete())
		pr.PATCH(":id", productHandler.Patch())
		pr.PUT(":id", productHandler.Put())
	}
	r.Run(":8080")
}

func loadFile(path string, list *[]domain.Product) {
	file, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal([]byte(file), &list)
	if err != nil {
		panic(err)
	}
}
