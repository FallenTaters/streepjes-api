package catalog

import (
	"errors"

	"github.com/PotatoesFall/bbucket"
)

var (
	ErrCategoryNotFound = errors.New("category not found")
	ErrNameTaken        = errors.New("this name is taken")
	ErrEmptyName        = errors.New("name may not be empty")
	ErrNoPrice          = errors.New("product must have at least one price")
)

func validateProduct(product Product) error {
	original, err := getProduct(product.ID)
	switch err {
	case nil:
		if original.Name != product.Name && productNameExists(product.Name) {
			return ErrNameTaken
		}
	case bbucket.ErrObjectNotFound:
		if productNameExists(product.Name) {
			return ErrNameTaken
		}
	default:
		panic(err)
	}

	if !categoryExists(product.CategoryID) {
		return ErrCategoryNotFound
	}
	if product.Name == `` {
		return ErrEmptyName
	}
	if product.PriceGladiators == 0 && product.PriceParabool == 0 {
		return ErrNoPrice
	}

	return nil
}

func productNameExists(name string) bool {
	products, err := getProducts()
	if err != nil {
		panic(err)
	}

	found := false
	for _, p := range products {
		if p.Name == name {
			found = true
		}
	}

	return found
}

func categoryExists(id int) bool {
	categories, err := getCategories()
	if err != nil {
		panic(err)
	}

	found := false
	for _, c := range categories {
		if c.ID == id {
			found = true
			break
		}
	}

	return found
}

func validateCategory(category Category) error {
	original, err := getCategory(category.ID)
	switch err {
	case nil:
		if original.Name != category.Name && categoryNameExists(category.Name) {
			return ErrNameTaken
		}
	case bbucket.ErrObjectNotFound:
		if categoryNameExists(category.Name) {
			return ErrNameTaken
		}
	default:
		panic(err)
	}

	if category.Name == `` {
		return ErrEmptyName
	}

	return nil
}

func categoryNameExists(name string) bool {
	categories, err := getCategories()
	if err != nil {
		panic(err)
	}

	found := false
	for _, c := range categories {
		if c.Name == name {
			found = true
		}
	}

	return found
}
