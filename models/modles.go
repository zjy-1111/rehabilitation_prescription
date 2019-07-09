package models

import (
	"fmt"
	"log"
	"rehabilitation_prescription/pkg/setting"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Model struct {
	ID         int `gorm:"primary_key" json:"id"`
	CreatedOn  int `json:"created_on"`
	ModifiedOn int `json:"modified_on"`
}

func init() {
	var err error
	var dbType, dbName, user, password, host, tablePrefix string

	sec, err := setting.Cfg.GetSection("database")
	if err != nil {
		log.Fatal(2, "Fail to get section 'database': %v", err)
	}

	dbType = sec.Key("TYPE").String()
	dbName = sec.Key("NAME").String()
	user = sec.Key("USER").String()
	password = sec.Key("PASSWORD").String()
	host = sec.Key("HOST").String()
	tablePrefix = sec.Key("TABLE_PREFIX").String()

	// 创建db，user:password@tcp(host)/dbName
	db, err = gorm.Open(dbType, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=local",
		user,
		password,
		host,
		dbName,
	))
	if err != nil {
		log.Println(err)
	}

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return tablePrefix + defaultTableName
	}

	db.SingularTable(true)
	db.DB().SetMaxIdleConns(10)  // 设置最大空闲连接数
	db.DB().SetMaxOpenConns(100) // 设置最大连接数，最大空闲连接数小于等于最大连接数
}

func CloseDB() {
	defer db.Close()
}
