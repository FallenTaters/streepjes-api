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
	if err != nil || bcrypt.CompareHashAndPassword(user.Password, []byte(c.Password)) != nil {
		return RoleNotAuthorized
	}
	return user.Role
}

func ValidateToken(username, token string) bool {
	panic(`not implemented`)
}
