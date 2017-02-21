// Code generated by thriftrw-plugin-yarpc
// @generated

package dosatest

import (
	"context"
	"go.uber.org/yarpc"
	"github.com/golang/mock/gomock"
	"github.com/uber/dosa-idl/.gen/dosa"
	"github.com/uber/dosa-idl/.gen/dosa/dosaclient"
)

// MockClient implements a gomock-compatible mock client for service
// Dosa.
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *_MockClientRecorder
}

var _ dosaclient.Interface = (*MockClient)(nil)

type _MockClientRecorder struct {
	mock *MockClient
}

// Build a new mock client for service Dosa.
//
// 	mockCtrl := gomock.NewController(t)
// 	client := dosatest.NewMockClient(mockCtrl)
//
// Use EXPECT() to set expectations on the mock.
func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &_MockClientRecorder{mock}
	return mock
}

// EXPECT returns an object that allows you to define an expectation on the
// Dosa mock client.
func (m *MockClient) EXPECT() *_MockClientRecorder {
	return m.recorder
}

// BatchRead responds to a BatchRead call based on the mock expectations. This
// call will fail if the mock does not expect this call. Use EXPECT to expect
// a call to this function.
//
// 	client.EXPECT().BatchRead(gomock.Any(), ...).Return(...)
// 	... := client.BatchRead(...)
func (m *MockClient) BatchRead(
	ctx context.Context,
	_Request *dosa.BatchReadRequest,
	opts ...yarpc.CallOption,
) (success *dosa.BatchReadResponse, err error) {

	args := []interface{}{ctx, _Request}
	for _, o := range opts {
		args = append(args, o)
	}
	i := 0
	ret := m.ctrl.Call(m, "BatchRead", args...)
	success, _ = ret[i].(*dosa.BatchReadResponse)
	i++
	err, _ = ret[i].(error)
	return
}

func (mr *_MockClientRecorder) BatchRead(
	ctx interface{},
	_Request interface{},
	opts ...interface{},
) *gomock.Call {
	args := append([]interface{}{ctx, _Request}, opts...)
	return mr.mock.ctrl.RecordCall(mr.mock, "BatchRead", args...)
}

// CheckSchema responds to a CheckSchema call based on the mock expectations. This
// call will fail if the mock does not expect this call. Use EXPECT to expect
// a call to this function.
//
// 	client.EXPECT().CheckSchema(gomock.Any(), ...).Return(...)
// 	... := client.CheckSchema(...)
func (m *MockClient) CheckSchema(
	ctx context.Context,
	_Request *dosa.CheckSchemaRequest,
	opts ...yarpc.CallOption,
) (success *dosa.CheckSchemaResponse, err error) {

	args := []interface{}{ctx, _Request}
	for _, o := range opts {
		args = append(args, o)
	}
	i := 0
	ret := m.ctrl.Call(m, "CheckSchema", args...)
	success, _ = ret[i].(*dosa.CheckSchemaResponse)
	i++
	err, _ = ret[i].(error)
	return
}

func (mr *_MockClientRecorder) CheckSchema(
	ctx interface{},
	_Request interface{},
	opts ...interface{},
) *gomock.Call {
	args := append([]interface{}{ctx, _Request}, opts...)
	return mr.mock.ctrl.RecordCall(mr.mock, "CheckSchema", args...)
}

// CreateIfNotExists responds to a CreateIfNotExists call based on the mock expectations. This
// call will fail if the mock does not expect this call. Use EXPECT to expect
// a call to this function.
//
// 	client.EXPECT().CreateIfNotExists(gomock.Any(), ...).Return(...)
// 	... := client.CreateIfNotExists(...)
func (m *MockClient) CreateIfNotExists(
	ctx context.Context,
	_Request *dosa.CreateRequest,
	opts ...yarpc.CallOption,
) (err error) {

	args := []interface{}{ctx, _Request}
	for _, o := range opts {
		args = append(args, o)
	}
	i := 0
	ret := m.ctrl.Call(m, "CreateIfNotExists", args...)
	err, _ = ret[i].(error)
	return
}

func (mr *_MockClientRecorder) CreateIfNotExists(
	ctx interface{},
	_Request interface{},
	opts ...interface{},
) *gomock.Call {
	args := append([]interface{}{ctx, _Request}, opts...)
	return mr.mock.ctrl.RecordCall(mr.mock, "CreateIfNotExists", args...)
}

// CreateScope responds to a CreateScope call based on the mock expectations. This
// call will fail if the mock does not expect this call. Use EXPECT to expect
// a call to this function.
//
// 	client.EXPECT().CreateScope(gomock.Any(), ...).Return(...)
// 	... := client.CreateScope(...)
func (m *MockClient) CreateScope(
	ctx context.Context,
	_Request *dosa.CreateScopeRequest,
	opts ...yarpc.CallOption,
) (err error) {

	args := []interface{}{ctx, _Request}
	for _, o := range opts {
		args = append(args, o)
	}
	i := 0
	ret := m.ctrl.Call(m, "CreateScope", args...)
	err, _ = ret[i].(error)
	return
}

func (mr *_MockClientRecorder) CreateScope(
	ctx interface{},
	_Request interface{},
	opts ...interface{},
) *gomock.Call {
	args := append([]interface{}{ctx, _Request}, opts...)
	return mr.mock.ctrl.RecordCall(mr.mock, "CreateScope", args...)
}

// DropScope responds to a DropScope call based on the mock expectations. This
// call will fail if the mock does not expect this call. Use EXPECT to expect
// a call to this function.
//
// 	client.EXPECT().DropScope(gomock.Any(), ...).Return(...)
// 	... := client.DropScope(...)
func (m *MockClient) DropScope(
	ctx context.Context,
	_Request *dosa.DropScopeRequest,
	opts ...yarpc.CallOption,
) (err error) {

	args := []interface{}{ctx, _Request}
	for _, o := range opts {
		args = append(args, o)
	}
	i := 0
	ret := m.ctrl.Call(m, "DropScope", args...)
	err, _ = ret[i].(error)
	return
}

func (mr *_MockClientRecorder) DropScope(
	ctx interface{},
	_Request interface{},
	opts ...interface{},
) *gomock.Call {
	args := append([]interface{}{ctx, _Request}, opts...)
	return mr.mock.ctrl.RecordCall(mr.mock, "DropScope", args...)
}

// Range responds to a Range call based on the mock expectations. This
// call will fail if the mock does not expect this call. Use EXPECT to expect
// a call to this function.
//
// 	client.EXPECT().Range(gomock.Any(), ...).Return(...)
// 	... := client.Range(...)
func (m *MockClient) Range(
	ctx context.Context,
	_Request *dosa.RangeRequest,
	opts ...yarpc.CallOption,
) (success *dosa.RangeResponse, err error) {

	args := []interface{}{ctx, _Request}
	for _, o := range opts {
		args = append(args, o)
	}
	i := 0
	ret := m.ctrl.Call(m, "Range", args...)
	success, _ = ret[i].(*dosa.RangeResponse)
	i++
	err, _ = ret[i].(error)
	return
}

func (mr *_MockClientRecorder) Range(
	ctx interface{},
	_Request interface{},
	opts ...interface{},
) *gomock.Call {
	args := append([]interface{}{ctx, _Request}, opts...)
	return mr.mock.ctrl.RecordCall(mr.mock, "Range", args...)
}

// Read responds to a Read call based on the mock expectations. This
// call will fail if the mock does not expect this call. Use EXPECT to expect
// a call to this function.
//
// 	client.EXPECT().Read(gomock.Any(), ...).Return(...)
// 	... := client.Read(...)
func (m *MockClient) Read(
	ctx context.Context,
	_Request *dosa.ReadRequest,
	opts ...yarpc.CallOption,
) (success *dosa.ReadResponse, err error) {

	args := []interface{}{ctx, _Request}
	for _, o := range opts {
		args = append(args, o)
	}
	i := 0
	ret := m.ctrl.Call(m, "Read", args...)
	success, _ = ret[i].(*dosa.ReadResponse)
	i++
	err, _ = ret[i].(error)
	return
}

func (mr *_MockClientRecorder) Read(
	ctx interface{},
	_Request interface{},
	opts ...interface{},
) *gomock.Call {
	args := append([]interface{}{ctx, _Request}, opts...)
	return mr.mock.ctrl.RecordCall(mr.mock, "Read", args...)
}

// Remove responds to a Remove call based on the mock expectations. This
// call will fail if the mock does not expect this call. Use EXPECT to expect
// a call to this function.
//
// 	client.EXPECT().Remove(gomock.Any(), ...).Return(...)
// 	... := client.Remove(...)
func (m *MockClient) Remove(
	ctx context.Context,
	_Request *dosa.RemoveRequest,
	opts ...yarpc.CallOption,
) (err error) {

	args := []interface{}{ctx, _Request}
	for _, o := range opts {
		args = append(args, o)
	}
	i := 0
	ret := m.ctrl.Call(m, "Remove", args...)
	err, _ = ret[i].(error)
	return
}

func (mr *_MockClientRecorder) Remove(
	ctx interface{},
	_Request interface{},
	opts ...interface{},
) *gomock.Call {
	args := append([]interface{}{ctx, _Request}, opts...)
	return mr.mock.ctrl.RecordCall(mr.mock, "Remove", args...)
}

// Scan responds to a Scan call based on the mock expectations. This
// call will fail if the mock does not expect this call. Use EXPECT to expect
// a call to this function.
//
// 	client.EXPECT().Scan(gomock.Any(), ...).Return(...)
// 	... := client.Scan(...)
func (m *MockClient) Scan(
	ctx context.Context,
	_Request *dosa.ScanRequest,
	opts ...yarpc.CallOption,
) (success *dosa.ScanResponse, err error) {

	args := []interface{}{ctx, _Request}
	for _, o := range opts {
		args = append(args, o)
	}
	i := 0
	ret := m.ctrl.Call(m, "Scan", args...)
	success, _ = ret[i].(*dosa.ScanResponse)
	i++
	err, _ = ret[i].(error)
	return
}

func (mr *_MockClientRecorder) Scan(
	ctx interface{},
	_Request interface{},
	opts ...interface{},
) *gomock.Call {
	args := append([]interface{}{ctx, _Request}, opts...)
	return mr.mock.ctrl.RecordCall(mr.mock, "Scan", args...)
}

// Search responds to a Search call based on the mock expectations. This
// call will fail if the mock does not expect this call. Use EXPECT to expect
// a call to this function.
//
// 	client.EXPECT().Search(gomock.Any(), ...).Return(...)
// 	... := client.Search(...)
func (m *MockClient) Search(
	ctx context.Context,
	_Request *dosa.SearchRequest,
	opts ...yarpc.CallOption,
) (success *dosa.SearchResponse, err error) {

	args := []interface{}{ctx, _Request}
	for _, o := range opts {
		args = append(args, o)
	}
	i := 0
	ret := m.ctrl.Call(m, "Search", args...)
	success, _ = ret[i].(*dosa.SearchResponse)
	i++
	err, _ = ret[i].(error)
	return
}

func (mr *_MockClientRecorder) Search(
	ctx interface{},
	_Request interface{},
	opts ...interface{},
) *gomock.Call {
	args := append([]interface{}{ctx, _Request}, opts...)
	return mr.mock.ctrl.RecordCall(mr.mock, "Search", args...)
}

// TruncateScope responds to a TruncateScope call based on the mock expectations. This
// call will fail if the mock does not expect this call. Use EXPECT to expect
// a call to this function.
//
// 	client.EXPECT().TruncateScope(gomock.Any(), ...).Return(...)
// 	... := client.TruncateScope(...)
func (m *MockClient) TruncateScope(
	ctx context.Context,
	_Request *dosa.TruncateScopeRequest,
	opts ...yarpc.CallOption,
) (err error) {

	args := []interface{}{ctx, _Request}
	for _, o := range opts {
		args = append(args, o)
	}
	i := 0
	ret := m.ctrl.Call(m, "TruncateScope", args...)
	err, _ = ret[i].(error)
	return
}

func (mr *_MockClientRecorder) TruncateScope(
	ctx interface{},
	_Request interface{},
	opts ...interface{},
) *gomock.Call {
	args := append([]interface{}{ctx, _Request}, opts...)
	return mr.mock.ctrl.RecordCall(mr.mock, "TruncateScope", args...)
}

// Upsert responds to a Upsert call based on the mock expectations. This
// call will fail if the mock does not expect this call. Use EXPECT to expect
// a call to this function.
//
// 	client.EXPECT().Upsert(gomock.Any(), ...).Return(...)
// 	... := client.Upsert(...)
func (m *MockClient) Upsert(
	ctx context.Context,
	_Request *dosa.UpsertRequest,
	opts ...yarpc.CallOption,
) (err error) {

	args := []interface{}{ctx, _Request}
	for _, o := range opts {
		args = append(args, o)
	}
	i := 0
	ret := m.ctrl.Call(m, "Upsert", args...)
	err, _ = ret[i].(error)
	return
}

func (mr *_MockClientRecorder) Upsert(
	ctx interface{},
	_Request interface{},
	opts ...interface{},
) *gomock.Call {
	args := append([]interface{}{ctx, _Request}, opts...)
	return mr.mock.ctrl.RecordCall(mr.mock, "Upsert", args...)
}

// UpsertSchema responds to a UpsertSchema call based on the mock expectations. This
// call will fail if the mock does not expect this call. Use EXPECT to expect
// a call to this function.
//
// 	client.EXPECT().UpsertSchema(gomock.Any(), ...).Return(...)
// 	... := client.UpsertSchema(...)
func (m *MockClient) UpsertSchema(
	ctx context.Context,
	_Request *dosa.UpsertSchemaRequest,
	opts ...yarpc.CallOption,
) (success *dosa.UpsertSchemaResponse, err error) {

	args := []interface{}{ctx, _Request}
	for _, o := range opts {
		args = append(args, o)
	}
	i := 0
	ret := m.ctrl.Call(m, "UpsertSchema", args...)
	success, _ = ret[i].(*dosa.UpsertSchemaResponse)
	i++
	err, _ = ret[i].(error)
	return
}

func (mr *_MockClientRecorder) UpsertSchema(
	ctx interface{},
	_Request interface{},
	opts ...interface{},
) *gomock.Call {
	args := append([]interface{}{ctx, _Request}, opts...)
	return mr.mock.ctrl.RecordCall(mr.mock, "UpsertSchema", args...)
}
