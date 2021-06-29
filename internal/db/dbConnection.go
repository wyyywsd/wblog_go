package db

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var W_Db *gorm.DB

func InitDb() {
	dsn := "host=localhost user=mac password=wyyywsd dbname=wblog port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	var err error
	W_Db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("数据库连接成功!")
	}
}
