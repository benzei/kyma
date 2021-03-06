// Code generated by mockery v1.0.0
package mocks

import mock "github.com/stretchr/testify/mock"

// Generator is an autogenerated mock type for the Generator type
type Generator struct {
	mock.Mock
}

// NewUUID provides a mock function with given fields:
func (_m *Generator) NewUUID() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}
