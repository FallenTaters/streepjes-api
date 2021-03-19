package main

import (
	"net/http"

	"git.fuyu.moe/Fuyu/router"
	"github.com/PotatoesFall/streepjes/domain/catalog"
	"github.com/PotatoesFall/streepjes/domain/user"
	"github.com/PotatoesFall/streepjes/shared"
)

func postLogin(c *router.Context, credentials user.Credentials) error {
	role := user.LogIn(c.Response, credentials)
	if role == user.RoleNotAuthorized {
		return c.String(http.StatusUnauthorized, `invalid username or password`)
	}

	shared.SetCookie(c.Response, authCookieName, `token`, 5*60)
	return c.JSON(http.StatusOK, role)
}

func getCatalog(c *router.Context) error {
	cat := catalog.Get()
	return c.JSON(http.StatusOK, cat)
}

func getMembers(c *router.Context) error {
	return c.NoContent(http.StatusNotImplemented)
}

func postOrder(c *router.Context) error {
	return c.NoContent(http.StatusNotImplemented)
}
