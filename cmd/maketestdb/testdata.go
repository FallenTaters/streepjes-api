package main

import (
	"github.com/FallenTaters/streepjes-api/domain/catalog"
	"github.com/FallenTaters/streepjes-api/domain/members"
	"github.com/FallenTaters/streepjes-api/domain/users"
	"github.com/FallenTaters/streepjes-api/shared"
	"github.com/FallenTaters/streepjes-api/shared/buckets"
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
				buckets.Itob(1),
				members.Member{
					ID:   1,
					Club: shared.ClubGladiators,
					Name: "Gladiator1",
				},
			}, {
				buckets.Itob(2),
				members.Member{
					ID:   2,
					Club: shared.ClubGladiators,
					Name: "Gladiator2",
				},
			}, {
				buckets.Itob(3),
				members.Member{
					ID:   3,
					Club: shared.ClubParabool,
					Name: "Parabool1",
				},
			}, {
				buckets.Itob(4),
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
				buckets.Itob(1),
				catalog.Category{
					ID:   1,
					Name: "Drinks",
				},
			}, {
				buckets.Itob(2),
				catalog.Category{
					ID:   2,
					Name: "Snacks",
				},
			},
		},
	}, {
		[]byte("products"),
		[]keyValuePair{
			{
				buckets.Itob(1),
				catalog.Product{
					ID:              1,
					CategoryID:      1,
					Name:            "Beer",
					PriceParabool:   120,
					PriceGladiators: 150,
				},
			}, {
				buckets.Itob(2),
				catalog.Product{
					ID:              2,
					CategoryID:      1,
					Name:            "Special Beer",
					PriceParabool:   180,
					PriceGladiators: 210,
				},
			}, {
				buckets.Itob(3),
				catalog.Product{
					ID:              3,
					CategoryID:      1,
					Name:            "Wine",
					PriceParabool:   150,
					PriceGladiators: 170,
				},
			}, {
				buckets.Itob(4),
				catalog.Product{
					ID:              4,
					CategoryID:      2,
					Name:            "Tosti Kaas",
					PriceParabool:   100,
					PriceGladiators: 120,
				},
			}, {
				buckets.Itob(5),
				catalog.Product{
					ID:              5,
					CategoryID:      2,
					Name:            "Tosti Ham & Kaas",
					PriceParabool:   120,
					PriceGladiators: 140,
				},
			}, {
				buckets.Itob(6),
				catalog.Product{
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
