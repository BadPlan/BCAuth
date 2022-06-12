package utils

import (
	"BCAuth/cmd"
	"BCAuth/configuration"
	"BCAuth/pkg/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

var (
	BCAUTH_CONFIG = "BCAUTH_CONFIG"
)

func init() {
	fmt.Println("Init tests package")
	cmd.ConfigPath = os.Getenv(BCAUTH_CONFIG)
	if cmd.ConfigPath == "" {
		fmt.Errorf("BCAUTH_CONFIG env variable was not set")
		os.Exit(1)
	}
	err := configuration.ParseConfig()
	if err != nil {
		log.Fatalln(err)
		return
	}
}

func TestSetCookie(t *testing.T) {
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	utils.SetCookie(ctx, "token")
	assert.Contains(t, ctx.Writer.Header().Get("set-cookie"), "token")
}

func TestGetTokenFromHeader(t *testing.T) {
	t.Run("GetTokenFromHeader with empty auth header", EmptyHeader)
	t.Run("GetTokenFromHeader with valid header", ValidHeader)
	t.Run("GetTokenFromHeader with no header", NoHeader)
}

func EmptyHeader(t *testing.T) {
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	ctx.Request = new(http.Request)
	ctx.Request.Header = make(map[string][]string)
	ctx.Request.Header.Add("authorization", "")
	token, err := utils.GetTokenFromHeader(ctx)
	assert.NotNil(t, err)
	require.Empty(t, token)
}

func ValidHeader(t *testing.T) {
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	ctx.Request = new(http.Request)
	ctx.Request.Header = make(map[string][]string)
	ctx.Request.Header.Add("authorization", "token")
	token, err := utils.GetTokenFromHeader(ctx)
	assert.Nil(t, err)
	require.NotEmpty(t, token)
}

func NoHeader(t *testing.T) {
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	ctx.Request = new(http.Request)
	ctx.Request.Header = make(map[string][]string)
	token, err := utils.GetTokenFromHeader(ctx)
	assert.NotNil(t, err)
	require.Empty(t, token)
}
