package repositories

import (
	"BCAuth/internal/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Users interface {
	Create(ctx *gin.Context, user models.User) (models.UserInfo, error)
	Browse(ctx *gin.Context, user models.User) ([]models.UserInfo, error)
	Find(ctx *gin.Context, id uint) (models.UserInfo, error)
	Update(ctx *gin.Context, user models.User) (models.UserInfo, error)
	Delete(ctx *gin.Context, id uint) (models.UserInfo, error)
}

type Roles interface {
	Create(ctx *gin.Context, role models.Role) (models.RoleInfo, error)
	Browse(ctx *gin.Context, role models.Role) ([]models.RoleInfo, error)
	Find(ctx *gin.Context, id uint) (models.RoleInfo, error)
	Update(ctx *gin.Context, role models.Role) (models.RoleInfo, error)
	Delete(ctx *gin.Context, id uint) (models.RoleInfo, error)
}

type UsersRoles interface {
	Create(ctx *gin.Context, userRole models.UserRole) (models.UserRole, error)
}

type Session interface {
	UserByToken(ctx *gin.Context, session models.Session) (models.Session, error)
	CreateSession(ctx *gin.Context, session models.Session) (models.Session, error)
}

type Repository struct {
	Users
	Roles
	UsersRoles
	Session
}

func RepositoryInit(db *gorm.DB) *Repository {
	return &Repository{
		Users:      InitUsersRepository(db),
		Roles:      InitRolesRepository(db),
		UsersRoles: InitUsersRolesRepository(db),
		Session:    InitSessionRepository(db),
	}
}
