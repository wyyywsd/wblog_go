package models

import (
	"fmt"
	"wblogApi/internal/db"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserName     string `gorm:"not null" form:"user_name" json:"user_name" binding:"required"`
	PassWord     string `gorm:"not null" form:"pass_word" json:"pass_word" binding:"required"`
	Email        string
	ProfilePhoto string     `gorm:"default: '/public/user_profile/default.png'"`
	Articles     []*Article `gorm:"foreignKey:UserId"`
}

//通过用户名查询用户,  目前先用于校验注册 用户名是否存在
func FindUserByUserName(userName string) (*gorm.DB, User) {
	var user User
	result := db.W_Db.Where("user_name = ?", userName).First(&user)
	//err := db.W_Db.Table("users").Where("user_name = ?", userName).Find(&user).Error
	fmt.Println(result)
	return result, user
}

func CreateUser(userName string, passWord string) *gorm.DB {
	user := User{UserName: userName, PassWord: passWord}
	result := db.W_Db.Create(&user) // 通过数据的指针来创建
	// fmt.Println(result)
	// fmt.Println(result.Error)
	return result
}

func UpdateUserInfo(user User, updateusers map[string]interface{}) *gorm.DB {
	result := db.W_Db.Model(&user).Updates(updateusers)
	return result
}

func FindUserById(userId uint) User {
	var user User
	db.W_Db.Where("id = ?", userId).First(&user)
	return user
}
