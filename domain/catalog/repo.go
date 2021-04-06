package catalog

import (
	"github.com/PotatoesFall/streepjes/shared/buckets"
)

func getCatalog() (Catalog, error) {
	catalog := Catalog{}

	err := buckets.Categories.GetAll(&Category{}, func(ptr interface{}) error {
		catalog.Categories = append(catalog.Categories, *ptr.(*Category))
		return nil
	})
	if err != nil {
		return catalog, err
	}

	return catalog, buckets.Products.GetAll(&Product{}, func(ptr interface{}) error {
		catalog.Products = append(catalog.Products, *ptr.(*Product))
		return nil
	})
}
