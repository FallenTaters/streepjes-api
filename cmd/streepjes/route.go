package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"git.fuyu.moe/Fuyu/router"
	"github.com/FallenTaters/bbucket"
	"github.com/FallenTaters/streepjes-api/domain/catalog"
	"github.com/FallenTaters/streepjes-api/domain/members"
	"github.com/FallenTaters/streepjes-api/domain/orders"
	"github.com/FallenTaters/streepjes-api/domain/streepjes"
	"github.com/FallenTaters/streepjes-api/domain/users"
	"github.com/FallenTaters/streepjes-api/shared"
	"github.com/FallenTaters/streepjes-api/shared/null"
)

func startServer() {
	r := router.New()

	r.Use(corsMiddleware)
	r.ErrorHandler = errorHandler
	r.Reader = reader

	r.POST(`/login`, postLogin)

	au := r.Group(`/`, authMiddleware)
	au.POST(`/logout`, postLogout)
	au.POST(`/active`, postActive)
	au.GET(`/catalog`, getCatalog)
	au.GET(`/members`, getMembers)
	au.POST(`/order`, postOrder)
	au.GET(`/orders`, getOrders)
	au.POST(`/order/delete/:id`, postOrderDelete)
	au.GET(`/club`, getClub)

	ad := au.Group(`/`, roleMiddleware(users.RoleAdmin))
	ad.GET(`/users`, getUsers)
	ad.POST(`/user`, postUser)
	ad.POST(`/user/delete/:username`, postUserDelete)

	ad.POST(`/member`, postMember)
	ad.POST(`/member/delete/:id`, postMemberDelete)

	ad.POST(`/category`, postCategory)
	ad.POST(`/category/delete/:id`, postCategoryDelete)
	ad.POST(`/product`, postProduct)
	ad.POST(`/product/delete/:id`, postProductDelete)

	ad.GET(`/orders/:year/:month`, getOrdersByMonth)
	ad.GET(`/orders/:year/:month/csv`, getOrdersByMonthCSV)

	panic(r.Start(`:` + settings.Port))
}

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

	err := users.Delete(username)
	switch err {
	case nil:
		return c.StatusText(http.StatusOK)

	case users.ErrUserHasOpenOrders, users.ErrUserNotFound:
		return c.String(http.StatusBadRequest, err.Error())
	}

	panic(err)
}

func getOrdersByMonth(c *router.Context) error {
	year, errYear := strconv.Atoi(c.Param(`year`))
	month, errMonth := strconv.Atoi(c.Param(`month`))
	if errYear != nil || errMonth != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	date := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.Local)
	orders, err := orders.Filter(orders.OrderFilter{
		Month: &date,
		Club:  null.NewInt(getUserFromContext(c).Club.Int()),
	})
	if err != nil {
		panic(err)
	}

	return c.JSON(http.StatusOK, orders)
}

func getOrdersByMonthCSV(c *router.Context) error {
	year, errYear := strconv.Atoi(c.Param(`year`))
	month, errMonth := strconv.Atoi(c.Param(`month`))
	if errYear != nil || errMonth != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	date := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.Local)
	ordersByMonth, err := orders.Filter(orders.OrderFilter{
		Month: &date,
		Club:  null.NewInt(getUserFromContext(c).Club.Int()),
	})
	if err != nil {
		panic(err)
	}

	data, err := orders.MakeCSV(ordersByMonth)
	if err != nil {
		panic(err)
	}

	filename := fmt.Sprintf(`orders-%s-%d.csv`, date.Month().String(), date.Year())

	c.Response.Header().Set(`Content-Disposition`, `attachment; filename=`+filename)
	http.ServeContent(c.Response, c.Request, filename, time.Now(), bytes.NewReader(data))

	return nil
}
