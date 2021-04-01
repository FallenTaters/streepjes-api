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

func LogOut(userID int) {
	removeToken(userID)
}

func RefreshToken(username string) {
	refreshToken(username)
}

func ValidateToken(username, token string) (User, bool) {
	user, err := getUserByUsername(username)
	if err != nil {
		return User{}, false
	}

	if time.Since(user.AuthDatetime) > loginTime || token != user.AuthToken {
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

	return insert(user)
}

func MustGetByUsername(username string) User {
	user, err := getUserByUsername(username)
	if err != nil {
		panic(err)
	}
	return user
}

func GetAll() []User {
	return getAll()
}
