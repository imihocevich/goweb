package handler_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/imihocevich/goweb/practice3/cmd/server/handler"
	"github.com/imihocevich/goweb/practice3/internal/domain"
	"github.com/imihocevich/goweb/practice3/internal/product"
	"github.com/stretchr/testify/assert"
)

func CreateServer(token string) *gin.Engine {

	err := os.Setenv("TOKEN", token)
	if err != nil {
		panic(err)
	}

	productListTest := []domain.Product{}
	LoadFile("/Users/imihocevich/Desktop/goweb/products.json", &productListTest)

	repositoryTest := product.NewRepository(productListTest)
	serviceTest := product.NewService(repositoryTest)
	productHandlerTest := handler.NewpHandler(serviceTest)

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	pr := r.Group("/products")
	{
		pr.GET("", productHandlerTest.GetAll())
		pr.GET(":id", productHandlerTest.GetByID())
		pr.GET("/search", productHandlerTest.Search())
		pr.POST("", productHandlerTest.Post())
		pr.DELETE(":id", productHandlerTest.Delete())
		pr.PATCH(":id", productHandlerTest.Patch())
		pr.PUT(":id", productHandlerTest.Put())
	}
	return r
}

func LoadFile(path string, list *[]domain.Product) {
	file, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal([]byte(file), &list)
	if err != nil {
		panic(err)
	}
}

func loadProducts(path string) ([]domain.Product, error) {
	var products []domain.Product
	file, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal([]byte(file), &products)
	if err != nil {
		return nil, err
	}
	return products, nil
}

func createRequestTest(method string, url string, body string, token string) (*http.Request, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(method, url, bytes.NewBuffer([]byte(body)))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("TOKEN", token)
	return req, httptest.NewRecorder()
}

func GetAllTest(t *testing.T) {
	var expectd = []domain.Product{}

	r := CreateServer("secret")
	req, rr := createRequestTest(http.MethodGet, "/products", "", "secret")

	LoadFile("/Users/imihocevich/Desktop/goweb/products.json", &expectd)

	r.ServeHTTP(rr, req)

	actual := map[string][]domain.Product{}

	assert.Equal(t, 200, rr.Code)
	assert.Equal(t, expectd, actual["data"])
}

func GetByIDTest(t *testing.T) {
	var expectd = domain.Product{
		Id:          7,
		Name:        "Melon - Honey Dew",
		Quantity:    165,
		CodeValue:   "S52381G",
		IsPublished: true,
		Expiration:  "01/06/2021",
		Price:       622.33,
	}
	r := CreateServer("secret")
	req, rr := createRequestTest(http.MethodGet, "/products/7", "", "secret")
	r.ServeHTTP(rr, req)

	p, err := loadProducts("/Users/imihocevich/Desktop/goweb/products.json")
	if err != nil {
		panic(err)
	}
	expectd = p[6]

	actual := map[string]domain.Product{}

	assert.Equal(t, 200, rr.Code)
	err = json.Unmarshal(rr.Body.Bytes(), &actual)
	assert.Nil(t, err)
	assert.Equal(t, expectd, actual["data"])

}
