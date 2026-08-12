package wallets_test

import (
	"app/internal/db/dto"
)

type CreateData struct {
	S string
	E error
}

type ReadUpdateData struct {
	W dto.WalletDTO
	E error
}

type ServiceMock struct {
	CurrentTest   int
	CreateReturns *CreateData
	ReadReturns   *ReadUpdateData
	UpdateReturns *ReadUpdateData
	UpdatedWallet *dto.WalletDTO
}

func (self ServiceMock) Create() (ret string, err error) {
	return self.CreateReturns.S, self.CreateReturns.E
}

func (self ServiceMock) Read(uuid *string) (ret dto.WalletDTO, err error) {
	return self.ReadReturns.W, self.ReadReturns.E
}
func (self ServiceMock) Update(new *dto.WalletDTO) (ret dto.WalletDTO, err error) {
	*self.UpdatedWallet = *new
	return self.UpdateReturns.W, self.UpdateReturns.E
}
