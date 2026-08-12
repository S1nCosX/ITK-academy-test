package wallets_test

import "net/http"

type ResponseWriterMock struct {
	Code *int
	Body *string
}

func (ResponseWriterMock) Header() http.Header {
	return http.Header{}
}

func (self ResponseWriterMock) Write(inp []byte) (int, error) {
	*self.Body = string(inp)
	return len(*self.Body), nil
}

func (self ResponseWriterMock) WriteHeader(statusCode int) {
	*self.Code = statusCode
}
