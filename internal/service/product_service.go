package service

import (
	"errors"
	"go-ecommerce/internal/model"
	"go-ecommerce/internal/repository"
)

type ProductService interface {
	GetAll() ([]model.Product, error)
	GetByID(id uint) (*model.Product, error)
	Create(p *model.Product) error
	Update(id uint, p *model.Product) error
	Delete(id uint) error
}

type productService struct {
	repo repository.ProductRepository
}

func NewProductService(repo repository.ProductRepository) ProductService {
	return &productService{repo: repo}
}

func (s *productService) GetAll() ([]model.Product, error) {
	return s.repo.FindAll()
}

func (s *productService) GetByID(id uint) (*model.Product, error) {
	return s.repo.FindByID(id)
}

func (s *productService) Create(p *model.Product) error {
	if p.Name == "" {
		return errors.New("tên sản phẩm không được rỗng")
	}
	if p.Price < 0 {
		return errors.New("giá không được âm")
	}
	return s.repo.Create(p)
}

func (s *productService) Update(id uint, p *model.Product) error {
	existing, err := s.repo.FindByID(id)
	if err != nil {
		return errors.New("sản phẩm không tồn tại")
	}
	existing.Name = p.Name
	existing.Price = p.Price
	existing.Stock = p.Stock
	existing.Description = p.Description
	return s.repo.Update(existing)
}

func (s *productService) Delete(id uint) error {
	return s.repo.Delete(id)
}
