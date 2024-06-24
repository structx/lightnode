// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// StateMachine is an autogenerated mock type for the StateMachine type
type StateMachine struct {
	mock.Mock
}

type StateMachine_Expecter struct {
	mock *mock.Mock
}

func (_m *StateMachine) EXPECT() *StateMachine_Expecter {
	return &StateMachine_Expecter{mock: &_m.Mock}
}

// Close provides a mock function with given fields:
func (_m *StateMachine) Close() error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Close")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// StateMachine_Close_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Close'
type StateMachine_Close_Call struct {
	*mock.Call
}

// Close is a helper method to define mock.On call
func (_e *StateMachine_Expecter) Close() *StateMachine_Close_Call {
	return &StateMachine_Close_Call{Call: _e.mock.On("Close")}
}

func (_c *StateMachine_Close_Call) Run(run func()) *StateMachine_Close_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *StateMachine_Close_Call) Return(_a0 error) *StateMachine_Close_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *StateMachine_Close_Call) RunAndReturn(run func() error) *StateMachine_Close_Call {
	_c.Call.Return(run)
	return _c
}

// Get provides a mock function with given fields: key
func (_m *StateMachine) Get(key string) ([]byte, error) {
	ret := _m.Called(key)

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 []byte
	var r1 error
	if rf, ok := ret.Get(0).(func(string) ([]byte, error)); ok {
		return rf(key)
	}
	if rf, ok := ret.Get(0).(func(string) []byte); ok {
		r0 = rf(key)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(key)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// StateMachine_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type StateMachine_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//   - key string
func (_e *StateMachine_Expecter) Get(key interface{}) *StateMachine_Get_Call {
	return &StateMachine_Get_Call{Call: _e.mock.On("Get", key)}
}

func (_c *StateMachine_Get_Call) Run(run func(key string)) *StateMachine_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *StateMachine_Get_Call) Return(_a0 []byte, _a1 error) *StateMachine_Get_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *StateMachine_Get_Call) RunAndReturn(run func(string) ([]byte, error)) *StateMachine_Get_Call {
	_c.Call.Return(run)
	return _c
}

// Put provides a mock function with given fields: key, value
func (_m *StateMachine) Put(key string, value []byte) error {
	ret := _m.Called(key, value)

	if len(ret) == 0 {
		panic("no return value specified for Put")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string, []byte) error); ok {
		r0 = rf(key, value)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// StateMachine_Put_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Put'
type StateMachine_Put_Call struct {
	*mock.Call
}

// Put is a helper method to define mock.On call
//   - key string
//   - value []byte
func (_e *StateMachine_Expecter) Put(key interface{}, value interface{}) *StateMachine_Put_Call {
	return &StateMachine_Put_Call{Call: _e.mock.On("Put", key, value)}
}

func (_c *StateMachine_Put_Call) Run(run func(key string, value []byte)) *StateMachine_Put_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].([]byte))
	})
	return _c
}

func (_c *StateMachine_Put_Call) Return(_a0 error) *StateMachine_Put_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *StateMachine_Put_Call) RunAndReturn(run func(string, []byte) error) *StateMachine_Put_Call {
	_c.Call.Return(run)
	return _c
}

// NewStateMachine creates a new instance of StateMachine. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewStateMachine(t interface {
	mock.TestingT
	Cleanup(func())
}) *StateMachine {
	mock := &StateMachine{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
