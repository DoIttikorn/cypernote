package https

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/Doittikorn/cypernote/internal/domain/finance"
	"github.com/Doittikorn/cypernote/pkg/errorc"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

type FinanceUsecaseMock struct {
}

func (f *FinanceUsecaseMock) Save(m *finance.M) error {
	return nil
}

func (f *FinanceUsecaseMock) ExecuteGetByUserID(userID int64, filter *finance.Filter) ([]finance.M, error) {
	return nil, errors.New("Not implemented")
}

func (f *FinanceUsecaseMock) Update() {

}
func (f *FinanceUsecaseMock) Delete() {

}

func TestFinanceHandler_Save(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", bytes.NewReader([]byte(`{}`)))
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	handler := New(&FinanceUsecaseMock{})
	err := handler.save(c)

	if assert.NoError(t, err) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		var m finance.M
		json.Unmarshal(rec.Body.Bytes(), &m)
		assert.Equal(t, "Y", m.Status)
	}
}

func TestFinanceHandler_GetByUserID(t *testing.T) {
	e := echo.New()
	userID := int64(123)
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues(strconv.FormatInt(userID, 10))

	handler := New(&FinanceUsecaseMock{})
	err := handler.getByUserID(c)

	if assert.Error(t, err) {
		httpError := err.(*echo.HTTPError)
		assert.Equal(t, http.StatusInternalServerError, httpError.Code)
		assert.Equal(t, errorc.ResponseErr("Not implemented").Error(), httpError.Message)
	}
}
