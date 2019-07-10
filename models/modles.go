package models

import (
	"fmt"
	"log"
	"rehabilitation_prescription/pkg/setting"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" // to remember driver's import path重要，没有会找不到driver "mysql" 导致panic
)

var db *gorm.DB

type Model struct {
	ID         int `gorm:"primary_key" json:"id"`
	CreatedOn  int `json:"created_on"`
	ModifiedOn int `json:"modified_on"`
}

// InitDB initializes the database instance
func InitDB() {
	var err error
	db, err = gorm.Open(setting.DatabaseSetting.Type, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		setting.DatabaseSetting.User,
		setting.DatabaseSetting.Password,
		setting.DatabaseSetting.Host,
		setting.DatabaseSetting.Name))

	if err != nil {
		log.Fatalf("models.Setup err: %v", err)
	}

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return setting.DatabaseSetting.TablePrefix + defaultTableName
	}

	db.SingularTable(true)
	db.Callback().Create().Replace("gorm:update_timestamp", updateTimestampForCreateCallback)
	db.Callback().Update().Replace("gorm:update_timestamp", updateTimestampForUpdateCallback)
	db.DB().SetMaxIdleConns(10)  // 设置最大空闲连接数
	db.DB().SetMaxOpenConns(100) // 设置最大连接数，最大空闲连接数小于等于最大连接数
}

// updateTimestampForCreateCallback will set `CreatedOn`, `ModifiedOn` when creating
func updateTimestampForCreateCallback(scope *gorm.Scope) {
	if !scope.HasError() {
		nowTime := time.Now().Unix()
		if createdOn, ok := scope.FieldByName("CreatedOn"); ok {
			if createdOn.IsBlank {
				createdOn.Set(nowTime)
			}
		}

		if modifyOn, ok := scope.FieldByName("ModifyOn"); ok {
			if modifyOn.IsBlank {
				modifyOn.Set(nowTime)
			}
		}
	}
}

// updateTimestampForUpdateCallback will set `ModifiedOn` when updating
func updateTimestampForUpdateCallback(scope *gorm.Scope) {
	if _, ok := scope.Get("gorm:update_column"); !ok {
		scope.SetColumn("ModifiedOn", time.Now().Unix())
	}
}

func CloseDB() {
	defer db.Close()
}
