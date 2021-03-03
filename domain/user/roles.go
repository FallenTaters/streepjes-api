package user

type Role struct {
	Role string `json:"role"`
}

var (
	RoleNone      = Role{``}
	RoleBartender = Role{`bartender`}
	RoleAdmin     = Role{`admin`}
)
