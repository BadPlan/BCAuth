package db

import (
	cfg "BCAuth/configuration"
	"BCAuth/internal/models"
	"database/sql"
	"fmt"
	"github.com/rs/zerolog/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

var Instance *gorm.DB = nil

func Init() *gorm.DB {
	uri := fmt.Sprintf("%s://%s:%s@%s:%s/%s",
		cfg.Instance.DB.Driver,
		cfg.Instance.DB.User,
		cfg.Instance.DB.Password,
		cfg.Instance.DB.Host,
		cfg.Instance.DB.Port,
		cfg.Instance.DB.Name)

	db, err := gorm.Open(postgres.Open(uri), &gorm.Config{})
	if err != nil {
		log.Error().Msgf("error: %s", err.Error())
		os.Exit(1)
	}

	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Role{})
	db.AutoMigrate(&models.Session{})

	return db
}

func ConnectTestDB() error {
	dbConn, _ := sql.Open("copyist_postgres", fmt.Sprintf("%s://%s:%s@%s:%s/%s",
		cfg.Instance.DB.Driver,
		cfg.Instance.DB.User,
		cfg.Instance.DB.Password,
		cfg.Instance.DB.Host,
		cfg.Instance.DB.Port,
		cfg.Instance.DB.Name))

	var err error = nil
	Instance, err = gorm.Open(postgres.New(postgres.Config{Conn: dbConn}), &gorm.Config{})
	if err != nil {
		return err
	}
	return nil
}
