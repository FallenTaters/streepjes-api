package orders

import (
	"encoding/json"
	"errors"

	"github.com/PotatoesFall/streepjes/shared"
	"go.etcd.io/bbolt"
)

var db *bbolt.DB

var (
	ErrOrderNotFound      = errors.New("order not found")
	ErrOrderAlreadyExists = errors.New("order already exists")
)

func filtered(filterFunc func(Order) bool) []Order {
	orders := []Order{}

	err := db.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket(shared.OrdersBucket)
		return b.ForEach(func(_, v []byte) error {
			order := _unmarshal(v)
			if filterFunc(order) {
				orders = append(orders, order)
			}

			return nil
		})
	})
	if err != nil {
		panic(err)
	}

	return orders
}

func create(o Order) error {
	return db.Update(func(tx *bbolt.Tx) error {
		_, err := _get(tx, o.ID)
		if err == nil {
			return ErrOrderAlreadyExists
		} else if err != ErrOrderNotFound {
			return err
		}

		return _put(tx, o)
	})
}

func _get(tx *bbolt.Tx, id int) (Order, error) {
	b := tx.Bucket(shared.OrdersBucket)

	data := b.Get(shared.Itob(id))
	if data == nil {
		return Order{}, ErrOrderNotFound
	}

	return _unmarshal(data), nil
}

func _put(tx *bbolt.Tx, order Order) error {
	b := tx.Bucket(shared.OrdersBucket)
	return b.Put(shared.Itob(order.ID), _marshal(order))
}

func _unmarshal(data []byte) Order {
	var o Order
	err := json.Unmarshal(data, &o)
	if err != nil {
		panic(err)
	}
	return o
}

func _marshal(o Order) []byte {
	data, err := json.Marshal(o)
	if err != nil {
		panic(err)
	}
	return data
}
