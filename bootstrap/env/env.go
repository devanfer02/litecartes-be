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
	AccessTokenExpyHour  int	 `mapstructure:"ACCESS_TOKEN_EXPY_HOUR"`
	RefreshTokenExpyHour int 	 `mapstructure:"REFRESH_TOKEN_EXPY_HOUR"`
	AccessTokenSecret    string `mapstructure:"ACCESS_TOKEN_SECRET"`
	RefreshTokenSecret   string `mapstructure:"REFRESH_TOKEN_SECRET"`
	ClientURL			 string `mapstructure:"CLIENT_URL"`
    FirebaseDBURL        string `mapstructure:"FIREBASE_DB_URL"`
}

var ProcEnv = GetEnv()

func GetEnv() *Env {
	env := Env{}

	viper.SetConfigFile(".env")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Can't find the env file. ERR: %s\n", err.Error())
	}

	if err := viper.Unmarshal(&env); err != nil {
		log.Fatalf("Env variabels can't be loaded. ERR: %s\n", err.Error())
	}

	if env.AppEnv == "development" {
		log.Println("Server application is running on development mode")
	}

	return &env
}