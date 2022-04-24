package models

import "gorm.io/gorm"

type Session struct {
	gorm.Model
	Token   *string `json:"token"`
	Expires *int64  `json:"expires"`
	UserID  *uint   `json:"userID"`
	User    User    `gorm:"foreignKey:UserID"`
}
