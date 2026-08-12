package wallets

type WalletModification struct {
	WalletId      string  `json:walletId`
	OperationType string  `json:operationType`
	Amount        float32 `json:amount`
}
