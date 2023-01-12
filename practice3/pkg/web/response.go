package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type errorResponse struct {
	Status  int    `json:"status"`
	Code    string `json:"code"`
	Message string `json:"message"`
}

type response struct {
	Data interface{} `json:"data"`
}

func Success(c *gin.Context, status int, data interface{}) {
	c.JSON(status, response{
		Data: data,
	})
}

func Failure(c *gin.Context, status int, err error) {
	c.JSON(status, errorResponse{
		Message: err.Error(),
		Status:  status,
		Code:    http.StatusText(status),
	})
}
