package main

import (
	"encoding/json"
	"net/http"

	"git.fuyu.moe/Fuyu/router"
	"github.com/PotatoesFall/streepjes/domain/user"
	"github.com/PotatoesFall/streepjes/shared"
)

const authCookieName = `auth_token`

type authCookie struct {
	Username  string `json:"username"`
	AuthToken string `json:"auth_token"`
}

func authMiddleware(next router.Handle) router.Handle {
	return func(c *router.Context) error {
		cookieValue, ok := shared.GetCookie(c.Request, authCookieName)
		if !ok {
			return c.StatusText(http.StatusUnauthorized)
		}

		var cookie authCookie
		err := json.Unmarshal([]byte(cookieValue), &cookie)

		if err != nil || !user.ValidateToken(cookie.Username, cookie.AuthToken) {
			shared.UnsetCookie(c.Response, authCookieName)
			return c.StatusText(http.StatusUnauthorized)
		}

		return next(c)
	}
}
