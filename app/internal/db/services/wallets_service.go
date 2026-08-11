package services

import (
	applogger "app/internal/app_logger"
	driver "app/internal/db/driver"
	"log"
	"sync"
)

var (
	postgres *driver.PostgresDriver
	err      error
	logger   *log.Logger
	once     sync.Once
)

func Create() {

}

func Read() {

}

func Update() {

}

func init() {
	postgres, err = driver.Get()
	logger = applogger.NewLogger("Subscription service")

	if err != nil {
		logger.Panic("Init error. Database driver getting got err: ", err)
	}
}
