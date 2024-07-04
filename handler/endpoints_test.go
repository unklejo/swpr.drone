package handler

import (
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/unklejo/swpr.drone/repository/mocks"
)

// 1. Create estate test files

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

	h.CreateEstate(c)

	assert.Equal(t, http.StatusCreated, rec.Code)
	assert.Contains(t, rec.Body.String(), "id")
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

	assert.Equal(t, http.StatusBadRequest, rec.Code)
	assert.Contains(t, rec.Body.String(), "Invalid input")
}

func TestCreateEstate_NegativeValues(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/estate", strings.NewReader(`{"width":-1, "length":-2}`))
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

	assert.Equal(t, http.StatusBadRequest, rec.Code)
	assert.Contains(t, rec.Body.String(), "Width and Length must be greater than 0")
}

func TestCreateEstate_InternalServerError(t *testing.T) {
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

	mockRepo.EXPECT().CreateEstate(gomock.Any(), 10, 10).Return(fmt.Errorf("some error"))

	h.CreateEstate(c)

	assert.Equal(t, http.StatusInternalServerError, rec.Code)
	assert.Contains(t, rec.Body.String(), "Failed to create estate")
}

// 2. Add tree test files
func TestAddTree_Success(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/estate/1/tree", strings.NewReader(`{"x": 1, "y": 1, "height": 10}`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockRepositoryInterface(ctrl)
	h := &Server{
		Repository: mockRepo,
	}

	mockRepo.EXPECT().AddTree(gomock.Any(), "1", 1, 1, 10).Return(nil)

	h.AddTreeToEstate(c)

	assert.Equal(t, http.StatusCreated, rec.Code)
	assert.Contains(t, rec.Body.String(), "id")
}

func TestAddTree_InvalidInput(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/estate/1/tree", strings.NewReader(`{"x": "invalid", "y": 1, "height": 10}`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockRepositoryInterface(ctrl)
	h := &Server{
		Repository: mockRepo,
	}

	h.AddTreeToEstate(c)

	assert.Equal(t, http.StatusBadRequest, rec.Code)
	assert.Contains(t, rec.Body.String(), "Invalid input")
}

func TestAddTree_EstateNotFound(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/estate/1/tree", strings.NewReader(`{"x": 1, "y": 1, "height": 10}`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockRepositoryInterface(ctrl)
	h := &Server{
		Repository: mockRepo,
	}

	mockRepo.EXPECT().AddTree(gomock.Any(), "1", 1, 1, 10).Return(errors.New("estate not found"))

	h.AddTreeToEstate(c)

	assert.Equal(t, http.StatusInternalServerError, rec.Code)
	assert.Contains(t, rec.Body.String(), "Failed to add tree")
}
