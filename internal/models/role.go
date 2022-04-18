package models

import (
	"gorm.io/gorm"
	"strings"
)

type Role struct {
	gorm.Model
	Name        *string `json:"name" form:"name" gorm:"unique, notNull"`
	Description *string `json:"description" form:"description"`
}

type RoleInfo struct {
	ID          *uint   `json:"id,omitempty" form:"id"`
	Name        *string `json:"name,omitempty" form:"name"`
	Description *string `json:"description,omitempty" form:"description"`
}

func (r *Role) BrowseQuery() string {
	query := "SELECT * FROM roles WHERE deleted_at is null"
	if r.Name != nil {
		query += strings.Replace(` AND name LIKE '%?%'`, "?", *r.Name, 1)
	}
	if r.Description != nil {
		query += strings.Replace(` AND description LIKE '%?%'`, "?", *r.Description, 1)
	}
	return query
}
