// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import context "context"
import mock "github.com/stretchr/testify/mock"
import models "github.com/felipe_rodrigues/poll-api/pkg/domain/models"

// PollService is an autogenerated mock type for the PollService type
type PollService struct {
	mock.Mock
}

// Count provides a mock function with given fields: ctx, id
func (_m *PollService) Count(ctx context.Context, id int64) (*models.PollResult, error) {
	ret := _m.Called(ctx, id)

	var r0 *models.PollResult
	if rf, ok := ret.Get(0).(func(context.Context, int64) *models.PollResult); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.PollResult)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CountByHour provides a mock function with given fields: ctx, id
func (_m *PollService) CountByHour(ctx context.Context, id int64) ([]models.PollResultByHour, error) {
	ret := _m.Called(ctx, id)

	var r0 []models.PollResultByHour
	if rf, ok := ret.Get(0).(func(context.Context, int64) []models.PollResultByHour); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.PollResultByHour)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CountByNominate provides a mock function with given fields: ctx, id
func (_m *PollService) CountByNominate(ctx context.Context, id int64) (*models.PollResultByNominates, error) {
	ret := _m.Called(ctx, id)

	var r0 *models.PollResultByNominates
	if rf, ok := ret.Get(0).(func(context.Context, int64) *models.PollResultByNominates); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.PollResultByNominates)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindById provides a mock function with given fields: ctx, id
func (_m *PollService) FindById(ctx context.Context, id int64) (*models.Poll, error) {
	ret := _m.Called(ctx, id)

	var r0 *models.Poll
	if rf, ok := ret.Get(0).(func(context.Context, int64) *models.Poll); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Poll)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IsPollClosed provides a mock function with given fields: ctx, poll
func (_m *PollService) IsPollClosed(ctx context.Context, poll models.Poll) bool {
	ret := _m.Called(ctx, poll)

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context, models.Poll) bool); ok {
		r0 = rf(ctx, poll)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// Create provides a mock function with given fields: ctx, poll
func (_m *PollService) Create(ctx context.Context, poll models.Poll) (*models.Poll, error) {
	ret := _m.Called(ctx, poll)

	var r0 *models.Poll
	if rf, ok := ret.Get(0).(func(context.Context, models.Poll) *models.Poll); ok {
		r0 = rf(ctx, poll)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Poll)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, models.Poll) error); ok {
		r1 = rf(ctx, poll)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IsValidPoll provides a mock function with given fields: ctx, poll
func (_m *PollService) IsValidPoll(ctx context.Context, poll models.Poll) bool {
	ret := _m.Called(ctx, poll)

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context, models.Poll) bool); ok {
		r0 = rf(ctx, poll)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}
