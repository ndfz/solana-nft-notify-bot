package config

import (
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	AppEnv      string `required:"true" envconfig:"APP_ENV"`
	TgBotToken  string `required:"true" envconfig:"TELEGRAM_BOT_TOKEN"`
	DatabaseUrl string `required:"true" envconfig:"DATABASE_URL"`
}

var getWd = os.Getwd
var processEnv = envconfig.Process

func New(customPath *string) (*Config, error) {
	var newCfg Config

	wd, err := getWd()
	if err != nil {
		return nil, err
	}

	envPath := filepath.Join(wd, ".env")

	if customPath != nil {
		envPath = *customPath
	}

	_ = godotenv.Overload(envPath)
	if err = processEnv("", &newCfg); err != nil {
		return nil, err
	}

	return &newCfg, nil
}
