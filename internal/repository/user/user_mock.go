// Code generated by MockGen. DO NOT EDIT.
// Source: types.go
//
// Generated by this command:
//
//	mockgen -destination=user_mock.go -source=types.go -package=user
//

// Package user is a generated GoMock package.
package user

import (
	context "context"
	reflect "reflect"

	pgx "github.com/jackc/pgx/v5"
	pgconn "github.com/jackc/pgx/v5/pgconn"
	gomock "go.uber.org/mock/gomock"
)

// MockdatabaseResource is a mock of databaseResource interface.
type MockdatabaseResource struct {
	ctrl     *gomock.Controller
	recorder *MockdatabaseResourceMockRecorder
}

// MockdatabaseResourceMockRecorder is the mock recorder for MockdatabaseResource.
type MockdatabaseResourceMockRecorder struct {
	mock *MockdatabaseResource
}

// NewMockdatabaseResource creates a new mock instance.
func NewMockdatabaseResource(ctrl *gomock.Controller) *MockdatabaseResource {
	mock := &MockdatabaseResource{ctrl: ctrl}
	mock.recorder = &MockdatabaseResourceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockdatabaseResource) EXPECT() *MockdatabaseResourceMockRecorder {
	return m.recorder
}

// Exec mocks base method.
func (m *MockdatabaseResource) Exec(ctx context.Context, sql string, arguments ...any) (pgconn.CommandTag, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, sql}
	for _, a := range arguments {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Exec", varargs...)
	ret0, _ := ret[0].(pgconn.CommandTag)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Exec indicates an expected call of Exec.
func (mr *MockdatabaseResourceMockRecorder) Exec(ctx, sql any, arguments ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, sql}, arguments...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exec", reflect.TypeOf((*MockdatabaseResource)(nil).Exec), varargs...)
}

// Query mocks base method.
func (m *MockdatabaseResource) Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, sql}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Query", varargs...)
	ret0, _ := ret[0].(pgx.Rows)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Query indicates an expected call of Query.
func (mr *MockdatabaseResourceMockRecorder) Query(ctx, sql any, args ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, sql}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Query", reflect.TypeOf((*MockdatabaseResource)(nil).Query), varargs...)
}

// QueryRow mocks base method.
func (m *MockdatabaseResource) QueryRow(ctx context.Context, sql string, args ...any) pgx.Row {
	m.ctrl.T.Helper()
	varargs := []any{ctx, sql}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "QueryRow", varargs...)
	ret0, _ := ret[0].(pgx.Row)
	return ret0
}

// QueryRow indicates an expected call of QueryRow.
func (mr *MockdatabaseResourceMockRecorder) QueryRow(ctx, sql any, args ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, sql}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryRow", reflect.TypeOf((*MockdatabaseResource)(nil).QueryRow), varargs...)
}