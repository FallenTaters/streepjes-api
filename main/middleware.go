package main

import (
	"net/http"

	"git.fuyu.moe/Fuyu/router"
	"github.com/PotatoesFall/streepjes/domain/user"
	"github.com/PotatoesFall/streepjes/shared"
)

func authMiddleware(next router.Handle) router.Handle {
	return func(c *router.Context) error {
		cookie, ok := shared.GetCookie(c.Request, user.AuthCookie)
		if !ok {
			return c.StatusText(http.StatusUnauthorized)
		}

		if cookie != `token` {
			shared.UnsetCookie(c.Response, user.AuthCookie)
			return c.StatusText(http.StatusUnauthorized)
		}

		return next(c)
	}
}
