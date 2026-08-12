package responsewriter

import (
	"net/http"
)

func WriteResponse(w http.ResponseWriter, message string, code int) {
	w.WriteHeader(code)
	w.Write([]byte(message))
}
