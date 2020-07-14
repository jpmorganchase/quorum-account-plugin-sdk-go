// Code generated by MockGen. DO NOT EDIT.
// Source: proto/account.pb.go

// Package mock_proto is a generated GoMock package.
package mock_proto

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	proto "github.com/jpmorganchase/quorum-account-plugin-sdk-go/proto"
	grpc "google.golang.org/grpc"
)

// MockAccountServiceClient is a mock of AccountServiceClient interface.
type MockAccountServiceClient struct {
	ctrl     *gomock.Controller
	recorder *MockAccountServiceClientMockRecorder
}

// MockAccountServiceClientMockRecorder is the mock recorder for MockAccountServiceClient.
type MockAccountServiceClientMockRecorder struct {
	mock *MockAccountServiceClient
}

// NewMockAccountServiceClient creates a new mock instance.
func NewMockAccountServiceClient(ctrl *gomock.Controller) *MockAccountServiceClient {
	mock := &MockAccountServiceClient{ctrl: ctrl}
	mock.recorder = &MockAccountServiceClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAccountServiceClient) EXPECT() *MockAccountServiceClientMockRecorder {
	return m.recorder
}

// Status mocks base method.
func (m *MockAccountServiceClient) Status(ctx context.Context, in *proto.StatusRequest, opts ...grpc.CallOption) (*proto.StatusResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Status", varargs...)
	ret0, _ := ret[0].(*proto.StatusResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Status indicates an expected call of Status.
func (mr *MockAccountServiceClientMockRecorder) Status(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Status", reflect.TypeOf((*MockAccountServiceClient)(nil).Status), varargs...)
}

// Open mocks base method.
func (m *MockAccountServiceClient) Open(ctx context.Context, in *proto.OpenRequest, opts ...grpc.CallOption) (*proto.OpenResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Open", varargs...)
	ret0, _ := ret[0].(*proto.OpenResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Open indicates an expected call of Open.
func (mr *MockAccountServiceClientMockRecorder) Open(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Open", reflect.TypeOf((*MockAccountServiceClient)(nil).Open), varargs...)
}

// Close mocks base method.
func (m *MockAccountServiceClient) Close(ctx context.Context, in *proto.CloseRequest, opts ...grpc.CallOption) (*proto.CloseResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Close", varargs...)
	ret0, _ := ret[0].(*proto.CloseResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Close indicates an expected call of Close.
func (mr *MockAccountServiceClientMockRecorder) Close(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockAccountServiceClient)(nil).Close), varargs...)
}

// Accounts mocks base method.
func (m *MockAccountServiceClient) Accounts(ctx context.Context, in *proto.AccountsRequest, opts ...grpc.CallOption) (*proto.AccountsResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Accounts", varargs...)
	ret0, _ := ret[0].(*proto.AccountsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Accounts indicates an expected call of Accounts.
func (mr *MockAccountServiceClientMockRecorder) Accounts(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Accounts", reflect.TypeOf((*MockAccountServiceClient)(nil).Accounts), varargs...)
}

// Contains mocks base method.
func (m *MockAccountServiceClient) Contains(ctx context.Context, in *proto.ContainsRequest, opts ...grpc.CallOption) (*proto.ContainsResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Contains", varargs...)
	ret0, _ := ret[0].(*proto.ContainsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Contains indicates an expected call of Contains.
func (mr *MockAccountServiceClientMockRecorder) Contains(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Contains", reflect.TypeOf((*MockAccountServiceClient)(nil).Contains), varargs...)
}

// Sign mocks base method.
func (m *MockAccountServiceClient) Sign(ctx context.Context, in *proto.SignRequest, opts ...grpc.CallOption) (*proto.SignResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Sign", varargs...)
	ret0, _ := ret[0].(*proto.SignResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Sign indicates an expected call of Sign.
func (mr *MockAccountServiceClientMockRecorder) Sign(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Sign", reflect.TypeOf((*MockAccountServiceClient)(nil).Sign), varargs...)
}

// UnlockAndSign mocks base method.
func (m *MockAccountServiceClient) UnlockAndSign(ctx context.Context, in *proto.UnlockAndSignRequest, opts ...grpc.CallOption) (*proto.SignResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UnlockAndSign", varargs...)
	ret0, _ := ret[0].(*proto.SignResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UnlockAndSign indicates an expected call of UnlockAndSign.
func (mr *MockAccountServiceClientMockRecorder) UnlockAndSign(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnlockAndSign", reflect.TypeOf((*MockAccountServiceClient)(nil).UnlockAndSign), varargs...)
}

// TimedUnlock mocks base method.
func (m *MockAccountServiceClient) TimedUnlock(ctx context.Context, in *proto.TimedUnlockRequest, opts ...grpc.CallOption) (*proto.TimedUnlockResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "TimedUnlock", varargs...)
	ret0, _ := ret[0].(*proto.TimedUnlockResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TimedUnlock indicates an expected call of TimedUnlock.
func (mr *MockAccountServiceClientMockRecorder) TimedUnlock(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TimedUnlock", reflect.TypeOf((*MockAccountServiceClient)(nil).TimedUnlock), varargs...)
}

// Lock mocks base method.
func (m *MockAccountServiceClient) Lock(ctx context.Context, in *proto.LockRequest, opts ...grpc.CallOption) (*proto.LockResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Lock", varargs...)
	ret0, _ := ret[0].(*proto.LockResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Lock indicates an expected call of Lock.
func (mr *MockAccountServiceClientMockRecorder) Lock(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Lock", reflect.TypeOf((*MockAccountServiceClient)(nil).Lock), varargs...)
}

// NewAccount mocks base method.
func (m *MockAccountServiceClient) NewAccount(ctx context.Context, in *proto.NewAccountRequest, opts ...grpc.CallOption) (*proto.NewAccountResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "NewAccount", varargs...)
	ret0, _ := ret[0].(*proto.NewAccountResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewAccount indicates an expected call of NewAccount.
func (mr *MockAccountServiceClientMockRecorder) NewAccount(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewAccount", reflect.TypeOf((*MockAccountServiceClient)(nil).NewAccount), varargs...)
}

// ImportRawKey mocks base method.
func (m *MockAccountServiceClient) ImportRawKey(ctx context.Context, in *proto.ImportRawKeyRequest, opts ...grpc.CallOption) (*proto.ImportRawKeyResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ImportRawKey", varargs...)
	ret0, _ := ret[0].(*proto.ImportRawKeyResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ImportRawKey indicates an expected call of ImportRawKey.
func (mr *MockAccountServiceClientMockRecorder) ImportRawKey(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ImportRawKey", reflect.TypeOf((*MockAccountServiceClient)(nil).ImportRawKey), varargs...)
}

// MockAccountServiceServer is a mock of AccountServiceServer interface.
type MockAccountServiceServer struct {
	ctrl     *gomock.Controller
	recorder *MockAccountServiceServerMockRecorder
}

// MockAccountServiceServerMockRecorder is the mock recorder for MockAccountServiceServer.
type MockAccountServiceServerMockRecorder struct {
	mock *MockAccountServiceServer
}

// NewMockAccountServiceServer creates a new mock instance.
func NewMockAccountServiceServer(ctrl *gomock.Controller) *MockAccountServiceServer {
	mock := &MockAccountServiceServer{ctrl: ctrl}
	mock.recorder = &MockAccountServiceServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAccountServiceServer) EXPECT() *MockAccountServiceServerMockRecorder {
	return m.recorder
}

// Status mocks base method.
func (m *MockAccountServiceServer) Status(arg0 context.Context, arg1 *proto.StatusRequest) (*proto.StatusResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Status", arg0, arg1)
	ret0, _ := ret[0].(*proto.StatusResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Status indicates an expected call of Status.
func (mr *MockAccountServiceServerMockRecorder) Status(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Status", reflect.TypeOf((*MockAccountServiceServer)(nil).Status), arg0, arg1)
}

// Open mocks base method.
func (m *MockAccountServiceServer) Open(arg0 context.Context, arg1 *proto.OpenRequest) (*proto.OpenResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Open", arg0, arg1)
	ret0, _ := ret[0].(*proto.OpenResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Open indicates an expected call of Open.
func (mr *MockAccountServiceServerMockRecorder) Open(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Open", reflect.TypeOf((*MockAccountServiceServer)(nil).Open), arg0, arg1)
}

// Close mocks base method.
func (m *MockAccountServiceServer) Close(arg0 context.Context, arg1 *proto.CloseRequest) (*proto.CloseResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close", arg0, arg1)
	ret0, _ := ret[0].(*proto.CloseResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Close indicates an expected call of Close.
func (mr *MockAccountServiceServerMockRecorder) Close(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockAccountServiceServer)(nil).Close), arg0, arg1)
}

// Accounts mocks base method.
func (m *MockAccountServiceServer) Accounts(arg0 context.Context, arg1 *proto.AccountsRequest) (*proto.AccountsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Accounts", arg0, arg1)
	ret0, _ := ret[0].(*proto.AccountsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Accounts indicates an expected call of Accounts.
func (mr *MockAccountServiceServerMockRecorder) Accounts(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Accounts", reflect.TypeOf((*MockAccountServiceServer)(nil).Accounts), arg0, arg1)
}

// Contains mocks base method.
func (m *MockAccountServiceServer) Contains(arg0 context.Context, arg1 *proto.ContainsRequest) (*proto.ContainsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Contains", arg0, arg1)
	ret0, _ := ret[0].(*proto.ContainsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Contains indicates an expected call of Contains.
func (mr *MockAccountServiceServerMockRecorder) Contains(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Contains", reflect.TypeOf((*MockAccountServiceServer)(nil).Contains), arg0, arg1)
}

// Sign mocks base method.
func (m *MockAccountServiceServer) Sign(arg0 context.Context, arg1 *proto.SignRequest) (*proto.SignResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Sign", arg0, arg1)
	ret0, _ := ret[0].(*proto.SignResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Sign indicates an expected call of Sign.
func (mr *MockAccountServiceServerMockRecorder) Sign(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Sign", reflect.TypeOf((*MockAccountServiceServer)(nil).Sign), arg0, arg1)
}

// UnlockAndSign mocks base method.
func (m *MockAccountServiceServer) UnlockAndSign(arg0 context.Context, arg1 *proto.UnlockAndSignRequest) (*proto.SignResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UnlockAndSign", arg0, arg1)
	ret0, _ := ret[0].(*proto.SignResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UnlockAndSign indicates an expected call of UnlockAndSign.
func (mr *MockAccountServiceServerMockRecorder) UnlockAndSign(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnlockAndSign", reflect.TypeOf((*MockAccountServiceServer)(nil).UnlockAndSign), arg0, arg1)
}

// TimedUnlock mocks base method.
func (m *MockAccountServiceServer) TimedUnlock(arg0 context.Context, arg1 *proto.TimedUnlockRequest) (*proto.TimedUnlockResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TimedUnlock", arg0, arg1)
	ret0, _ := ret[0].(*proto.TimedUnlockResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TimedUnlock indicates an expected call of TimedUnlock.
func (mr *MockAccountServiceServerMockRecorder) TimedUnlock(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TimedUnlock", reflect.TypeOf((*MockAccountServiceServer)(nil).TimedUnlock), arg0, arg1)
}

// Lock mocks base method.
func (m *MockAccountServiceServer) Lock(arg0 context.Context, arg1 *proto.LockRequest) (*proto.LockResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Lock", arg0, arg1)
	ret0, _ := ret[0].(*proto.LockResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Lock indicates an expected call of Lock.
func (mr *MockAccountServiceServerMockRecorder) Lock(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Lock", reflect.TypeOf((*MockAccountServiceServer)(nil).Lock), arg0, arg1)
}

// NewAccount mocks base method.
func (m *MockAccountServiceServer) NewAccount(arg0 context.Context, arg1 *proto.NewAccountRequest) (*proto.NewAccountResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewAccount", arg0, arg1)
	ret0, _ := ret[0].(*proto.NewAccountResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewAccount indicates an expected call of NewAccount.
func (mr *MockAccountServiceServerMockRecorder) NewAccount(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewAccount", reflect.TypeOf((*MockAccountServiceServer)(nil).NewAccount), arg0, arg1)
}

// ImportRawKey mocks base method.
func (m *MockAccountServiceServer) ImportRawKey(arg0 context.Context, arg1 *proto.ImportRawKeyRequest) (*proto.ImportRawKeyResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ImportRawKey", arg0, arg1)
	ret0, _ := ret[0].(*proto.ImportRawKeyResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ImportRawKey indicates an expected call of ImportRawKey.
func (mr *MockAccountServiceServerMockRecorder) ImportRawKey(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ImportRawKey", reflect.TypeOf((*MockAccountServiceServer)(nil).ImportRawKey), arg0, arg1)
}
