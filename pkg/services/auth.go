package services

import (
	"BCAuth/internal/models"
	"BCAuth/pkg/repositories"
	"github.com/gin-gonic/gin"
)

type AuthService struct {
	AuthRepo repositories.Auth
}

func (a AuthService) UpdateUser(ctx *gin.Context, user models.User) (models.UserInfo, error) {
	found, err := a.AuthRepo.Find(ctx, user.ID)
	if err != nil {
		return models.UserInfo{}, err
	}
	if user.Name == nil {
		user.Name = found.Name
	}
	if user.Email == nil {
		user.Email = found.Email
	}
	if user.Password == nil {
		user.Password = []byte(*found.Password)
	}
	value, err := a.AuthRepo.Update(ctx, user)
	if err != nil {
		return models.UserInfo{}, err
	}
	return value, nil
}

func (a AuthService) DeleteUser(ctx *gin.Context, id uint) (models.UserInfo, error) {
	value, err := a.AuthRepo.Delete(ctx, id)
	if err != nil {
		return models.UserInfo{}, err
	}
	return value, nil
}

func (a AuthService) BrowseUser(ctx *gin.Context, user models.User) ([]models.UserInfo, error) {
	value, err := a.AuthRepo.Browse(ctx, user)
	if err != nil {
		return nil, err
	}
	return value, nil
}

func (a AuthService) FindUserByID(ctx *gin.Context, id uint) (models.UserInfo, error) {
	value, err := a.AuthRepo.Find(ctx, id)
	if err != nil {
		return models.UserInfo{}, err
	}
	return value, nil
}

func (a AuthService) RegisterUser(ctx *gin.Context, user models.User) (models.UserInfo, error) {
	value, err := a.AuthRepo.Create(ctx, user)
	if err != nil {
		return models.UserInfo{}, err
	}
	return value, nil
}

func InitAuthService(repository *repositories.Repository) *AuthService {
	return &AuthService{
		AuthRepo: repository,
	}
}
