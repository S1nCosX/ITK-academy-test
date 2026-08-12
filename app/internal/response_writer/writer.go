package responsewriter

import (
	"log"
	"net/http"
)

func WriteResponse(w http.ResponseWriter, message string, code int) {
	w.WriteHeader(code)
	log.New(log.Default().Writer(), "da", 0).Print(message, code)
	w.Write([]byte(message))
}
