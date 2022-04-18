package models

import (
	"gorm.io/gorm"
	"strings"
)

type UserInfo struct {
	ID       *uint      `json:"id,omitempty" form:"id"`
	Name     *string    `json:"name,omitempty" form:"name"`
	Email    *string    `json:"email,omitempty" form:"email"`
	Password *string    `json:"password" form:"password"`
	Roles    []RoleInfo `json:"roles" form:"roles"`
}

type User struct {
	gorm.Model
	Name     *string `json:"name" form:"name"`
	Email    *string `json:"email" gorm:"unique, notNull" form:"email"`
	Password []byte  `json:"password" gorm:"notNull" form:"password"`
	Roles    []Role  `json:"roles" gorm:"many2many:user_roles;"`
}

func (u *User) BrowseQuery() string {
	query := "SELECT * FROM users WHERE deleted_at is null"
	if u.Name != nil {
		query += strings.Replace(` AND name LIKE '%?%'`, "?", *u.Name, 1)
	}
	if u.Email != nil {
		query += strings.Replace(` AND email LIKE '%?%'`, "?", *u.Email, 1)
	}
	return query
}
