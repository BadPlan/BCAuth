package repositories

import (
	"BCAuth/internal/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Auth interface {
	Create(ctx *gin.Context, user models.User) (models.UserInfo, error)
	Browse(ctx *gin.Context, user models.User) ([]models.UserInfo, error)
	Find(ctx *gin.Context, id uint) (models.UserInfo, error)
	Update(ctx *gin.Context, user models.User) (models.UserInfo, error)
	Delete(ctx *gin.Context, id uint) (models.UserInfo, error)
}

type Repository struct {
	Auth
}

func RepositoryInit(db *gorm.DB) *Repository {
	return &Repository{
		Auth: InitAuthRepository(db),
	}
}
