package main

import (
	"github.com/FallenTaters/streepjes-api/domain/members"
	"github.com/FallenTaters/streepjes-api/domain/users"
	"github.com/FallenTaters/streepjes-api/model"
	"github.com/FallenTaters/streepjes-api/shared"
)

var testUsers = []users.User{
	{
		Username: "adminG",
		Club:     shared.ClubGladiators,
		Name:     "adminG",
		Role:     users.RoleAdmin,
		Password: []byte("admin"),
	}, {
		Username: "adminP",
		Club:     shared.ClubParabool,
		Name:     "adminP",
		Role:     users.RoleAdmin,
		Password: []byte("admin"),
	}, {
		Username: "bar",
		Name:     "bar",
		Role:     users.RoleBartender,
		Password: []byte("bar"),
	},
}

var testData = []bucketData{
	{
		[]byte("members"),
		[]keyValuePair{
			{
				shared.Itob(1),
				members.Member{
					ID:   1,
					Club: shared.ClubGladiators,
					Name: "Gladiator1",
				},
			}, {
				shared.Itob(2),
				members.Member{
					ID:   2,
					Club: shared.ClubGladiators,
					Name: "Gladiator2",
				},
			}, {
				shared.Itob(3),
				members.Member{
					ID:   3,
					Club: shared.ClubParabool,
					Name: "Parabool1",
				},
			}, {
				shared.Itob(4),
				members.Member{
					ID:   4,
					Club: shared.ClubParabool,
					Name: "Parabool2",
				},
			},
		},
	}, {
		[]byte("categories"),
		[]keyValuePair{
			{
				shared.Itob(1),
				model.Category{
					ID:   1,
					Name: "Drinks",
				},
			}, {
				shared.Itob(2),
				model.Category{
					ID:   2,
					Name: "Snacks",
				},
			},
		},
	}, {
		[]byte("products"),
		[]keyValuePair{
			{
				shared.Itob(1),
				model.Product{
					ID:              1,
					CategoryID:      1,
					Name:            "Beer",
					PriceParabool:   120,
					PriceGladiators: 150,
				},
			}, {
				shared.Itob(2),
				model.Product{
					ID:              2,
					CategoryID:      1,
					Name:            "Special Beer",
					PriceParabool:   180,
					PriceGladiators: 210,
				},
			}, {
				shared.Itob(3),
				model.Product{
					ID:              3,
					CategoryID:      1,
					Name:            "Wine",
					PriceParabool:   150,
					PriceGladiators: 170,
				},
			}, {
				shared.Itob(4),
				model.Product{
					ID:              4,
					CategoryID:      2,
					Name:            "Tosti Kaas",
					PriceParabool:   100,
					PriceGladiators: 120,
				},
			}, {
				shared.Itob(5),
				model.Product{
					ID:              5,
					CategoryID:      2,
					Name:            "Tosti Ham & Kaas",
					PriceParabool:   120,
					PriceGladiators: 140,
				},
			}, {
				shared.Itob(6),
				model.Product{
					ID:              6,
					CategoryID:      2,
					Name:            "Kaassoufflé",
					PriceParabool:   70,
					PriceGladiators: 90,
				},
			},
		},
	},
}
