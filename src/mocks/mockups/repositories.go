// Code generated by MockGen. DO NOT EDIT.
// Source: src/internal/core/ports/repositories.go

// Package mock_ports is a generated GoMock package.
package mock_ports

import (
	reflect "reflect"

	domain "github.com/D-D-EINA-Calendar/CalendarServer/src/internal/core/domain"
	gomock "github.com/golang/mock/gomock"
)

// MockHorarioRepositorio is a mock of HorarioRepositorio interface.
type MockHorarioRepositorio struct {
	ctrl     *gomock.Controller
	recorder *MockHorarioRepositorioMockRecorder
}

// MockHorarioRepositorioMockRecorder is the mock recorder for MockHorarioRepositorio.
type MockHorarioRepositorioMockRecorder struct {
	mock *MockHorarioRepositorio
}

// NewMockHorarioRepositorio creates a new mock instance.
func NewMockHorarioRepositorio(ctrl *gomock.Controller) *MockHorarioRepositorio {
	mock := &MockHorarioRepositorio{ctrl: ctrl}
	mock.recorder = &MockHorarioRepositorioMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockHorarioRepositorio) EXPECT() *MockHorarioRepositorioMockRecorder {
	return m.recorder
}

// GetAvailableHours mocks base method.
func (m *MockHorarioRepositorio) GetAvailableHours(arg0 domain.Terna) ([]domain.AvailableHours, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAvailableHours", arg0)
	ret0, _ := ret[0].([]domain.AvailableHours)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAvailableHours indicates an expected call of GetAvailableHours.
func (mr *MockHorarioRepositorioMockRecorder) GetAvailableHours(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAvailableHours", reflect.TypeOf((*MockHorarioRepositorio)(nil).GetAvailableHours), arg0)
}
