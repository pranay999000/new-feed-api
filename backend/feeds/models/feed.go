package models

import (
	"github.com/jinzhu/gorm"
	"github.com/pranay999000/feeds/configs"
)

var readdb *gorm.DB
var writedb *gorm.DB

type Feed struct {
	gorm.Model
	Title		string		`json:"title"`
	Body		string		`json:"body"`
	Image		string		`json:"image"`
	UserId		string		`json:"user_id"`
	LikeCount	int64		`json:"like_count"`
	ViewCount	int64		`json:"view_count"`
}

func init() {
	configs.WriteConnect()
	configs.ReadConnect()
	writedb = configs.GetWriteDB()
	readdb = configs.GetReadDB()
	writedb.AutoMigrate(&Feed{})
}

func GetFeeds(limit int64, page int64) []Feed {
	var feeds []Feed
	readdb.Order("").Offset((page - 1) * limit).Limit(limit).Find(&feeds)
	return feeds
}

func (f *Feed) CreateFeed() *Feed {
	writedb.NewRecord(f)
	writedb.Create(f)
	return f
}

func GetFeedById(id int64) (*Feed, *gorm.DB) {
	var feed Feed
	f := readdb.Where("id=?", id).Find(&feed)
	return &feed, f
}

func GetFeedByUser(user_id int64) []Feed {
	var feeds []Feed
	readdb.Where("user_id=?", user_id).Find(&feeds)
	return feeds
}

func CheckFeed(feed_id int64, channel chan Feed) {
	var feed Feed
	readdb.Where("ID=?", feed_id).First(&feed)

	channel <- feed
}