// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	entity "github.com/costaconrado/services-csm/internal/entity"
	mock "github.com/stretchr/testify/mock"
)

// ProposalWebAPI is an autogenerated mock type for the ProposalWebAPI type
type ProposalWebAPI struct {
	mock.Mock
}

// Translate provides a mock function with given fields: _a0
func (_m *ProposalWebAPI) Translate(_a0 entity.Proposal) (entity.Proposal, error) {
	ret := _m.Called(_a0)

	var r0 entity.Proposal
	if rf, ok := ret.Get(0).(func(entity.Proposal) entity.Proposal); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(entity.Proposal)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(entity.Proposal) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewProposalWebAPI interface {
	mock.TestingT
	Cleanup(func())
}

// NewProposalWebAPI creates a new instance of ProposalWebAPI. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewProposalWebAPI(t mockConstructorTestingTNewProposalWebAPI) *ProposalWebAPI {
	mock := &ProposalWebAPI{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
