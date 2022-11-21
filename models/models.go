package models

import (
	"fmt"

	"github.com/jary-287/web-demo/pkg/logging"
	"github.com/jary-287/web-demo/pkg/setting"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

type Model struct {
	ID         int `gorm:"primary_key" json:"id"`
	CreatedOn  int `json:"created_on"`
	ModifiedOn int `json:"modified_on"`
}

func init() {
	var (
		err                                               error
		dbType, host, user, password, tablePrefix, dbName string
	)
	// sec, err := setting.Cfg.GetSection("database")
	// if err != nil {
	// 	logging.Error("Fail to get section 'database':", err)
	// }
	dbType = setting.CFG.GetString("TYPE")
	dbName = setting.CFG.GetString("NAME")
	user = setting.CFG.GetString("USER")
	password = setting.CFG.GetString("PASSWORD")
	host = setting.CFG.GetString("HOST")
	tablePrefix = setting.CFG.GetString("TABLE_PREFIX")
	// dbType = sec.Key("TYPE").String()
	// dbName = sec.Key("NAME").String()
	// user = sec.Key("USER").String()
	// password = sec.Key("PASSWORD").String()
	// host = sec.Key("HOST").String()
	// tablePrefix = sec.Key("TABLE_PREFIX").String()
	db, err = gorm.Open(dbType, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user, password, host, dbName))
	if err != nil {
		logging.Error(err.Error())
		panic(err.Error())
	}
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return tablePrefix + defaultTableName
	}

	db.SingularTable(true)
	db.LogMode(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
}

func CloseDB() {
	defer db.Close()
}
