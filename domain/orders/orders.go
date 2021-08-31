package orders

import (
	"errors"
	"time"

	"github.com/FallenTaters/bbucket"
	"github.com/FallenTaters/streepjes-api/domain/users"
	"github.com/FallenTaters/streepjes-api/shared/null"
)

var (
	ErrOrderNotFound      = errors.New("order not found")
	ErrOrderAlreadyExists = errors.New("order already exists")
	ErrNoPermission       = errors.New("user not permitted to modify this order")
)

func Get(id int) (Order, error) {
	return get(id)
}

func Add(order Order, user users.User) error {
	order.Bartender = user.Username

	if err := validateAddOrder(order); err != nil {
		return err
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

		if filter.Month != nil && !sameMonthAndYear(*filter.Month, o.OrderTime) {
			return false
		}

		if len(filter.Status) == 0 {
			return true
		}

		for _, status := range filter.Status {
			if int(o.Status) == status {
				return true
			}
		}

		return false
	})
}

func sameMonthAndYear(t1 time.Time, t2 time.Time) bool {
	return t1.Month() == t2.Month() && t1.Year() == t2.Year()
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
		return order.Bartender == user.Username && order.IsEditable(), nil
	}

	return false, users.ErrUserNotFound
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
