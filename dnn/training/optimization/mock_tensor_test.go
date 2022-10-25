// Code generated by MockGen. DO NOT EDIT.
// Source: gitlab.com/akita/dnn/tensor (interfaces: Tensor,Operator)

package optimization

import (
	gomock "github.com/golang/mock/gomock"
	tensor "gitlab.com/akita/dnn/tensor"
	reflect "reflect"
)

// MockTensor is a mock of Tensor interface
type MockTensor struct {
	ctrl     *gomock.Controller
	recorder *MockTensorMockRecorder
}

// MockTensorMockRecorder is the mock recorder for MockTensor
type MockTensorMockRecorder struct {
	mock *MockTensor
}

// NewMockTensor creates a new mock instance
func NewMockTensor(ctrl *gomock.Controller) *MockTensor {
	mock := &MockTensor{ctrl: ctrl}
	mock.recorder = &MockTensorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockTensor) EXPECT() *MockTensorMockRecorder {
	return m.recorder
}

// Descriptor mocks base method
func (m *MockTensor) Descriptor() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Descriptor")
	ret0, _ := ret[0].(string)
	return ret0
}

// Descriptor indicates an expected call of Descriptor
func (mr *MockTensorMockRecorder) Descriptor() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Descriptor", reflect.TypeOf((*MockTensor)(nil).Descriptor))
}

// Dim mocks base method
func (m *MockTensor) Dim() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Dim")
	ret0, _ := ret[0].(int)
	return ret0
}

// Dim indicates an expected call of Dim
func (mr *MockTensorMockRecorder) Dim() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Dim", reflect.TypeOf((*MockTensor)(nil).Dim))
}

// NumElement mocks base method
func (m *MockTensor) NumElement() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NumElement")
	ret0, _ := ret[0].(int)
	return ret0
}

// NumElement indicates an expected call of NumElement
func (mr *MockTensorMockRecorder) NumElement() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NumElement", reflect.TypeOf((*MockTensor)(nil).NumElement))
}

// SetDescriptor mocks base method
func (m *MockTensor) SetDescriptor(arg0 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetDescriptor", arg0)
}

// SetDescriptor indicates an expected call of SetDescriptor
func (mr *MockTensorMockRecorder) SetDescriptor(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetDescriptor", reflect.TypeOf((*MockTensor)(nil).SetDescriptor), arg0)
}

// SetSize mocks base method
func (m *MockTensor) SetSize(arg0 []int) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetSize", arg0)
}

// SetSize indicates an expected call of SetSize
func (mr *MockTensorMockRecorder) SetSize(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetSize", reflect.TypeOf((*MockTensor)(nil).SetSize), arg0)
}

// Size mocks base method
func (m *MockTensor) Size() []int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Size")
	ret0, _ := ret[0].([]int)
	return ret0
}

// Size indicates an expected call of Size
func (mr *MockTensorMockRecorder) Size() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Size", reflect.TypeOf((*MockTensor)(nil).Size))
}

// Vector mocks base method
func (m *MockTensor) Vector() []float64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Vector")
	ret0, _ := ret[0].([]float64)
	return ret0
}

// Vector indicates an expected call of Vector
func (mr *MockTensorMockRecorder) Vector() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Vector", reflect.TypeOf((*MockTensor)(nil).Vector))
}

// MockOperator is a mock of Operator interface
type MockOperator struct {
	ctrl     *gomock.Controller
	recorder *MockOperatorMockRecorder
}

// MockOperatorMockRecorder is the mock recorder for MockOperator
type MockOperatorMockRecorder struct {
	mock *MockOperator
}

// NewMockOperator creates a new mock instance
func NewMockOperator(ctrl *gomock.Controller) *MockOperator {
	mock := &MockOperator{ctrl: ctrl}
	mock.recorder = &MockOperatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockOperator) EXPECT() *MockOperatorMockRecorder {
	return m.recorder
}

// Adam mocks base method
func (m *MockOperator) Adam(arg0, arg1, arg2, arg3 tensor.Tensor, arg4, arg5, arg6 float64) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Adam", arg0, arg1, arg2, arg3, arg4, arg5, arg6)
}

// Adam indicates an expected call of Adam
func (mr *MockOperatorMockRecorder) Adam(arg0, arg1, arg2, arg3, arg4, arg5, arg6 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Adam", reflect.TypeOf((*MockOperator)(nil).Adam), arg0, arg1, arg2, arg3, arg4, arg5, arg6)
}

// AvgPoolingBackward mocks base method
func (m *MockOperator) AvgPoolingBackward(arg0, arg1 tensor.Tensor, arg2, arg3, arg4 []int) tensor.Tensor {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AvgPoolingBackward", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(tensor.Tensor)
	return ret0
}

// AvgPoolingBackward indicates an expected call of AvgPoolingBackward
func (mr *MockOperatorMockRecorder) AvgPoolingBackward(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AvgPoolingBackward", reflect.TypeOf((*MockOperator)(nil).AvgPoolingBackward), arg0, arg1, arg2, arg3, arg4)
}

// AvgPoolingForward mocks base method
func (m *MockOperator) AvgPoolingForward(arg0 tensor.Tensor, arg1, arg2, arg3 []int) tensor.Tensor {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AvgPoolingForward", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(tensor.Tensor)
	return ret0
}

// AvgPoolingForward indicates an expected call of AvgPoolingForward
func (mr *MockOperatorMockRecorder) AvgPoolingForward(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AvgPoolingForward", reflect.TypeOf((*MockOperator)(nil).AvgPoolingForward), arg0, arg1, arg2, arg3)
}

// Clear mocks base method
func (m *MockOperator) Clear(arg0 tensor.Tensor) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Clear", arg0)
}

// Clear indicates an expected call of Clear
func (mr *MockOperatorMockRecorder) Clear(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Clear", reflect.TypeOf((*MockOperator)(nil).Clear), arg0)
}

// Clone mocks base method
func (m *MockOperator) Clone(arg0 tensor.Tensor) tensor.Tensor {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Clone", arg0)
	ret0, _ := ret[0].(tensor.Tensor)
	return ret0
}

// Clone indicates an expected call of Clone
func (mr *MockOperatorMockRecorder) Clone(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Clone", reflect.TypeOf((*MockOperator)(nil).Clone), arg0)
}

// Copy mocks base method
func (m *MockOperator) Copy(arg0, arg1 tensor.Tensor) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Copy", arg0, arg1)
}

// Copy indicates an expected call of Copy
func (mr *MockOperatorMockRecorder) Copy(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Copy", reflect.TypeOf((*MockOperator)(nil).Copy), arg0, arg1)
}

// Create mocks base method
func (m *MockOperator) Create(arg0 []int) tensor.Tensor {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0)
	ret0, _ := ret[0].(tensor.Tensor)
	return ret0
}

// Create indicates an expected call of Create
func (mr *MockOperatorMockRecorder) Create(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockOperator)(nil).Create), arg0)
}

// CreateWithData mocks base method
func (m *MockOperator) CreateWithData(arg0 []float64, arg1 []int, arg2 string) tensor.Tensor {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateWithData", arg0, arg1, arg2)
	ret0, _ := ret[0].(tensor.Tensor)
	return ret0
}

// CreateWithData indicates an expected call of CreateWithData
func (mr *MockOperatorMockRecorder) CreateWithData(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateWithData", reflect.TypeOf((*MockOperator)(nil).CreateWithData), arg0, arg1, arg2)
}

// CrossEntropy mocks base method
func (m *MockOperator) CrossEntropy(arg0 tensor.Tensor, arg1 []int) float64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CrossEntropy", arg0, arg1)
	ret0, _ := ret[0].(float64)
	return ret0
}

// CrossEntropy indicates an expected call of CrossEntropy
func (mr *MockOperatorMockRecorder) CrossEntropy(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CrossEntropy", reflect.TypeOf((*MockOperator)(nil).CrossEntropy), arg0, arg1)
}

// CrossEntropyDerivative mocks base method
func (m *MockOperator) CrossEntropyDerivative(arg0 tensor.Tensor, arg1 []int) tensor.Tensor {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CrossEntropyDerivative", arg0, arg1)
	ret0, _ := ret[0].(tensor.Tensor)
	return ret0
}

// CrossEntropyDerivative indicates an expected call of CrossEntropyDerivative
func (mr *MockOperatorMockRecorder) CrossEntropyDerivative(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CrossEntropyDerivative", reflect.TypeOf((*MockOperator)(nil).CrossEntropyDerivative), arg0, arg1)
}

// Dilate mocks base method
func (m *MockOperator) Dilate(arg0 tensor.Tensor, arg1 []int) tensor.Tensor {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Dilate", arg0, arg1)
	ret0, _ := ret[0].(tensor.Tensor)
	return ret0
}

// Dilate indicates an expected call of Dilate
func (mr *MockOperatorMockRecorder) Dilate(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Dilate", reflect.TypeOf((*MockOperator)(nil).Dilate), arg0, arg1)
}

// Dump mocks base method
func (m *MockOperator) Dump(arg0 tensor.Tensor) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Dump", arg0)
	ret0, _ := ret[0].(string)
	return ret0
}

// Dump indicates an expected call of Dump
func (mr *MockOperatorMockRecorder) Dump(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Dump", reflect.TypeOf((*MockOperator)(nil).Dump), arg0)
}

// ElementWiseMul mocks base method
func (m *MockOperator) ElementWiseMul(arg0, arg1 tensor.Tensor) tensor.Tensor {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ElementWiseMul", arg0, arg1)
	ret0, _ := ret[0].(tensor.Tensor)
	return ret0
}

// ElementWiseMul indicates an expected call of ElementWiseMul
func (mr *MockOperatorMockRecorder) ElementWiseMul(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ElementWiseMul", reflect.TypeOf((*MockOperator)(nil).ElementWiseMul), arg0, arg1)
}

// Free mocks base method
func (m *MockOperator) Free(arg0 tensor.Tensor) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Free", arg0)
}

// Free indicates an expected call of Free
func (mr *MockOperatorMockRecorder) Free(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Free", reflect.TypeOf((*MockOperator)(nil).Free), arg0)
}

// Gemm mocks base method
func (m *MockOperator) Gemm(arg0, arg1 bool, arg2, arg3 float64, arg4, arg5, arg6 tensor.Tensor) tensor.Tensor {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Gemm", arg0, arg1, arg2, arg3, arg4, arg5, arg6)
	ret0, _ := ret[0].(tensor.Tensor)
	return ret0
}

// Gemm indicates an expected call of Gemm
func (mr *MockOperatorMockRecorder) Gemm(arg0, arg1, arg2, arg3, arg4, arg5, arg6 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Gemm", reflect.TypeOf((*MockOperator)(nil).Gemm), arg0, arg1, arg2, arg3, arg4, arg5, arg6)
}

// Im2Col mocks base method
func (m *MockOperator) Im2Col(arg0 tensor.Tensor, arg1, arg2, arg3, arg4 []int) tensor.Tensor {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Im2Col", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(tensor.Tensor)
	return ret0
}

// Im2Col indicates an expected call of Im2Col
func (mr *MockOperatorMockRecorder) Im2Col(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Im2Col", reflect.TypeOf((*MockOperator)(nil).Im2Col), arg0, arg1, arg2, arg3, arg4)
}

// Init mocks base method
func (m *MockOperator) Init(arg0 tensor.Tensor, arg1 []float64) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Init", arg0, arg1)
}

// Init indicates an expected call of Init
func (mr *MockOperatorMockRecorder) Init(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Init", reflect.TypeOf((*MockOperator)(nil).Init), arg0, arg1)
}

// MaxPoolingBackward mocks base method
func (m *MockOperator) MaxPoolingBackward(arg0, arg1, arg2 tensor.Tensor, arg3, arg4, arg5 []int) tensor.Tensor {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MaxPoolingBackward", arg0, arg1, arg2, arg3, arg4, arg5)
	ret0, _ := ret[0].(tensor.Tensor)
	return ret0
}

// MaxPoolingBackward indicates an expected call of MaxPoolingBackward
func (mr *MockOperatorMockRecorder) MaxPoolingBackward(arg0, arg1, arg2, arg3, arg4, arg5 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MaxPoolingBackward", reflect.TypeOf((*MockOperator)(nil).MaxPoolingBackward), arg0, arg1, arg2, arg3, arg4, arg5)
}

// MaxPoolingForward mocks base method
func (m *MockOperator) MaxPoolingForward(arg0 tensor.Tensor, arg1, arg2, arg3 []int) (tensor.Tensor, tensor.Tensor) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MaxPoolingForward", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(tensor.Tensor)
	ret1, _ := ret[1].(tensor.Tensor)
	return ret0, ret1
}

// MaxPoolingForward indicates an expected call of MaxPoolingForward
func (mr *MockOperatorMockRecorder) MaxPoolingForward(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MaxPoolingForward", reflect.TypeOf((*MockOperator)(nil).MaxPoolingForward), arg0, arg1, arg2, arg3)
}

// RMSProp mocks base method
func (m *MockOperator) RMSProp(arg0, arg1, arg2 tensor.Tensor, arg3, arg4 float64) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RMSProp", arg0, arg1, arg2, arg3, arg4)
}

// RMSProp indicates an expected call of RMSProp
func (mr *MockOperatorMockRecorder) RMSProp(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RMSProp", reflect.TypeOf((*MockOperator)(nil).RMSProp), arg0, arg1, arg2, arg3, arg4)
}

// ReluBackward mocks base method
func (m *MockOperator) ReluBackward(arg0, arg1 tensor.Tensor) tensor.Tensor {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReluBackward", arg0, arg1)
	ret0, _ := ret[0].(tensor.Tensor)
	return ret0
}

// ReluBackward indicates an expected call of ReluBackward
func (mr *MockOperatorMockRecorder) ReluBackward(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReluBackward", reflect.TypeOf((*MockOperator)(nil).ReluBackward), arg0, arg1)
}

// ReluForward mocks base method
func (m *MockOperator) ReluForward(arg0 tensor.Tensor) tensor.Tensor {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReluForward", arg0)
	ret0, _ := ret[0].(tensor.Tensor)
	return ret0
}

// ReluForward indicates an expected call of ReluForward
func (mr *MockOperatorMockRecorder) ReluForward(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReluForward", reflect.TypeOf((*MockOperator)(nil).ReluForward), arg0)
}

// Repeat mocks base method
func (m *MockOperator) Repeat(arg0 tensor.Tensor, arg1 int) tensor.Tensor {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Repeat", arg0, arg1)
	ret0, _ := ret[0].(tensor.Tensor)
	return ret0
}

// Repeat indicates an expected call of Repeat
func (mr *MockOperatorMockRecorder) Repeat(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Repeat", reflect.TypeOf((*MockOperator)(nil).Repeat), arg0, arg1)
}

// Reshape mocks base method
func (m *MockOperator) Reshape(arg0 tensor.Tensor, arg1 []int) tensor.Tensor {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Reshape", arg0, arg1)
	ret0, _ := ret[0].(tensor.Tensor)
	return ret0
}

// Reshape indicates an expected call of Reshape
func (mr *MockOperatorMockRecorder) Reshape(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Reshape", reflect.TypeOf((*MockOperator)(nil).Reshape), arg0, arg1)
}

// Rotate180 mocks base method
func (m *MockOperator) Rotate180(arg0 tensor.Tensor) tensor.Tensor {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Rotate180", arg0)
	ret0, _ := ret[0].(tensor.Tensor)
	return ret0
}

// Rotate180 indicates an expected call of Rotate180
func (mr *MockOperatorMockRecorder) Rotate180(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Rotate180", reflect.TypeOf((*MockOperator)(nil).Rotate180), arg0)
}

// ScaleAdd mocks base method
func (m *MockOperator) ScaleAdd(arg0, arg1 float64, arg2, arg3 tensor.Tensor) tensor.Tensor {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ScaleAdd", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(tensor.Tensor)
	return ret0
}

// ScaleAdd indicates an expected call of ScaleAdd
func (mr *MockOperatorMockRecorder) ScaleAdd(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ScaleAdd", reflect.TypeOf((*MockOperator)(nil).ScaleAdd), arg0, arg1, arg2, arg3)
}

// Slice mocks base method
func (m *MockOperator) Slice(arg0 tensor.Tensor, arg1, arg2 int) tensor.Tensor {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Slice", arg0, arg1, arg2)
	ret0, _ := ret[0].(tensor.Tensor)
	return ret0
}

// Slice indicates an expected call of Slice
func (mr *MockOperatorMockRecorder) Slice(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Slice", reflect.TypeOf((*MockOperator)(nil).Slice), arg0, arg1, arg2)
}

// Softmax mocks base method
func (m *MockOperator) Softmax(arg0 tensor.Tensor) tensor.Tensor {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Softmax", arg0)
	ret0, _ := ret[0].(tensor.Tensor)
	return ret0
}

// Softmax indicates an expected call of Softmax
func (mr *MockOperatorMockRecorder) Softmax(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Softmax", reflect.TypeOf((*MockOperator)(nil).Softmax), arg0)
}

// SoftmaxCrossEntropyDerivative mocks base method
func (m *MockOperator) SoftmaxCrossEntropyDerivative(arg0 tensor.Tensor, arg1 []int) tensor.Tensor {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SoftmaxCrossEntropyDerivative", arg0, arg1)
	ret0, _ := ret[0].(tensor.Tensor)
	return ret0
}

// SoftmaxCrossEntropyDerivative indicates an expected call of SoftmaxCrossEntropyDerivative
func (mr *MockOperatorMockRecorder) SoftmaxCrossEntropyDerivative(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SoftmaxCrossEntropyDerivative", reflect.TypeOf((*MockOperator)(nil).SoftmaxCrossEntropyDerivative), arg0, arg1)
}

// Sum mocks base method
func (m *MockOperator) Sum(arg0 tensor.Tensor, arg1 []int) tensor.Tensor {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Sum", arg0, arg1)
	ret0, _ := ret[0].(tensor.Tensor)
	return ret0
}

// Sum indicates an expected call of Sum
func (mr *MockOperatorMockRecorder) Sum(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Sum", reflect.TypeOf((*MockOperator)(nil).Sum), arg0, arg1)
}

// Transpose mocks base method
func (m *MockOperator) Transpose(arg0 tensor.Tensor, arg1 []int) tensor.Tensor {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Transpose", arg0, arg1)
	ret0, _ := ret[0].(tensor.Tensor)
	return ret0
}

// Transpose indicates an expected call of Transpose
func (mr *MockOperatorMockRecorder) Transpose(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Transpose", reflect.TypeOf((*MockOperator)(nil).Transpose), arg0, arg1)
}

// Zeros mocks base method
func (m *MockOperator) Zeros(arg0 []int) tensor.Tensor {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Zeros", arg0)
	ret0, _ := ret[0].(tensor.Tensor)
	return ret0
}

// Zeros indicates an expected call of Zeros
func (mr *MockOperatorMockRecorder) Zeros(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Zeros", reflect.TypeOf((*MockOperator)(nil).Zeros), arg0)
}
