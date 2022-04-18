package services

import (
	"BCAuth/internal/models"
	"BCAuth/pkg/repositories"
	"github.com/gin-gonic/gin"
)

type RolesService struct {
	RolesRepo repositories.Roles
}

func (r RolesService) BrowseRole(ctx *gin.Context, role models.Role) ([]models.RoleInfo, error) {
	value, err := r.RolesRepo.Browse(ctx, role)
	if err != nil {
		return nil, err
	}
	return value, nil
}

func (r RolesService) FindRoleByID(ctx *gin.Context, id uint) (models.RoleInfo, error) {
	value, err := r.RolesRepo.Find(ctx, id)
	if err != nil {
		return models.RoleInfo{}, err
	}
	return value, nil
}

func (r RolesService) UpdateRole(ctx *gin.Context, role models.Role) (models.RoleInfo, error) {
	value, err := r.RolesRepo.Update(ctx, role)
	if err != nil {
		return models.RoleInfo{}, err
	}
	return value, nil
}

func (r RolesService) DeleteRole(ctx *gin.Context, id uint) (models.RoleInfo, error) {
	value, err := r.RolesRepo.Delete(ctx, id)
	if err != nil {
		return models.RoleInfo{}, err
	}
	return value, nil
}

func (r RolesService) CreateRole(ctx *gin.Context, role models.Role) (models.RoleInfo, error) {
	if value, err := r.RolesRepo.Create(ctx, role); err != nil {
		return models.RoleInfo{}, err
	} else {
		return value, nil
	}
}

func InitRolesService(repository *repositories.Roles) *RolesService {
	return &RolesService{
		RolesRepo: *repository,
	}
}
