package configs

import "github.com/jinzhu/gorm"

var (
	transactionDB *gorm.DB
)

func TransactionConnect() {
	transactiondb, err := gorm.Open("mysql", "root:masterpassword@tcp(127.0.0.1:3306)/masterdatabase?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		panic(err.Error())
	}

	transactionDB = transactiondb
}

func GetTransactionDB() *gorm.DB {
	return transactionDB
}
