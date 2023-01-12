package product

import (
	"errors"

	"github.com/imihocevich/goweb/practice3/internal/domain"
)

type Service interface {
	GetAll() ([]domain.Product, error)
	GetByID(id int) (domain.Product, error)
	SearchPriceGt(price float64) ([]domain.Product, error)
	Create(p domain.Product) (domain.Product, error)
	Update(id int, p domain.Product) (domain.Product, error)
	Delete(id int) error
}

type service struct {
	r Repository
}

func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) GetAll() ([]domain.Product, error) {
	todos := s.r.GetAll()
	return todos, nil
}

func (s *service) GetByID(id int) (domain.Product, error) {
	p, err := s.r.GetByID(id)
	if err != nil {
		return domain.Product{}, err
	}
	return p, nil
}

func (s *service) SearchPriceGt(price float64) ([]domain.Product, error) {
	producto := s.r.SearchPriceGt(price)
	if len(producto) == 0 {
		return []domain.Product{}, errors.New("not product found")
	}
	return producto, nil
}

func (s *service) Create(producto domain.Product) (domain.Product, error) {
	p, err := s.r.Create(producto)
	if err != nil {
		return domain.Product{}, err
	}
	return p, nil
}

func (s *service) Delete(id int) error {
	err := s.r.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

func (s *service) Update(id int, product domain.Product) (domain.Product, error) {
	p, err := s.r.GetByID(id)
	if err != nil {
		return domain.Product{}, err
	}
	if product.Name != "" {
		p.Name = product.Name
	}
	if product.CodeValue != "" {
		p.CodeValue = product.CodeValue
	}
	if product.Expiration != "" {
		p.Expiration = product.Expiration
	}
	if product.Quantity > 0 {
		p.Quantity = product.Quantity
	}
	if product.Price > 0 {
		p.Price = product.Price
	}
	r, err := s.r.Update(id, product)
	if err != nil {
		return domain.Product{}, err
	}
	return r, nil
}
