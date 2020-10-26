package dbutil

import (
	"fmt"

	"github.com/finalone/easy-blog/app/config"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var (
	// SQLXDB is sqlx instance
	SQLXDB *sqlx.DB
)

func init() {
	if SQLXDB != nil {
		return
	}
	SQLXDB = NewSQLXDB()
}

// NewSQLXDB get a *sqlx.DB
func NewSQLXDB() *sqlx.DB {
	cfg := &config.Cfg
	if cfg == nil {
		return nil
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/easyblog?charset=utf8&parseTime=True", cfg.User, cfg.Passwd, cfg.Addr, cfg.Port)

	db, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		fmt.Println("sql connect error", err)
		panic("sql connect error")
	}

	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(10)

	return db
}
