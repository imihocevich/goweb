package handler

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/imihocevich/goweb/practice3/internal/domain"
	"github.com/imihocevich/goweb/practice3/internal/product"
)

type pHandler struct {
	s product.Service
}

func NewpHandler(s product.Service) *pHandler {
	return &pHandler{
		s: s,
	}
}

func (h *pHandler) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		p, _ := h.s.GetAll()
		c.JSON(200, p)
	}
}

func (h *pHandler) GetByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		idp := c.Param("id")
		id, err := strconv.Atoi(idp)
		if err != nil {
			c.JSON(400, gin.H{"error": "invalid id"})
			return
		}
		product, err := h.s.GetByID(id)
		if err != nil {
			c.JSON(404, gin.H{"error": "product not found"})
		}
		c.JSON(200, product)
	}
}

func (h *pHandler) Post() gin.HandlerFunc {
	return func(c *gin.Context) {
		var p domain.Product
		err := c.ShouldBindJSON(&p)
		if err != nil {
			c.JSON(400, gin.H{"error": "invalid product"})
			return
		}
		product, err := h.s.Create(p)
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		c.JSON(201, product)
	}

}

func (h *pHandler) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		idp := c.Param("id")
		id, err := strconv.Atoi(idp)
		if err != nil {
			c.JSON(400, gin.H{"error": "invalid id"})
			return
		}
		err = h.s.Delete(id)
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, gin.H{"success": "deleted successfully"})

	}
}

func (h *pHandler) Put() gin.HandlerFunc {
	return func(c *gin.Context) {
		idp := c.Param("id")
		id, err := strconv.Atoi(idp)
		if err != nil {
			c.JSON(400, gin.H{"error": "invalid id"})
			return
		}
		var product domain.Product
		err = c.ShouldBindJSON(&product)
		if err != nil {
			c.JSON(400, gin.H{"error": "invalid product"})
			return
		}
		p, err := h.s.Update(id, product)
		if err != nil {
			c.JSON(409, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, p)

	}
}

func (h *pHandler) Patch() gin.HandlerFunc {
	type Request struct {
		Name        string  `json:"name,omitempty"`
		Quantity    int     `json:"quantity,omitempty"`
		CodeValue   string  `json:"code_value,omitempty"`
		IsPublished bool    `json:"is_published,omitempty"`
		Expiration  string  `json:"expiration,omitempty"`
		Price       float64 `json:"price,omitempty"`
	}
	return func(c *gin.Context) {
		var r Request
		idp := c.Param("id")
		id, err := strconv.Atoi(idp)
		if err != nil {
			c.JSON(400, gin.H{"error": "invalid id"})
			return
		}
		err = c.ShouldBindJSON(&r)
		if err != nil {
			c.JSON(400, gin.H{"error": "invalid request"})
			return
		}
		update := domain.Product{
			Name:        r.Name,
			Quantity:    r.Quantity,
			CodeValue:   r.CodeValue,
			IsPublished: r.IsPublished,
			Expiration:  r.Expiration,
			Price:       r.Price,
		}
		p, err := h.s.Update(id, update)
		if err != nil {
			c.JSON(409, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, p)

	}
}

func (h *pHandler) Search() gin.HandlerFunc {
	return func(c *gin.Context) {
		priceP := c.Query("priceGt")
		price, err := strconv.ParseFloat(priceP, 64)
		if err != nil {
			c.JSON(400, gin.H{"error": "invalid price"})
			return
		}
		products, err := h.s.SearchPriceGt(price)
		if err != nil {
			c.JSON(404, gin.H{"error": "no products found"})
			return
		}
		c.JSON(200, products)
	}
}
