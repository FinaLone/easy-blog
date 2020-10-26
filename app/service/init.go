package service

import (
	"fmt"

	"github.com/finalone/easy-blog/app/dbutil"
	"github.com/jmoiron/sqlx"
)

var (
	SqlxDB *sqlx.DB
)

func init() {
	SqlxDB = dbutil.SQLXDB
	fmt.Println(SqlxDB)
	if SqlxDB == nil {
		SqlxDB = dbutil.NewSQLXDB()
	}
}
