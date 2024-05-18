// Code generated by MockGen. DO NOT EDIT.
// Source: service/user_service/interfaces.go
//
// Generated by this command:
//
//	mockgen -source=service/user_service/interfaces.go -destination=service/user_service/interfaces.mock.gen.go -package=user_service
//

// Package user_service is a generated GoMock package.
package user_service

import (
	reflect "reflect"

	request "github.com/DangerZombie/case-study-dealls/model/request"
	response "github.com/DangerZombie/case-study-dealls/model/response"
	gomock "go.uber.org/mock/gomock"
)

// MockUserService is a mock of UserService interface.
type MockUserService struct {
	ctrl     *gomock.Controller
	recorder *MockUserServiceMockRecorder
}

// MockUserServiceMockRecorder is the mock recorder for MockUserService.
type MockUserServiceMockRecorder struct {
	mock *MockUserService
}

// NewMockUserService creates a new mock instance.
func NewMockUserService(ctrl *gomock.Controller) *MockUserService {
	mock := &MockUserService{ctrl: ctrl}
	mock.recorder = &MockUserServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserService) EXPECT() *MockUserServiceMockRecorder {
	return m.recorder
}

// BuySubscription mocks base method.
func (m *MockUserService) BuySubscription(req request.BuySubscriptionRequestBody) (response.BuySubscriptionResponse, int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BuySubscription", req)
	ret0, _ := ret[0].(response.BuySubscriptionResponse)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// BuySubscription indicates an expected call of BuySubscription.
func (mr *MockUserServiceMockRecorder) BuySubscription(req any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BuySubscription", reflect.TypeOf((*MockUserService)(nil).BuySubscription), req)
}

// CreateUserInteraction mocks base method.
func (m *MockUserService) CreateUserInteraction(req request.CreateUserInteractionRequest) (response.CreateUserInteractionResponse, int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUserInteraction", req)
	ret0, _ := ret[0].(response.CreateUserInteractionResponse)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// CreateUserInteraction indicates an expected call of CreateUserInteraction.
func (mr *MockUserServiceMockRecorder) CreateUserInteraction(req any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUserInteraction", reflect.TypeOf((*MockUserService)(nil).CreateUserInteraction), req)
}

// GetUserToSwipe mocks base method.
func (m *MockUserService) GetUserToSwipe(req request.GetUserToSwipeRequest) (response.GetUserToSwipeResponse, int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserToSwipe", req)
	ret0, _ := ret[0].(response.GetUserToSwipeResponse)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetUserToSwipe indicates an expected call of GetUserToSwipe.
func (mr *MockUserServiceMockRecorder) GetUserToSwipe(req any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserToSwipe", reflect.TypeOf((*MockUserService)(nil).GetUserToSwipe), req)
}

// Login mocks base method.
func (m *MockUserService) Login(req request.LoginRequestBody) (response.LoginResponse, int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Login", req)
	ret0, _ := ret[0].(response.LoginResponse)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Login indicates an expected call of Login.
func (mr *MockUserServiceMockRecorder) Login(req any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Login", reflect.TypeOf((*MockUserService)(nil).Login), req)
}

// RegisterUser mocks base method.
func (m *MockUserService) RegisterUser(req request.RegisterUserRequestBody) (response.RegisterUserResponse, int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegisterUser", req)
	ret0, _ := ret[0].(response.RegisterUserResponse)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// RegisterUser indicates an expected call of RegisterUser.
func (mr *MockUserServiceMockRecorder) RegisterUser(req any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterUser", reflect.TypeOf((*MockUserService)(nil).RegisterUser), req)
}

// ResetSwipeCount mocks base method.
func (m *MockUserService) ResetSwipeCount(req request.ResetSwipeCountRequest) (response.ResetSwipeCountResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResetSwipeCount", req)
	ret0, _ := ret[0].(response.ResetSwipeCountResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ResetSwipeCount indicates an expected call of ResetSwipeCount.
func (mr *MockUserServiceMockRecorder) ResetSwipeCount(req any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResetSwipeCount", reflect.TypeOf((*MockUserService)(nil).ResetSwipeCount), req)
}
