package dao

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

var err error

func init() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/bubble?&parseTime=True&loc=Local" //我把uft8mb删掉了，否则会报错
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}
}
