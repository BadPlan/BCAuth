package main

import (
	"BCAuth/cmd"
	"BCAuth/configuration"
	"BCAuth/pkg"
	"BCAuth/pkg/db"
	"BCAuth/pkg/handlers"
	"BCAuth/pkg/repositories"
	"BCAuth/pkg/services"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

var engine *gin.Engine

func run() {
	err := initLogger()
	if err != nil {
		log.Error().Msg("Cannot run web server due to logger error")
		panic(err.Error())
	}
	db.Instance = db.Init()
	repos := repositories.RepositoryInit(db.Instance)
	serv := services.ServiceInit(repos)
	handle := handlers.HandlerInit(serv)
	engine = handle.InitRoutes()

	err = engine.Run(configuration.Instance.App.Host + ":" + configuration.Instance.App.Port)
	if err != nil {
		log.Error().Msg("Cannot run web server")
		panic(err.Error())
	}
}

func main() {
	cmd.Execute()
	err := configuration.ParseConfig()
	if err != nil {
		panic(err)
	}

	run()
}

func initLogger() error {
	var err error
	if log.Logger, err = pkg.InitLogger(); err != nil {
		return err
	}
	log.Info().Msg("Logger initialized!")
	return nil
}
