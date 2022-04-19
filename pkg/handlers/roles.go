package handlers

import (
	"BCAuth/internal/models"
	"BCAuth/pkg/services"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"net/http"
	"strconv"
)

type RolesService struct {
	Roles services.Roles
}

func (h *Handler) CreateRole(ctx *gin.Context) {
	var model models.RoleInfo
	err := ctx.ShouldBindJSON(&model)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, BindJsonError)
		return
	}
	value, err := h.services.Roles.CreateRole(ctx, models.Role{Name: model.Name, Description: model.Description})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusCreated, value)
}

func (h *Handler) BrowseRole(ctx *gin.Context) {
	var model models.Role
	err := ctx.ShouldBindQuery(&model)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, BindParamsError)
		return
	}
	value, err := h.services.Roles.BrowseRole(ctx, model)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"roles": value})
}

func (h *Handler) RoleInfo(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, BindIdError)
		return
	}
	value, err := h.services.Roles.FindRoleByID(ctx, cast.ToUint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, value)
}

func (h *Handler) UpdateRole(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, BindIdError)
		return
	}
	var model models.Role
	err = ctx.ShouldBindJSON(&model)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, BindJsonError)
		return
	}
	model.ID = cast.ToUint(id)

	value, err := h.services.Roles.UpdateRole(ctx, model)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, value)
}

func (h *Handler) DeleteRole(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, BindIdError)
		return
	}
	value, err := h.services.Roles.DeleteRole(ctx, cast.ToUint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, value)
}
