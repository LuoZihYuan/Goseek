package repository

import (
	_ "embed"
	"encoding/json"
	"log"
	"math/rand/v2"
	"os"
	"regexp"

	"github.com/LuoZihYuan/Goseek/internal/models"
	"github.com/LuoZihYuan/Goseek/internal/shared"
)

type ProductsRepository struct {
	products [100_000]*models.Product
}

func NewProductsRepository(dataPath string) *ProductsRepository {
	repo := &ProductsRepository{}
	if err := repo.LoadProducts(dataPath); err != nil {
		log.Fatal(err)
	}
	return repo
}

func (r *ProductsRepository) LoadProducts(dataPath string) error {
	file, err := os.Open(dataPath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	if err := json.NewDecoder(file).Decode(&r.products); err != nil {
		log.Fatal(err)
	}
	return err
}

func (r *ProductsRepository) SearchByQuery(query string, maxSearches int, maxResults int) ([]*models.Product, error) {
	re, err := regexp.Compile("(?i)" + query)
	if err != nil {
		return nil, err
	}

	start_index := rand.IntN(len(r.products))

	var found []*models.Product
	for i := range maxSearches {
		index := (start_index + i) % len(r.products)
		if re.MatchString(r.products[index].Name) || re.MatchString(r.products[index].Category) {
			found = append(found, r.products[index])
		}
		if len(found) == maxResults {
			break
		}
	}
	if len(found) == 0 {
		return nil, shared.ErrProductNotFound
	}
	return found, nil
}
