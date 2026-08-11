package healthcheck

import (
	responsewriter "app/internal/response_writer"
	"net/http"
)

func Handle(mux *http.ServeMux) {
	mux.HandleFunc("GET /health",
		healthcheck)
}

func healthcheck(w http.ResponseWriter, r *http.Request) {
	responsewriter.WriteResponse(w, "App is healthy", http.StatusOK)
}
