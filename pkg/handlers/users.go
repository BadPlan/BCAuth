package handlers

import (
	"BCAuth/internal/models"
	"BCAuth/pkg/services"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"strconv"
)

type UsersHandler struct {
	UsersService services.Users
}

func (h *Handler) Register(ctx *gin.Context) {
	var userData models.UserInfo

	if err := ctx.ShouldBindJSON(&userData); err != nil {
		ctx.JSON(http.StatusBadRequest, BindJsonError)
		return
	}

	if userData.Password == nil || userData.Name == nil || userData.Email == nil {
		ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": "one field is empty, enter full data",
		})
		return
	}

	user, err := h.services.Users.BrowseUser(ctx, models.User{
		Email: userData.Email,
	})
	if user != nil {
		ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": "email " + *userData.Email + " already in use",
		})
		return
	}

	password, _ := bcrypt.GenerateFromPassword([]byte(*userData.Password), 12)
	var model models.User
	model.Password = password
	model.Name = userData.Name
	model.Email = userData.Email

	value, err := h.services.RegisterUser(ctx, model)
	if err != nil {
		ctx.JSON(http.StatusBadRequest,
			map[string]string{
				"message": err.Error(),
			})
		return
	}
	jwt, err := GenerateJWT(*value.Email)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}
	ctx.SetCookie("bc_auth", jwt, 300*60, "/", viper.GetString("domain"), true, true)
	ctx.JSON(http.StatusCreated, value)
}

func (h *Handler) UserInfo(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, BindIdError)
		return
	}
	value, err := h.services.Users.FindUserByID(ctx, cast.ToUint(id))
	if err != nil {
		return
	}

	if err != nil {
		ctx.JSON(http.StatusInternalServerError,
			map[string]string{
				"message": err.Error(),
			})
		return
	}
	value.Password = nil

	ctx.JSON(http.StatusOK, value)
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func (h *Handler) SignIn(ctx *gin.Context) {
	var data models.Authentication
	err := ctx.ShouldBindJSON(&data)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, BindJsonError)
		return
	}
	value, err := h.services.Users.BrowseUser(ctx, models.User{Email: &data.Email})
	if err != nil {
		ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": "wrong login or password",
		})
		return
	}
	if len(value) > 0 {
		ok := CheckPasswordHash(data.Password, *value[0].Password)
		if ok {
			token, err := GenerateJWT(data.Email)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, map[string]string{
					"message": err.Error(),
				})
				return
			}
			ctx.SetCookie("bc_auth", token, 300*60, "/", viper.GetString("host"), true, true)
		} else {
			ctx.JSON(http.StatusBadRequest, map[string]string{
				"message": "wrong login or password",
			})
			return
		}
	} else {
		ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": "wrong login or password",
		})
		return
	}
	value[0].Password = nil
	ctx.JSON(http.StatusOK, value[0])
}

func (h *Handler) BrowseUsers(ctx *gin.Context) {
	var model models.User
	err := ctx.ShouldBindQuery(&model)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, BindParamsError)
		return
	}
	model.Password = nil

	users, err := h.services.Users.BrowseUser(ctx, model)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
		})
		return
	}
	for i := range users {
		users[i].Password = nil
	}
	ctx.JSON(http.StatusOK, gin.H{"users": users})

}

func (h *Handler) UpdateUser(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, BindIdError)
		return
	}
	var model models.UserInfo
	err = ctx.ShouldBindJSON(&model)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": "could not bind params",
		})
		return
	}
	var user models.User
	if model.Password != nil {
		password, _ := bcrypt.GenerateFromPassword([]byte(*model.Password), 12)
		user.Password = password
	}
	user.ID = cast.ToUint(id)
	user.Name = model.Name
	user.Email = model.Email
	value, err := h.services.UpdateUser(ctx, user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
		return
	}
	value.Password = nil
	ctx.JSON(http.StatusOK, value)
}

func (h *Handler) DeleteUser(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, BindIdError)
		return
	}

	value, err := h.services.Users.DeleteUser(ctx, cast.ToUint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, value)
}
