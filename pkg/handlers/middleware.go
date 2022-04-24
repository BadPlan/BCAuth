package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/spf13/viper"
	"net/http"
	"time"
)

func GenerateJWT(email string) (string, error) {
	var signingKey = []byte(viper.GetString("token"))
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["email"] = email
	claims["exp"] = time.Now().Add(time.Minute * 300).Unix()

	tokenString, err := token.SignedString(signingKey)

	if err != nil {
		return "", fmt.Errorf("signed key corrupted")
	}
	return tokenString, nil
}

func (h *Handler) ValidateToken(ctx *gin.Context) {
	token, err := ctx.Cookie("bc_auth")
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, map[string]string{
			"message": "not found token",
		})
		ctx.Abort()
		return
	}
	var signingKey = []byte(viper.GetString("token"))

	tok, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("there was an error in parsing")
		}
		return signingKey, nil
	})

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, map[string]string{
			"message": err.Error(),
		})
		ctx.Abort()
		return
	}
	if !tok.Valid {
		ctx.JSON(http.StatusUnauthorized, map[string]string{
			"message": "token not valid",
		})
		ctx.Abort()
		return
	}
	ctx.Set("token", token)
}
