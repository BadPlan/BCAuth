package pkg

import (
	"BCAuth/configuration"
	"fmt"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/diode"
	"os"
	"runtime"
	"time"
)

var loggerInstance *zerolog.Logger

func initLogger() error {
	level := configuration.Instance.Log.Level

	logPath := func() string {
		if configuration.Instance.Log.Path == "" {
			opSystem := runtime.GOOS
			switch opSystem {
			case "linux":
				return "/var/log/" + configuration.Instance.App.Name + "/error.log"
			case "windows":
				return "c:\\temp\\" + configuration.Instance.App.Name + "\\error.log"
			}
			return "error.log"
		}
		return configuration.Instance.Log.Path
	}

	file, err := os.OpenFile(logPath(), os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		err := fmt.Errorf("cannot open the log file %s, Reason : %s", logPath(), err.Error())
		return err
	}

	writer := diode.NewWriter(file, 10000, 10*time.Microsecond, func(missed int) {
		fmt.Printf("Logger dropped %d messages", missed)
	})

	zLogger := zerolog.New(writer).With().Caller().Timestamp().Logger().Output(file)
	loggerInstance = &zLogger

	switch level {
	case "ERROR":
		zerolog.SetGlobalLevel(zerolog.ErrorLevel)
		break
	case "WARNING":
		zerolog.SetGlobalLevel(zerolog.WarnLevel)
		break
	case "INFO":
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
		break
	case "DEBUG":
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
		break
	case "TRACE":
		zerolog.SetGlobalLevel(zerolog.TraceLevel)
		break
	default:
		zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	}
	return nil
}

func InitLogger() (zerolog.Logger, error) {
	if err := initLogger(); err != nil {
		return zerolog.Logger{}, err
	}
	loggerInstance.Info().Msg("Starting logger...")
	return *loggerInstance, nil
}
