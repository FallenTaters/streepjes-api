package orders

import (
	"errors"

	"github.com/PotatoesFall/bbucket"
	"github.com/PotatoesFall/streepjes/shared/buckets"
)

var (
	ErrOrderNotFound      = errors.New("order not found")
	ErrOrderAlreadyExists = errors.New("order already exists")
)

func filtered(filterFunc func(Order) bool) ([]Order, error) {
	orders := []Order{}

	return orders, buckets.Orders.GetAll(&Order{}, func(ptr interface{}) error {
		o := *ptr.(*Order)
		if filterFunc(o) {
			orders = append(orders, o)
		}
		return nil
	})
}

func create(o Order) error {
	o.ID = buckets.Orders.NextSequence()
	return buckets.Orders.Create(o.Key(), o)
}

func deleteByID(id int) error {
	return buckets.Orders.Delete(bbucket.Itob(id))
}
