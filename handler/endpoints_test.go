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

func TestCreateEstate(t *testing.T) {
	// Setup
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

	// Expectations
	mockRepo.EXPECT().CreateEstate(gomock.Any(), 10, 10).Return(nil)

	// Assertions
	if assert.NoError(t, h.CreateEstate(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.Contains(t, rec.Body.String(), "id")
	}
}
