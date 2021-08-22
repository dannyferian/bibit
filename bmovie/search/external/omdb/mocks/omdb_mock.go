// Code generated by MockGen. DO NOT EDIT.
// Source: client.go

// Package omdb_mock is a generated GoMock package.
package omdb_mock

import (
	context "context"
	reflect "reflect"

	omdb "github.com/danClauz/bibit/bmovie/search/external/omdb"
	gomock "github.com/golang/mock/gomock"
)

// MockClient is a mock of Client interface.
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *MockClientMockRecorder
}

// MockClientMockRecorder is the mock recorder for MockClient.
type MockClientMockRecorder struct {
	mock *MockClient
}

// NewMockClient creates a new mock instance.
func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &MockClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClient) EXPECT() *MockClientMockRecorder {
	return m.recorder
}

// SearchMovie mocks base method.
func (m *MockClient) SearchMovie(ctx context.Context, reqId, searchKey string, page int) (*omdb.SearchResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchMovie", ctx, reqId, searchKey, page)
	ret0, _ := ret[0].(*omdb.SearchResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchMovie indicates an expected call of SearchMovie.
func (mr *MockClientMockRecorder) SearchMovie(ctx, reqId, searchKey, page interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchMovie", reflect.TypeOf((*MockClient)(nil).SearchMovie), ctx, reqId, searchKey, page)
}

// SearchMovieByImdbId mocks base method.
func (m *MockClient) SearchMovieByImdbId(ctx context.Context, reqId, searchKey string) (*omdb.SearchByTitleResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchMovieByImdbId", ctx, reqId, searchKey)
	ret0, _ := ret[0].(*omdb.SearchByTitleResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchMovieByImdbId indicates an expected call of SearchMovieByImdbId.
func (mr *MockClientMockRecorder) SearchMovieByImdbId(ctx, reqId, searchKey interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchMovieByImdbId", reflect.TypeOf((*MockClient)(nil).SearchMovieByImdbId), ctx, reqId, searchKey)
}