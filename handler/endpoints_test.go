package handler

import (
	"database/sql"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	"github.com/unklejo/swpr.drone/repository"
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

	mockRepo := repository.NewMockRepositoryInterface(ctrl)
	h := &Server{
		Repository: mockRepo,
	}

	mockRepo.EXPECT().CreateEstate(10, 10).Return("1", nil)

	h.PostEstate(c)

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

	mockRepo := repository.NewMockRepositoryInterface(ctrl)
	h := &Server{
		Repository: mockRepo,
	}

	h.PostEstate(c)

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

	mockRepo := repository.NewMockRepositoryInterface(ctrl)
	h := &Server{
		Repository: mockRepo,
	}

	h.PostEstate(c)

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

	mockRepo := repository.NewMockRepositoryInterface(ctrl)
	h := &Server{
		Repository: mockRepo,
	}

	mockRepo.EXPECT().CreateEstate(10, 10).Return("", repository.ErrDatabaseError)

	h.PostEstate(c)

	assert.Equal(t, http.StatusInternalServerError, rec.Code)
	assert.Contains(t, rec.Body.String(), "Failed to create estate")
}

// 2. Add tree test files
func TestAddTree_Success(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/estate/1/tree", strings.NewReader(`{"x": 1, "y": 10, "height": 10}`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := repository.NewMockRepositoryInterface(ctrl)
	h := &Server{
		Repository: mockRepo,
	}

	mockRepo.EXPECT().GetEstateById("1").Return(repository.Estate{Id: "1", Width: 10, Length: 10}, nil)
	mockRepo.EXPECT().AddTree(gomock.Any(), 1, 10, 10).Return("1", nil)

	h.PostEstateIdTree(c, uuid.Nil)

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

	mockRepo := repository.NewMockRepositoryInterface(ctrl)
	h := &Server{
		Repository: mockRepo,
	}

	h.PostEstateIdTree(c, uuid.Nil)

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

	mockRepo := repository.NewMockRepositoryInterface(ctrl)
	h := &Server{
		Repository: mockRepo,
	}

	mockRepo.EXPECT().GetEstateById("1").Return(repository.Estate{Id: "", Width: 0, Length: 0}, sql.ErrNoRows)

	h.PostEstateIdTree(c, uuid.Nil)

	assert.Equal(t, http.StatusNotFound, rec.Code)
	assert.Contains(t, rec.Body.String(), "Estate not found")
}

func TestAddTree_DatabaseError(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/estate/1/tree", strings.NewReader(`{"x": 1, "y": 1, "height": 10}`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := repository.NewMockRepositoryInterface(ctrl)
	h := &Server{
		Repository: mockRepo,
	}

	mockRepo.EXPECT().GetEstateById("1").Return(repository.Estate{Id: "1", Width: 10, Length: 10}, nil)
	mockRepo.EXPECT().AddTree(gomock.Any(), 1, 1, 10).Return("", repository.ErrDatabaseError)

	h.PostEstateIdTree(c, uuid.Nil)

	assert.Equal(t, http.StatusInternalServerError, rec.Code)
	assert.Contains(t, rec.Body.String(), "Failed to add tree")
}

func TestAddTree_PlotAlreadyHasTree(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/estate/1/tree", strings.NewReader(`{"x": 1, "y": 1, "height": 10}`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := repository.NewMockRepositoryInterface(ctrl)
	h := &Server{
		Repository: mockRepo,
	}

	mockRepo.EXPECT().GetEstateById("1").Return(repository.Estate{Id: "1", Width: 10, Length: 10}, nil)
	mockRepo.EXPECT().AddTree(gomock.Any(), 1, 1, 10).Return("", &pq.Error{Code: "23505"})

	h.PostEstateIdTree(c, uuid.Nil)

	assert.Equal(t, http.StatusBadRequest, rec.Code)
	assert.Contains(t, rec.Body.String(), "Plot already has a tree")
}

func TestAddTree_CoordinatesOutOfBounds(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/estate/1/tree", strings.NewReader(`{"x": 11, "y": 8, "height": 10}`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := repository.NewMockRepositoryInterface(ctrl)
	mockRepo.EXPECT().GetEstateById("1").Return(repository.Estate{Id: "1", Width: 10, Length: 10}, nil)

	h := &Server{Repository: mockRepo}

	h.PostEstateIdTree(c, uuid.Nil)

	assert.Equal(t, http.StatusBadRequest, rec.Code)
	assert.Contains(t, rec.Body.String(), "Coordinates out of bounds")
}

// 3. Get Estate test files
func TestGetEstateStats_Success(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/estate/1/stats", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := repository.NewMockRepositoryInterface(ctrl)
	h := &Server{
		Repository: mockRepo,
	}

	mockRepo.EXPECT().GetEstateById("1").Return(repository.Estate{Id: "1", Width: 10, Length: 10}, nil)
	mockRepo.EXPECT().GetEstateStatsById("1").Return(repository.EstateStats{Count: 3, MaxHeight: 20, MinHeight: 5, MedianHeight: 15}, nil)

	h.GetEstateIdStats(c, uuid.Nil)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Contains(t, rec.Body.String(), `"count":3`)
	assert.Contains(t, rec.Body.String(), `"max":20`)
	assert.Contains(t, rec.Body.String(), `"min":5`)
	assert.Contains(t, rec.Body.String(), `"median":15`)
}

func TestGetEstateStats_NoTreesFound(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/estate/1/stats", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := repository.NewMockRepositoryInterface(ctrl)
	h := &Server{
		Repository: mockRepo,
	}

	mockRepo.EXPECT().GetEstateById("1").Return(repository.Estate{Id: "1", Width: 10, Length: 10}, nil)
	mockRepo.EXPECT().GetEstateStatsById("1").Return(repository.EstateStats{Count: 0, MaxHeight: 0, MinHeight: 0, MedianHeight: 0}, nil)

	h.GetEstateIdStats(c, uuid.Nil)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Contains(t, rec.Body.String(), `"count":0`)
	assert.Contains(t, rec.Body.String(), `"max":0`)
	assert.Contains(t, rec.Body.String(), `"min":0`)
	assert.Contains(t, rec.Body.String(), `"median":0`)
}

func TestGetEstateStats_EstateNotFound(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/estate/1/stats", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := repository.NewMockRepositoryInterface(ctrl)
	h := &Server{
		Repository: mockRepo,
	}

	mockRepo.EXPECT().GetEstateById("1").Return(repository.Estate{Id: "", Width: 0, Length: 0}, sql.ErrNoRows)

	h.GetEstateIdStats(c, uuid.Nil)

	assert.Equal(t, http.StatusNotFound, rec.Code)
	assert.Contains(t, rec.Body.String(), "Estate not found")
}

func TestGetEstateStats_DatabaseError(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/estate/1/stats", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := repository.NewMockRepositoryInterface(ctrl)
	h := &Server{
		Repository: mockRepo,
	}

	mockRepo.EXPECT().GetEstateById("1").Return(repository.Estate{Id: "1", Width: 10, Length: 10}, nil)
	mockRepo.EXPECT().GetEstateStatsById("1").Return(repository.EstateStats{}, repository.ErrDatabaseError)

	h.GetEstateIdStats(c, uuid.Nil)

	assert.Equal(t, http.StatusInternalServerError, rec.Code)
	assert.Contains(t, rec.Body.String(), "Failed to retrieve estate stats")
}

// 4. Get drone plan test files
func TestGetDronePlan_Success(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/estate/1/drone-plan", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := repository.NewMockRepositoryInterface(ctrl)
	h := &Server{
		Repository: mockRepo,
	}

	mockRepo.EXPECT().GetEstateById("1").Return(repository.Estate{Id: "1", Width: 10, Length: 10}, nil)
	mockRepo.EXPECT().GetDronePlanByEstateId("1").Return(repository.DronePlan{Distance: 200}, nil)

	h.GetEstateIdDronePlan(c, uuid.Nil)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Contains(t, rec.Body.String(), `"distance":200`)
}

func TestGetDronePlan_EstateNotFound(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/estate/1/drone-plan", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := repository.NewMockRepositoryInterface(ctrl)
	h := &Server{
		Repository: mockRepo,
	}

	mockRepo.EXPECT().GetEstateById("1").Return(repository.Estate{}, sql.ErrNoRows)

	h.GetEstateIdDronePlan(c, uuid.Nil)

	assert.Equal(t, http.StatusNotFound, rec.Code)
	assert.Contains(t, rec.Body.String(), "Estate not found")
}

func TestGetDronePlan_DronePlanNotFound(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/estate/1/drone-plan", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := repository.NewMockRepositoryInterface(ctrl)
	h := &Server{
		Repository: mockRepo,
	}

	mockRepo.EXPECT().GetEstateById("1").Return(repository.Estate{Id: "1", Width: 10, Length: 5}, nil)
	mockRepo.EXPECT().GetDronePlanByEstateId("1").Return(repository.DronePlan{}, sql.ErrNoRows)

	h.GetEstateIdDronePlan(c, uuid.Nil)

	assert.Equal(t, http.StatusNotFound, rec.Code)
	assert.Contains(t, rec.Body.String(), "Drone plan not found")
}

func TestGetDronePlan_DatabaseError(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/estate/1/drone-plan", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := repository.NewMockRepositoryInterface(ctrl)
	h := &Server{
		Repository: mockRepo,
	}

	mockRepo.EXPECT().GetEstateById("1").Return(repository.Estate{Id: "1", Width: 10, Length: 10}, nil)
	mockRepo.EXPECT().GetDronePlanByEstateId("1").Return(repository.DronePlan{}, repository.ErrDatabaseError)

	h.GetEstateIdDronePlan(c, uuid.Nil)

	assert.Equal(t, http.StatusInternalServerError, rec.Code)
	assert.Contains(t, rec.Body.String(), "Failed to retrieve drone plans")
}
