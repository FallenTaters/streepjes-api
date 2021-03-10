package sldb

import (
	"database/sql"
	"strconv"
	"strings"

	"git.ultraware.nl/NiseVoid/qb"
	"git.ultraware.nl/NiseVoid/qb/qbdb"
)

// Driver implements SQLite-specific features
type Driver struct{}

// New returns the driver
func New(db *sql.DB) qbdb.DB {
	return qbdb.New(Driver{}, db)
}

// ValueString returns a the SQL for a parameter value
func (d Driver) ValueString(i int) string {
	return `@p` + strconv.Itoa(i)
}

// BoolString formats a boolean in a format supported by SQLite
func (d Driver) BoolString(v bool) string {
	if v {
		return `1`
	}
	return `0`
}

// EscapeCharacter returns the correct escape character for SQLite
func (d Driver) EscapeCharacter() string {
	return `"`
}

// UpsertSQL implements qb.Driver
func (d Driver) UpsertSQL(t *qb.Table, conflict []qb.Field, q qb.Query) (string, []interface{}) {
	c := qb.NewContext(d, qb.NoAlias())
	sql := ``
	for k, v := range conflict {
		if k > 0 {
			sql += qb.COMMA
		}
		sql += v.QueryString(c)
	}

	usql, values := q.SQL(qb.NewSQLBuilder(d))
	if !strings.HasPrefix(usql, `UPDATE `+t.Name) {
		panic(`Update does not update the correct table`)
	}
	usql = strings.ReplaceAll(usql, `UPDATE `+t.Name, `UPDATE`)

	return `ON CONFLICT (` + sql + `) DO ` + usql, values
}

// IgnoreConflictSQL implements qb.Driver
func (d Driver) IgnoreConflictSQL(t *qb.Table, conflict []qb.Field) (string, []interface{}) {
	c := qb.NewContext(d, qb.NoAlias())
	sql := ``
	for k, v := range conflict {
		if k > 0 {
			sql += qb.COMMA
		}
		sql += v.QueryString(c)
	}

	return `ON CONFLICT (` + sql + ") DO NOTHING\n", *c.Values
}

// LimitOffset implements qb.Driver
func (d Driver) LimitOffset(sql qb.SQL, limit, offset int) {
	if limit <= 0 {
		panic(`SQLite does not support offset without limit`)
	}
	sql.WriteLine(`LIMIT ` + strconv.Itoa(limit))
	if offset > 0 {
		sql.WriteLine(`OFFSET ` + strconv.Itoa(offset))
	}
}

// Returning implements qb.Driver
func (d Driver) Returning(b qb.SQLBuilder, q qb.Query, f []qb.Field) (string, []interface{}) {
	panic(`returning not currently supported`)
}

var types = map[qb.DataType]string{
	qb.Int:    `INTEGER`,
	qb.String: `TEXT`,
	qb.Bool:   `TINYINT`,
	qb.Float:  `REAL`,
	qb.Date:   `DATE`,
	qb.Time:   `DATETIME`,
}

// TypeName implements qb.Driver
func (d Driver) TypeName(t qb.DataType) string {
	if s, ok := types[t]; ok {
		return s
	}
	panic(`Unknown type`)
}

var override = qb.OverrideMap{}

func init() {
}

// Override implements qb.Driver
func (d Driver) Override() qb.OverrideMap {
	return override
}
