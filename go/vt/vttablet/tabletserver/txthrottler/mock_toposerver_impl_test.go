// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/youtube/vitess/go/vt/topo (interfaces: Impl)

// Package txthrottler is a generated GoMock package.
package txthrottler

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	topo "github.com/youtube/vitess/go/vt/topo"
	context "golang.org/x/net/context"
)

// MockImpl is a mock of Impl interface
type MockImpl struct {
	ctrl     *gomock.Controller
	recorder *MockImplMockRecorder
}

// MockImplMockRecorder is the mock recorder for MockImpl
type MockImplMockRecorder struct {
	mock *MockImpl
}

// NewMockImpl creates a new mock instance
func NewMockImpl(ctrl *gomock.Controller) *MockImpl {
	mock := &MockImpl{ctrl: ctrl}
	mock.recorder = &MockImplMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockImpl) EXPECT() *MockImplMockRecorder {
	return m.recorder
}

// Close mocks base method
func (m *MockImpl) Close() {
	m.ctrl.Call(m, "Close")
}

// Close indicates an expected call of Close
func (mr *MockImplMockRecorder) Close() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockImpl)(nil).Close))
}

// Create mocks base method
func (m *MockImpl) Create(arg0 context.Context, arg1, arg2 string, arg3 []byte) (topo.Version, error) {
	ret := m.ctrl.Call(m, "Create", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(topo.Version)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *MockImplMockRecorder) Create(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockImpl)(nil).Create), arg0, arg1, arg2, arg3)
}

// Delete mocks base method
func (m *MockImpl) Delete(arg0 context.Context, arg1, arg2 string, arg3 topo.Version) error {
	ret := m.ctrl.Call(m, "Delete", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockImplMockRecorder) Delete(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockImpl)(nil).Delete), arg0, arg1, arg2, arg3)
}

// Get mocks base method
func (m *MockImpl) Get(arg0 context.Context, arg1, arg2 string) ([]byte, topo.Version, error) {
	ret := m.ctrl.Call(m, "Get", arg0, arg1, arg2)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(topo.Version)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Get indicates an expected call of Get
func (mr *MockImplMockRecorder) Get(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockImpl)(nil).Get), arg0, arg1, arg2)
}

// GetKnownCells mocks base method
func (m *MockImpl) GetKnownCells(arg0 context.Context) ([]string, error) {
	ret := m.ctrl.Call(m, "GetKnownCells", arg0)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetKnownCells indicates an expected call of GetKnownCells
func (mr *MockImplMockRecorder) GetKnownCells(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetKnownCells", reflect.TypeOf((*MockImpl)(nil).GetKnownCells), arg0)
}

// ListDir mocks base method
func (m *MockImpl) ListDir(arg0 context.Context, arg1, arg2 string) ([]string, error) {
	ret := m.ctrl.Call(m, "ListDir", arg0, arg1, arg2)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListDir indicates an expected call of ListDir
func (mr *MockImplMockRecorder) ListDir(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListDir", reflect.TypeOf((*MockImpl)(nil).ListDir), arg0, arg1, arg2)
}

// Lock mocks base method
func (m *MockImpl) Lock(arg0 context.Context, arg1, arg2 string) (topo.LockDescriptor, error) {
	ret := m.ctrl.Call(m, "Lock", arg0, arg1, arg2)
	ret0, _ := ret[0].(topo.LockDescriptor)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Lock indicates an expected call of Lock
func (mr *MockImplMockRecorder) Lock(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Lock", reflect.TypeOf((*MockImpl)(nil).Lock), arg0, arg1, arg2)
}

// NewMasterParticipation mocks base method
func (m *MockImpl) NewMasterParticipation(arg0, arg1 string) (topo.MasterParticipation, error) {
	ret := m.ctrl.Call(m, "NewMasterParticipation", arg0, arg1)
	ret0, _ := ret[0].(topo.MasterParticipation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewMasterParticipation indicates an expected call of NewMasterParticipation
func (mr *MockImplMockRecorder) NewMasterParticipation(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewMasterParticipation", reflect.TypeOf((*MockImpl)(nil).NewMasterParticipation), arg0, arg1)
}

// Update mocks base method
func (m *MockImpl) Update(arg0 context.Context, arg1, arg2 string, arg3 []byte, arg4 topo.Version) (topo.Version, error) {
	ret := m.ctrl.Call(m, "Update", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(topo.Version)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update
func (mr *MockImplMockRecorder) Update(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockImpl)(nil).Update), arg0, arg1, arg2, arg3, arg4)
}

// Watch mocks base method
func (m *MockImpl) Watch(arg0 context.Context, arg1, arg2 string) (*topo.WatchData, <-chan *topo.WatchData, topo.CancelFunc) {
	ret := m.ctrl.Call(m, "Watch", arg0, arg1, arg2)
	ret0, _ := ret[0].(*topo.WatchData)
	ret1, _ := ret[1].(<-chan *topo.WatchData)
	ret2, _ := ret[2].(topo.CancelFunc)
	return ret0, ret1, ret2
}

// Watch indicates an expected call of Watch
func (mr *MockImplMockRecorder) Watch(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Watch", reflect.TypeOf((*MockImpl)(nil).Watch), arg0, arg1, arg2)
}
