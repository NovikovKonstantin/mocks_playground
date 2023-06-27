// Code generated by mockery v2.30.15. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// Get provides a mock function with given fields: keys
func (_m *Repository) Get(keys []string) ([]int64, error) {
	ret := _m.Called(keys)

	var r0 []int64
	var r1 error
	if rf, ok := ret.Get(0).(func([]string) ([]int64, error)); ok {
		return rf(keys)
	}
	if rf, ok := ret.Get(0).(func([]string) []int64); ok {
		r0 = rf(keys)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]int64)
		}
	}

	if rf, ok := ret.Get(1).(func([]string) error); ok {
		r1 = rf(keys)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Store provides a mock function with given fields: values
func (_m *Repository) Store(values []int64) ([]string, error) {
	ret := _m.Called(values)

	var r0 []string
	var r1 error
	if rf, ok := ret.Get(0).(func([]int64) ([]string, error)); ok {
		return rf(values)
	}
	if rf, ok := ret.Get(0).(func([]int64) []string); ok {
		r0 = rf(values)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	if rf, ok := ret.Get(1).(func([]int64) error); ok {
		r1 = rf(values)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewRepository creates a new instance of Repository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *Repository {
	mock := &Repository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
