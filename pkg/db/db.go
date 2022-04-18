package db

import (
	"BCAuth/internal/models"
	"fmt"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

func Init() *gorm.DB {
	uri := fmt.Sprintf("postgres://%s:%s@%s:%d/%s",
		viper.GetString("db_user"),
		viper.GetString("db_password"),
		viper.GetString("db_host"),
		viper.GetInt("db_port"),
		viper.GetString("db_name"))

	db, err := gorm.Open(postgres.Open(uri), &gorm.Config{})
	if err != nil {
		log.Error().Msgf("error: %s", err.Error())
		os.Exit(1)
	}

	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Role{})

	return db
}
