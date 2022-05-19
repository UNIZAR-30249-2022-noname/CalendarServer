// Code generated by MockGen. DO NOT EDIT.
// Source: src/internal/core/ports/repositories.go

// Package mock_ports is a generated GoMock package.
package mock_ports

import (
	reflect "reflect"

	domain "github.com/D-D-EINA-Calendar/CalendarServer/src/internal/core/domain"
	gomock "github.com/golang/mock/gomock"
)

// MockSchedulerRepository is a mock of SchedulerRepository interface.
type MockSchedulerRepository struct {
	ctrl     *gomock.Controller
	recorder *MockSchedulerRepositoryMockRecorder
}

// MockSchedulerRepositoryMockRecorder is the mock recorder for MockSchedulerRepository.
type MockSchedulerRepositoryMockRecorder struct {
	mock *MockSchedulerRepository
}

// NewMockSchedulerRepository creates a new mock instance.
func NewMockSchedulerRepository(ctrl *gomock.Controller) *MockSchedulerRepository {
	mock := &MockSchedulerRepository{ctrl: ctrl}
	mock.recorder = &MockSchedulerRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSchedulerRepository) EXPECT() *MockSchedulerRepositoryMockRecorder {
	return m.recorder
}

// DeleteAllEntries mocks base method.
func (m *MockSchedulerRepository) DeleteAllEntries(terna domain.DegreeSet) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAllEntries", terna)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAllEntries indicates an expected call of DeleteAllEntries.
func (mr *MockSchedulerRepositoryMockRecorder) DeleteAllEntries(terna interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAllEntries", reflect.TypeOf((*MockSchedulerRepository)(nil).DeleteAllEntries), terna)
}

// DeleteEntry mocks base method.
func (m *MockSchedulerRepository) DeleteEntry(req domain.Entry) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteEntry", req)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteEntry indicates an expected call of DeleteEntry.
func (mr *MockSchedulerRepositoryMockRecorder) DeleteEntry(req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteEntry", reflect.TypeOf((*MockSchedulerRepository)(nil).DeleteEntry), req)
}

// GetAvailableHours mocks base method.
func (m *MockSchedulerRepository) GetAvailableHours(req domain.DegreeSet) ([]domain.AvailableHours, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAvailableHours", req)
	ret0, _ := ret[0].([]domain.AvailableHours)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAvailableHours indicates an expected call of GetAvailableHours.
func (mr *MockSchedulerRepositoryMockRecorder) GetAvailableHours(req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAvailableHours", reflect.TypeOf((*MockSchedulerRepository)(nil).GetAvailableHours), req)
}

// GetEntries mocks base method.
func (m *MockSchedulerRepository) GetEntries(req domain.DegreeSet) ([]domain.Entry, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEntries", req)
	ret0, _ := ret[0].([]domain.Entry)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEntries indicates an expected call of GetEntries.
func (mr *MockSchedulerRepositoryMockRecorder) GetEntries(req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEntries", reflect.TypeOf((*MockSchedulerRepository)(nil).GetEntries), req)
}

// ListAllDegrees mocks base method.
func (m *MockSchedulerRepository) ListAllDegrees() ([]domain.DegreeDescription, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAllDegrees")
	ret0, _ := ret[0].([]domain.DegreeDescription)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAllDegrees indicates an expected call of ListAllDegrees.
func (mr *MockSchedulerRepositoryMockRecorder) ListAllDegrees() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAllDegrees", reflect.TypeOf((*MockSchedulerRepository)(nil).ListAllDegrees))
}

// UpdateScheduler mocks base method.
func (m *MockSchedulerRepository) UpdateScheduler(entries []domain.Entry, terna domain.DegreeSet) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateScheduler", entries, terna)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateScheduler indicates an expected call of UpdateScheduler.
func (mr *MockSchedulerRepositoryMockRecorder) UpdateScheduler(entries, terna interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateScheduler", reflect.TypeOf((*MockSchedulerRepository)(nil).UpdateScheduler), entries, terna)
}

// MockUploadDataRepository is a mock of UploadDataRepository interface.
type MockUploadDataRepository struct {
	ctrl     *gomock.Controller
	recorder *MockUploadDataRepositoryMockRecorder
}

// MockUploadDataRepositoryMockRecorder is the mock recorder for MockUploadDataRepository.
type MockUploadDataRepositoryMockRecorder struct {
	mock *MockUploadDataRepository
}

// NewMockUploadDataRepository creates a new mock instance.
func NewMockUploadDataRepository(ctrl *gomock.Controller) *MockUploadDataRepository {
	mock := &MockUploadDataRepository{ctrl: ctrl}
	mock.recorder = &MockUploadDataRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUploadDataRepository) EXPECT() *MockUploadDataRepositoryMockRecorder {
	return m.recorder
}

// CreateNewDegree mocks base method.
func (m *MockUploadDataRepository) CreateNewDegree(id int, name string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateNewDegree", id, name)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateNewDegree indicates an expected call of CreateNewDegree.
func (mr *MockUploadDataRepositoryMockRecorder) CreateNewDegree(id, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateNewDegree", reflect.TypeOf((*MockUploadDataRepository)(nil).CreateNewDegree), id, name)
}

// CreateNewGroup mocks base method.
func (m *MockUploadDataRepository) CreateNewGroup(group, yearCode int) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateNewGroup", group, yearCode)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateNewGroup indicates an expected call of CreateNewGroup.
func (mr *MockUploadDataRepositoryMockRecorder) CreateNewGroup(group, yearCode interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateNewGroup", reflect.TypeOf((*MockUploadDataRepository)(nil).CreateNewGroup), group, yearCode)
}

// CreateNewHour mocks base method.
func (m *MockUploadDataRepository) CreateNewHour(available, total, subjectCode, groupCode, kind int, group, week string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateNewHour", available, total, subjectCode, groupCode, kind, group, week)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateNewHour indicates an expected call of CreateNewHour.
func (mr *MockUploadDataRepositoryMockRecorder) CreateNewHour(available, total, subjectCode, groupCode, kind, group, week interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateNewHour", reflect.TypeOf((*MockUploadDataRepository)(nil).CreateNewHour), available, total, subjectCode, groupCode, kind, group, week)
}

// CreateNewSubject mocks base method.
func (m *MockUploadDataRepository) CreateNewSubject(id int, name string, degreeCode int) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateNewSubject", id, name, degreeCode)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateNewSubject indicates an expected call of CreateNewSubject.
func (mr *MockUploadDataRepositoryMockRecorder) CreateNewSubject(id, name, degreeCode interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateNewSubject", reflect.TypeOf((*MockUploadDataRepository)(nil).CreateNewSubject), id, name, degreeCode)
}

// CreateNewYear mocks base method.
func (m *MockUploadDataRepository) CreateNewYear(year, degreeCode int) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateNewYear", year, degreeCode)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateNewYear indicates an expected call of CreateNewYear.
func (mr *MockUploadDataRepositoryMockRecorder) CreateNewYear(year, degreeCode interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateNewYear", reflect.TypeOf((*MockUploadDataRepository)(nil).CreateNewYear), year, degreeCode)
}

// RawExec mocks base method.
func (m *MockUploadDataRepository) RawExec(exec string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RawExec", exec)
	ret0, _ := ret[0].(error)
	return ret0
}

// RawExec indicates an expected call of RawExec.
func (mr *MockUploadDataRepositoryMockRecorder) RawExec(exec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RawExec", reflect.TypeOf((*MockUploadDataRepository)(nil).RawExec), exec)
}

// MockMonitoringRepository is a mock of MonitoringRepository interface.
type MockMonitoringRepository struct {
	ctrl     *gomock.Controller
	recorder *MockMonitoringRepositoryMockRecorder
}

// MockMonitoringRepositoryMockRecorder is the mock recorder for MockMonitoringRepository.
type MockMonitoringRepositoryMockRecorder struct {
	mock *MockMonitoringRepository
}

// NewMockMonitoringRepository creates a new mock instance.
func NewMockMonitoringRepository(ctrl *gomock.Controller) *MockMonitoringRepository {
	mock := &MockMonitoringRepository{ctrl: ctrl}
	mock.recorder = &MockMonitoringRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMonitoringRepository) EXPECT() *MockMonitoringRepositoryMockRecorder {
	return m.recorder
}

// Ping mocks base method.
func (m *MockMonitoringRepository) Ping() (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Ping")
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Ping indicates an expected call of Ping.
func (mr *MockMonitoringRepositoryMockRecorder) Ping() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Ping", reflect.TypeOf((*MockMonitoringRepository)(nil).Ping))
}

// MockUsersRepository is a mock of UsersRepository interface.
type MockUsersRepository struct {
	ctrl     *gomock.Controller
	recorder *MockUsersRepositoryMockRecorder
}

// MockUsersRepositoryMockRecorder is the mock recorder for MockUsersRepository.
type MockUsersRepositoryMockRecorder struct {
	mock *MockUsersRepository
}

// NewMockUsersRepository creates a new mock instance.
func NewMockUsersRepository(ctrl *gomock.Controller) *MockUsersRepository {
	mock := &MockUsersRepository{ctrl: ctrl}
	mock.recorder = &MockUsersRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUsersRepository) EXPECT() *MockUsersRepositoryMockRecorder {
	return m.recorder
}

// GetCredentials mocks base method.
func (m *MockUsersRepository) GetCredentials(username string) (domain.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCredentials", username)
	ret0, _ := ret[0].(domain.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCredentials indicates an expected call of GetCredentials.
func (mr *MockUsersRepositoryMockRecorder) GetCredentials(username interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCredentials", reflect.TypeOf((*MockUsersRepository)(nil).GetCredentials), username)
}

// MockSpaceRepository is a mock of SpaceRepository interface.
type MockSpaceRepository struct {
	ctrl     *gomock.Controller
	recorder *MockSpaceRepositoryMockRecorder
}

// MockSpaceRepositoryMockRecorder is the mock recorder for MockSpaceRepository.
type MockSpaceRepositoryMockRecorder struct {
	mock *MockSpaceRepository
}

// NewMockSpaceRepository creates a new mock instance.
func NewMockSpaceRepository(ctrl *gomock.Controller) *MockSpaceRepository {
	mock := &MockSpaceRepository{ctrl: ctrl}
	mock.recorder = &MockSpaceRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSpaceRepository) EXPECT() *MockSpaceRepositoryMockRecorder {
	return m.recorder
}

// CancelReserve mocks base method.
func (m *MockSpaceRepository) CancelReserve(key string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CancelReserve", key)
	ret0, _ := ret[0].(error)
	return ret0
}

// CancelReserve indicates an expected call of CancelReserve.
func (mr *MockSpaceRepositoryMockRecorder) CancelReserve(key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CancelReserve", reflect.TypeOf((*MockSpaceRepository)(nil).CancelReserve), key)
}

// FilterBy mocks base method.
func (m *MockSpaceRepository) FilterBy(arg0 domain.SpaceFilterParams) ([]domain.Space, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FilterBy", arg0)
	ret0, _ := ret[0].([]domain.Space)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FilterBy indicates an expected call of FilterBy.
func (mr *MockSpaceRepositoryMockRecorder) FilterBy(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FilterBy", reflect.TypeOf((*MockSpaceRepository)(nil).FilterBy), arg0)
}

// GetReservesOwner mocks base method.
func (m *MockSpaceRepository) GetReservesOwner(owner string) ([]domain.Reserve, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetReservesOwner", owner)
	ret0, _ := ret[0].([]domain.Reserve)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetReservesOwner indicates an expected call of GetReservesOwner.
func (mr *MockSpaceRepositoryMockRecorder) GetReservesOwner(owner interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetReservesOwner", reflect.TypeOf((*MockSpaceRepository)(nil).GetReservesOwner), owner)
}

// RequestInfoSlots mocks base method.
func (m *MockSpaceRepository) RequestInfoSlots(req domain.ReqInfoSlot) (domain.AllInfoSlot, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RequestInfoSlots", req)
	ret0, _ := ret[0].(domain.AllInfoSlot)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RequestInfoSlots indicates an expected call of RequestInfoSlots.
func (mr *MockSpaceRepositoryMockRecorder) RequestInfoSlots(req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RequestInfoSlots", reflect.TypeOf((*MockSpaceRepository)(nil).RequestInfoSlots), req)
}

// Reserve mocks base method.
func (m *MockSpaceRepository) Reserve(sp string, init, end domain.Hour, date, person, event string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Reserve", sp, init, end, date, person, event)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Reserve indicates an expected call of Reserve.
func (mr *MockSpaceRepositoryMockRecorder) Reserve(sp, init, end, date, person, event interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Reserve", reflect.TypeOf((*MockSpaceRepository)(nil).Reserve), sp, init, end, date, person, event)
}

// ReserveBatch mocks base method.
func (m *MockSpaceRepository) ReserveBatch(spaces []string, init, end domain.Hour, dates []string, person string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReserveBatch", spaces, init, end, dates, person)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReserveBatch indicates an expected call of ReserveBatch.
func (mr *MockSpaceRepositoryMockRecorder) ReserveBatch(spaces, init, end, dates, person interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReserveBatch", reflect.TypeOf((*MockSpaceRepository)(nil).ReserveBatch), spaces, init, end, dates, person)
}

// MockIssueRepository is a mock of IssueRepository interface.
type MockIssueRepository struct {
	ctrl     *gomock.Controller
	recorder *MockIssueRepositoryMockRecorder
}

// MockIssueRepositoryMockRecorder is the mock recorder for MockIssueRepository.
type MockIssueRepositoryMockRecorder struct {
	mock *MockIssueRepository
}

// NewMockIssueRepository creates a new mock instance.
func NewMockIssueRepository(ctrl *gomock.Controller) *MockIssueRepository {
	mock := &MockIssueRepository{ctrl: ctrl}
	mock.recorder = &MockIssueRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIssueRepository) EXPECT() *MockIssueRepositoryMockRecorder {
	return m.recorder
}

// ChangeState mocks base method.
func (m *MockIssueRepository) ChangeState(key string, state int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChangeState", key, state)
	ret0, _ := ret[0].(error)
	return ret0
}

// ChangeState indicates an expected call of ChangeState.
func (mr *MockIssueRepositoryMockRecorder) ChangeState(key, state interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangeState", reflect.TypeOf((*MockIssueRepository)(nil).ChangeState), key, state)
}

// Create mocks base method.
func (m *MockIssueRepository) Create(issue domain.Issue) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", issue)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockIssueRepositoryMockRecorder) Create(issue interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockIssueRepository)(nil).Create), issue)
}

// Delete mocks base method.
func (m *MockIssueRepository) Delete(key string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", key)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockIssueRepositoryMockRecorder) Delete(key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockIssueRepository)(nil).Delete), key)
}

// GetAll mocks base method.
func (m *MockIssueRepository) GetAll() ([]domain.Issue, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll")
	ret0, _ := ret[0].([]domain.Issue)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll.
func (mr *MockIssueRepositoryMockRecorder) GetAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockIssueRepository)(nil).GetAll))
}

func (m *MockSchedulerRepository) GetICS(terna domain.DegreeSet) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetICS", terna)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteAllEntries indicates an expected call of DeleteAllEntries.
func (mr *MockSchedulerRepositoryMockRecorder) GetICS(terna interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetICS", reflect.TypeOf((*MockSchedulerRepository)(nil).GetICS), terna)
}