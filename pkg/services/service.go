package services

import (
	"BCAuth/internal/models"
	"BCAuth/pkg/repositories"
	"github.com/gin-gonic/gin"
)

type Auth interface {
	RegisterUser(ctx *gin.Context, user models.User) (models.UserInfo, error)
	FindUserByID(ctx *gin.Context, id uint) (models.UserInfo, error)
	BrowseUser(ctx *gin.Context, user models.User) ([]models.UserInfo, error)
	UpdateUser(ctx *gin.Context, user models.User) (models.UserInfo, error)
	DeleteUser(ctx *gin.Context, id uint) (models.UserInfo, error)
}

type Service struct {
	Auth
}

func ServiceInit(repository *repositories.Repository) *Service {
	return &Service{
		Auth: InitAuthService(repository),
	}
}
