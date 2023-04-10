package configs

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

func ReadConnect() {
	d, err := gorm.Open("mysql", "root:masterpassword@tcp(127.0.0.1:3300)/masterdatabase?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		panic(err.Error())
	}
	db = d
}

func GetReadDB() *gorm.DB {
	return db
}