package config

import (
	"github.com/joho/godotenv"
	"github.com/jrarkaan/go-cronjob/dto"
	"github.com/jrarkaan/go-cronjob/pkg/helpers"
	"os"
)

func Init() *dto.Conf {
	errEnv := godotenv.Load()
	helpers.PanicIfError(errEnv)
	appName := os.Getenv("service.APP_NAME")
	appNameApi := os.Getenv("service.APP_NAME_API")
	appPort := os.Getenv("service.APP_PORT")

	conf := dto.Conf{
		App: struct {
			Name       string `env:"APP_NAME"`
			Name_api   string `env:"APP_NAME_API"`
			Port       string `env:"APP_PORT"`
			Mode       string `env:"APP_MODE"`
			Url        string `env:"APP_URL"`
			Secret_key string `env:"APP_SECRET"`
		}(struct {
			Name       string
			Name_api   string
			Port       string
			Mode       string
			Url        string
			Secret_key string
		}{Name: appName, Name_api: appNameApi, Port: appPort, Mode: "", Url: "", Secret_key: ""}),
	}
	return &conf
}
