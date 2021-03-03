package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"git.fuyu.moe/Fuyu/router"
	"github.com/PotatoesFall/streepjes/domain/user"
)

func main() {
	readSettings()

	r := router.New()

	r.ErrorHandler = errorHandler
	r.Reader = reader

	r.POST(`/login`, login)

	a := r.Group(`/`, authMiddleware)
	a.GET(`/order`, getOrder)

	panic(r.Start(`:` + Settings.Port))
}

func errorHandler(c *router.Context, v interface{}) {
	fmt.Printf("panic: %#v\n", v)
}

func reader(c *router.Context, dst interface{}) (bool, error) {
	defer c.Request.Body.Close()
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		return false, c.String(http.StatusBadRequest, err.Error())
	}

	err = json.Unmarshal(body, dst)
	if err != nil {
		return false, c.String(http.StatusBadRequest, err.Error())
	}

	return true, nil
}

func login(c *router.Context, credentials user.Credentials) error {
	role, ok := user.LogIn(credentials)
	if !ok {
		return c.NoContent(http.StatusUnauthorized)
	}

	return c.JSON(http.StatusOK, role)
}

func getOrder(c *router.Context) error {
	return c.NoContent(http.StatusNotImplemented)
}
