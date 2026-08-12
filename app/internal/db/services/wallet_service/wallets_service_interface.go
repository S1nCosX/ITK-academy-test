package wallet_service

import "app/internal/db/dto"

type WalletServiceInterface interface {
	Create() (ret string, err error)
	Read(uuid *string) (ret dto.WalletDTO, err error)
	Update(new *dto.WalletDTO) (ret dto.WalletDTO, err error)
}
