package dto

import (
	"app/internal/enums"
)

type WalletDTO struct {
	valletId      string              `json:"valletId"`
	operationType enums.OperationType `json:"operationType"`
	amount        uint64              `json:"amount"`
}
