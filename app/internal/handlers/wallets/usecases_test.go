package wallets

import (
	"app/internal/db/dto"
	wallets_test "app/internal/handlers/wallets/test"
	"errors"
	"net/http"
	"testing"
)

func TestCreate(t *testing.T) {
	code := 0
	body := ""
	w := wallets_test.MockRW{
		Code: &code, Body: &body,
	}
	r := http.Request{}
	mockService := wallets_test.ServiceMock{}
	mockService.CurrentTest = 0
	var testData wallets_test.CreateData
	mockService.CreateReturns = &testData

	Init(mockService)

	testData = wallets_test.CreateData{S: "228", E: nil}

	createWallet(w, &r)
	mockService.CurrentTest += 1
	if *w.Body != "228" {
		t.Errorf("Test %d: Failed. Expected body contains 228, got: %s", mockService.CurrentTest, *w.Body)
	}
	if *w.Code != http.StatusOK {
		t.Errorf("Test %d: Failed. Expected status %d, got: %d", mockService.CurrentTest, http.StatusOK, *w.Code)
	}

	testData = wallets_test.CreateData{S: "", E: errors.New("boba")}
	createWallet(w, &r)
	mockService.CurrentTest += 1
	if *w.Code != http.StatusInternalServerError {
		t.Errorf("Test %d: Failed. Expected status %d, got: %d", mockService.CurrentTest, http.StatusInternalServerError, *w.Code)
	}
}
func TestRead(t *testing.T) {
	code := 0
	body := ""
	w := wallets_test.MockRW{
		Code: &code, Body: &body,
	}
	r := http.Request{}
	mockService := wallets_test.ServiceMock{}
	mockService.CurrentTest = 0
	var testData wallets_test.ReadUpdateData
	mockService.ReadReturns = &testData

	Init(mockService)

	r.SetPathValue("WALLET_UUID", "228")
	testData = wallets_test.ReadUpdateData{W: dto.WalletDTO{"228", 0.}, E: nil}
	readWallet(w, &r)
	mockService.CurrentTest += 1
	if len([]byte(*w.Body)) == 0 {
		t.Errorf("Test %d: Failed. Expected body contains data, got nothing", mockService.CurrentTest)
	}
	if *w.Code != http.StatusOK {
		t.Errorf("Test %d: Failed. Expected status %d, got: %d", mockService.CurrentTest, http.StatusOK, *w.Code)
	}

	testData = wallets_test.ReadUpdateData{W: dto.WalletDTO{"228", 0.}, E: errors.New("boba")}
	readWallet(w, &r)
	mockService.CurrentTest += 1
	if *w.Code != http.StatusNotFound {
		t.Errorf(" Test %d: When service read returns error, expecting code %d, but got %d", mockService.CurrentTest, http.StatusNotFound, *w.Code)
	}
}

func TestUpdate(t *testing.T) {
	code := 0
	body := ""
	w := wallets_test.MockRW{
		Code: &code, Body: &body,
	}
	r := http.Request{}
	mockService := wallets_test.ServiceMock{}
	mockService.CurrentTest = 0
	var testData wallets_test.ReadUpdateData
	mockService.UpdateReturns = &testData

	updateWallet(w, &r)
	mockService.CurrentTest += 1
	if *w.Code != http.StatusNotAcceptable {
		t.Errorf("Test %d: Failed. Expected error cause of inacceptable empty or non json|text content type header with code %d but got %d", mockService.CurrentTest, http.StatusNotAcceptable, *w.Code)
	}
}
