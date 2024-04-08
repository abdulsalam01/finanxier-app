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
	redis "github.com/finanxier-app/pkg/redis"
	gomock "go.uber.org/mock/gomock"
)

// MockproductUsecase is a mock of productUsecase interface.
type MockproductUsecase struct {
	ctrl     *gomock.Controller
	recorder *MockproductUsecaseMockRecorder
}

// MockproductUsecaseMockRecorder is the mock recorder for MockproductUsecase.
type MockproductUsecaseMockRecorder struct {
	mock *MockproductUsecase
}

// NewMockproductUsecase creates a new mock instance.
func NewMockproductUsecase(ctrl *gomock.Controller) *MockproductUsecase {
	mock := &MockproductUsecase{ctrl: ctrl}
	mock.recorder = &MockproductUsecaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockproductUsecase) EXPECT() *MockproductUsecaseMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockproductUsecase) Create(ctx context.Context, params entity.Product) (entity.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, params)
	ret0, _ := ret[0].(entity.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockproductUsecaseMockRecorder) Create(ctx, params any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockproductUsecase)(nil).Create), ctx, params)
}

// GetByID mocks base method.
func (m *MockproductUsecase) GetByID(ctx context.Context, id string) (entity.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", ctx, id)
	ret0, _ := ret[0].(entity.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID.
func (mr *MockproductUsecaseMockRecorder) GetByID(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockproductUsecase)(nil).GetByID), ctx, id)
}

// GetByParams mocks base method.
func (m *MockproductUsecase) GetByParams(ctx context.Context, params base.PaginationRequest) (entity.ProductBulkResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByParams", ctx, params)
	ret0, _ := ret[0].(entity.ProductBulkResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByParams indicates an expected call of GetByParams.
func (mr *MockproductUsecaseMockRecorder) GetByParams(ctx, params any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByParams", reflect.TypeOf((*MockproductUsecase)(nil).GetByParams), ctx, params)
}

// MockbaseAppInitializerResource is a mock of baseAppInitializerResource interface.
type MockbaseAppInitializerResource struct {
	ctrl     *gomock.Controller
	recorder *MockbaseAppInitializerResourceMockRecorder
}

// MockbaseAppInitializerResourceMockRecorder is the mock recorder for MockbaseAppInitializerResource.
type MockbaseAppInitializerResourceMockRecorder struct {
	mock *MockbaseAppInitializerResource
}

// NewMockbaseAppInitializerResource creates a new mock instance.
func NewMockbaseAppInitializerResource(ctrl *gomock.Controller) *MockbaseAppInitializerResource {
	mock := &MockbaseAppInitializerResource{ctrl: ctrl}
	mock.recorder = &MockbaseAppInitializerResourceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockbaseAppInitializerResource) EXPECT() *MockbaseAppInitializerResourceMockRecorder {
	return m.recorder
}

// Lock mocks base method.
func (m *MockbaseAppInitializerResource) Lock(ctx context.Context, key string) (redis.RedisLockResource, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Lock", ctx, key)
	ret0, _ := ret[0].(redis.RedisLockResource)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Lock indicates an expected call of Lock.
func (mr *MockbaseAppInitializerResourceMockRecorder) Lock(ctx, key any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Lock", reflect.TypeOf((*MockbaseAppInitializerResource)(nil).Lock), ctx, key)
}

// Struct mocks base method.
func (m *MockbaseAppInitializerResource) Struct(s any) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Struct", s)
	ret0, _ := ret[0].(error)
	return ret0
}

// Struct indicates an expected call of Struct.
func (mr *MockbaseAppInitializerResourceMockRecorder) Struct(s any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Struct", reflect.TypeOf((*MockbaseAppInitializerResource)(nil).Struct), s)
}