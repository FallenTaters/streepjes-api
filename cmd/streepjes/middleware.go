package main

import (
	"encoding/json"
	"net/http"

	"git.fuyu.moe/Fuyu/router"
	"github.com/FallenTaters/streepjes-api/domain/users"
	"github.com/FallenTaters/streepjes-api/shared/cookies"
)

const (
	authCookieName     = `auth_token`
	authCookieDuration = 5 * 60
)

type authCookie struct {
	Username  string `json:"username"`
	AuthToken string `json:"auth_token"`
}

func roleMiddleware(role users.Role) router.Middleware {
	return func(next router.Handle) router.Handle {
		return func(c *router.Context) error {
			user := getUserFromContext(c)
			if user.Role != role {
				return c.String(http.StatusUnauthorized, `admin role required`)
			}

			return next(c)
		}
	}
}

func authMiddleware(next router.Handle) router.Handle {
	return func(c *router.Context) error {
		cookieValue, ok := cookies.Get(c.Request, authCookieName)
		if !ok {
			return c.StatusText(http.StatusUnauthorized)
		}

		var cookie authCookie
		err := json.Unmarshal([]byte(cookieValue), &cookie)
		if err != nil {
			return authFail(c)
		}

		user, ok := users.ValidateToken(cookie.Username, cookie.AuthToken)
		if !ok {
			return authFail(c)
		}

		c.Set(`user`, user)

		// refresh cookie duration
		cookies.Set(c.Response, authCookieName, cookieValue, authCookieDuration)
		err = users.RefreshToken(users.User{Username: cookie.Username})
		if err != nil {
			panic(err)
		}

		return next(c)
	}
}

func authFail(c *router.Context) error {
	cookies.Unset(c.Response, authCookieName)
	return c.StatusText(http.StatusUnauthorized)
}

func getUserFromContext(c *router.Context) users.User {
	user, ok := c.Get(`user`).(users.User)
	if !ok {
		panic(`user not found on context`)
	}
	return user
}

func corsMiddleware(next router.Handle) router.Handle {
	return func(c *router.Context) error {
		c.Response.Header().Set(`Access-Control-Allow-Origin`, c.Request.Header.Get(`Origin`))
		c.Response.Header().Set(`Access-Control-Allow-Credentials`, `true`)
		return next(c)
	}
}
