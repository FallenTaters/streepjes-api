package users

import (
	"database/sql"
	"time"
)

var db *sql.DB

var selectBase = `SELECT id, club, name, username, password, role, auth_token, auth_datetime FROM user `

func getUserByUsername(username string) (User, error) {
	row := db.QueryRow(selectBase+`WHERE username = $username;`, username)
	return scanUser(row)
}

func insert(user User) error {
	_, err := db.Exec(`INSERT INTO user(name, club, username, password, role) VALUES($1, $2, $3, $4, $5)`, user.Name, user.Club, user.Username, user.Password, user.Role)
	return err
}

func removeToken(id int) {
	_, err := db.Exec(`UPDATE user SET auth_datetime = $time WHERE id = $id;`, time.Now().Add(-loginTime), id)
	if err != nil {
		panic(err)
	}
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

func getAll() []User {
	rows, err := db.Query(selectBase + `;`)
	if err != nil {
		panic(err)
	}
	defer rows.Close() //nolint: errcheck

	users := []User{}
	for rows.Next() {
		users = append(users, mustScanUser(rows))
	}

	return users
}

func mustScanUser(rows Scanner) User {
	u, err := scanUser(rows)
	if err != nil {
		panic(err)
	}
	return u
}

func scanUser(rows Scanner) (User, error) {
	var u User
	return u, rows.Scan(&u.ID, &u.Club, &u.Name, &u.Username, &u.Password, &u.Role, &u.AuthToken, &u.AuthDatetime)
}

type Scanner interface {
	Scan(...interface{}) error
}
