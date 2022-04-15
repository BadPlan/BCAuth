package repositories

import (
	"BCAuth/internal/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AuthRepository struct {
	tx *gorm.DB
}

func (a AuthRepository) Update(ctx *gin.Context, user models.User) (models.UserInfo, error) {
	if err := a.tx.Save(user).Error; err != nil {
		return models.UserInfo{}, err
	}
	newPass := string(user.Password)
	return models.UserInfo{ID: &user.ID, Name: user.Name, Email: user.Email, Password: &newPass}, nil
}

func (a AuthRepository) Delete(ctx *gin.Context, id uint) (models.UserInfo, error) {
	var model models.User
	if err := a.tx.Where(map[string]interface{}{
		"id": id,
	}).First(&model).Error; err != nil {
		return models.UserInfo{}, err
	}
	if err := a.tx.Delete(&models.User{}, map[string]interface{}{
		"id": id,
	}).Error; err != nil {
		return models.UserInfo{}, err
	}
	return models.UserInfo{ID: &model.ID, Name: model.Name, Email: model.Email}, nil
}

func (a AuthRepository) Browse(ctx *gin.Context, user models.User) ([]models.UserInfo, error) {
	var values []models.User
	if err := a.tx.Where(&user).Find(&values).Error; err != nil {
		return []models.UserInfo{}, err
	}

	var result []models.UserInfo
	for _, v := range values {
		pass := string(v.Password)
		model := models.UserInfo{}
		model.ID = new(uint)
		*model.ID = v.ID
		model.Name = v.Name
		model.Password = &pass
		model.Email = v.Email
		result = append(result, model)
	}

	return result, nil
}

func (a AuthRepository) Find(ctx *gin.Context, id uint) (models.UserInfo, error) {
	var model models.User
	if err := a.tx.Where(map[string]interface{}{
		"id": id,
	}).First(&model).Error; err != nil {
		return models.UserInfo{}, err
	}
	return models.UserInfo{ID: &model.ID, Name: model.Name, Email: model.Email}, nil
}

func (a AuthRepository) Create(ctx *gin.Context, user models.User) (models.UserInfo, error) {
	if err := a.tx.Create(&user).Error; err != nil {
		return models.UserInfo{}, err
	}
	return models.UserInfo{Name: user.Name, Email: user.Email, ID: &user.ID}, nil
}

func InitAuthRepository(database *gorm.DB) *AuthRepository {
	return &AuthRepository{
		tx: database.Session(&gorm.Session{SkipDefaultTransaction: true}),
	}
}
