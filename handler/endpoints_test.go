package handler

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/unklejo/swpr.drone/repository/mocks"
)

func TestCreateEstate_Success(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/estate", strings.NewReader(`{"width":10, "length":10}`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockRepositoryInterface(ctrl)
	h := &Server{
		Repository: mockRepo,
	}

	mockRepo.EXPECT().CreateEstate(gomock.Any(), 10, 10).Return(nil)

	if assert.NoError(t, h.CreateEstate(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.Contains(t, rec.Body.String(), "id")
	}
}

func TestCreateEstate_InvalidInput(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/estate", strings.NewReader(`{"width":"xxx", "length":"yyy"}`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockRepositoryInterface(ctrl)
	h := &Server{
		Repository: mockRepo,
	}

	h.CreateEstate(c)

	// // Print debug output
	// fmt.Printf("Response Code: %d, Response Body: %s\n", rec.Code, rec.Body.String())

	assert.Equal(t, http.StatusBadRequest, rec.Code)
	assert.Contains(t, rec.Body.String(), "Invalid input")
}
