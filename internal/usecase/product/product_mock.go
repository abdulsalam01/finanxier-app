// Code generated by MockGen. DO NOT EDIT.
// Source: types.go
//
// Generated by this command:
//
//	mockgen -destination=product_mock.go -source=types.go -package=product
//

// Package product is a generated GoMock package.
package product

import (
	context "context"
	reflect "reflect"

	entity "github.com/finanxier-app/internal/entity"
	base "github.com/finanxier-app/internal/entity/base"
	gomock "go.uber.org/mock/gomock"
)

// MockproductResource is a mock of productResource interface.
type MockproductResource struct {
	ctrl     *gomock.Controller
	recorder *MockproductResourceMockRecorder
}

// MockproductResourceMockRecorder is the mock recorder for MockproductResource.
type MockproductResourceMockRecorder struct {
	mock *MockproductResource
}

// NewMockproductResource creates a new mock instance.
func NewMockproductResource(ctrl *gomock.Controller) *MockproductResource {
	mock := &MockproductResource{ctrl: ctrl}
	mock.recorder = &MockproductResourceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockproductResource) EXPECT() *MockproductResourceMockRecorder {
	return m.recorder
}

// Count mocks base method.
func (m *MockproductResource) Count(ctx context.Context) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Count", ctx)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Count indicates an expected call of Count.
func (mr *MockproductResourceMockRecorder) Count(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Count", reflect.TypeOf((*MockproductResource)(nil).Count), ctx)
}

// Create mocks base method.
func (m *MockproductResource) Create(ctx context.Context, params entity.Product) (entity.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, params)
	ret0, _ := ret[0].(entity.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockproductResourceMockRecorder) Create(ctx, params any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockproductResource)(nil).Create), ctx, params)
}

// GetByID mocks base method.
func (m *MockproductResource) GetByID(ctx context.Context, id string) (entity.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", ctx, id)
	ret0, _ := ret[0].(entity.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID.
func (mr *MockproductResourceMockRecorder) GetByID(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockproductResource)(nil).GetByID), ctx, id)
}

// GetByParams mocks base method.
func (m *MockproductResource) GetByParams(ctx context.Context, params base.PaginationRequest) ([]entity.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByParams", ctx, params)
	ret0, _ := ret[0].([]entity.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByParams indicates an expected call of GetByParams.
func (mr *MockproductResourceMockRecorder) GetByParams(ctx, params any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByParams", reflect.TypeOf((*MockproductResource)(nil).GetByParams), ctx, params)
}
