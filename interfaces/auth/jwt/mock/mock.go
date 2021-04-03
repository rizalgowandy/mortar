// Code generated by MockGen. DO NOT EDIT.
// Source: interfaces.go

// Package mock_jwt is a generated GoMock package.
package mock_jwt

import (
	context "context"
	base64 "encoding/base64"
	reflect "reflect"

	jwt "github.com/go-masonry/mortar/interfaces/auth/jwt"
	gomock "github.com/golang/mock/gomock"
)

// MockExtractorBuilder is a mock of ExtractorBuilder interface.
type MockExtractorBuilder struct {
	ctrl     *gomock.Controller
	recorder *MockExtractorBuilderMockRecorder
}

// MockExtractorBuilderMockRecorder is the mock recorder for MockExtractorBuilder.
type MockExtractorBuilderMockRecorder struct {
	mock *MockExtractorBuilder
}

// NewMockExtractorBuilder creates a new mock instance.
func NewMockExtractorBuilder(ctrl *gomock.Controller) *MockExtractorBuilder {
	mock := &MockExtractorBuilder{ctrl: ctrl}
	mock.recorder = &MockExtractorBuilderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockExtractorBuilder) EXPECT() *MockExtractorBuilderMockRecorder {
	return m.recorder
}

// Build mocks base method.
func (m *MockExtractorBuilder) Build() jwt.TokenExtractor {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Build")
	ret0, _ := ret[0].(jwt.TokenExtractor)
	return ret0
}

// Build indicates an expected call of Build.
func (mr *MockExtractorBuilderMockRecorder) Build() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Build", reflect.TypeOf((*MockExtractorBuilder)(nil).Build))
}

// SetBase64Decoder mocks base method.
func (m *MockExtractorBuilder) SetBase64Decoder(dec *base64.Encoding) jwt.ExtractorBuilder {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetBase64Decoder", dec)
	ret0, _ := ret[0].(jwt.ExtractorBuilder)
	return ret0
}

// SetBase64Decoder indicates an expected call of SetBase64Decoder.
func (mr *MockExtractorBuilderMockRecorder) SetBase64Decoder(dec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetBase64Decoder", reflect.TypeOf((*MockExtractorBuilder)(nil).SetBase64Decoder), dec)
}

// SetContextExtractor mocks base method.
func (m *MockExtractorBuilder) SetContextExtractor(extractor jwt.ContextExtractor) jwt.ExtractorBuilder {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetContextExtractor", extractor)
	ret0, _ := ret[0].(jwt.ExtractorBuilder)
	return ret0
}

// SetContextExtractor indicates an expected call of SetContextExtractor.
func (mr *MockExtractorBuilderMockRecorder) SetContextExtractor(extractor interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetContextExtractor", reflect.TypeOf((*MockExtractorBuilder)(nil).SetContextExtractor), extractor)
}

// SetDecoder mocks base method.
func (m *MockExtractorBuilder) SetDecoder(dec jwt.JSONDecoder) jwt.ExtractorBuilder {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetDecoder", dec)
	ret0, _ := ret[0].(jwt.ExtractorBuilder)
	return ret0
}

// SetDecoder indicates an expected call of SetDecoder.
func (mr *MockExtractorBuilderMockRecorder) SetDecoder(dec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetDecoder", reflect.TypeOf((*MockExtractorBuilder)(nil).SetDecoder), dec)
}

// MockTokenExtractor is a mock of TokenExtractor interface.
type MockTokenExtractor struct {
	ctrl     *gomock.Controller
	recorder *MockTokenExtractorMockRecorder
}

// MockTokenExtractorMockRecorder is the mock recorder for MockTokenExtractor.
type MockTokenExtractorMockRecorder struct {
	mock *MockTokenExtractor
}

// NewMockTokenExtractor creates a new mock instance.
func NewMockTokenExtractor(ctrl *gomock.Controller) *MockTokenExtractor {
	mock := &MockTokenExtractor{ctrl: ctrl}
	mock.recorder = &MockTokenExtractorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTokenExtractor) EXPECT() *MockTokenExtractorMockRecorder {
	return m.recorder
}

// FromContext mocks base method.
func (m *MockTokenExtractor) FromContext(ctx context.Context) (jwt.Token, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FromContext", ctx)
	ret0, _ := ret[0].(jwt.Token)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FromContext indicates an expected call of FromContext.
func (mr *MockTokenExtractorMockRecorder) FromContext(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FromContext", reflect.TypeOf((*MockTokenExtractor)(nil).FromContext), ctx)
}

// FromString mocks base method.
func (m *MockTokenExtractor) FromString(str string) (jwt.Token, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FromString", str)
	ret0, _ := ret[0].(jwt.Token)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FromString indicates an expected call of FromString.
func (mr *MockTokenExtractorMockRecorder) FromString(str interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FromString", reflect.TypeOf((*MockTokenExtractor)(nil).FromString), str)
}

// MockToken is a mock of Token interface.
type MockToken struct {
	ctrl     *gomock.Controller
	recorder *MockTokenMockRecorder
}

// MockTokenMockRecorder is the mock recorder for MockToken.
type MockTokenMockRecorder struct {
	mock *MockToken
}

// NewMockToken creates a new mock instance.
func NewMockToken(ctrl *gomock.Controller) *MockToken {
	mock := &MockToken{ctrl: ctrl}
	mock.recorder = &MockTokenMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockToken) EXPECT() *MockTokenMockRecorder {
	return m.recorder
}

// Decode mocks base method.
func (m *MockToken) Decode(target interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Decode", target)
	ret0, _ := ret[0].(error)
	return ret0
}

// Decode indicates an expected call of Decode.
func (mr *MockTokenMockRecorder) Decode(target interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Decode", reflect.TypeOf((*MockToken)(nil).Decode), target)
}

// Map mocks base method.
func (m *MockToken) Map() (map[string]interface{}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Map")
	ret0, _ := ret[0].(map[string]interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Map indicates an expected call of Map.
func (mr *MockTokenMockRecorder) Map() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Map", reflect.TypeOf((*MockToken)(nil).Map))
}

// Payload mocks base method.
func (m *MockToken) Payload() []byte {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Payload")
	ret0, _ := ret[0].([]byte)
	return ret0
}

// Payload indicates an expected call of Payload.
func (mr *MockTokenMockRecorder) Payload() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Payload", reflect.TypeOf((*MockToken)(nil).Payload))
}

// Raw mocks base method.
func (m *MockToken) Raw() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Raw")
	ret0, _ := ret[0].(string)
	return ret0
}

// Raw indicates an expected call of Raw.
func (mr *MockTokenMockRecorder) Raw() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Raw", reflect.TypeOf((*MockToken)(nil).Raw))
}
