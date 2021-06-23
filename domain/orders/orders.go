package orders

import (
	"errors"

	"github.com/PotatoesFall/bbucket"
	"github.com/PotatoesFall/streepjes/domain/users"
	"github.com/PotatoesFall/streepjes/shared/null"
)

var (
	ErrOrderNotFound       = errors.New("order not found")
	ErrOrderAlreadyExists  = errors.New("order already exists")
	ErrNoPermission        = errors.New("user not permitted to modify this order")
	ErrStatusNotOpenOrPaid = errors.New("order status must be open or paid")
)

func Get(id int) (Order, error) {
	return get(id)
}

func Add(order Order, user users.User) error {
	order.Bartender = user.Username
	if order.Status != OrderStatusOpen && order.Status != OrderStatusPaid {
		return ErrStatusNotOpenOrPaid
	}

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

func Delete(id int, user users.User) error {
	allowed, err := HasPermissions(id, user)
	switch {
	case err == nil && allowed:
		return deleteByID(id)

	case err == nil && !allowed:
		return ErrNoPermission

	case err == bbucket.ErrObjectNotFound:
		return ErrOrderNotFound

	default:
		panic(err)
	}
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

func GetForUser(user users.User) ([]Order, error) {
	filter := OrderFilter{}

	switch user.Role {
	case users.RoleAdmin:
		filter.Club = null.NewInt(user.Club.Int())

	case users.RoleBartender:
		filter.Bartender = null.NewString(user.Username)

	default:
		return nil, ErrNoPermission
	}

	return Filter(filter)
}
