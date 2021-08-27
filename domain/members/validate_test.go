package members

import (
	"testing"

	"git.fuyu.moe/Fuyu/assert"
	"github.com/FallenTaters/bbucket"
	"github.com/FallenTaters/streepjes-api/shared"
)

func prepareRepoMock() {
	tr := &testRepo{}
	repo = tr

	member1 := Member{
		ID:   1,
		Name: "Name",
		Club: shared.ClubGladiators,
	}
	member2 := Member{
		ID:   2,
		Name: "Name2",
		Club: shared.ClubParabool,
	}

	tr._get = func(id int) (Member, error) {
		switch id {
		case 1:
			return member1, nil
		case 2:
			return member2, nil
		}

		return Member{}, bbucket.ErrObjectNotFound
	}

	tr._getAll = func() ([]Member, error) {
		return []Member{member1, member2}, nil
	}
}

var testValues = []struct {
	text     string
	member   Member
	expected error
}{
	{
		`allow save of new member`,
		Member{
			Name: "NewName",
			Club: shared.ClubGladiators,
		},
		nil,
	}, {
		`allow save of new member with duplicate name but different club`,
		Member{
			Name: "Name",
			Club: shared.ClubParabool,
		},
		nil,
	}, {
		`disallow save of new member with duplicate name-club combination`,
		Member{
			Name: "Name",
			Club: shared.ClubGladiators,
		},
		ErrNameTaken,
	}, {
		`disallow save of new member with empty name`,
		Member{
			Name: "",
			Club: shared.ClubGladiators,
		},
		ErrEmptyName,
	}, {
		`disallow save of new member with unknown club`,
		Member{
			Name: "Name",
			Club: shared.ClubUnknown,
		},
		ErrUnknownClub,
	}, {
		`allow update of member`,
		Member{
			ID:   1,
			Name: "NewName",
			Club: shared.ClubGladiators,
		},
		nil,
	}, {
		`disallow update of non-existing member`,
		Member{
			ID:   69,
			Name: "NewName",
			Club: shared.ClubGladiators,
		},
		ErrMemberNotFound,
	}, {
		`disallow update with empty name`,
		Member{
			ID:   1,
			Name: "",
			Club: shared.ClubGladiators,
		},
		ErrEmptyName,
	}, {
		`disallow update with unknown club`,
		Member{
			ID:   1,
			Name: "NewName",
			Club: shared.ClubUnknown,
		},
		ErrUnknownClub,
	}, {
		`allow update without name and club change`,
		Member{
			ID:   1,
			Name: "Name",
			Club: shared.ClubGladiators,
		},
		nil,
	}, {
		`disallow update to existing name-club combination`,
		Member{
			ID:   1,
			Name: "Name2",
			Club: shared.ClubParabool,
		},
		ErrNameTaken,
	},
}

func TestValidateMember(t *testing.T) {
	prepareRepoMock()

	for _, testCase := range testValues {
		t.Run(testCase.text, func(t *testing.T) {
			assert := assert.New(t)

			assert.Eq(testCase.expected, validateMember(testCase.member))
		})
	}
}
