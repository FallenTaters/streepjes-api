package orders

import (
	"github.com/PotatoesFall/streepjes/domain/users"
	"github.com/PotatoesFall/streepjes/shared/null"
)

func Get(id int) (Order, error) {
	return get(id)
}

func AddOrder(order Order) error {
	return create(order)
}

func Filter(filter OrderFilter) ([]Order, error) {
	return filtered(func(o Order) bool {
		if filter.Club.Valid && filter.Club.Int64 != int64(o.Club) {
			return false
		}

		if filter.Bartender.Valid && filter.Bartender.String != o.Bartender {
			return false
		}

		if filter.Member.Valid && filter.Member.Int64 != int64(o.MemberID) {
			return false
		}

		if filter.Status.Valid && filter.Status.Int64 != int64(o.Status) {
			return false
		}

		return true
	})
}

func Delete(id int) error {
	return deleteByID(id)
}

func HasPermissions(id int, user users.User) (bool, error) {
	order, err := get(id)
	if err != nil {
		return false, err
	}

	switch user.Role {
	case users.RoleAdmin:
		return true, nil
	case users.RoleBartender:
		editable := order.MemberID == 0 || (order.Status == OrderStatusOpen)
		return order.Bartender == user.Username && editable, nil
	}

	return false, users.ErrUserNotFound
}

func MemberHasUnpaidOrders(id int) bool {
	orders, err := Filter(OrderFilter{
		Member: null.NewInt(id),
	})
	if err != nil {
		panic(err)
	}

	for _, o := range orders {
		if o.Status == OrderStatusOpen {
			return true
		}
	}

	return false
}
