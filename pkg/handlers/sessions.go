package handlers

import (
	"BCAuth/pkg/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type SessionService struct {
	service services.Session
}

func (h *Handler) UserByToken(ctx *gin.Context) {
	cookie, exists := ctx.Get("token")
	if !exists {
		ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": "no cookie header",
		})
		return
	}
	value, err := h.services.Session.UserByToken(ctx, cookie.(string))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, value)
}
