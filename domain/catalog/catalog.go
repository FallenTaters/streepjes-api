package catalog

import (
	"time"

	"github.com/FallenTaters/streepjes-api/model"
	"github.com/FallenTaters/streepjes-api/repo"
)

var (
	productRepo  repo.ProductRepo
	categoryRepo repo.CategoryRepo
)

func Init(pr repo.ProductRepo, cr repo.CategoryRepo) {
	productRepo = pr
	categoryRepo = cr
}

func Get() (model.Catalog, error) {
	if time.Since(lastCatalog.Time) > cacheTime || !cacheValid {
		c, err := getCatalog()
		if err != nil {
			return model.Catalog{}, err
		}
		lastCatalog.Catalog = c
		lastCatalog.Time = time.Now()

		cacheValid = true
	}

	return lastCatalog.Catalog, nil
}

func PutProduct(product model.Product) error {
	if err := validateProduct(product); err != nil {
		return err
	}

	cacheValid = false

	if product.ID != 0 {
		return productRepo.Update(product)
	}
	return productRepo.Add(product)
}

func DeleteProduct(id int) error {
	cacheValid = false
	return productRepo.Delete(id)
}

func PutCategory(category model.Category) error {
	if err := validateCategory(category); err != nil {
		return err
	}

	cacheValid = false

	if category.ID != 0 {
		return categoryRepo.Update(category)
	}
	return categoryRepo.Add(category)
}

func DeleteCategory(id int) error {
	if err := validateDeleteCategory(id); err != nil {
		return err
	}

	cacheValid = false

	return categoryRepo.Delete(id)
}
