package services

import (
	"BCAuth/internal/models"
	"BCAuth/pkg/repositories"
	"github.com/gin-gonic/gin"
)

type UsersService struct {
	UsersRepo repositories.Users
}

func (a UsersService) UpdateUser(ctx *gin.Context, user models.User) (models.UserInfo, error) {
	found, err := a.UsersRepo.Find(ctx, user.ID)
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
	value, err := a.UsersRepo.Update(ctx, user)
	if err != nil {
		return models.UserInfo{}, err
	}
	return value, nil
}

func (a UsersService) DeleteUser(ctx *gin.Context, id uint) (models.UserInfo, error) {
	value, err := a.UsersRepo.Delete(ctx, id)
	if err != nil {
		return models.UserInfo{}, err
	}
	return value, nil
}

func (a UsersService) BrowseUser(ctx *gin.Context, user models.User) ([]models.UserInfo, error) {
	value, err := a.UsersRepo.Browse(ctx, user)
	if err != nil {
		return nil, err
	}
	return value, nil
}

func (a UsersService) FindUserByID(ctx *gin.Context, id uint) (models.UserInfo, error) {
	value, err := a.UsersRepo.Find(ctx, id)
	if err != nil {
		return models.UserInfo{}, err
	}
	return value, nil
}

func (a UsersService) RegisterUser(ctx *gin.Context, user models.User) (models.UserInfo, error) {
	value, err := a.UsersRepo.Create(ctx, user)
	if err != nil {
		return models.UserInfo{}, err
	}
	return value, nil
}

func InitUsersService(repository *repositories.Users) *UsersService {
	return &UsersService{
		UsersRepo: *repository,
	}
}
