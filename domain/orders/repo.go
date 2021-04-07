package orders

import (
	"errors"

	"github.com/PotatoesFall/bbucket"
	"github.com/PotatoesFall/streepjes/domain/members"
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
	err := buckets.Orders.Create(o.Key(), o)
	if err != nil || o.MemberID == 0 {
		return err
	}

	return buckets.Members.Update(bbucket.Itob(o.MemberID), &members.Member{}, func(ptr interface{}) (object interface{}, err error) {
		member := *ptr.(*members.Member)

		member.Debt += o.Price

		return member, nil
	})
}

func deleteByID(id int) error {
	return buckets.Orders.Delete(bbucket.Itob(id))
}
