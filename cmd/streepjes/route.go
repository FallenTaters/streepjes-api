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
	"github.com/PotatoesFall/streepjes/shared/null"
)

func postActive(c *router.Context) error {
	return c.NoContent(http.StatusOK)
}

func postLogin(c *router.Context, credentials users.Credentials) error {
	user, err := users.LogIn(c.Response, credentials)
	if err != nil || user.Role == users.RoleNotAuthorized {
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
	err := users.LogOut(getUserFromContext(c))
	if err != nil {
		panic(err)
	}
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
	order.Bartender = getUserFromContext(c).Username
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

func getUsers(c *router.Context) error {
	users := users.GetAll()
	return c.JSON(http.StatusOK, users)
}

func getOrders(c *router.Context) error {
	user := getUserFromContext(c)
	filter := orders.OrderFilter{}

	switch user.Role {
	case users.RoleAdmin:
		filter.Club = null.NewInt(user.Club.Int())
	case users.RoleBartender:
		filter.Bartender = null.NewString(user.Username)
	default:
		return c.StatusText(http.StatusUnauthorized)
	}

	orders := orders.Filter(filter)
	return c.JSON(http.StatusOK, orders)
}
