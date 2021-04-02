package orders

import (
	"go.etcd.io/bbolt"
)

func Init(database *bbolt.DB) {
	db = database
}

func AddOrder(order Order) error {
	return create(order)
}

func Get(filter Filter) []Order {
	return filtered(func(o Order) bool {
		if filter.Club.Valid && filter.Club.Int64 != int64(o.Club) {
			return false
		}
		if filter.Bartender.Valid && filter.Bartender.String != o.Bartender {
			return false
		}
		return true
	})
}
