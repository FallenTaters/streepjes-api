package user

type Role struct {
	Role string `json:"role"`
}

var (
	RoleNotAuthorized = Role{``}
	RoleBartender     = Role{`bartender`}
	RoleAdmin         = Role{`admin`}
)
