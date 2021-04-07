package users

import (
	"net/http"
	"time"

	"github.com/PotatoesFall/streepjes/shared"
	"golang.org/x/crypto/bcrypt"
)

const (
	tokenLength = 32
	loginTime   = 5 * time.Minute
)

func LogIn(w http.ResponseWriter, c Credentials) (User, error) {
	var u User
	return u, update(c.Username, func(user User) (User, error) {
		err := bcrypt.CompareHashAndPassword(user.Password, []byte(c.Password))
		if err != nil {
			return User{}, err
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

func Insert(user User) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.Password = hash

	return create(user)
}

func MustGetByUsername(username string) User {
	user, err := get(username)
	if err != nil {
		panic(err)
	}

	return user
}

func GetAll() ([]User, error) {
	return getAll()
}
