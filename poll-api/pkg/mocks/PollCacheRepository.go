// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"
import models "github.com/felipe_rodrigues/poll-api/pkg/domain/models"

// PollCacheRepository is an autogenerated mock type for the PollCacheRepository type
type PollCacheRepository struct {
	mock.Mock
}

// Get provides a mock function with given fields: k
func (_m *PollCacheRepository) Get(k int64) (*models.Poll, error) {
	ret := _m.Called(k)

	var r0 *models.Poll
	if rf, ok := ret.Get(0).(func(int64) *models.Poll); ok {
		r0 = rf(k)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Poll)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int64) error); ok {
		r1 = rf(k)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Set provides a mock function with given fields: k, poll
func (_m *PollCacheRepository) Set(k int64, poll *models.Poll) {
	_m.Called(k, poll)
}
