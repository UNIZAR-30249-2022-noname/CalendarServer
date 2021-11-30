// Code generated by MockGen. DO NOT EDIT.
// Source: src/internal/core/ports/services.go

// Package mock_ports is a generated GoMock package.
package mock_ports

import (
	reflect "reflect"

	domain "github.com/D-D-EINA-Calendar/CalendarServer/src/internal/core/domain"
	gomock "github.com/golang/mock/gomock"
)

// MockHorarioService is a mock of HorarioService interface.
type MockHorarioService struct {
	ctrl     *gomock.Controller
	recorder *MockHorarioServiceMockRecorder
}

// MockHorarioServiceMockRecorder is the mock recorder for MockHorarioService.
type MockHorarioServiceMockRecorder struct {
	mock *MockHorarioService
}

// NewMockHorarioService creates a new mock instance.
func NewMockHorarioService(ctrl *gomock.Controller) *MockHorarioService {
	mock := &MockHorarioService{ctrl: ctrl}
	mock.recorder = &MockHorarioServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockHorarioService) EXPECT() *MockHorarioServiceMockRecorder {
	return m.recorder
}

// GetAvailableHours mocks base method.
func (m *MockHorarioService) GetAvailableHours(terna domain.Terna) ([]domain.AvailableHours, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAvailableHours", terna)
	ret0, _ := ret[0].([]domain.AvailableHours)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAvailableHours indicates an expected call of GetAvailableHours.
func (mr *MockHorarioServiceMockRecorder) GetAvailableHours(terna interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAvailableHours", reflect.TypeOf((*MockHorarioService)(nil).GetAvailableHours), terna)
}

// GetEntries mocks base method.
func (m *MockHorarioService) GetEntries(terna domain.Terna) ([]domain.Entry, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEntries", terna)
	ret0, _ := ret[0].([]domain.Entry)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEntries indicates an expected call of GetEntries.
func (mr *MockHorarioServiceMockRecorder) GetEntries(terna interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEntries", reflect.TypeOf((*MockHorarioService)(nil).GetEntries), terna)
}

// ListAllDegrees mocks base method.
func (m *MockHorarioService) ListAllDegrees() ([]domain.DegreeDescription, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAllDegrees")
	ret0, _ := ret[0].([]domain.DegreeDescription)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAllDegrees indicates an expected call of ListAllDegrees.
func (mr *MockHorarioServiceMockRecorder) ListAllDegrees() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAllDegrees", reflect.TypeOf((*MockHorarioService)(nil).ListAllDegrees))
}

// UpdateScheduler mocks base method.
func (m *MockHorarioService) UpdateScheduler(entries []domain.Entry) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateScheduler", entries)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateScheduler indicates an expected call of UpdateScheduler.
func (mr *MockHorarioServiceMockRecorder) UpdateScheduler(entries interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateScheduler", reflect.TypeOf((*MockHorarioService)(nil).UpdateScheduler), entries)
}
