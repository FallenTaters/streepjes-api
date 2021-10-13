package catalog

import (
	"time"

	"github.com/FallenTaters/streepjes-api/model"
)

const cacheTime = time.Minute

var (
	cacheValid  = false
	lastCatalog struct {
		model.Catalog
		Time time.Time
	}
)

func getCatalog() (model.Catalog, error) {
	catalog := model.Catalog{}

	categories, err := categoryRepo.GetAll()
	if err != nil {
		return catalog, err
	}
	catalog.Categories = categories

	products, err := productRepo.GetAll()
	if err != nil {
		return catalog, err
	}
	catalog.Products = products

	return catalog, nil
}
