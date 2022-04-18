package repositories

import (
	"BCAuth/internal/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type RolesRepository struct {
	tx *gorm.DB
}

func (r RolesRepository) Create(ctx *gin.Context, role models.Role) (models.RoleInfo, error) {
	if err := r.tx.Create(&role).Error; err != nil {
		return models.RoleInfo{}, err
	}
	return models.RoleInfo{ID: &role.ID, Name: role.Name, Description: role.Description}, nil
}

func (r RolesRepository) Browse(ctx *gin.Context, role models.Role) ([]models.RoleInfo, error) {
	var roles []models.RoleInfo
	var values []models.Role
	if err := r.tx.Table("role").Raw(role.BrowseQuery()).Scan(&values).Error; err != nil {
		return roles, nil
	}
	for _, v := range values {
		model := models.RoleInfo{}
		model.ID = new(uint)
		*model.ID = v.ID
		model.Name = v.Name
		model.Description = v.Description
		roles = append(roles, model)
	}
	return roles, nil
}

func (r RolesRepository) Find(ctx *gin.Context, id uint) (models.RoleInfo, error) {
	var role models.Role
	if err := r.tx.Where(map[string]interface{}{
		"id": id,
	}).First(&role).Error; err != nil {
		return models.RoleInfo{}, err
	}
	return models.RoleInfo{ID: &role.ID, Name: role.Name, Description: role.Description}, nil
}

func (r RolesRepository) Update(ctx *gin.Context, role models.Role) (models.RoleInfo, error) {
	if err := r.tx.Save(&role).Error; err != nil {
		return models.RoleInfo{}, err
	}
	return models.RoleInfo{ID: &role.ID, Name: role.Name, Description: role.Description}, nil
}

func (r RolesRepository) Delete(ctx *gin.Context, id uint) (models.RoleInfo, error) {
	var role models.Role
	if err := r.tx.Where(map[string]interface{}{
		"id": id,
	}).First(&role).Error; err != nil {
		return models.RoleInfo{}, err
	}
	if err := r.tx.Delete(&models.User{}, map[string]interface{}{
		"id": id,
	}).Error; err != nil {
		return models.RoleInfo{}, err
	}
	return models.RoleInfo{ID: &role.ID, Name: role.Name, Description: role.Description}, nil
}

func InitRolesRepository(database *gorm.DB) *RolesRepository {
	return &RolesRepository{
		tx: database.Session(&gorm.Session{SkipDefaultTransaction: true}),
	}
}
