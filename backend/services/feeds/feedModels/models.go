package feedmodels

import (
	"github.com/jinzhu/gorm"
	usermodel "github.com/pranay999000/social-minor/services/users/userModel"
	"github.com/pranay999000/social-minor/utils/config"
)

var db *gorm.DB

type Feed struct {
	gorm.Model
	Title		string			`json:"title"`
	Description	string			`json:"description"`
	Image		string			`json:"image"`
	UserId		int64			`json:"user_id"`
	User		usermodel.User	`json:"user"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Feed{})
}

func GetAllFeed() []Feed {
	var feeds []Feed
	db.Preload("User").Find(&feeds)
	return feeds
}

func GetFeedById(id int64) (*Feed, *gorm.DB) {
	var feed Feed
	db := db.Where("id=?", id).Preload("User").Find(&feed)
	return &feed, db
}

func (f *Feed) CreateFeed() *Feed {
	db.NewRecord(f)
	db.Create(f)
	return f
}

func GetFeedByUser(user_id int64) []Feed {
	var feeds []Feed
	db.Where("user_id=?", user_id).Preload("User").Find(&feeds)
	return feeds
}