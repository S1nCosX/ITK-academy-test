package wallets

import (
	applogger "app/internal/app_logger"
	"app/internal/db/dto"
	"app/internal/db/services/wallet_service"
	operationtype "app/internal/enums/operationType"
	responsewriter "app/internal/response_writer"
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"strings"
)

var (
	baseAddr = "/api/v1/wallets"
	logger   applogger.Apploger
	service  wallet_service.WalletServiceInterface
)

func Init(new_service wallet_service.WalletServiceInterface) {
	logger = applogger.NewLogger("Wallet handler")
	service = new_service
}

func Handle(mux *http.ServeMux) {
	mux.HandleFunc("PUT "+baseAddr, createWallet)
	logger.Info("Listening PUT on addr: " + baseAddr)

	mux.HandleFunc("GET "+baseAddr+"/{WALLET_UUID}", readWallet)
	logger.Info("Listening GET on addr: " + baseAddr + "/{WALLET_UUID}")

	mux.HandleFunc("POST "+baseAddr, updateWallet)
	logger.Info("Listening POST on addr: " + baseAddr)

}

func createWallet(w http.ResponseWriter, r *http.Request) {
	id, err := service.Create()
	if err != nil {
		responsewriter.WriteResponse(w, fmt.Sprintf("DB error: %s", err), http.StatusInternalServerError)
		return
	}
	logger.Info(fmt.Sprint("Created new wallet", id))
	responsewriter.WriteResponse(w, id, http.StatusOK)
}

func readWallet(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("WALLET_UUID")
	wallet, err := service.Read(&id)
	if err != nil {
		responsewriter.WriteResponse(w, "No such wallet", http.StatusNotFound)
		return
	}
	encodeAndResponseWalletsDTO(&wallet, w)
}

func updateWallet(w http.ResponseWriter, r *http.Request) {
	containsJson := false
	for _, contentType := range strings.Split(r.Header.Get("Content-Type"), ";") {
		if match, err := regexp.Match("application/(json|text)", []byte(contentType)); match && err == nil {
			containsJson = true
		}
	}

	if !containsJson {
		responsewriter.WriteResponse(w, "Does not contains json or text", http.StatusNotAcceptable)
		return
	}

	var modification WalletModification
	if err := json.NewDecoder(r.Body).Decode(&modification); err != nil {
		responsewriter.WriteResponse(w, fmt.Sprint("Got error in decoding json: ", err), http.StatusBadRequest)
		return
	}

	wallet, err := service.Read(&modification.WalletId)
	if err != nil {
		responsewriter.WriteResponse(w, "No such wallet", http.StatusBadRequest)
		return
	}

	operation, err := operationtype.FromString(modification.OperationType)
	if err != nil {
		responsewriter.WriteResponse(w, "invalid operation type", http.StatusBadRequest)
		return
	}

	switch operation {
	case operationtype.DEPOSIT:
		wallet.Balance += modification.Amount
		break
	case operationtype.WITHDRAW:
		if wallet.Balance < modification.Amount {
			responsewriter.WriteResponse(w, "Not enough money to withdraw", http.StatusBadRequest)
			return
		}
		wallet.Balance -= modification.Amount
		break
	}
	if updatedWallet, err := service.Update(&wallet); err != nil {
		responsewriter.WriteResponse(w, fmt.Sprint("DB error: ", err), http.StatusInternalServerError)
		return
	} else {
		encodeAndResponseWalletsDTO(&updatedWallet, w)
	}
}

func encodeAndResponseWalletsDTO(wallet *dto.WalletDTO, w http.ResponseWriter) {
	encoded, err := json.Marshal(wallet)

	if err != nil {
		responsewriter.WriteResponse(w, "Encoding answer error", http.StatusInternalServerError)
		return
	}
	responsewriter.WriteResponse(w, string(encoded), http.StatusOK)
}
