package postgres_driver

import (
	applogger "app/internal/app_logger"
	"app/internal/config"
	"database/sql"
	"fmt"
	"sync"

	_ "github.com/lib/pq"
)

type PostgresDriver struct {
	Conn *sql.DB
}

var (
	once     sync.Once
	instance PostgresDriver
	initErr  error
	logger   applogger.Apploger
)

func (*PostgresDriver) init() {
	logger = applogger.NewLogger("Postgres driver")

	conf, err := config.Get()
	if err != nil {
		logger.Error("Config initiated with error")
	}

	conn_str := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", conf.DB_HOST, conf.DB_PORT, conf.POSTRGES_USER, conf.POSTGRES_PASSWORD, conf.POSTGRES_DB)

	instance.Conn, err = sql.Open("postgres", conn_str)
	if err != nil {
		logger.Error(fmt.Sprintf("Got error during db connection: %s", err))
	}
}

func Get() (*PostgresDriver, error) {
	once.Do(instance.init)
	return &instance, initErr
}
