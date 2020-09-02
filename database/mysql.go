package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init()  {
	var err error
	dsn := "root:123456@tcp(127.0.0.1:3306)/tech?charset=utf8&parseTime=True&loc=Local&timeout=10ms"
	DB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Printf("mysql connect error %v", err)
	}

	if DB.Error != nil {
		fmt.Printf("database error %v", DB.Error)
	}
}