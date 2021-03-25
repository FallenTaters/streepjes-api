package users

import "time"

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type User struct {
	ID           int       `json:"ID"`
	Name         string    `json:"name"`
	Username     string    `json:"username"`
	Password     []byte    `json:"-"`
	Role         Role      `json:"role"`
	AuthToken    string    `json:"auth_token"`
	AuthDatetime time.Time `json:"auth_datetime"`
}
