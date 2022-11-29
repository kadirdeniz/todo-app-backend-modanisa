// Code generated by MockGen. DO NOT EDIT.
// Source: todo_repository.go

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"
	model "todo/model"

	gomock "github.com/golang/mock/gomock"
)

// MockITodoRepository is a mock of ITodoRepository interface.
type MockITodoRepository struct {
	ctrl     *gomock.Controller
	recorder *MockITodoRepositoryMockRecorder
}

// MockITodoRepositoryMockRecorder is the mock recorder for MockITodoRepository.
type MockITodoRepositoryMockRecorder struct {
	mock *MockITodoRepository
}

// NewMockITodoRepository creates a new mock instance.
func NewMockITodoRepository(ctrl *gomock.Controller) *MockITodoRepository {
	mock := &MockITodoRepository{ctrl: ctrl}
	mock.recorder = &MockITodoRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockITodoRepository) EXPECT() *MockITodoRepositoryMockRecorder {
	return m.recorder
}

// CreateTodo mocks base method.
func (m *MockITodoRepository) CreateTodo() model.Todo {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTodo")
	ret0, _ := ret[0].(model.Todo)
	return ret0
}

// CreateTodo indicates an expected call of CreateTodo.
func (mr *MockITodoRepositoryMockRecorder) CreateTodo() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTodo", reflect.TypeOf((*MockITodoRepository)(nil).CreateTodo))
}
