package members

import "database/sql"

func Init(database *sql.DB) {
	db = database
}

func GetAll() []Member {
	return getAll()
}
