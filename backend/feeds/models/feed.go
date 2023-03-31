package models

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/pranay999000/feeds/configs"
)

var db *gorm.DB
var writedb *gorm.DB

type Feed struct {
	gorm.Model
	Title		string		`json:"title"`
	Body		string		`json:"body"`
	Image		string		`json:"image"`
	TimeStamp	time.Time	`json:"timestamp"`
}

func init() {
	configs.WriteConnect()
	configs.Connect()
	writedb = configs.GetWriteDB()
	db = configs.GetDB()
	writedb.AutoMigrate(&Feed{})
}

func GetFeeds() []Feed {
	var feeds []Feed
	db.Find(&feeds)
	return feeds
}