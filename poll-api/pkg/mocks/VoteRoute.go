// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// VoteRoute is an autogenerated mock type for the VoteRoute type
type VoteRoute struct {
	mock.Mock
}

// PublishVote provides a mock function with given fields: body, queueName
func (_m *VoteRoute) PublishVote(body []byte, queueName string) (bool, error) {
	ret := _m.Called(body, queueName)

	var r0 bool
	if rf, ok := ret.Get(0).(func([]byte, string) bool); ok {
		r0 = rf(body, queueName)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]byte, string) error); ok {
		r1 = rf(body, queueName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
