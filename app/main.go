package main

import (
	applogger "app/internal/app_logger"
	"app/internal/config"
	"app/internal/handlers/healthcheck"
	"app/internal/handlers/wallets"
	"fmt"
	"net/http"
)

func main() {
	logger := applogger.NewLogger("main")

	mux := http.NewServeMux()

	healthcheck.Handle(mux)
	wallets.Handle(mux)

	config, err := config.Get()
	if err != nil {
		logger.Error(fmt.Sprintf("Got error during geting config: %s", err))
	}

	addr := fmt.Sprintf("%s:%d", config.HOST, config.PORT)
	logger.Info(fmt.Sprintf("Starting server on addr: %s", addr))
	http.ListenAndServe(addr, mux)
}
