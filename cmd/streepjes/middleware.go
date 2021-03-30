package main

import (
	"encoding/json"
	"net/http"

	"git.fuyu.moe/Fuyu/router"
	"github.com/PotatoesFall/streepjes/domain/users"
	"github.com/PotatoesFall/streepjes/shared"
)

const (
	authCookieName     = `auth_token`
	authCookieDuration = 5 * 60
)

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

		if err != nil || !users.ValidateToken(cookie.Username, cookie.AuthToken) {
			shared.UnsetCookie(c.Response, authCookieName)
			return c.StatusText(http.StatusUnauthorized)
		}

		// refresh cookie duration
		shared.SetCookie(c.Response, authCookieName, cookieValue, authCookieDuration)
		users.RefreshToken(cookie.Username)
		c.Set(`username`, cookie.Username)

		return next(c)
	}
}

func getUsernameFromContext(c *router.Context) string {
	username, ok := c.Get(`username`).(string)
	if !ok {
		return ``
	}
	return username
}

func corsMiddleware(next router.Handle) router.Handle {
	return func(c *router.Context) error {
		c.Response.Header().Set(`Access-Control-Allow-Origin`, c.Request.Header.Get(`Origin`))
		c.Response.Header().Set(`Access-Control-Allow-Credentials`, `true`)
		return next(c)
	}
}
