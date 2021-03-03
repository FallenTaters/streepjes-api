package user

import (
	"net/http"
)

const AuthCookie = `auth_token`

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func LogIn(w http.ResponseWriter, c Credentials) Role {
	switch c {
	case Credentials{Username: "admin", Password: "password"}:
		return RoleAdmin
	case Credentials{Username: "bartender", Password: "password"}:
		return RoleBartender
	}
	return RoleNotAuthorized
}
