package user

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func LogIn(c Credentials) (Role, bool) {
	testCredentials := Credentials{
		Username: "username",
		Password: "password",
	}
	if c == testCredentials {
		return RoleAdmin, true
	}
	return RoleNone, false
}
