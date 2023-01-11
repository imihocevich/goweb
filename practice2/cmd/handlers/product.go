package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/imihocevich/goweb/practice2/pkg/response"
	"github.com/imihocevich/goweb/practice2/services"
	"github.com/imihocevich/goweb/practice2/services/models"
)

func GetAll(c *gin.Context) {
	c.JSON(http.StatusOK, response.Ok("succeed request", services.Products))
}

func AddProduct(c *gin.Context) {
	var addproduct models.Product

	err := c.ShouldBindJSON(&addproduct)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	addproduct.Id = services.Products[len(services.Products)-1].Id + 1
	services.Products = append(services.Products, addproduct)
	c.JSON(http.StatusOK, response.Ok("added product", services.Products))
}
