package users

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/PotatoesFall/streepjes/shared"
	"golang.org/x/crypto/bcrypt"
)

const (
	tokenLength = 32
	loginTime   = 5 * time.Minute
)

func Init(database *sql.DB) {
	db = database
}

func LogIn(w http.ResponseWriter, c Credentials) User {
	user, err := getUserByUsername(c.Username)
	if err != nil {
		return User{Role: RoleNotAuthorized}
	}

	err = bcrypt.CompareHashAndPassword(user.Password, []byte(c.Password))
	if err != nil {
		return User{Role: RoleNotAuthorized}
	}

	user.AuthToken = shared.GenerateToken(tokenLength)
	err = setToken(user)
	if err != nil {
		panic(err)
	}

	return user
}

func LogOut(username string) {
	removeToken(username)
}

func RefreshToken(username string) {
	refreshToken(username)
}

func ValidateToken(username, token string) bool {
	user, err := getUserByUsername(username)
	if err != nil {
		return false
	}

	if time.Since(user.AuthDatetime) > loginTime || token != user.AuthToken {
		return false
	}

	return true
}

func Insert(name, username, password string, role Role) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	return insert(name, username, hash, role)
}

func MustGetByUsername(username string) User {
	user, err := getUserByUsername(username)
	if err != nil {
		panic(err)
	}
	return user
}
