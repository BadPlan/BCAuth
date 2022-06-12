package handlers

import (
	"BCAuth/configuration"
	"BCAuth/pkg/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"net/http"
	"time"
)

func GenerateJWT(email string) (string, error) {
	var signingKey = []byte(configuration.Instance.App.SecretToken)
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["email"] = email
	claims["exp"] = time.Now().Add(time.Minute * time.Duration(configuration.Instance.Cookie.TTL)).Unix()

	tokenString, err := token.SignedString(signingKey)

	if err != nil {
		return "", fmt.Errorf("signed key corrupted")
	}
	return tokenString, nil
}

func (h *Handler) ValidateToken(ctx *gin.Context) {
	token, err := utils.GetTokenFromHeader(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, map[string]string{
			"message": err.Error(),
		})
		ctx.Abort()
		return
	}
	var signingKey = []byte(configuration.Instance.App.SecretToken)

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
