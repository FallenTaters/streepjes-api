package users

import (
	"encoding/json"
	"errors"

	"github.com/PotatoesFall/streepjes/shared"
	"go.etcd.io/bbolt"
)

var db *bbolt.DB

var (
	ErrUserNotFound      = errors.New("user not found")
	ErrUserAlreadyExists = errors.New("user already exists")
)

func getAll() []User {
	users := []User{}

	err := db.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket(shared.UsersBucket)
		return b.ForEach(func(_, v []byte) error {
			users = append(users, _unmarshal(v))
			return nil
		})
	})
	if err != nil {
		panic(err)
	}

	return users
}

func get(username string) (User, error) {
	var user User

	err := db.View(func(tx *bbolt.Tx) error {
		var err error
		user, err = _get(tx, username)
		return err
	})

	return user, err
}

func create(user User) error {
	return db.Update(func(tx *bbolt.Tx) error {
		_, err := _get(tx, user.Username)
		if err == nil {
			return ErrUserAlreadyExists
		} else if err != ErrUserNotFound {
			return err
		}

		return _put(tx, user)
	})
}

func update(username string, mutate func(User) (User, error)) error {
	return db.Update(func(tx *bbolt.Tx) error {
		user, err := _get(tx, username)
		if err != nil {
			return err
		}

		newUser, err := mutate(user)
		if err != nil {
			return err
		}
		if newUser.Username != user.Username {
			panic("key change not implemented!")
		}

		return _put(tx, newUser)
	})
}

func _get(tx *bbolt.Tx, username string) (User, error) {
	b := tx.Bucket(shared.UsersBucket)

	data := b.Get([]byte(username))
	if data == nil {
		return User{}, ErrUserNotFound
	}

	return _unmarshal(data), nil
}

func _put(tx *bbolt.Tx, user User) error {
	b := tx.Bucket(shared.UsersBucket)
	return b.Put([]byte(user.Username), _marshal(user))
}

func _unmarshal(data []byte) User {
	var u User
	err := json.Unmarshal(data, &u)
	if err != nil {
		panic(err)
	}
	return u
}

func _marshal(user User) []byte {
	data, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}
	return data
}
