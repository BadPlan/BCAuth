package handlers

import (
	"BCAuth/internal/models"
	"BCAuth/pkg/services"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"net/http"
	"strconv"
)

type UsersRolesHandler struct {
	services.UsersRoles
}

func (h *Handler) AddRole(ctx *gin.Context) {
	var model models.UserRole
	err := ctx.ShouldBindJSON(&model)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, BindJsonError)
		return
	}

	value, err := h.services.UsersRoles.CreateUserRole(ctx, model)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusCreated, value)

}

func (h *Handler) RolesByUser(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("user_id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, BindIdError)
		return
	}
	value, err := h.services.Users.FindUserByID(ctx, cast.ToUint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"user_roles": value.Roles})
}
