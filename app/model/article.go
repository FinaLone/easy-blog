package model

import (
	"database/sql"
)

// Article represent sql table article
type Article struct {
	ID       int            `db:"id"`
	Title    string         `db:"title"`
	ClassID  int            `db:"classId"`
	Abstract sql.NullString `db:"abstract"`
	Main     sql.NullString `db:"main"`
	Ctime    string         `db:"ctime"`
	Mtime    string         `db:"mtime"`
}
