package repo

import (
	"github.com/FallenTaters/streepjes-api/model"
	"github.com/FallenTaters/streepjes-api/repo/buckets"
	"github.com/FallenTaters/streepjes-api/shared"
)

type ProductRepo interface {
	GetAll() ([]model.Product, error)
	Get(id int) (model.Product, error)
	Add(product model.Product) error
	Update(product model.Product) error
	Delete(id int) error
}

func NewProductRepo() Product {
	return Product{}
}

type Product struct{}

func (Product) GetAll() ([]model.Product, error) {
	products := []model.Product{}
	return products, buckets.Products.GetAll(&model.Product{}, func(ptr interface{}) error {
		products = append(products, *ptr.(*model.Product))
		return nil
	})
}

func (Product) Get(id int) (model.Product, error) {
	var product model.Product
	return product, buckets.Products.Get(shared.Itob(id), &product)
}

func (Product) Add(product model.Product) error {
	product.ID = buckets.Products.NextSequence()
	return buckets.Products.Create(product.Key(), product)
}

func (Product) Update(product model.Product) error {
	return buckets.Products.Update(product.Key(), &model.Product{}, func(ptr interface{}) (object interface{}, err error) {
		return product, nil
	})
}

func (Product) Delete(id int) error {
	return buckets.Products.Delete(shared.Itob(id))
}
