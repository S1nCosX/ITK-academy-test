package healthcheck

import (
	applogger "app/internal/app_logger"
	responsewriter "app/internal/response_writer"
	"net/http"
)

var logger applogger.Apploger

func Handle(mux *http.ServeMux) {
	logger = applogger.NewLogger("Health")
	mux.HandleFunc("GET /health", healthcheck)
	logger.Info("Listening GET on addr: /health")
}

func healthcheck(w http.ResponseWriter, r *http.Request) {
	responsewriter.WriteResponse(w, "App is healthy", http.StatusOK)
}
