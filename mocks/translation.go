// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"

	entity "github.com/costaconrado/services-csm/internal/entity"
	mock "github.com/stretchr/testify/mock"
)

// Translation is an autogenerated mock type for the Translation type
type Translation struct {
	mock.Mock
}

// History provides a mock function with given fields: _a0
func (_m *Translation) History(_a0 context.Context) ([]entity.Translation, error) {
	ret := _m.Called(_a0)

	var r0 []entity.Translation
	if rf, ok := ret.Get(0).(func(context.Context) []entity.Translation); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entity.Translation)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Translate provides a mock function with given fields: _a0, _a1
func (_m *Translation) Translate(_a0 context.Context, _a1 entity.Translation) (entity.Translation, error) {
	ret := _m.Called(_a0, _a1)

	var r0 entity.Translation
	if rf, ok := ret.Get(0).(func(context.Context, entity.Translation) entity.Translation); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(entity.Translation)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, entity.Translation) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewTranslation interface {
	mock.TestingT
	Cleanup(func())
}

// NewTranslation creates a new instance of Translation. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewTranslation(t mockConstructorTestingTNewTranslation) *Translation {
	mock := &Translation{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
