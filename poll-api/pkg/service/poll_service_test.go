package service

import (
	"context"
	"errors"
	"github.com/felipe_rodrigues/poll-api/pkg/domain/models"
	"github.com/felipe_rodrigues/poll-api/pkg/mocks"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)


var (
	pollServiceUnderTest PollService
	mockPollRepository *mocks.PollRepository
	mockPollCacheRepository *mocks.PollCacheRepository
)

func ConfigPollService(){
	mockPollRepository = &mocks.PollRepository{}
	mockPollCacheRepository = &mocks.PollCacheRepository{}

	pollServiceUnderTest = &pollService{
		PollRepository:          mockPollRepository,
		PollCacheRepository:     mockPollCacheRepository,
	}
}

func TestGivenIdThenReturnPoll(t *testing.T){
	ConfigPollService()
	var nominates []models.Nominate

	nominate1 := models.Nominate{Id:1, Name:"Teste1", Picture:"ddasd"}
	nominate2 := models.Nominate{Id:2, Name:"Teste2", Picture:"ddasd"}

	nominates = append(nominates, nominate1)
	nominates = append(nominates, nominate2)

	poll := &models.Poll{
		Id: 156,
		Nominates: nominates,
		Year: 2019,
		Duration: 10,
	}

	mockPollRepository.On("FindById", context.Background(), int64(156)).Return(poll, nil)
	mockPollCacheRepository.On("Get", int64(156)).Return(poll, nil)

	p, _ := pollServiceUnderTest.FindById(context.Background(), int64(156))
	assert.Equal(t, 156, p.Id)
	assert.Equal(t, len(p.Nominates), 2)

}

func TestGivenIdAndPollIsOutOfPeriodThenNotReturn(t *testing.T){
	ConfigPollService()
	var nominates []models.Nominate

	nominate1 := models.Nominate{Id:1, Name:"Teste1", Picture:"ddasd"}
	nominate2 := models.Nominate{Id:2, Name:"Teste2", Picture:"ddasd"}

	nominates = append(nominates, nominate1)
	nominates = append(nominates, nominate2)

	poll := &models.Poll{
		Id: 156,
		Nominates: nominates,
		Year: 2019,
		Duration: 0,
		StartAt:time.Now(),
	}

	mockPollRepository.On("FindById", context.Background(), int64(156)).Return(poll, nil)
	mockPollCacheRepository.On("Get", int64(156)).Return(poll, nil)

	result := pollServiceUnderTest.IsPollClosed(context.Background(), *poll)
	assert.True(t, result)
}


func TestGivenIdThenNotExistPollThenReturnNil(t *testing.T){
	ConfigPollService()
	mockPollRepository.On("FindById", context.Background(), int64(156)).Return(nil, errors.New("not found"))
	mockPollCacheRepository.On("Get", int64(156)).Return(nil, errors.New("not found"))

	p, err := pollServiceUnderTest.FindById(context.Background(), int64(156))
	assert.Nil(t, p)
	assert.Error(t, err)
}

func TestGivenPollThenAlreadyExistThenReturnConflict(t *testing.T){
	ConfigPollService()

	var nominates []models.Nominate

	nominate1 := models.Nominate{Id:1, Name:"Teste1", Picture:"ddasd"}
	nominate2 := models.Nominate{Id:2, Name:"Teste2", Picture:"ddasd"}

	nominates = append(nominates, nominate1)
	nominates = append(nominates, nominate2)

	poll := &models.Poll{
		Id: 156,
		Nominates: nominates,
		Year: 2019,
		Duration: 0,
		StartAt:time.Now(),
	}

	mockPollRepository.On("FindById", context.Background(), int64(156)).Return(poll, nil)
	mockPollCacheRepository.On("Get", int64(156)).Return(poll, nil)
	p, err := pollServiceUnderTest.Create(context.Background(), *poll)

	assert.Nil(t, p)
	assert.Equal(t,"conflict", err.Error())
}

func TestGivenPollThenCreateIt(t *testing.T){
	ConfigPollService()

	var nominates []models.Nominate

	nominate1 := models.Nominate{Id:1, Name:"Teste1", Picture:"ddasd"}
	nominate2 := models.Nominate{Id:2, Name:"Teste2", Picture:"ddasd"}

	nominates = append(nominates, nominate1)
	nominates = append(nominates, nominate2)

	poll := &models.Poll{
		Id: 156,
		Nominates: nominates,
		Year: 2019,
		Duration: 3,
		StartAt:time.Now(),
	}

	mockPollRepository.On("FindById", context.Background(), int64(156)).Return(nil, nil)
	mockPollRepository.On("Create", context.Background(), *poll).Return(poll, nil)
	mockPollCacheRepository.On("Get", int64(156)).Return(nil, nil)
	p, err := pollServiceUnderTest.Create(context.Background(), *poll)

	assert.Nil(t, err)
	assert.NotNil(t, p)
	assert.Equal(t,156, p.Id)
}