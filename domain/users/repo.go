package users

import (
	"database/sql"
	"time"
)

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

func setToken(user User) error {
	res, err := db.Exec(`UPDATE user SET auth_token = $token, auth_datetime = $time WHERE username = $username;`, user.AuthToken, time.Now(), user.Username)
	if err != nil {
		return err
	}

	affected, err := res.RowsAffected()
	if err != nil || affected != 1 {
		return err
	}

	return nil
}

func refreshToken(username string) {
	_, err := db.Exec(`UPDATE user SET auth_datetime = $time WHERE username = $username;`, time.Now(), username)
	if err != nil {
		panic(err)
	}
}
