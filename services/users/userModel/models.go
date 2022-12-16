package usermodel

import (
	"github.com/jinzhu/gorm"
	"github.com/pranay999000/social-minor/utils/config"
)

var db *gorm.DB

type User struct {
	gorm.Model
	Name		string	`json:"name"`
	Email		string	`json:"email"`
	Password	string	`json:"password"`
	Image		string	`json:"image"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&User{})
}

func GetAllUsers() []User {
	var users []User
	db.Find(&users)
	return users
}

func (u *User) CreateUser() *User {
	db.NewRecord(u)
	db.Create(&u)
	return u
}

func GetUserByEmail(Email string) (*User, *gorm.DB) {
	var user User
	db := db.Where("email=?", Email).Find(&user)
	return &user, db
}