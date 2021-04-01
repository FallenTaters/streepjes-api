package users

import (
	"time"

	"github.com/PotatoesFall/streepjes/shared"
)

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type User struct {
	ID           int         `json:"ID"`
	Club         shared.Club `json:"club"`
	Name         string      `json:"name"`
	Username     string      `json:"username"`
	Password     []byte      `json:"-"`
	Role         Role        `json:"role"`
	AuthToken    string      `json:"-"`
	AuthDatetime time.Time   `json:"-"`
}
