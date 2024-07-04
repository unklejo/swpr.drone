// Code generated by MockGen. DO NOT EDIT.
// Source: repository/interfaces.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	repository "github.com/unklejo/swpr.drone/repository"
)

// MockRepositoryInterface is a mock of RepositoryInterface interface.
type MockRepositoryInterface struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryInterfaceMockRecorder
}

// MockRepositoryInterfaceMockRecorder is the mock recorder for MockRepositoryInterface.
type MockRepositoryInterfaceMockRecorder struct {
	mock *MockRepositoryInterface
}

// NewMockRepositoryInterface creates a new mock instance.
func NewMockRepositoryInterface(ctrl *gomock.Controller) *MockRepositoryInterface {
	mock := &MockRepositoryInterface{ctrl: ctrl}
	mock.recorder = &MockRepositoryInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepositoryInterface) EXPECT() *MockRepositoryInterfaceMockRecorder {
	return m.recorder
}

// AddTree mocks base method.
func (m *MockRepositoryInterface) AddTree(id, estateId string, x, y, height int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddTree", id, estateId, x, y, height)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddTree indicates an expected call of AddTree.
func (mr *MockRepositoryInterfaceMockRecorder) AddTree(id, estateId, x, y, height interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddTree", reflect.TypeOf((*MockRepositoryInterface)(nil).AddTree), id, estateId, x, y, height)
}

// CreateEstate mocks base method.
func (m *MockRepositoryInterface) CreateEstate(id string, width, length int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateEstate", id, width, length)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateEstate indicates an expected call of CreateEstate.
func (mr *MockRepositoryInterfaceMockRecorder) CreateEstate(id, width, length interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateEstate", reflect.TypeOf((*MockRepositoryInterface)(nil).CreateEstate), id, width, length)
}

// GetEstateById mocks base method.
func (m *MockRepositoryInterface) GetEstateById(id string) (repository.Estate, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEstateById", id)
	ret0, _ := ret[0].(repository.Estate)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEstateById indicates an expected call of GetEstateById.
func (mr *MockRepositoryInterfaceMockRecorder) GetEstateById(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEstateById", reflect.TypeOf((*MockRepositoryInterface)(nil).GetEstateById), id)
}

// GetTestById mocks base method.
func (m *MockRepositoryInterface) GetTestById(ctx context.Context, input repository.GetTestByIdInput) (repository.GetTestByIdOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTestById", ctx, input)
	ret0, _ := ret[0].(repository.GetTestByIdOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTestById indicates an expected call of GetTestById.
func (mr *MockRepositoryInterfaceMockRecorder) GetTestById(ctx, input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTestById", reflect.TypeOf((*MockRepositoryInterface)(nil).GetTestById), ctx, input)
}
