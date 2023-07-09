package config

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

const (
	userName = "root"
	password = "123456"
	host     = "localhost:3306"
	dbName   = "tjtc_account"
	//mysqlDsn = "userName:password@tcp(host)/dbName?parseTime=true"
	mysqlDsn = "root:123456@tcp(127.0.0.1:3306)/localhost"
)

func InitDB() *sqlx.DB {
	db, err := sqlx.Connect("mysql", mysqlDsn)
	if err != nil {
		logrus.Panic(err)
	}
	db.SetMaxIdleConns(100)
	db.SetMaxOpenConns(200)
	logrus.Info("初始化数据库成功")
	return db
}

var DB = InitDB()
