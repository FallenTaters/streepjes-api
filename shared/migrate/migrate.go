package migrate

import (
	"database/sql"
	"fmt"
)

//go:generate b-data -pkg $GOPACKAGE -prefix files files/...

// Migrate ...
func Migrate(db *sql.DB) error {
	migrated, _ := getMigrations(db)

	for _, file := range GetFileNames() {
		if isMigrated(migrated, file) {
			continue
		}
		q, err := GetFile(file)
		if err != nil {
			panic(err)
		}
		_, err = db.Exec(string(q))
		if err != nil {
			panic(err)
		}
		_, err = db.Exec(fmt.Sprintf(`INSERT INTO migration(filename) VALUES('%s');`, file))
		if err != nil {
			panic(err)
		}
	}

	return nil
}

func getMigrations(db *sql.DB) ([]string, error) {
	rows, err := db.Query(`SELECT filename FROM migration;`)
	if err != nil {
		return nil, err
	}
	filenames := []string{}
	var filename string
	for rows.Next() {
		err = rows.Scan(&filename)
		if err != nil {
			return nil, err
		}
		filenames = append(filenames, filename)
	}
	return filenames, nil
}

func isMigrated(migrated []string, filename string) bool {
	for _, file := range migrated {
		if file == filename {
			return true
		}
	}
	return false
}
