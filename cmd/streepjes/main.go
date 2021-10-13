package main

import (
	"github.com/FallenTaters/streepjes-api/domain/catalog"
	"github.com/FallenTaters/streepjes-api/domain/users"
	"github.com/FallenTaters/streepjes-api/repo"
	"github.com/FallenTaters/streepjes-api/repo/buckets"
	"github.com/FallenTaters/streepjes-api/shared"
	"github.com/FallenTaters/streepjes-api/shared/cookies"
)

func main() {
	readSettings()

	close := initPackages()
	defer close()

	startupChecks()

	r := makeRouter()
	panic(r.Start(`:` + settings.Port))
}

func initPackages() func() {
	cookies.Init(settings.DisableSecure)
	closeDB := buckets.Init()

	catalog.Init(repo.NewProductRepo(), repo.NewCategoryRepo())

	return closeDB
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
