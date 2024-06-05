package config

import (
	"log/slog"

	"github.com/knadh/koanf/parsers/dotenv"
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
)

type Config struct {
	Port string `koanf:"PORT"`
}

var Conf *Config

var k = koanf.New(".")

func ProvideConfig() (*Config, error) {
	if Conf == nil {
		// try to load from .env files
		if err := k.Load(file.Provider(".env"), dotenv.Parser()); err != nil {
			slog.Error("Unable to find .env file", err)
			slog.Info("falling back to env variables")

			// fall back to env variables
			if err := k.Load(env.Provider("", ".", nil), nil); err != nil {
				slog.Error("error loading config", err)
				return nil, err
			}
		}

		// populate config struct
		Conf = &Config{}
		if err := k.Unmarshal("", Conf); err != nil {
			slog.Error("error loading config", err)
			return nil, err
		}
	}

	return Conf, nil
}
