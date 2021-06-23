package orders

import (
	"time"

	"github.com/PotatoesFall/bbucket"
	"github.com/PotatoesFall/streepjes/domain/members"
	"github.com/PotatoesFall/streepjes/shared/buckets"
)

func get(id int) (Order, error) {
	var o Order
	return o, buckets.Orders.Get(bbucket.Itob(id), &o)
}

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
	o.OrderTime = time.Now()
	o.StatusTime = o.OrderTime

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
	var o Order

	err := buckets.Orders.Update(bbucket.Itob(id), &Order{}, func(ptr interface{}) (object interface{}, err error) {
		order := *ptr.(*Order)
		o = order

		order.Status = OrderStatusCancelled
		order.StatusTime = time.Now()

		return order, nil
	})
	if err != nil || o.MemberID == 0 || o.Status != OrderStatusOpen {
		return err
	}

	return buckets.Members.Update(bbucket.Itob(o.MemberID), &members.Member{}, func(ptr interface{}) (object interface{}, err error) {
		member := *ptr.(*members.Member)

		member.Debt -= o.Price

		return member, nil
	})
}
