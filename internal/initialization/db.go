package initialization

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// var db *gorm.DB

func InitDB() *gorm.DB {
	// 数据库地址 账号 密码，
	dsn := "root:pepsi1145ai@tcp(172.23.84.152:3306)/go-flash-sale?charset=utf8mb4&parseTime=True&loc=Local"
	// gorm连接到mysql数据库
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}
	fmt.Println("Database connected successfully!")
	return db
}
