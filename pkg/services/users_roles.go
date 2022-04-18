package services

import (
	"BCAuth/internal/models"
	"BCAuth/pkg/repositories"
	"github.com/gin-gonic/gin"
)

type UsersRolesService struct {
	UsersRolesRepo repositories.UsersRoles
}

func (u UsersRolesService) CreateUserRole(ctx *gin.Context, userRole models.UserRole) (models.UserRole, error) {
	value, err := u.UsersRolesRepo.Create(ctx, userRole)
	if err != nil {
		return models.UserRole{}, err
	}
	return value, nil
}

func InitUsersRolesService(repository *repositories.UsersRoles) *UsersRolesService {
	return &UsersRolesService{
		UsersRolesRepo: *repository,
	}
}
