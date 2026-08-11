package config

import (
	applogger "app/src/app_logger"
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"
)

type Config struct {
	POSTGRES_DB        string
	POSTRGES_USER      string
	POSTGRES_PASS_FILE string

	DB_HOST string
	DB_PORT uint

	POSTGRES_PASS string

	HOST string
	PORT uint
}

var (
	instance Config
	once     sync.Once
	initErr  error
	logger   *log.Logger
)

func getStringEnv(name string, defaultValue string) string {
	value, isExist := os.LookupEnv(name)
	if isExist {
		applogger.Warning(logger, fmt.Sprintf("Cannot read environment %s, so using default value: %s", name, defaultValue))
		return value
	}
	return defaultValue
}

func getUintEnv(name string, defaultValue uint) uint {
	value, isExist := os.LookupEnv(name)
	if isExist {
		applogger.Warning(logger, fmt.Sprintf("Cannot read environment %s. So using default value: %s", name, defaultValue))

		parsedValue, err := strconv.ParseUint(value, 10, 32)

		if err != nil {
			applogger.Warning(logger, fmt.Sprintf("Environment variable %s contains not uint value: %s. So using default value instead: %s", name, value, defaultValue))
			return defaultValue
		}
		return uint(parsedValue)
	}
	return defaultValue
}

func (*Config) init() {
	logger = applogger.NewLogger("Config")

	instance = Config{
		POSTGRES_DB:        getStringEnv("POSGRES_DB", "default"),
		POSTRGES_USER:      getStringEnv("POSGRES_USER", "postgres"),
		POSTGRES_PASS_FILE: getStringEnv("POSTGRES_PASS_FILE", "/run/secrets/db_pw"),
		DB_HOST:            getStringEnv("DB_HOST", "db"),
		DB_PORT:            getUintEnv("DB_PORT", 5432),
		HOST:               getStringEnv("HOST", "0.0.0.0"),
		PORT:               getUintEnv("PORT", 8080),
	}

	password_bytes, err := os.ReadFile(instance.POSTGRES_PASS_FILE)

	if err != nil {
		applogger.Warning(logger, fmt.Sprintf("Got error: %s", err.Error()))
		instance.POSTGRES_PASS = "defaultpass"
		initErr = err
	} else {
		instance.POSTGRES_PASS = string(password_bytes)
	}
}

func Get() (*Config, error) {
	once.Do(instance.init)
	return &instance, initErr
}
