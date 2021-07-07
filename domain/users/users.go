package users

import (
	"errors"
	"net/http"
	"time"

	"github.com/PotatoesFall/streepjes/shared"
	"golang.org/x/crypto/bcrypt"
)

const (
	tokenLength = 32
	loginTime   = 5 * time.Minute
)

var (
	ErrUserNotFound      = errors.New("user not found")
	ErrUserAlreadyExists = errors.New("user already exists")
	ErrInvalidLogin      = errors.New("invalid username or password")
)

func LogIn(w http.ResponseWriter, c Credentials) (User, error) {
	var u User
	return u, update(c.Username, func(user User) (User, error) {
		err := bcrypt.CompareHashAndPassword(user.Password, []byte(c.Password))
		if err != nil {
			return User{}, ErrInvalidLogin
		}

		user.AuthToken = shared.GenerateToken(tokenLength)
		user.AuthTime = time.Now()
		u = user

		return user, nil
	})
}

func LogOut(user User) error {
	return update(user.Username, func(user User) (User, error) {
		user.AuthTime = time.Now().Add(-loginTime)
		return user, nil
	})
}

func RefreshToken(user User) error {
	return update(user.Username, func(user User) (User, error) {
		user.AuthTime = time.Now()
		return user, nil
	})
}

func ValidateToken(username, token string) (User, bool) {
	user, err := get(username)
	if err != nil {
		return User{}, false
	}

	if time.Since(user.AuthTime) > loginTime || token != user.AuthToken {
		return User{}, false
	}

	return user, true
}

func Put(user User) error {
	user, err := validatePutUser(user)
	if err != nil {
		return err
	}

	return create(user)
}

func Delete(username string) error {
	err := validateDeleteUser(username)
	if err != nil {
		return err
	}

	return delete(username)
}

func MustGetByUsername(username string) User {
	user, err := get(username)
	if err != nil {
		panic(err)
	}

	return user
}

func GetAll() ([]UserPayload, error) {
	users, err := getAll()
	if err != nil {
		return nil, err
	}

	payloads := make([]UserPayload, len(users))
	for i, user := range users {
		payloads[i] = user.AsPayload()
	}

	return payloads, nil
}

func Get(username string) (User, error) {
	return get(username)
}
