package repositories

import (
	"BCAuth/internal/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UsersRepository struct {
	tx *gorm.DB
}

func (a UsersRepository) Update(ctx *gin.Context, user models.User) (models.UserInfo, error) {
	if err := a.tx.Save(user).Error; err != nil {
		return models.UserInfo{}, err
	}
	newPass := string(user.Password)
	return models.UserInfo{ID: &user.ID, Name: user.Name, Email: user.Email, Password: &newPass}, nil
}

func (a UsersRepository) Delete(ctx *gin.Context, id uint) (models.UserInfo, error) {
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

func (a UsersRepository) Browse(ctx *gin.Context, user models.User) ([]models.UserInfo, error) {
	var values []models.User
	if err := a.tx.Raw(user.BrowseQuery()).Scan(&values).Error; err != nil {
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

func (a UsersRepository) Find(ctx *gin.Context, id uint) (models.UserInfo, error) {
	var model models.User
	if err := a.tx.Where(map[string]interface{}{
		"id": id,
	}).First(&model).Error; err != nil {
		return models.UserInfo{}, err
	}
	if err := a.tx.Model(&model).Association("Roles").Find(&model.Roles); err != nil {
		return models.UserInfo{}, err
	}
	pass := string(model.Password)
	var roles []models.RoleInfo
	println(len(model.Roles))
	for _, r := range model.Roles {
		roles = append(roles, models.RoleInfo{ID: &r.ID, Name: r.Name, Description: r.Description})
	}
	return models.UserInfo{ID: &model.ID, Name: model.Name, Email: model.Email, Password: &pass, Roles: roles}, nil
}

func (a UsersRepository) Create(ctx *gin.Context, user models.User) (models.UserInfo, error) {
	if err := a.tx.Create(&user).Error; err != nil {
		return models.UserInfo{}, err
	}
	return models.UserInfo{Name: user.Name, Email: user.Email, ID: &user.ID}, nil
}

func InitUsersRepository(database *gorm.DB) *UsersRepository {
	return &UsersRepository{
		tx: database.Session(&gorm.Session{SkipDefaultTransaction: true}),
	}
}
