package config

import (
	"os"
	"path/filepath"
	"time"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	AppEnv            string        `required:"true" envconfig:"APP_ENV"`
	CollectionSleep   time.Duration `required:"true" envconfig:"COLLECTION_SLEEP"`
	CycleSleep        time.Duration `required:"true" envconfig:"CYCLE_SLEEP"`
	TgBotToken        string        `required:"true" envconfig:"TELEGRAM_BOT_TOKEN"`
	DatabaseUrl       string        `required:"true" envconfig:"DATABASE_URL"`
	MagicEdenEndpoint string        `required:"true" envconfig:"MAGIC_EDEN_ENDPOINT"`
}

var getWd = os.Getwd
var processEnv = envconfig.Process

func New() (*Config, error) {
	var newCfg Config

	wd, err := getWd()
	if err != nil {
		return nil, err
	}

	envPath := filepath.Join(wd, ".env")

	_ = godotenv.Overload(envPath)
	if err = processEnv("", &newCfg); err != nil {
		return nil, err
	}

	return &newCfg, nil
}
