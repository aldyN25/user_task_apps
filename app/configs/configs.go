package configs

import (
	"os"
	"sync"
)

type (
	AppConfig struct {
		Name string
		Env  string
		Port string
		Host string
	}

	DbConfig struct {
		Host     string
		Port     string
		Dbname   string
		Username string
		Password string
	}

	OauthConfig struct {
		ClientID     string
		ClientSecret string
		TokenURL     string
	}

	Configs struct {
		Mode        string
		Appconfig   AppConfig
		Dbconfig    DbConfig
		OauthConfig OauthConfig
	}
)

var lock = &sync.Mutex{}
var configs *Configs

func GetInstance() *Configs {
	configs = &Configs{
		Appconfig: AppConfig{
			Name: os.Getenv("APP_NAME"),
			Env:  os.Getenv("APP_ENV"),
			Port: os.Getenv("APP_PORT"),
			Host: os.Getenv("APP_HOST"),
		},
		Dbconfig: DbConfig{
			Host:     os.Getenv("MYSQL_HOST"),
			Port:     os.Getenv("MYSQL_PORT"),
			Dbname:   os.Getenv("MYSQL_DBNAME"),
			Username: os.Getenv("MYSQL_USER"),
			Password: os.Getenv("MYSQL_PASSWORD"),
		},
		Mode: os.Getenv("MODE"),
		OauthConfig: OauthConfig{
			ClientID:     os.Getenv("CLIENT_ID"),
			ClientSecret: os.Getenv("CLIENT_SECRET"),
			TokenURL:     os.Getenv("TOKEN_URL"),
		},
	}

	return configs
}
