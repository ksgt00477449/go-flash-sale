package model

import (
	"time"
)

type User struct {
	// gorm.Model
	ID        uint      `gorm:"primaryKey" json:"id"`
	Email     string    `gorm:"type:varchar(255);uniqueIndex;not null" json:"email"`
	Password  string    `gorm:"type:varchar(255);not null" json:"-"` // 不返回密码
	CreatedAt time.Time `json:"-"`
}
