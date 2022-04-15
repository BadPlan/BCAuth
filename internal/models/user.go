package models

import "gorm.io/gorm"

type UserInfo struct {
	ID       *uint   `json:"id,omitempty" form:"id"`
	Name     *string `json:"name,omitempty" form:"name"`
	Email    *string `json:"email,omitempty" form:"email"`
	Password *string `json:"password" form:"password"`
}

type User struct {
	gorm.Model
	Name     *string `json:"name" form:"name"`
	Email    *string `json:"email" gorm:"unique, notNull" form:"email"`
	Password []byte  `json:"password" gorm:"notNull" form:"password"`
}
