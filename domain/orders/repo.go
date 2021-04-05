package orders

import (
	"errors"

	"github.com/PotatoesFall/bbucket"
)

var ordersBucket bbucket.Bucket

var (
	ErrOrderNotFound      = errors.New("order not found")
	ErrOrderAlreadyExists = errors.New("order already exists")
)

func filtered(filterFunc func(Order) bool) []Order {
	orders := []Order{}

	err := ordersBucket.GetAll(&Order{}, func(ptr interface{}) {
		o := *ptr.(*Order)
		if filterFunc(o) {
			orders = append(orders, o)
		}
	})
	if err != nil {
		panic(err)
	}

	return orders
}

func create(o Order) error {
	o.ID = ordersBucket.NextSequence()
	return ordersBucket.Create(o)
}
