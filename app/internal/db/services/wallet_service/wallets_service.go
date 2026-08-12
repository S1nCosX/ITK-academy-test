package wallet_service

import (
	applogger "app/internal/app_logger"
	driver "app/internal/db/driver"
	"app/internal/db/dto"
	"fmt"
	"sync"
)

var (
	postgres *driver.PostgresDriver
	err      error
	logger   applogger.Apploger
	once     sync.Once
)

func Create() (ret string, err error) {
	once.Do(Init)
	query := "INSERT INTO wallets (walletId, balance) VALUES (uuidv4(), 0) RETURNING walletId"
	err = postgres.Conn.QueryRow(query).Scan(&ret)
	return ret, err
}

func Read(uuid *string) (ret dto.WalletDTO, err error) {
	once.Do(Init)
	query :=
		`SELECT 
	walletId,
	balance
FROM wallets
WHERE
	walletId = $1`

	err = postgres.Conn.QueryRow(query, uuid).Scan(
		&ret.WalletId,
		&ret.Balance)
	return ret, err
}

func Update(new *dto.WalletDTO) (ret dto.WalletDTO, err error) {
	once.Do(Init)
	query :=
		`UPDATE wallets  
	SET balance = $2
WHERE
	walletId = $1
RETURNING
	walletId,
	balance`

	err = postgres.Conn.QueryRow(query, new.WalletId, new.Balance).Scan(
		&ret.WalletId,
		&ret.Balance)
	return ret, err

}

func Init() {
	postgres, err = driver.Get()
	logger = applogger.NewLogger("Subscription service")

	if err != nil {
		logger.Error(fmt.Sprintf("Init error. Database driver getting got err: %s", err))
	}
	logger.Info("Initializated")
}
