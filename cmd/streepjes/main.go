package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"runtime/debug"

	"git.fuyu.moe/Fuyu/router"
	"github.com/PotatoesFall/streepjes/domain/users"
	"github.com/PotatoesFall/streepjes/shared"
	"github.com/PotatoesFall/streepjes/shared/buckets"
)

func main() {
	readSettings()

	shared.Init(settings.DisableSecure)
	close := buckets.Init()
	defer close() //nolint: errcheck

	startupChecks()

	startServer()
}

func errorHandler(c *router.Context, v interface{}) {
	fmt.Fprintf(os.Stderr, "panic: %s\n", v)
	fmt.Fprintln(os.Stderr, string(debug.Stack()))
	_ = c.NoContent(http.StatusInternalServerError)
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

func startupChecks() {
	u, err := users.GetAll()
	if err != nil {
		panic(err)
	}

	if len(u) == 0 {
		err = users.Put(users.User{
			Username: `adminGladiators`,
			Club:     shared.ClubGladiators,
			Name:     `Gladiators Admin`,
			Role:     users.RoleAdmin,
			Password: []byte(`playlacrossebecauseitsfun`),
		})
		if err != nil {
			panic(err)
		}

		err = users.Put(users.User{
			Username: `adminParabool`,
			Club:     shared.ClubParabool,
			Name:     `Parabool Admin`,
			Role:     users.RoleAdmin,
			Password: []byte(`groningerstudentenkorfbalcommissie`),
		})
		if err != nil {
			panic(err)
		}
	}
}
