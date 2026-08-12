package wallets

import (
	"app/internal/db/dto"
	operationtype "app/internal/enums/operationType"
	wallets_test "app/internal/handlers/wallets/test"
	"encoding/json"
	"errors"
	"net/http"
	"testing"
)

func TestCreate(t *testing.T) {
	code := 0
	body := ""
	w := wallets_test.ResponseWriterMock{
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
		t.Log(*w.Body)
	}
	if *w.Code != http.StatusOK {
		t.Errorf("Test %d: Failed. Expected status %d, got: %d", mockService.CurrentTest, http.StatusOK, *w.Code)
		t.Log(*w.Body)
	}

	testData = wallets_test.CreateData{S: "", E: errors.New("boba")}
	createWallet(w, &r)
	mockService.CurrentTest += 1
	if *w.Code != http.StatusInternalServerError {
		t.Errorf("Test %d: Failed. Expected status %d, got: %d", mockService.CurrentTest, http.StatusInternalServerError, *w.Code)
		t.Log(*w.Body)
	}
}
func TestRead(t *testing.T) {
	code := 0
	body := ""
	w := wallets_test.ResponseWriterMock{
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
		t.Log(*w.Body)
	}
	if *w.Code != http.StatusOK {
		t.Errorf("Test %d: Failed. Expected status %d, got: %d", mockService.CurrentTest, http.StatusOK, *w.Code)
		t.Log(*w.Body)
	}

	testData = wallets_test.ReadUpdateData{W: dto.WalletDTO{"228", 0.}, E: errors.New("boba")}
	readWallet(w, &r)
	mockService.CurrentTest += 1
	if *w.Code != http.StatusNotFound {
		t.Errorf(" Test %d: When service read returns error, expecting code %d, but got %d", mockService.CurrentTest, http.StatusNotFound, *w.Code)
		t.Log(*w.Body)
	}
}

func TestUpdate(t *testing.T) {
	code := 0
	body := ""
	w := wallets_test.ResponseWriterMock{
		Code: &code, Body: &body,
	}
	r := http.Request{}
	mockService := wallets_test.ServiceMock{}
	mockService.CurrentTest = 0
	var updateTestData wallets_test.ReadUpdateData
	mockService.UpdateReturns = &updateTestData
	var readTestData wallets_test.ReadUpdateData
	mockService.UpdateReturns = &readTestData
	var RequestBody wallets_test.RequestBodyMock
	buffer := []byte{}
	RequestBody.Buffer = &buffer
	r.Body = RequestBody
	r.Header = make(http.Header)

	updateWallet(w, &r)
	mockService.CurrentTest += 1
	if *w.Code != http.StatusNotAcceptable {
		t.Errorf("Test %d: Failed. Expected error caused by inacceptable empty or non json|text content type header with code %d but got %d", mockService.CurrentTest, http.StatusNotAcceptable, *w.Code)
	}

	r.Header["Content-Type"] = []string{"application/json"}
	updateWallet(w, &r)
	mockService.CurrentTest += 1
	if *w.Code != http.StatusBadRequest {
		t.Errorf("Test %d: Failed. Expected error caused by inacceptable empty or unparsable body with code %d but got %d", mockService.CurrentTest, http.StatusBadRequest, *w.Code)
		t.Log(*w.Body)
	}

	r.Header["Content-Type"] = []string{"application/json"}
	updateWallet(w, &r)
	mockService.CurrentTest += 1
	if *w.Code != http.StatusBadRequest {
		t.Errorf("Test %d: Failed. Expected error caused by inacceptable empty or unparsable body with code %d but got %d", mockService.CurrentTest, http.StatusBadRequest, *w.Code)
		t.Log(*w.Body)
	}

	modification := WalletModification{
		WalletId:      "blah",
		OperationType: "Blahblah",
		Amount:        1,
	}
	{
		buf := []byte{}
		buf, err := json.Marshal(modification)
		if err != nil {
			panic(err)
		}
		RequestBody.Write(buf)
	}

	updateWallet(w, &r)
	mockService.CurrentTest += 1
	if *w.Code != http.StatusBadRequest {
		t.Errorf("Test %d: Failed. Expected error caused by wrong operation type in body with code %d but got %d", mockService.CurrentTest, http.StatusBadRequest, *w.Code)
		t.Log(*w.Body)
	}

	readTestData = wallets_test.ReadUpdateData{
		W: dto.WalletDTO{
			WalletId: "blah",
			Balance:  0,
		},
		E: errors.New("blah"),
	}
	{
		buf := []byte{}
		buf, err := json.Marshal(modification)
		if err != nil {
			panic(err)
		}
		RequestBody.Write(buf)
	}
	updateWallet(w, &r)
	mockService.CurrentTest += 1
	if *w.Code != http.StatusBadRequest {
		t.Errorf("Test %d: Failed. Expected error caused by error from service read with code %d but got %d", mockService.CurrentTest, http.StatusBadRequest, *w.Code)
		t.Log(*w.Body)
	}

	var err error
	modification.OperationType, err = operationtype.ToString(operationtype.WITHDRAW)

	if err != nil {
		panic(err)
	}
	modification.Amount = 100
	{
		buf := []byte{}
		buf, err := json.Marshal(modification)
		if err != nil {
			panic(err)
		}
		RequestBody.Write(buf)
	}
	updateWallet(w, &r)
	mockService.CurrentTest += 1
	t.Log(*w.Body)
	if *w.Code != http.StatusBadRequest {
		t.Errorf("Test %d: Failed. Expected error caused by not enough money header with code %d but got %d", mockService.CurrentTest, http.StatusBadRequest, *w.Code)
		t.Log(*w.Body)
	}

	modification.OperationType, err = operationtype.ToString(operationtype.DEPOSIT)
	if err != nil {
		panic(err)
	}

	updateTestData = wallets_test.ReadUpdateData{
		W: dto.WalletDTO{
			WalletId: "blah",
			Balance:  0,
		},
		E: errors.New("blah"),
	}

	{
		buf := []byte{}
		buf, err := json.Marshal(modification)
		if err != nil {
			panic(err)
		}
		RequestBody.Write(buf)
	}
	updateWallet(w, &r)
	mockService.CurrentTest += 1
	if *w.Code != http.StatusInternalServerError {
		t.Errorf("Test %d: Failed. Expected error caused by error from service update with code %d but got %d", mockService.CurrentTest, http.StatusInternalServerError, *w.Code)
		t.Log(*w.Body)
	}

	updateTestData = wallets_test.ReadUpdateData{
		W: dto.WalletDTO{
			WalletId: "blah",
			Balance:  0,
		},
		E: nil,
	}

	{
		buf := []byte{}
		buf, err := json.Marshal(modification)
		if err != nil {
			panic(err)
		}
		RequestBody.Write(buf)
	}
	updateWallet(w, &r)
	mockService.CurrentTest += 1
	if *w.Code != http.StatusOK {
		t.Errorf("Test %d: Failed. That if everything is ok, we'll get code %d but got %d", mockService.CurrentTest, http.StatusOK, *w.Code)
		t.Log(*w.Body)
	}
	{
		buf := []byte{}
		buf, err := json.Marshal(modification)
		if err != nil {
			panic(err)
		}
		RequestBody.Write(buf)
	}
	updateWallet(w, &r)
	mockService.CurrentTest += 1
	if mockService.UpdatedWallet.Balance != 101 {
		t.Errorf("Test %d: Failed. That if previous we had 1 and deposit 100, we'll get 101 but got %f", mockService.CurrentTest, mockService.UpdatedWallet.Balance)
		t.Log(*w.Body)
	}
}
