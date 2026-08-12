package wallets_test

import "io"

type RequestBodyMock struct {
	Buffer *[]byte
}

func (self RequestBodyMock) Write(b []byte) {
	*self.Buffer = append(*self.Buffer, b...)
}

func (self RequestBodyMock) Read(b []byte) (n int, err error) {
	if len(*self.Buffer) == 0 {
		return 0, io.EOF
	}
	b = append(b, []byte(*self.Buffer)...)
	n = len(*self.Buffer)
	*self.Buffer = []byte{}
	return n, nil
}

func (self RequestBodyMock) Close() error {
	return nil
}
