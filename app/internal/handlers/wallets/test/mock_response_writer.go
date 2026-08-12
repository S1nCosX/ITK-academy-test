package wallets_test

import "net/http"

type MockRW struct {
	Code *int
	Body *string
}

func (MockRW) Header() http.Header {
	return http.Header{}
}

func (self MockRW) Write(inp []byte) (int, error) {
	*self.Body = string(inp)
	return len(*self.Body), nil
}

func (self MockRW) WriteHeader(statusCode int) {
	*self.Code = statusCode
}
