package configs

import "github.com/jinzhu/gorm"

var (
	writedb *gorm.DB
)

func WriteConnect() {
	writed, err := gorm.Open("mysql", "root:masterpassword@tcp(127.0.0.1:3306)/masterdatabase?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		panic(err.Error())
	}

	writedb = writed
}

func GetWriteDB() *gorm.DB {
	return writedb
}