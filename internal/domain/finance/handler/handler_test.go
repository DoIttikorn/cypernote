package handler

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/Doittikorn/cypernote/internal/domain/finance"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

type FinanceUsecaseMock struct {
}

func (f *FinanceUsecaseMock) ExecuteSave(m *finance.FinanceRequest) (*finance.FinanceResponse, error) {
	return &finance.FinanceResponse{
		ID:         1,
		Amount:     1,
		Type:       "income",
		Note:       "จ่ายค่าข้าว",
		DateTimeAt: time.Now(),
		Status:     "Y",
	}, nil
}

func (f *FinanceUsecaseMock) ExecuteGetByUserID(userID int64, filter *finance.Filter) ([]finance.FinanceResponse, error) {
	return nil, errors.New("Not implemented")
}

func (f *FinanceUsecaseMock) ExecuteUpdate() {

}
func (f *FinanceUsecaseMock) ExecuteDelete() {

}

func TestFinanceHandler_Save(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", bytes.NewReader([]byte(`{
		"user_id": 1,
		"amount": 1,
		"type": "income",
		"note": "จ่ายค่าข้าว",
		"datetime_at": "2023-02-17T09:18:57.829Z"
	}`)))
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	handler := New(&FinanceUsecaseMock{})
	err := handler.Save(c)

	if assert.NoError(t, err) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		var m finance.M
		json.Unmarshal(rec.Body.Bytes(), &m)
		assert.Equal(t, "Y", m.Status)
	}
}

func TestFinanceHandler_GetByUserID(t *testing.T) {

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	handler := New(&FinanceUsecaseMock{})
	err := handler.GetByUserID(c)

	if assert.NoError(t, err) {
		assert.Equal(t, http.StatusOK, rec.Code)

		var m []finance.M
		json.Unmarshal(rec.Body.Bytes(), &m)
		assert.Equal(t, 1, len(m))

	}
}
