package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"git.fuyu.moe/Fuyu/router"
	"github.com/PotatoesFall/bbucket"
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

func getClub(c *router.Context) error {
	user := c.Get("user").(users.User)
	return c.JSON(http.StatusOK, user.Club)
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
	cat, err := catalog.Get()
	if err != nil {
		panic(err)
	}
	return c.JSON(http.StatusOK, cat)
}

func getMembers(c *router.Context) error {
	members, err := members.GetAll()
	if err != nil {
		panic(err)
	}

	return c.JSON(http.StatusOK, members)
}

func postOrder(c *router.Context, order orders.Order) error {
	order.Bartender = getUserFromContext(c).Username
	if order.Status != orders.OrderStatusOpen && order.Status != orders.OrderStatusPaid {
		return c.String(http.StatusBadRequest, `Status must be "Open" or "Paid".`)
	}

	err := orders.AddOrder(order)
	if err != nil {
		panic(err)
	}

	return c.NoContent(http.StatusOK)
}

func getUsers(c *router.Context) error {
	users, err := users.GetAll()
	if err != nil {
		panic(err)
	}

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

	orders, err := orders.Filter(filter)
	if err != nil {
		panic(err)
	}

	return c.JSON(http.StatusOK, orders)
}

func postOrderDelete(c *router.Context) error {
	id, err := strconv.Atoi(c.Param(`id`))
	if err != nil {
		return c.StatusText(http.StatusBadRequest)
	}

	allowed, err := orders.HasPermissions(id, c.Get("user").(users.User))
	switch {
	case err == nil && allowed:
		err = orders.Delete(id)
		if err != nil {
			panic(err)
		}
		return c.StatusText(http.StatusOK)

	case err == nil && !allowed:
		return c.StatusText(http.StatusUnauthorized)

	case err == bbucket.ErrObjectNotFound:
		return c.StatusText(http.StatusNotFound)

	default:
		panic(err)
	}
}

func postProduct(c *router.Context, product catalog.Product) error {
	err := catalog.PutProduct(product)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	return c.StatusText(http.StatusOK)
}

func postCategory(c *router.Context, category catalog.Category) error {
	err := catalog.PutCategory(category)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	return c.StatusText(http.StatusOK)
}

func postProductDelete(c *router.Context) error {
	id, err := strconv.Atoi(c.Param(`id`))
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	err = catalog.DeleteProduct(id)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	return c.StatusText(http.StatusOK)
}

func postCategoryDelete(c *router.Context) error {
	id, err := strconv.Atoi(c.Param(`id`))
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	err = catalog.DeleteCategory(id)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	return c.StatusText(http.StatusOK)
}
