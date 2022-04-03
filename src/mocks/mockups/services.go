// Code generated by MockGen. DO NOT EDIT.
// Source: src/internal/core/ports/services.go

// Package mock_ports is a generated GoMock package.
package mock_ports

import (
	reflect "reflect"

	domain "github.com/D-D-EINA-Calendar/CalendarServer/src/internal/core/domain"
	gomock "github.com/golang/mock/gomock"
)

// MockSchedulerService is a mock of SchedulerService interface.
type MockSchedulerService struct {
	ctrl     *gomock.Controller
	recorder *MockSchedulerServiceMockRecorder
}

// MockSchedulerServiceMockRecorder is the mock recorder for MockSchedulerService.
type MockSchedulerServiceMockRecorder struct {
	mock *MockSchedulerService
}

// NewMockSchedulerService creates a new mock instance.
func NewMockSchedulerService(ctrl *gomock.Controller) *MockSchedulerService {
	mock := &MockSchedulerService{ctrl: ctrl}
	mock.recorder = &MockSchedulerServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSchedulerService) EXPECT() *MockSchedulerServiceMockRecorder {
	return m.recorder
}

// GetAvailableHours mocks base method.
func (m *MockSchedulerService) GetAvailableHours(terna domain.DegreeSet) ([]domain.AvailableHours, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAvailableHours", terna)
	ret0, _ := ret[0].([]domain.AvailableHours)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAvailableHours indicates an expected call of GetAvailableHours.
func (mr *MockSchedulerServiceMockRecorder) GetAvailableHours(terna interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAvailableHours", reflect.TypeOf((*MockSchedulerService)(nil).GetAvailableHours), terna)
}

// GetEntries mocks base method.
func (m *MockSchedulerService) GetEntries(terna domain.DegreeSet) ([]domain.Entry, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEntries", terna)
	ret0, _ := ret[0].([]domain.Entry)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEntries indicates an expected call of GetEntries.
func (mr *MockSchedulerServiceMockRecorder) GetEntries(terna interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEntries", reflect.TypeOf((*MockSchedulerService)(nil).GetEntries), terna)
}

// GetICS mocks base method.
func (m *MockSchedulerService) GetICS(terna domain.DegreeSet) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetICS", terna)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetICS indicates an expected call of GetICS.
func (mr *MockSchedulerServiceMockRecorder) GetICS(terna interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetICS", reflect.TypeOf((*MockSchedulerService)(nil).GetICS), terna)
}

// ListAllDegrees mocks base method.
func (m *MockSchedulerService) ListAllDegrees() ([]domain.DegreeDescription, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAllDegrees")
	ret0, _ := ret[0].([]domain.DegreeDescription)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAllDegrees indicates an expected call of ListAllDegrees.
func (mr *MockSchedulerServiceMockRecorder) ListAllDegrees() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAllDegrees", reflect.TypeOf((*MockSchedulerService)(nil).ListAllDegrees))
}

// UpdateScheduler mocks base method.
func (m *MockSchedulerService) UpdateScheduler(entries []domain.Entry, terna domain.DegreeSet) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateScheduler", entries, terna)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateScheduler indicates an expected call of UpdateScheduler.
func (mr *MockSchedulerServiceMockRecorder) UpdateScheduler(entries, terna interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateScheduler", reflect.TypeOf((*MockSchedulerService)(nil).UpdateScheduler), entries, terna)
}

// MockUploadDataService is a mock of UploadDataService interface.
type MockUploadDataService struct {
	ctrl     *gomock.Controller
	recorder *MockUploadDataServiceMockRecorder
}

// MockUploadDataServiceMockRecorder is the mock recorder for MockUploadDataService.
type MockUploadDataServiceMockRecorder struct {
	mock *MockUploadDataService
}

// NewMockUploadDataService creates a new mock instance.
func NewMockUploadDataService(ctrl *gomock.Controller) *MockUploadDataService {
	mock := &MockUploadDataService{ctrl: ctrl}
	mock.recorder = &MockUploadDataServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUploadDataService) EXPECT() *MockUploadDataServiceMockRecorder {
	return m.recorder
}

// UpdateByCSV mocks base method.
func (m *MockUploadDataService) UpdateByCSV(csv string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateByCSV", csv)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateByCSV indicates an expected call of UpdateByCSV.
func (mr *MockUploadDataServiceMockRecorder) UpdateByCSV(csv interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateByCSV", reflect.TypeOf((*MockUploadDataService)(nil).UpdateByCSV), csv)
}

// MockMonitoringService is a mock of MonitoringService interface.
type MockMonitoringService struct {
	ctrl     *gomock.Controller
	recorder *MockMonitoringServiceMockRecorder
}

// MockMonitoringServiceMockRecorder is the mock recorder for MockMonitoringService.
type MockMonitoringServiceMockRecorder struct {
	mock *MockMonitoringService
}

// NewMockMonitoringService creates a new mock instance.
func NewMockMonitoringService(ctrl *gomock.Controller) *MockMonitoringService {
	mock := &MockMonitoringService{ctrl: ctrl}
	mock.recorder = &MockMonitoringServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMonitoringService) EXPECT() *MockMonitoringServiceMockRecorder {
	return m.recorder
}

// Ping mocks base method.
func (m *MockMonitoringService) Ping() (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Ping")
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Ping indicates an expected call of Ping.
func (mr *MockMonitoringServiceMockRecorder) Ping() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Ping", reflect.TypeOf((*MockMonitoringService)(nil).Ping))
}

// MockUsersService is a mock of UsersService interface.
type MockUsersService struct {
	ctrl     *gomock.Controller
	recorder *MockUsersServiceMockRecorder
}

// MockUsersServiceMockRecorder is the mock recorder for MockUsersService.
type MockUsersServiceMockRecorder struct {
	mock *MockUsersService
}

// NewMockUsersService creates a new mock instance.
func NewMockUsersService(ctrl *gomock.Controller) *MockUsersService {
	mock := &MockUsersService{ctrl: ctrl}
	mock.recorder = &MockUsersServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUsersService) EXPECT() *MockUsersServiceMockRecorder {
	return m.recorder
}

// GetCredentials mocks base method.
func (m *MockUsersService) GetCredentials(username string) (domain.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCredentials", username)
	ret0, _ := ret[0].(domain.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCredentials indicates an expected call of GetCredentials.
func (mr *MockUsersServiceMockRecorder) GetCredentials(username interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCredentials", reflect.TypeOf((*MockUsersService)(nil).GetCredentials), username)
}

// MockSpacesService is a mock of SpacesService interface.
type MockSpacesService struct {
	ctrl     *gomock.Controller
	recorder *MockSpacesServiceMockRecorder
}

// MockSpacesServiceMockRecorder is the mock recorder for MockSpacesService.
type MockSpacesServiceMockRecorder struct {
	mock *MockSpacesService
}

// NewMockSpacesService creates a new mock instance.
func NewMockSpacesService(ctrl *gomock.Controller) *MockSpacesService {
	mock := &MockSpacesService{ctrl: ctrl}
	mock.recorder = &MockSpacesServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSpacesService) EXPECT() *MockSpacesServiceMockRecorder {
	return m.recorder
}

// CancelReserve mocks base method.
func (m *MockSpacesService) CancelReserve(key string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CancelReserve", key)
	ret0, _ := ret[0].(error)
	return ret0
}

// CancelReserve indicates an expected call of CancelReserve.
func (mr *MockSpacesServiceMockRecorder) CancelReserve(key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CancelReserve", reflect.TypeOf((*MockSpacesService)(nil).CancelReserve), key)
}

// FilterBy mocks base method.
func (m *MockSpacesService) FilterBy(arg0 domain.SpaceFilterParams) ([]domain.Spaces, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FilterBy", arg0)
	ret0, _ := ret[0].([]domain.Spaces)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FilterBy indicates an expected call of FilterBy.
func (mr *MockSpacesServiceMockRecorder) FilterBy(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FilterBy", reflect.TypeOf((*MockSpacesService)(nil).FilterBy), arg0)
}

// RequestInfoSlots mocks base method.
func (m *MockSpacesService) RequestInfoSlots(req domain.ReqInfoSlot) (domain.AllInfoSlot, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RequestInfoSlots", req)
	ret0, _ := ret[0].(domain.AllInfoSlot)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RequestInfoSlots indicates an expected call of RequestInfoSlots.
func (mr *MockSpacesServiceMockRecorder) RequestInfoSlots(req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RequestInfoSlots", reflect.TypeOf((*MockSpacesService)(nil).RequestInfoSlots), req)
}

// Reserve mocks base method.
func (m *MockSpacesService) Reserve(sp domain.Space, init, end domain.Hour, date, person string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Reserve", sp, init, end, date, person)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Reserve indicates an expected call of Reserve.
func (mr *MockSpacesServiceMockRecorder) Reserve(sp, init, end, date, person interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Reserve", reflect.TypeOf((*MockSpacesService)(nil).Reserve), sp, init, end, date, person)
}

// ReserveBatch mocks base method.
func (m *MockSpacesService) ReserveBatch(spaces []domain.Space, init, end domain.Hour, dates []string, person string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReserveBatch", spaces, init, end, dates, person)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReserveBatch indicates an expected call of ReserveBatch.
func (mr *MockSpacesServiceMockRecorder) ReserveBatch(spaces, init, end, dates, person interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReserveBatch", reflect.TypeOf((*MockSpacesService)(nil).ReserveBatch), spaces, init, end, dates, person)
}

// MockIssueService is a mock of IssueService interface.
type MockIssueService struct {
	ctrl     *gomock.Controller
	recorder *MockIssueServiceMockRecorder
}

// MockIssueServiceMockRecorder is the mock recorder for MockIssueService.
type MockIssueServiceMockRecorder struct {
	mock *MockIssueService
}

// NewMockIssueService creates a new mock instance.
func NewMockIssueService(ctrl *gomock.Controller) *MockIssueService {
	mock := &MockIssueService{ctrl: ctrl}
	mock.recorder = &MockIssueServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIssueService) EXPECT() *MockIssueServiceMockRecorder {
	return m.recorder
}

// ChangeState mocks base method.
func (m *MockIssueService) ChangeState(key string, state int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChangeState", key, state)
	ret0, _ := ret[0].(error)
	return ret0
}

// ChangeState indicates an expected call of ChangeState.
func (mr *MockIssueServiceMockRecorder) ChangeState(key, state interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangeState", reflect.TypeOf((*MockIssueService)(nil).ChangeState), key, state)
}

// Create mocks base method.
func (m *MockIssueService) Create(issue domain.Issue) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", issue)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockIssueServiceMockRecorder) Create(issue interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockIssueService)(nil).Create), issue)
}

// Delete mocks base method.
func (m *MockIssueService) Delete(key string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", key)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockIssueServiceMockRecorder) Delete(key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockIssueService)(nil).Delete), key)
}

// GetAll mocks base method.
func (m *MockIssueService) GetAll() ([]domain.Issue, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll")
	ret0, _ := ret[0].([]domain.Issue)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll.
func (mr *MockIssueServiceMockRecorder) GetAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockIssueService)(nil).GetAll))
}
