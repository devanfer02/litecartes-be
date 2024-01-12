package env

import (
	"log"

	"github.com/spf13/viper"
)

type Env struct {
	AppEnv               string `mapstructure:"APP_ENV"`
	ServerAddress        string `mapstructure:"SERVER_ADDRESS"`
	ApiKey				 string	`mapstructure:"API_KEY"`
	DBHost               string `mapstructure:"DB_HOST"`
	DBPort               string `mapstructure:"DB_PORT"`
	DBUser               string `mapstructure:"DB_USER"`
	DBPassword           string `mapstructure:"DB_PASSWORD"`
	DBName               string `mapstructure:"DB_NAME"`
	ClientURL			 string `mapstructure:"CLIENT_URL"`
    FireConfigPath       string `mapstructure:"FIREBASE_SDK_CONFIG_PATH"`
}

var ProcEnv = GetEnv()

func GetEnv() *Env {
	env := Env{}

	viper.SetConfigFile(".env")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("[ENV] Can't find the env file. ERR: %s\n", err.Error())
	}

	if err := viper.Unmarshal(&env); err != nil {
		log.Fatalf("[ENV] Env variabels can't be loaded. ERR: %s\n", err.Error())
	}

	if env.AppEnv == "development" {
		log.Println("[ENV] Server application is running on development mode")
	}

	return &env
}