package utils

import (
	"BCAuth/configuration"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"strings"
)

var cookieConfig = &configuration.Instance.Cookie

func SetCookie(ctx *gin.Context, token string) {
	ctx.SetCookie(cookieConfig.Name, token, cast.ToInt(cookieConfig.TTL), cookieConfig.Path, cookieConfig.Domain, cookieConfig.Secure, cookieConfig.HttpOnly)
}

func GetTokenFromHeader(ctx *gin.Context) (string, error) {
	header := ctx.GetHeader("authorization")
	if header == "" {
		return "", fmt.Errorf(`header 'authorization' not found`)
	}
	token := strings.TrimLeft(strings.Replace(header, configuration.Instance.App.Authorization, "", 1), " \t")
	return token, nil
}
