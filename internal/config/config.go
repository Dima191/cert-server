package config

import (
	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Port      int    `env:"PORT" env-required:"true"`
	DBConnStr string `env:"DB_URL" env-required:"true"`
}

func New(configPath string) (*Config, error) {
	var cfg Config
	err := cleanenv.ReadConfig(configPath, &cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
