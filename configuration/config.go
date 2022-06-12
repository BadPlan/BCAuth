package configuration

import (
	"BCAuth/cmd"
	"github.com/spf13/viper"
	"strings"
)

var Instance Config

type Config struct {
	App           Application `mapstructure:"app"`
	Cookie        Cookie      `mapstructure:"cookie"`
	DB            Database    `mapstructure:"db"`
	ServicesHosts Services    `mapstructure:"services_hosts"`
	Log           Logger      `mapstructure:"log"`
}

type Application struct {
	Name           string          `mapstructure:"name"`
	Host           string          `mapstructure:"host"`
	Port           string          `mapstructure:"port"`
	Version        string          `mapstructure:"version"`
	SecretToken    string          `mapstructure:"secret_token"`
	AllowedOrigins map[string]bool `mapstructure:"allowed_origins"`
	Authorization  string          `mapstructure:"authorization"`
}

type Database struct {
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	Name     string `mapstructure:"name"`
	Driver   string `mapstructure:"driver"`
}

type Cookie struct {
	Name     string `mapstructure:"name"`
	Domain   string `mapstructure:"domain"`
	Path     string `mapstructure:"path"`
	TTL      uint64 `mapstructure:"ttl"`
	Secure   bool   `mapstructure:"secure"`
	HttpOnly bool   `mapstructure:"http_only"`
}

type Services struct {
}

type Logger struct {
	Path  string `mapstructure:"path"`
	Level string `mapstructure:"level"`
}

// PreparePath gets file path, divide path to dir and file name for viper set functions
func PreparePath(file string) (string, string) {
	return file[0:strings.LastIndex(file, "/")], file[strings.LastIndex(file, "/")+1 : strings.LastIndex(file, ".")]
}

func ParseConfig() error {
	path, name := PreparePath(cmd.ConfigPath)
	viper.SetConfigName(name)
	viper.AddConfigPath(path)
	err := viper.ReadInConfig()
	if err != nil {
		return err
	}

	err = viper.Unmarshal(&Instance)
	if err != nil {
		return err
	}

	return nil
}
