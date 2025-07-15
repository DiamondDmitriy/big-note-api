package config

import (
	"github.com/joho/godotenv"
)

type (
	Config struct {
		HTTP HTTP
		DB   DB
	}

	HTTP struct {
		Port             string `env:"PORT" envDefault:"8080"`
		CorsAllowOrigins string `env:"CORS_ALLOW_ORIGINS" envDefault:""`
	}

	DB struct {
		DriverName string `env:"DB_DRIVER_NAME" envDefault:"postgres"`
		Host       string `env:"DB_HOST" envDefault:"localhost"`
		Port       string `env:"DB_PORT" envDefault:"5432"`
		Username   string `env:"DB_USERNAME" envDefault:"user"`
		Password   string `env:"DB_PASSWORD" envDefault:"password"`
		Database   string `env:"DB_DATABASE" envDefault:"postgres"`
	}
)

func NewConfig() (*Config, error) {
	envMap, err := godotenv.Read("config/.env")
	if err != nil {
		panic(err)
	}

	cfg := &Config{
		HTTP: HTTP{
			Port:             envMap["SERVER_PORT"],
			CorsAllowOrigins: envMap["CORS_ALLOW_ORIGINS"],
		},
		DB: DB{
			DriverName: envMap["DB_DRIVER_NAME"],
			Host:       envMap["DB_HOST"],
			Port:       envMap["DB_PORT"],
			Username:   envMap["DB_USERNAME"],
			Password:   envMap["DB_PASSWORD"],
			Database:   envMap["DB_DATABASE"],
		},
	}

	return cfg, nil
}
