package model

import "time"

type User struct {
	ID    uint   `gorm:"primaryKey" json:"-"`
	Email string `gorm:"uniqueIndex;not null" json:"email"`
	Password  string    `gorm:"not null" json:"-"` // 不返回密码
	CreatedAt time.Time `json:"-"`
}
