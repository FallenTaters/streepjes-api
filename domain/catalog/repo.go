package catalog

import (
	"encoding/json"

	"github.com/PotatoesFall/streepjes/shared"
	"go.etcd.io/bbolt"
)

var db *bbolt.DB

func getCatalog() Catalog {
	c := Catalog{}
	err := db.View(func(tx *bbolt.Tx) error {
		c.Categories = _getCategories(tx)
		c.Products = _getProducts(tx)
		return nil
	})
	if err != nil {
		panic(err)
	}
	return c
}

func _getCategories(tx *bbolt.Tx) []Category {
	b := tx.Bucket(shared.CategoriesBucket)

	categories := []Category{}
	err := b.ForEach(func(_, v []byte) error {
		categories = append(categories, _makeCategory(v))
		return nil
	})
	if err != nil {
		panic(err)
	}

	return categories
}

func _makeCategory(data []byte) Category {
	var c Category
	err := json.Unmarshal(data, &c)
	if err != nil {
		panic(err)
	}
	return c
}

func _getProducts(tx *bbolt.Tx) []Product {
	b := tx.Bucket(shared.ProductsBucket)

	products := []Product{}
	err := b.ForEach(func(_, v []byte) error {
		products = append(products, _makeProduct(v))
		return nil
	})
	if err != nil {
		panic(err)
	}

	return products
}

func _makeProduct(data []byte) Product {
	var c Product
	err := json.Unmarshal(data, &c)
	if err != nil {
		panic(err)
	}
	return c
}
