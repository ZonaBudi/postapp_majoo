// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/port/repository.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	domain "postapp/internal/domain"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockUserRepository is a mock of UserRepository interface.
type MockUserRepository struct {
	ctrl     *gomock.Controller
	recorder *MockUserRepositoryMockRecorder
}

// MockUserRepositoryMockRecorder is the mock recorder for MockUserRepository.
type MockUserRepositoryMockRecorder struct {
	mock *MockUserRepository
}

// NewMockUserRepository creates a new mock instance.
func NewMockUserRepository(ctrl *gomock.Controller) *MockUserRepository {
	mock := &MockUserRepository{ctrl: ctrl}
	mock.recorder = &MockUserRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserRepository) EXPECT() *MockUserRepositoryMockRecorder {
	return m.recorder
}

// FindOneByUsername mocks base method.
func (m *MockUserRepository) FindOneByUsername(ctx context.Context, username string) (*domain.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindOneByUsername", ctx, username)
	ret0, _ := ret[0].(*domain.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindOneByUsername indicates an expected call of FindOneByUsername.
func (mr *MockUserRepositoryMockRecorder) FindOneByUsername(ctx, username interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindOneByUsername", reflect.TypeOf((*MockUserRepository)(nil).FindOneByUsername), ctx, username)
}

// MockTransactionRepository is a mock of TransactionRepository interface.
type MockTransactionRepository struct {
	ctrl     *gomock.Controller
	recorder *MockTransactionRepositoryMockRecorder
}

// MockTransactionRepositoryMockRecorder is the mock recorder for MockTransactionRepository.
type MockTransactionRepositoryMockRecorder struct {
	mock *MockTransactionRepository
}

// NewMockTransactionRepository creates a new mock instance.
func NewMockTransactionRepository(ctrl *gomock.Controller) *MockTransactionRepository {
	mock := &MockTransactionRepository{ctrl: ctrl}
	mock.recorder = &MockTransactionRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTransactionRepository) EXPECT() *MockTransactionRepositoryMockRecorder {
	return m.recorder
}

// TransactionByMerchant mocks base method.
func (m *MockTransactionRepository) TransactionByMerchant(ctx context.Context, filter *domain.TransactionMerchantFilter) ([]*domain.TransactionMerchant, *domain.TransactionMerchantFilter, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TransactionByMerchant", ctx, filter)
	ret0, _ := ret[0].([]*domain.TransactionMerchant)
	ret1, _ := ret[1].(*domain.TransactionMerchantFilter)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// TransactionByMerchant indicates an expected call of TransactionByMerchant.
func (mr *MockTransactionRepositoryMockRecorder) TransactionByMerchant(ctx, filter interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TransactionByMerchant", reflect.TypeOf((*MockTransactionRepository)(nil).TransactionByMerchant), ctx, filter)
}

// TransactionByOutlet mocks base method.
func (m *MockTransactionRepository) TransactionByOutlet(ctx context.Context, filter *domain.TransactionOutletFilter) ([]*domain.TransactionOutlet, *domain.TransactionOutletFilter, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TransactionByOutlet", ctx, filter)
	ret0, _ := ret[0].([]*domain.TransactionOutlet)
	ret1, _ := ret[1].(*domain.TransactionOutletFilter)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// TransactionByOutlet indicates an expected call of TransactionByOutlet.
func (mr *MockTransactionRepositoryMockRecorder) TransactionByOutlet(ctx, filter interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TransactionByOutlet", reflect.TypeOf((*MockTransactionRepository)(nil).TransactionByOutlet), ctx, filter)
}

// MockMerchantRepository is a mock of MerchantRepository interface.
type MockMerchantRepository struct {
	ctrl     *gomock.Controller
	recorder *MockMerchantRepositoryMockRecorder
}

// MockMerchantRepositoryMockRecorder is the mock recorder for MockMerchantRepository.
type MockMerchantRepositoryMockRecorder struct {
	mock *MockMerchantRepository
}

// NewMockMerchantRepository creates a new mock instance.
func NewMockMerchantRepository(ctrl *gomock.Controller) *MockMerchantRepository {
	mock := &MockMerchantRepository{ctrl: ctrl}
	mock.recorder = &MockMerchantRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMerchantRepository) EXPECT() *MockMerchantRepositoryMockRecorder {
	return m.recorder
}

// FindOneByUserID mocks base method.
func (m *MockMerchantRepository) FindOneByUserID(ctx context.Context, merchantID uint64) (*domain.Merchant, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindOneByUserID", ctx, merchantID)
	ret0, _ := ret[0].(*domain.Merchant)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindOneByUserID indicates an expected call of FindOneByUserID.
func (mr *MockMerchantRepositoryMockRecorder) FindOneByUserID(ctx, merchantID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindOneByUserID", reflect.TypeOf((*MockMerchantRepository)(nil).FindOneByUserID), ctx, merchantID)
}

// MockOutletRepository is a mock of OutletRepository interface.
type MockOutletRepository struct {
	ctrl     *gomock.Controller
	recorder *MockOutletRepositoryMockRecorder
}

// MockOutletRepositoryMockRecorder is the mock recorder for MockOutletRepository.
type MockOutletRepositoryMockRecorder struct {
	mock *MockOutletRepository
}

// NewMockOutletRepository creates a new mock instance.
func NewMockOutletRepository(ctrl *gomock.Controller) *MockOutletRepository {
	mock := &MockOutletRepository{ctrl: ctrl}
	mock.recorder = &MockOutletRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOutletRepository) EXPECT() *MockOutletRepositoryMockRecorder {
	return m.recorder
}

// FindOneByOutletID mocks base method.
func (m *MockOutletRepository) FindOneByOutletID(ctx context.Context, outletID uint64) (*domain.Outlet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindOneByOutletID", ctx, outletID)
	ret0, _ := ret[0].(*domain.Outlet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindOneByOutletID indicates an expected call of FindOneByOutletID.
func (mr *MockOutletRepositoryMockRecorder) FindOneByOutletID(ctx, outletID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindOneByOutletID", reflect.TypeOf((*MockOutletRepository)(nil).FindOneByOutletID), ctx, outletID)
}
