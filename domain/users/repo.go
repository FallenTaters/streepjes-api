package users

import (
	"errors"

	"github.com/PotatoesFall/streepjes/shared/buckets"
)

var (
	ErrUserNotFound      = errors.New("user not found")
	ErrUserAlreadyExists = errors.New("user already exists")
)

func getAll() ([]User, error) {
	users := []User{}

	return users, buckets.Users.GetAll(&User{}, func(ptr interface{}) error {
		users = append(users, *ptr.(*User))
		return nil
	})
}

func get(username string) (User, error) {
	var user User

	return user, buckets.Users.Get([]byte(username), &user)
}

func create(user User) error {
	return buckets.Users.Create(user.Key(), user)
}

func update(username string, mutate func(User) (User, error)) error {
	return buckets.Users.Update([]byte(username), &User{}, func(ptr interface{}) (interface{}, error) {
		return mutate(*ptr.(*User))
	})
}
