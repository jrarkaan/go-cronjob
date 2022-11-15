package dto

type Conf struct {
	App struct {
		Name       string `env:"APP_NAME"`
		Name_api   string `env:"APP_NAME_API"`
		Port       string `env:"APP_PORT"`
		Mode       string `env:"APP_MODE"`
		Url        string `env:"APP_URL"`
		Secret_key string `env:"APP_SECRET"`
	}
}
