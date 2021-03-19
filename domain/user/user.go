package user

import (
	"database/sql"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func Init(database *sql.DB) {
	db = database
}

func LogIn(w http.ResponseWriter, c Credentials) Role {
	user, err := getUserByUsername(c.Username)
	if err != nil {
		return RoleNotAuthorized
	}
	err = bcrypt.CompareHashAndPassword(user.Password, []byte(c.Password))
	if err != nil {
		return RoleNotAuthorized
	}

	return user.Role
}

func ValidateToken(username, token string) bool {
	panic(`not implemented`)
}

func Insert(name, username, password string, role Role) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	return insert(name, username, hash, role)
}
