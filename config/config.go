package config

import (
	"github.com/joho/godotenv"
	"strconv"
)

type (
	Config struct {
		APP  APP
		HTTP HTTP
		DB   DB
		JWT  JWT
	}

	APP struct {
		Name    string `env:"APP_NAME"`
		Version string `env:"APP_VERSION"`
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

	JWT struct {
		TokenPassword          []byte `env:"JWT_TOKEN_PASSWORD" envDefault:"secret"`
		TokenValidityPeriodMin int    `env:"JWT_TOKEN_VALIDITY_PERIOD" envDefault:"15"`
	}
)

func NewConfig() (*Config, error) {
	// todo: автоматизировать
	envMap, err := godotenv.Read("config/.env")
	if err != nil {
		panic(err)
	}

	tokenValidityPeriodMin, err := strconv.Atoi(envMap["JWT_TOKEN_VALIDITY_PERIOD_MIN"])
	if err != nil {
		panic(err)
	}

	cfg := &Config{
		APP: APP{
			Name:    envMap["APP_NAME"],
			Version: envMap["APP_VERSION"],
		},
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
		JWT: JWT{
			TokenPassword:          []byte(envMap["JWT_TOKEN_PASSWORD"]),
			TokenValidityPeriodMin: tokenValidityPeriodMin,
		},
	}

	return cfg, nil
}
