package main

import (
	applogger "app/internal/app_logger"
	"app/internal/config"
	"app/internal/handlers/healthcheck"
	"fmt"
	"net/http"
)

func main() {
	logger := applogger.NewLogger("main")

	mux := http.NewServeMux()

	healthcheck.Handle(mux)

	config, err := config.Get()
	if err != nil {
		applogger.Error(logger, fmt.Sprintf("Got error during geting config: %s", err))
	}

	addr := fmt.Sprintf("%s:%d", config.HOST, config.PORT)
	http.ListenAndServe(addr, mux)
	applogger.Info(logger, fmt.Sprintf("Started server on addr: %s", addr))
}
