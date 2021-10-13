package repo

import (
	"github.com/FallenTaters/streepjes-api/model"
	"github.com/FallenTaters/streepjes-api/repo/buckets"
	"github.com/FallenTaters/streepjes-api/shared"
)

type CategoryRepo interface {
	GetAll() ([]model.Category, error)
	Get(id int) (model.Category, error)
	Add(category model.Category) error
	Update(category model.Category) error
	Delete(id int) error
}

func NewCategoryRepo() Category {
	return Category{}
}

type Category struct{}

func (Category) GetAll() ([]model.Category, error) {
	categories := []model.Category{}
	return categories, buckets.Categories.GetAll(&model.Category{}, func(ptr interface{}) error {
		categories = append(categories, *ptr.(*model.Category))
		return nil
	})
}

func (Category) Get(id int) (model.Category, error) {
	var category model.Category
	return category, buckets.Categories.Get(shared.Itob(id), &category)
}

func (Category) Add(category model.Category) error {
	category.ID = buckets.Categories.NextSequence()
	return buckets.Categories.Create(category.Key(), category)
}

func (Category) Update(category model.Category) error {
	return buckets.Categories.Update(category.Key(), &model.Category{}, func(ptr interface{}) (object interface{}, err error) {
		return category, nil
	})
}

func (Category) Delete(id int) error {
	return buckets.Categories.Delete(shared.Itob(id))
}
