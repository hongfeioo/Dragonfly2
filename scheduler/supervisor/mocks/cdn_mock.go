// Code generated by MockGen. DO NOT EDIT.
// Source: d7y.io/dragonfly/v2/scheduler/supervisor (interfaces: CDNDynmaicClient)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	dfnet "d7y.io/dragonfly/v2/pkg/basic/dfnet"
	base "d7y.io/dragonfly/v2/pkg/rpc/base"
	cdnsystem "d7y.io/dragonfly/v2/pkg/rpc/cdnsystem"
	client "d7y.io/dragonfly/v2/pkg/rpc/cdnsystem/client"
	config "d7y.io/dragonfly/v2/scheduler/config"
	supervisor "d7y.io/dragonfly/v2/scheduler/supervisor"
	gomock "github.com/golang/mock/gomock"
	grpc "google.golang.org/grpc"
)

// MockCDNDynmaicClient is a mock of CDNDynmaicClient interface.
type MockCDNDynmaicClient struct {
	ctrl     *gomock.Controller
	recorder *MockCDNDynmaicClientMockRecorder
}

// MockCDNDynmaicClientMockRecorder is the mock recorder for MockCDNDynmaicClient.
type MockCDNDynmaicClientMockRecorder struct {
	mock *MockCDNDynmaicClient
}

// NewMockCDNDynmaicClient creates a new mock instance.
func NewMockCDNDynmaicClient(ctrl *gomock.Controller) *MockCDNDynmaicClient {
	mock := &MockCDNDynmaicClient{ctrl: ctrl}
	mock.recorder = &MockCDNDynmaicClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCDNDynmaicClient) EXPECT() *MockCDNDynmaicClientMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockCDNDynmaicClient) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockCDNDynmaicClientMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockCDNDynmaicClient)(nil).Close))
}

// GetHost mocks base method.
func (m *MockCDNDynmaicClient) GetHost(arg0 string) (*supervisor.Host, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetHost", arg0)
	ret0, _ := ret[0].(*supervisor.Host)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// GetHost indicates an expected call of GetHost.
func (mr *MockCDNDynmaicClientMockRecorder) GetHost(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHost", reflect.TypeOf((*MockCDNDynmaicClient)(nil).GetHost), arg0)
}

// GetPieceTasks mocks base method.
func (m *MockCDNDynmaicClient) GetPieceTasks(arg0 context.Context, arg1 dfnet.NetAddr, arg2 *base.PieceTaskRequest, arg3 ...grpc.CallOption) (*base.PiecePacket, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2}
	for _, a := range arg3 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetPieceTasks", varargs...)
	ret0, _ := ret[0].(*base.PiecePacket)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPieceTasks indicates an expected call of GetPieceTasks.
func (mr *MockCDNDynmaicClientMockRecorder) GetPieceTasks(arg0, arg1, arg2 interface{}, arg3 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2}, arg3...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPieceTasks", reflect.TypeOf((*MockCDNDynmaicClient)(nil).GetPieceTasks), varargs...)
}

// ObtainSeeds mocks base method.
func (m *MockCDNDynmaicClient) ObtainSeeds(arg0 context.Context, arg1 *cdnsystem.SeedRequest, arg2 ...grpc.CallOption) (*client.PieceSeedStream, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ObtainSeeds", varargs...)
	ret0, _ := ret[0].(*client.PieceSeedStream)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ObtainSeeds indicates an expected call of ObtainSeeds.
func (mr *MockCDNDynmaicClientMockRecorder) ObtainSeeds(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ObtainSeeds", reflect.TypeOf((*MockCDNDynmaicClient)(nil).ObtainSeeds), varargs...)
}

// OnNotify mocks base method.
func (m *MockCDNDynmaicClient) OnNotify(arg0 *config.DynconfigData) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "OnNotify", arg0)
}

// OnNotify indicates an expected call of OnNotify.
func (mr *MockCDNDynmaicClientMockRecorder) OnNotify(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OnNotify", reflect.TypeOf((*MockCDNDynmaicClient)(nil).OnNotify), arg0)
}

// UpdateState mocks base method.
func (m *MockCDNDynmaicClient) UpdateState(arg0 []dfnet.NetAddr) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "UpdateState", arg0)
}

// UpdateState indicates an expected call of UpdateState.
func (mr *MockCDNDynmaicClientMockRecorder) UpdateState(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateState", reflect.TypeOf((*MockCDNDynmaicClient)(nil).UpdateState), arg0)
}
