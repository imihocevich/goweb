package product

import (
	"errors"

	"github.com/imihocevich/goweb/practice3/internal/domain"
)

type Repository interface {
	GetAll() []domain.Product
	GetByID(id int) (domain.Product, error)
	SearchPriceGt(price float64) []domain.Product
	Create(p domain.Product) (domain.Product, error)
	Update(id int, p domain.Product) (domain.Product, error)
	Delete(id int) error
}

type repository struct {
	list []domain.Product
}

func NewRepository(list []domain.Product) Repository {
	return &repository{list}
}

func (r *repository) GetAll() []domain.Product {
	return r.list
}

func (r *repository) GetByID(id int) (domain.Product, error) {
	for _, product := range r.list {
		if product.Id == id {
			return product, nil
		}
	}
	return domain.Product{}, errors.New("product not found")

}

func (r *repository) SearchPriceGt(price float64) []domain.Product {
	var products []domain.Product
	for _, v := range r.list {
		if v.Price > price {
			products = append(products, v)

		}
	}
	return products
}

func (r *repository) Create(product domain.Product) (domain.Product, error) {
	if !r.validateCode(product.CodeValue) {
		return domain.Product{}, errors.New("code already set")
	}
	product.Id = len(r.list) + 1
	r.list = append(r.list, product)
	return product, nil
}

func (r *repository) validateCode(codeValue string) bool {
	for _, v := range r.list {
		if v.CodeValue == codeValue {
			return false

		}
	}
	return true
}

func (r *repository) Delete(id int) error {
	for i, v := range r.list {
		if v.Id == id {
			r.list = append(r.list[:i], r.list[i+1:]...)
			return nil
		}
	}
	return errors.New("product not found")
}

func (r *repository) Update(id int, product domain.Product) (domain.Product, error) {
	for i, v := range r.list {
		if v.Id == id {
			if !r.validateCode(product.CodeValue) && v.CodeValue != product.CodeValue {
				return domain.Product{}, errors.New("code value set")
			}
			r.list[i] = product
			return product, nil
		}
	}
	return domain.Product{}, errors.New("product not found")
}
