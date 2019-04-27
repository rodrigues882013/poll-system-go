package service

import (
	"context"
	"encoding/json"
	"github.com/felipe_rodrigues/poll-api/pkg/domain/models"
	"github.com/felipe_rodrigues/poll-api/pkg/mocks"
	"github.com/stretchr/testify/assert"
	test "github.com/vektra/mockery/mockery/fixtures"
	"net/http"
	"testing"
	"time"
)

var (
	service VoteService
	mockVoteRepository *mocks.VoteRepository
	mockVoteRoute *mocks.VoteRoute
	mockPollService *mocks.PollService
)

func ConfigVoteService(){
	mockVoteRepository = &mocks.VoteRepository{}
	mockVoteRoute = &mocks.VoteRoute{}
	mockPollService = &mocks.PollService{}

	service = &voteService{
		Repository:      mockVoteRepository,
		Route:           mockVoteRoute,
		PollService:     mockPollService,
	}
}

func TestGivenPollIdAndVoteValidThenQueued(t *testing.T){
	ConfigVoteService()
	nominate := models.Nominate{Id:2, Name:"Teste2", Picture:"ddasd"}

	poll := models.Poll{
		Id: 156,
	}

	vote := models.Vote{Name: "Teste", Email:"dasd@gmail.com", Nominate:nominate, Poll:poll}
	b, _ := json.Marshal(vote)

	mockVoteRepository.On("FindById", context.Background(), int64(156)).Return(poll, nil)
	mockVoteRoute.On("PublishVote", []byte(b), "queueTest").Return(true, nil)

	c := service.Vote(vote, "queueTest")

	if <- c {
		mockVoteRoute.AssertNumberOfCalls(t, "PublishVote", 1)
	}
}

func TestGivenValidVoteThenProceed(t *testing.T){
	ConfigVoteService()
	var ctx = context.Background()
	var nominates []models.Nominate

	nominate1 := models.Nominate{Id:1, Name:"Teste1", Picture:"ddasd"}
	nominate2 := models.Nominate{Id:2, Name:"Teste2", Picture:"ddasd"}

	nominates = append(nominates, nominate1)
	nominates = append(nominates, nominate2)


	poll := &models.Poll{
		Id: 156,
		Nominates:nominates,
		Duration:4,
		StartAt: time.Now(),
		Year: 2019,
	}

	vote := models.Vote{Name: "Teste", Email:"dasd@gmail.com", Nominate:nominate1, Poll:*poll}


	mockPollService.On("FindById", ctx, int64(156)).Return(poll, nil)


	params := make(map[string]string)
	params["pollId"] = "156"

	canVote, response, status := service.CanVote(ctx, params, vote)

	assert.True(t, canVote)
	assert.Equal(t, status, http.StatusAccepted)
	assert.Equal(t, response.Message, "Your vote was registered with successful.")

}

func TestGivenPollThatNotExistThenNotProceed(t *testing.T){
	ConfigVoteService()
	var ctx = context.Background()
	var nominates []models.Nominate
	params := make(map[string]string)
	params["pollId"] = "156"

	nominate1 := models.Nominate{Id:1, Name:"Teste1", Picture:"ddasd"}
	nominate2 := models.Nominate{Id:2, Name:"Teste2", Picture:"ddasd"}
	nominates = append(nominates, nominate1)
	nominates = append(nominates, nominate2)

	poll := &models.Poll{
		Id: 155,
		Nominates:nominates,
		Duration:4,
		StartAt: time.Now(),
		Year: 2019,
	}

	params["pollId"] = "156"

	vote := models.Vote{Name: "Teste", Email:"dasd@gmail.com", Nominate:nominate1, Poll:*poll}

	// If poll doesn't exist
	poll = &models.Poll{
		Id: 156,
		Nominates:nominates,
		Duration:4,
		StartAt: time.Now(),
		Year: 2019,
	}
	vote = models.Vote{Name: "Teste", Email:"dasd@gmail.com", Nominate:nominate1, Poll:*poll}
	mockPollService.On("FindById", ctx, int64(156)).Return(nil, new(test.Err))
	params["pollId"] = "156"


	canVote, response, status := service.CanVote(ctx, params, vote)
	assert.False(t, canVote)
	assert.Equal(t, status, http.StatusNotFound)
	assert.Equal(t, response.Message, "This poll doesn't exist.")
}

func TestGivenDifferentIdsThenNotProceed(t *testing.T) {
	ConfigVoteService()
	var ctx = context.Background()
	var nominates []models.Nominate
	params := make(map[string]string)
	params["pollId"] = "156"

	nominate1 := models.Nominate{Id:1, Name:"Teste1", Picture:"ddasd"}
	nominate2 := models.Nominate{Id:2, Name:"Teste2", Picture:"ddasd"}
	nominates = append(nominates, nominate1)
	nominates = append(nominates, nominate2)

	poll := &models.Poll{
		Id: 155,
		Nominates:nominates,
		Duration:4,
		StartAt: time.Now(),
		Year: 2019,
	}

	params["pollId"] = "156"

	vote := models.Vote{Name: "Teste", Email:"dasd@gmail.com", Nominate:nominate1, Poll:*poll}
	canVote, response, status := service.CanVote(ctx, params, vote)
	assert.False(t, canVote)
	assert.Equal(t, status, http.StatusPreconditionFailed)
	assert.Equal(t, response.Message, "This poll doesn't exist.")
}

func TestGivenInvalidIdThenNotProceed(t *testing.T){
	ConfigVoteService()
	var ctx = context.Background()
	var nominates []models.Nominate

	nominate1 := models.Nominate{Id:1, Name:"Teste1", Picture:"ddasd"}
	nominate2 := models.Nominate{Id:2, Name:"Teste2", Picture:"ddasd"}
	nominates = append(nominates, nominate1)
	nominates = append(nominates, nominate2)

	//::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
	// Sem passar o id como argumento
	poll := &models.Poll{
		Id: 155,
		Nominates:nominates,
		Duration:4,
		StartAt: time.Now(),
		Year: 2019,
	}

	vote := models.Vote{Name: "Teste", Email:"dasd@gmail.com", Nominate:nominate1, Poll:*poll}
	params := make(map[string]string)

	canVote, response, status := service.CanVote(ctx, params, vote)
	assert.False(t, canVote)
	assert.Equal(t, status, http.StatusBadRequest)
	assert.Equal(t, response.Message, "This pollId was not given.")
}

func TestGivenValidPollButWithoutChoiceAnyNominateThenNotProceed(t *testing.T){
	ConfigVoteService()
	var ctx = context.Background()
	var nominates []models.Nominate

	nominate1 := models.Nominate{Id:1, Name:"Teste1", Picture:"ddasd"}
	nominate2 := models.Nominate{Id:2, Name:"Teste2", Picture:"ddasd"}
	nominates = append(nominates, nominate1)
	nominates = append(nominates, nominate2)

	poll := &models.Poll{
		Id: 156,
		Duration:4,
		StartAt: time.Now(),
		Year: 2019,
	}
	vote := models.Vote{Name: "Teste", Email:"dasd@gmail.com", Nominate:nominate1, Poll:*poll}

	mockPollService.On("FindById", ctx, int64(156)).Return(poll, nil)
	params := make(map[string]string)
	params["pollId"] = "156"

	canVote, response, status := service.CanVote(ctx, params, vote)
	assert.False(t, canVote)
	assert.Equal(t,  http.StatusPreconditionFailed, status)
	assert.Equal(t,"You do not chosen any nominate.", response.Message)
}

func TestGivenValidPollButWithChoiceOutThenNotProceed(t *testing.T){
	ConfigVoteService()
	var ctx = context.Background()
	var nominates []models.Nominate

	nominate1 := models.Nominate{Id:1, Name:"Teste1", Picture:"ddasd"}
	nominate2 := models.Nominate{Id:2, Name:"Teste2", Picture:"ddasd"}
	nominate3 := models.Nominate{Id:3, Name:"Teste2", Picture:"ddasd"}

	nominates = append(nominates, nominate1)
	nominates = append(nominates, nominate2)

	poll := &models.Poll{
		Id: 156,
		Duration:4,
		StartAt: time.Now(),
		Year: 2019,
		Nominates:nominates,
	}
	vote := models.Vote{Name: "Teste", Email:"dasd@gmail.com", Nominate:nominate3, Poll:*poll}

	mockPollService.On("FindById", ctx, int64(156)).Return(poll, nil)
	params := make(map[string]string)
	params["pollId"] = "156"

	canVote, response, status := service.CanVote(ctx, params, vote)
	assert.False(t, canVote)
	assert.Equal(t,  http.StatusPreconditionFailed, status)
	assert.Equal(t,"The nominate there isn't belong this poll.", response.Message)
}

func TestGivenValidPollButOutOfPeriodThenNotProceed(t *testing.T){
	ConfigVoteService()
	var ctx = context.Background()
	var nominates []models.Nominate

	nominate1 := models.Nominate{Id:1, Name:"Teste1", Picture:"ddasd"}
	nominate2 := models.Nominate{Id:2, Name:"Teste2", Picture:"ddasd"}

	nominates = append(nominates, nominate1)
	nominates = append(nominates, nominate2)

	poll := &models.Poll{
		Id: 156,
		Duration:0,
		StartAt: time.Now(),
		Year: 2019,
		Nominates:nominates,
	}
	vote := models.Vote{Name: "Teste", Email:"dasd@gmail.com", Nominate:nominate2, Poll:*poll}

	mockPollService.On("FindById", ctx, int64(156)).Return(poll, nil)
	params := make(map[string]string)
	params["pollId"] = "156"

	canVote, response, status := service.CanVote(ctx, params, vote)
	assert.False(t, canVote)
	assert.Equal(t,  http.StatusForbidden, status)
	assert.Equal(t,"The poll is closed.", response.Message)
}


