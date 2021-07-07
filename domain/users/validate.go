package users

import (
	"errors"

	"github.com/PotatoesFall/streepjes/shared"
)

var (
	ErrClubUnknown       = errors.New(`club unkown`)
	ErrNotAuthorized     = errors.New(`not authorized`)
	ErrEmptyPassword     = errors.New(`empty password`)
	ErrCannotChangeClub  = errors.New(`cannot change club`)
	ErrUserHasOpenOrders = errors.New(`cannot delete user with open orders`)
)

func validatePutUser(u User) (User, error) {
	user := User{
		Username: u.Username,
		Club:     u.Club,
		Name:     u.Name,
		Role:     u.Role,
		Password: u.Password,
	}

	original, err := get(u.Username)
	switch err {
	case nil:
		if u.Club != original.Club {
			return User{}, ErrCannotChangeClub
		}

	case ErrUserNotFound:
	default:
		panic(err)
	}

	if user.Club == shared.ClubUnknown {
		return User{}, ErrClubUnknown
	}

	if user.Role == RoleNotAuthorized {
		return User{}, ErrNotAuthorized
	}

	if len(user.Password) == 0 {
		return User{}, ErrEmptyPassword
	}

	return user, nil
}

func validateDeleteUser(username string) error {
	_, err := get(username)
	switch {
	case err == ErrUserNotFound:
		return err
	case err != nil:
		panic(err)
	}

	return nil
}
