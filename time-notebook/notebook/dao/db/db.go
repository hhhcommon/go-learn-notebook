package db

import (
	"github.com/jmoiron/sqlx"
)

var (
	// DB 全局可用的数据库连接
	DB *sqlx.DB
)

// Init 初始化数据库连接
func Init(dns string) error {
	var err error
	DB, err = sqlx.Open("mysql", dns)
	if err != nil {
		return err
	}
	err = DB.Ping()
	if err != nil {
		return err
	}
	DB.SetMaxOpenConns(100)
	DB.SetMaxIdleConns(16)
	return nil
}
