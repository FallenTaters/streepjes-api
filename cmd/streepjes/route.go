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
	"github.com/PotatoesFall/streepjes/domain/streepjes"
	"github.com/PotatoesFall/streepjes/domain/users"
	"github.com/PotatoesFall/streepjes/shared"
)

func postActive(c *router.Context) error {
	return c.NoContent(http.StatusOK)
}

func getClub(c *router.Context) error {
	return c.JSON(http.StatusOK, getUserFromContext(c).Club)
}

func postLogin(c *router.Context, credentials users.Credentials) error {
	user, err := users.LogIn(c.Response, credentials)
	switch {
	case err == users.ErrInvalidLogin || user.Role == users.RoleNotAuthorized:
		return c.StatusText(http.StatusUnauthorized)

	case err != nil:
		panic(err)
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
	err := orders.Add(order, getUserFromContext(c))
	switch err {
	case nil:
		return c.NoContent(http.StatusOK)

	case users.ErrUserNotFound, members.ErrMemberNotFound, orders.ErrStatusNotOpenOrPaid, orders.ErrUnknownClub, orders.ErrNoContents:
		return c.String(http.StatusBadRequest, err.Error())
	}

	panic(err)
}

func getUsers(c *router.Context) error {
	u, err := users.GetAll()
	if err != nil {
		panic(err)
	}

	return c.JSON(http.StatusOK, u)
}

func getOrders(c *router.Context) error {
	o, err := orders.GetForUser(getUserFromContext(c))
	switch err {
	case nil:
		return c.JSON(http.StatusOK, o)

	case orders.ErrNoPermission:
		return c.String(http.StatusBadRequest, err.Error())
	}

	panic(err)
}

func postOrderDelete(c *router.Context) error {
	id, err := strconv.Atoi(c.Param(`id`))
	if err != nil {
		return c.StatusText(http.StatusBadRequest)
	}

	err = orders.Delete(id, getUserFromContext(c))
	switch err {
	case nil:
		return c.StatusText(http.StatusOK)

	case orders.ErrNoPermission, orders.ErrOrderNotFound:
		return c.String(http.StatusBadRequest, err.Error())
	}

	panic(err)
}

func postProduct(c *router.Context, product catalog.Product) error {
	err := catalog.PutProduct(product)
	switch err {
	case nil:
		return c.StatusText(http.StatusOK)

	case catalog.ErrEmptyName, catalog.ErrNoPrice, catalog.ErrCategoryNotFound, catalog.ErrNameTaken, catalog.ErrProductNotFound:
		return c.String(http.StatusBadRequest, err.Error())
	}

	panic(err)
}

func postCategory(c *router.Context, category catalog.Category) error {
	err := catalog.PutCategory(category)
	switch err {
	case nil:
		return c.StatusText(http.StatusOK)

	case catalog.ErrEmptyName, catalog.ErrNameTaken:
		return c.String(http.StatusBadRequest, err.Error())
	}

	panic(err)
}

func postProductDelete(c *router.Context) error {
	id, err := strconv.Atoi(c.Param(`id`))
	if err != nil {
		return c.StatusText(http.StatusUnprocessableEntity)
	}

	err = catalog.DeleteProduct(id)
	switch err {
	case nil:
		return c.StatusText(http.StatusOK)

	case bbucket.ErrObjectNotFound:
		return c.StatusText(http.StatusNotFound)
	}

	panic(err)
}

func postCategoryDelete(c *router.Context) error {
	id, err := strconv.Atoi(c.Param(`id`))
	if err != nil {
		return c.StatusText(http.StatusUnprocessableEntity)
	}

	err = catalog.DeleteCategory(id)
	switch err {
	case nil:
		return c.StatusText(http.StatusOK)

	case catalog.ErrCategoryHasProduct, catalog.ErrCategoryNotFound:
		return c.String(http.StatusBadRequest, err.Error())
	}

	panic(err)
}

func postMember(c *router.Context, member members.Member) error {
	err := members.Put(member)
	switch err {
	case nil:
		return c.StatusText(http.StatusOK)

	case members.ErrEmptyName, members.ErrNameTaken, members.ErrUnknownClub, members.ErrMemberNotFound:
		return c.String(http.StatusBadRequest, err.Error())
	}

	panic(err)
}

func postMemberDelete(c *router.Context) error {
	id, err := strconv.Atoi(c.Param(`id`))
	if err != nil {
		return c.StatusText(http.StatusUnprocessableEntity)
	}

	err = streepjes.DeleteMember(id)
	switch err {
	case nil:
		return c.StatusText(http.StatusOK)

	case members.ErrUnpaidOrders, members.ErrMemberNotFound:
		return c.String(http.StatusBadRequest, err.Error())
	}

	panic(err)
}

func postUser(c *router.Context, user users.User) error {
	if user.Username == getUserFromContext(c).Username {
		return c.String(http.StatusBadRequest, users.ErrCannotChangeOwnAccount.Error())
	}

	err := users.Put(user)
	switch err {
	case nil:
		return c.StatusText(http.StatusOK)

	case users.ErrClubUnknown, users.ErrNotAuthorized, users.ErrEmptyPassword, users.ErrCannotChangeClub:
		return c.String(http.StatusBadRequest, err.Error())
	}

	panic(err)
}

func postUserDelete(c *router.Context) error {
	username := c.Param(`username`)
	if username == getUserFromContext(c).Username {
		return c.String(http.StatusBadRequest, users.ErrCannotChangeOwnAccount.Error())
	}

	err := streepjes.DeleteUser(username)
	switch err {
	case nil:
		return c.StatusText(http.StatusOK)

	case users.ErrUserHasOpenOrders, users.ErrUserNotFound:
		return c.String(http.StatusBadRequest, err.Error())
	}

	panic(err)
}
