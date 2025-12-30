package initialization

import (
	"go-flash-sale/internal/model"

	"gorm.io/gorm"
)

func InitTableAutoMigrate(db *gorm.DB) {
	// 自动迁移模式，创建或更新表结构
	db.AutoMigrate(&model.User{})
}
