package wallet_service

import (
	applogger "app/internal/app_logger"
	driver "app/internal/db/driver"
	"app/internal/db/dto"
	"fmt"
	"sync"
)

type WalletService struct {
	postgres *driver.PostgresDriver
	err      error
	logger   applogger.Apploger
	once     sync.Once
}

var (
	instance WalletService
)

func (self WalletService) Create() (ret string, err error) {
	self.once.Do(self.init)
	query := "INSERT INTO wallets (walletId, balance) VALUES (uuidv4(), 0) RETURNING walletId"
	err = self.postgres.Conn.QueryRow(query).Scan(&ret)
	return ret, err
}

func (self WalletService) Read(uuid *string) (ret dto.WalletDTO, err error) {
	self.once.Do(self.init)
	query :=
		`SELECT 
	walletId,
	balance
FROM wallets
WHERE
	walletId = $1`

	err = self.postgres.Conn.QueryRow(query, uuid).Scan(
		&ret.WalletId,
		&ret.Balance)
	return ret, err
}

func (self WalletService) Update(new *dto.WalletDTO) (ret dto.WalletDTO, err error) {
	self.once.Do(self.init)
	query :=
		`UPDATE wallets  
	SET balance = $2
WHERE
	walletId = $1
RETURNING
	walletId,
	balance`

	err = self.postgres.Conn.QueryRow(query, new.WalletId, new.Balance).Scan(
		&ret.WalletId,
		&ret.Balance)
	return ret, err

}

func (self WalletService) init() {
	self.postgres, self.err = driver.Get()
	self.logger = applogger.NewLogger("Subscription service")

	if self.err != nil {
		self.logger.Error(fmt.Sprintf("Init error. Database driver getting got err: %s", self.err))
	}
	self.logger.Info("Initializated")
}

func GetWalletService() *WalletService {
	return &instance
}
