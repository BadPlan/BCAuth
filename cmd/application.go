package main

import (
	"BCAuth/pkg"
	"BCAuth/pkg/db"
	"BCAuth/pkg/handlers"
	"BCAuth/pkg/repositories"
	"BCAuth/pkg/services"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

var engine *gin.Engine

func run() {
	err := initLogger()
	if err != nil {
		log.Error().Msg("Cannot run web server due to logger error")
		panic(err.Error())
	}
	database := db.Init()
	repos := repositories.RepositoryInit(database)
	serv := services.ServiceInit(repos)
	handle := handlers.HandlerInit(serv)
	engine = handle.InitRoutes()

	err = engine.Run(viper.GetString("host") + ":" + viper.GetString("port"))
	if err != nil {
		log.Error().Msg("Cannot run web server")
		panic(err.Error())
	}

}

func main() {
	err := parseConfig()
	if err != nil {
		panic(err)
	}

	run()
}

func parseConfig() error {
	configPath := "../configuration"

	viper.AddConfigPath(configPath)
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

func initLogger() error {
	var err error
	if log.Logger, err = pkg.InitLogger(); err != nil {
		return err
	}
	log.Info().Msg("Logger initialized!")
	return nil
}
