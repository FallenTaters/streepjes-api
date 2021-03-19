package user

import "database/sql"

var db *sql.DB

func getUserByUsername(username string) (User, error) {
	row := db.QueryRow(`SELECT id, name, username, password, role, auth_token, auth_datetime FROM user WHERE username = $username;`, username)
	var user User
	return user, row.Scan(&user.ID, &user.Name, &user.Username, &user.Password, &user.Role, &user.AuthToken, &user.AuthDatetime)
}

func insert(name, username string, password []byte, role Role) error {
	_, err := db.Exec(`INSERT INTO user(name, username, password, role) VALUES($1, $2, $3, $4)`, name, username, password, role)
	return err
}
