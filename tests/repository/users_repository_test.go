package repository

import (
	"BCAuth/cmd"
	"BCAuth/configuration"
	"BCAuth/internal/models"
	"BCAuth/pkg/db"
	"BCAuth/pkg/repositories"
	"fmt"
	"github.com/cockroachdb/copyist"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io"
	"log"
	"os"
	"testing"
)

var (
	BCAUTH_CONFIG = "BCAUTH_CONFIG"
)

func init() {
	fmt.Println("Init users repository test package")
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

var (
	closer          io.Closer                = nil
	addedUserId     uint                     = 0
	usersRepository *repositories.Repository = nil
)

func TestUsersRepository(t *testing.T) {
	copyist.Register("postgres")

	closer = copyist.Open(t)
	err := db.ConnectTestDB()
	assert.Nil(t, err)
	usersRepository = repositories.RepositoryInit(db.Instance)

}

func TestUsersRepository_Create(t *testing.T) {
	email := "admin@admin.ru"
	password := []byte("123456")
	name := "admin"
	user := models.User{
		Email:    &email,
		Name:     &name,
		Password: password,
	}
	u, err := usersRepository.Users.Create(&gin.Context{}, user)
	assert.Nil(t, err)
	addedUserId = *u.ID
}

func TestUsersRepository_Browse(t *testing.T) {
	name := "admin"
	output, err := usersRepository.Users.Browse(&gin.Context{}, models.User{Name: &name})
	require.Nil(t, err)
	assert.Equal(t, *output[0].Name, name, "user was added to database, so expected to browse it by name")
}

func TestUsersRepository_BrowseLike(t *testing.T) {
	name := "ad"
	output, err := usersRepository.Users.Browse(&gin.Context{}, models.User{Name: &name})
	require.Nil(t, err)
	assert.Equal(t, len(output), 1, "check LIKE filter")
}

func TestUsersRepository_BrowseLikeEmpty(t *testing.T) {
	name := ""
	output, err := usersRepository.Users.Browse(&gin.Context{}, models.User{Name: &name})
	require.Nil(t, err)
	assert.Greater(t, len(output), 0, "check LIKE filter")
}

func TestUsersRepository_Update(t *testing.T) {

}

func TestUsersRepository_Delete(t *testing.T) {
	_, err := usersRepository.Users.Delete(&gin.Context{}, addedUserId)
	assert.Nil(t, err)
	closer.Close()
}
