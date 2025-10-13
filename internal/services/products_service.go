package services

import (
	"log"

	"github.com/LuoZihYuan/Goseek/internal/models"
	"github.com/LuoZihYuan/Goseek/internal/repository"
	"github.com/LuoZihYuan/Goseek/internal/shared"
)

type ProductsService struct {
	repo *repository.ProductsRepository
}

func NewProductService(repo *repository.ProductsRepository) *ProductsService {
	return &ProductsService{repo: repo}
}

func (s *ProductsService) SearchByQuery(query string, maxSearches int, maxResults int) ([]*models.Product, error) {
	products, err := s.repo.SearchByQuery(query, maxSearches, maxResults)
	if err == shared.ErrProductNotFound {
		return nil, err
	} else if err != nil {
		log.Fatal(err)
	}
	return products, nil
}
