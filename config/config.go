package config

import (
	"errors"
	"github.com/caarlos0/env/v6"
)

type Config struct {
	HTTP     HTTP
	Database Database
}
type HTTP struct {
	PORT string `env:"PORT" envDefault:"8080"`
	URL  string `env:"URL" envDefault:"localhost"`
}
type Database struct {
	HOST     string `env:"DB_HOST" envDefault:"localhost"`
	PORT     string `env:"DB_PORT" envDefault:"5436"`
	USER     string `env:"DB_USER" envDefault:"postgres"`
	PASSWORD string `env:"DB_PASSWORD" envDefault:"qwerty"`
	DB       string `env:"DB_NAME" envDefault:"postgres"`
	PgUrl    string `env:"PG_URL" envDefault:"host=localhost port=5436 user=postgres password=qwerty dbname=postgres sslmode=disable"`
}

func New() (*Config, error) {
	cfg := Config{}
	if err := env.Parse(&cfg); err != nil {
		return nil, errors.New("cfg not created")
	}
	return &cfg, nil
}
