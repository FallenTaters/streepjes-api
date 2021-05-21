package catalog

import (
	"github.com/PotatoesFall/streepjes/shared/buckets"
)

func getCatalog() (Catalog, error) {
	catalog := Catalog{}

	categories, err := getCategories()
	if err != nil {
		return catalog, err
	}
	catalog.Categories = categories

	products, err := getProducts()
	if err != nil {
		return catalog, err
	}
	catalog.Products = products

	return catalog, nil
}

func addProduct(product Product) error {
	product.ID = buckets.Products.NextSequence()
	return buckets.Products.Create(buckets.Itob(product.ID), product)
}

func getCategories() ([]Category, error) {
	categories := []Category{}
	return categories, buckets.Categories.GetAll(&Category{}, func(ptr interface{}) error {
		categories = append(categories, *ptr.(*Category))
		return nil
	})
}

func getProducts() ([]Product, error) {
	products := []Product{}
	return products, buckets.Products.GetAll(&Product{}, func(ptr interface{}) error {
		products = append(products, *ptr.(*Product))
		return nil
	})
}

func getProduct(id int) (Product, error) {
	var product Product
	return product, buckets.Products.Get(buckets.Itob(id), &product)
}
