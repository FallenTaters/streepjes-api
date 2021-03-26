package members

import "database/sql"

var db *sql.DB

func getAll() []Member {
	rows, err := db.Query(`SELECT id, name, club, balance FROM members;`)
	if err != nil {
		panic(err)
	}

	members := []Member{}
	for rows.Next() {
		members = append(members, scanMember(rows))
	}

	return members
}

func scanMember(rows *sql.Rows) Member {
	var member Member
	err := rows.Scan(&member.ID, &member.Name, &member.Club, &member.Balance)
	if err != nil {
		panic(err)
	}
	member.Debt = -member.Balance
	return member
}
