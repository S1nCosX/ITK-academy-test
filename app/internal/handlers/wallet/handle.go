package wallet

import (
	responsewriter "app/internal/response_writer"
	"net/http"
)

var base_addr = "api/v1/wallet"

func Handle(mux *http.ServeMux) {
	mux.HandleFunc("GET "+base_addr, postWallet)
}

func postWallet(w http.ResponseWriter, r *http.Request) {
	responsewriter.WriteResponse(w, "got u", http.StatusOK)
}
