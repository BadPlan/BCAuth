package services

import (
	"BCAuth/internal/models"
	"BCAuth/pkg/repositories"
	"github.com/gin-gonic/gin"
)

type Users interface {
	RegisterUser(ctx *gin.Context, user models.User) (models.UserInfo, error)
	FindUserByID(ctx *gin.Context, id uint) (models.UserInfo, error)
	BrowseUser(ctx *gin.Context, user models.User) ([]models.UserInfo, error)
	UpdateUser(ctx *gin.Context, user models.User) (models.UserInfo, error)
	DeleteUser(ctx *gin.Context, id uint) (models.UserInfo, error)
}

type Roles interface {
	CreateRole(ctx *gin.Context, role models.Role) (models.RoleInfo, error)
	BrowseRole(ctx *gin.Context, role models.Role) ([]models.RoleInfo, error)
	FindRoleByID(ctx *gin.Context, id uint) (models.RoleInfo, error)
	UpdateRole(ctx *gin.Context, role models.Role) (models.RoleInfo, error)
	DeleteRole(ctx *gin.Context, id uint) (models.RoleInfo, error)
}

type UsersRoles interface {
	CreateUserRole(ctx *gin.Context, userRole models.UserRole) (models.UserRole, error)
}

type Service struct {
	Users
	Roles
	UsersRoles
}

func ServiceInit(repository *repositories.Repository) *Service {
	return &Service{
		Users:      InitUsersService(&repository.Users),
		Roles:      InitRolesService(&repository.Roles),
		UsersRoles: InitUsersRolesService(&repository.UsersRoles),
	}
}
