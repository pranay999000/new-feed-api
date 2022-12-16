package followmodels

import (
	"github.com/jinzhu/gorm"
	usermodel "github.com/pranay999000/social-minor/services/users/userModel"
	"github.com/pranay999000/social-minor/utils/config"
)

var db *gorm.DB

type Follow struct {
	gorm.Model
	FollowerId		int64	`json:"follower_id"`
	FollowingId		int64	`json:"following_id"`
	Follower		usermodel.User	`json:"follower"`
	Following		usermodel.User	`json:"following"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Follow{})
}

func (f *Follow) CreateFollow() *Follow {
	db.NewRecord(f)
	db.Create(&f)
	return f
}

func CheckFollow(follower_id int64, following_id int64) bool {
	var follow Follow
	db := db.Where("follower_id=?", follower_id).Where("following_id=?", following_id).Find(&follow)

	if db.RecordNotFound() {
		return false
	} else {
		return true
	}
}


func GetFollowers(user_id int64) []Follow {
	var follows []Follow
	db.Where("following_id=?", user_id).Preload("Follower").Find(&follows)
	return follows
}

func GetFollowing(user_id int64) []Follow {
	var follows []Follow
	db.Where("follower_id=?", user_id).Preload("Following").Find(&follows)
	return follows
}

func DeleteFollow(follower_id int64, following_id int64) bool {
	var follow Follow
	db := db.Where("follower_id=?", follower_id).Where("following_id=?", following_id).Delete(&follow)

	if db.RecordNotFound() {
		return false
	} else {
		return true
	}
}