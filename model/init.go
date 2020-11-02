package model

import (
	"fmt"
	"gin-web-skeleton/logger"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)

type Database struct {
	Self *gorm.DB
}

var DB *Database

func InitSelf() *gorm.DB {
	return openDB(
		viper.GetString("db.username"),
		viper.GetString("db.password"),
		viper.GetString("db.addr"),
		viper.GetString("db.port"),
		viper.GetString("db.name"),
	)
}

func openDB(username, password, addr, port, name string) *gorm.DB {
	fmt.Println(username, "---", password, "---", addr, "---", port, "---", name)
	config := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=%t&loc=%s",
		username,
		password,
		addr,
		port,
		name,
		true,
		"Local",
	)
	db, err := gorm.Open("mysql", config)
	if err != nil {
		fmt.Println("Database err------", err)
		logger.Debugf("Database connection failed.Database name: %s", name)
	}
	setupDB(db)
	return db
}

func setupDB(db *gorm.DB) {
	db.LogMode(viper.GetBool("gormlog"))
	db.DB().SetMaxIdleConns(0)
	db.DB().SetMaxOpenConns(100)
	db.SingularTable(true)
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return "gws_" + defaultTableName
	}
}

func GetSelfDB() *gorm.DB {
	return InitSelf()
}

func (db *Database) Init() {
	DB = &Database{
		Self: GetSelfDB(),
	}
}

func (db *Database) Close() {
	_ = DB.Self.Close()
}
