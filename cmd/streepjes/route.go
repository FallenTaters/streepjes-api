package main

import (
	"encoding/json"
	"net/http"
	"time"

	"git.fuyu.moe/Fuyu/router"
	"github.com/PotatoesFall/streepjes/domain/catalog"
	"github.com/PotatoesFall/streepjes/domain/members"
	"github.com/PotatoesFall/streepjes/domain/orders"
	"github.com/PotatoesFall/streepjes/domain/users"
	"github.com/PotatoesFall/streepjes/shared"
)

func postActive(c *router.Context) error {
	return c.NoContent(http.StatusOK)
}

func postLogin(c *router.Context, credentials users.Credentials) error {
	user := users.LogIn(c.Response, credentials)
	if user.Role == users.RoleNotAuthorized {
		return c.String(http.StatusUnauthorized, `invalid username or password`)
	}

	cookieValue, err := json.Marshal(authCookie{Username: user.Username, AuthToken: user.AuthToken})
	if err != nil {
		panic(err)
	}

	shared.SetCookie(c.Response, authCookieName, cookieValue, authCookieDuration)
	return c.JSON(http.StatusOK, user.Role)
}

func postLogout(c *router.Context) error {
	shared.UnsetCookie(c.Response, authCookieName)
	users.LogOut(getUsernameFromContext(c))
	return c.NoContent(http.StatusOK)
}

func getCatalog(c *router.Context) error {
	cat := catalog.Get()
	return c.JSON(http.StatusOK, cat)
}

func getMembers(c *router.Context) error {
	members := members.GetAll()
	return c.JSON(http.StatusOK, members)
}

func postOrder(c *router.Context, order orders.Order) error {
	order.BartenderID = users.MustGetByUsername(getUsernameFromContext(c)).ID
	order.OrderTime = time.Now()
	if order.Status != orders.OrderStatusOpen && order.Status != orders.OrderStatusPaid {
		return c.String(http.StatusBadRequest, `Status must be "Open" or "Paid".`)
	}
	order.StatusTime = time.Now()

	err := orders.AddOrder(order)
	if err != nil {
		panic(err)
	}

	return c.NoContent(http.StatusOK)
}
