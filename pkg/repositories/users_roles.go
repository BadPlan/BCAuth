package repositories

import (
	"BCAuth/internal/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UsersRolesRepository struct {
	tx *gorm.DB
}

func (u UsersRolesRepository) Create(ctx *gin.Context, userRole models.UserRole) (models.UserRole, error) {
	if err := u.tx.Table("user_roles").Create(&userRole).Error; err != nil {
		return models.UserRole{}, err
	}
	return userRole, nil
}

func InitUsersRolesRepository(database *gorm.DB) *UsersRolesRepository {
	return &UsersRolesRepository{
		tx: database.Session(&gorm.Session{SkipDefaultTransaction: true}),
	}
}
