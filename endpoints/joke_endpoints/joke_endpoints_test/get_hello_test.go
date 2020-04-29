package joke_endpoints_test

import (
	"context"
	"errors"
	"github.com/bmsandoval/apple-coding-challenge/api/transport_http/codecs/joke_codecs"
	"github.com/bmsandoval/apple-coding-challenge/endpoints/joke_endpoints"
	"github.com/bmsandoval/apple-coding-challenge/library/appcontext"
	"github.com/bmsandoval/apple-coding-challenge/mocks/services_mocks"
	"github.com/bmsandoval/apple-coding-challenge/services"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetHelloServer(t *testing.T) {
	var GetHelloUnitTestData = []GetJokeTestData{
		{
			BaseTestData: BaseTestData{
				Description: "happy path",
				Request:     joke_codecs.GetJokeRequest{},
				Response:    joke_codecs.GetJokeResponse("sample"),
			},
			MockGetJoke: &MockGetJoke{
				OutReply: stringPointer("John Doe"),
				OutError:     nil,
			},
			MockGetName: &MockGetName{
				OutReply: stringPointer("sample"),
				OutError: nil,
			},
			ResponseError: nil,
		},
		{
			BaseTestData: BaseTestData{
				Description: "assigned test condition",
				Request:     joke_codecs.GetJokeRequest{},
				Response:    joke_codecs.GetJokeResponse("Δαμέας's OSI network model has only one layer - Physical."),
			},
			MockGetJoke: &MockGetJoke{
				OutReply: stringPointer("John Doe's OSI network model has only one layer - Physical."),
				OutError:     nil,
			},
			MockGetName: &MockGetName{
				OutReply: stringPointer("Δαμέας"),
				OutError: nil,
			},
			ResponseError: nil,
		},
		{
			BaseTestData: BaseTestData{
				Description: "joke api fails",
				Request:     joke_codecs.GetJokeRequest{},
				Response:    nil,
			},
			MockGetJoke: &MockGetJoke{
				OutReply: stringPointer("John Doe's OSI network model has only one layer - Physical."),
				OutError:     errors.New("joke api failure"),
			},
			MockGetName: nil,
			ResponseError: errors.New("joke api failure"),
		},
		{
			BaseTestData: BaseTestData{
				Description: "person api fails",
				Request:     joke_codecs.GetJokeRequest{},
				Response:    nil,
			},
			MockGetJoke: &MockGetJoke{
				OutReply: stringPointer("John Doe's OSI network model has only one layer - Physical."),
				OutError:     nil,
			},
			MockGetName: &MockGetName{
				OutReply: nil,
				OutError: errors.New("person api failure"),
			},
			ResponseError: errors.New("person api failure"),
		},
	}

	_ = GetHelloUnitTestData

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	for _, data := range GetHelloUnitTestData {
		t.Run(data.Description, func(t *testing.T) {
			serviceBundle := MockGetHelloRequiredServices(mockCtrl, data)

			f := joke_endpoints.MakeGetHelloEndpoint(appcontext.Context{}, *serviceBundle)
			requestTestData := data.Request.(joke_codecs.GetJokeRequest)

			var responseTestData joke_codecs.GetJokeResponse
			if data.Response != nil {
				responseTestData = data.Response.(joke_codecs.GetJokeResponse)
			}
			responseData, responseError := f(context.Background(), requestTestData)

			//assert results
			assert.Equal(t, data.ResponseError, responseError)
			if data.Response != nil {
				assert.Equal(t, responseTestData, responseData)
			}
		})
	}
}

func MockGetHelloRequiredServices(mockCtrl *gomock.Controller, data GetJokeTestData) *services.Bundle {
	jokesMock := services_mocks.NewMock_jokes(mockCtrl)
	jokesExpect := jokesMock.EXPECT()
	namesMock := services_mocks.NewMock_names(mockCtrl)
	namesExpect := namesMock.EXPECT()

	if data.MockGetJoke != nil {
		jokesExpect.Get().Return(
				data.MockGetJoke.OutReply,
				data.MockGetJoke.OutError)
	}
	if data.MockGetName != nil {
		namesExpect.Get().Return(
			data.MockGetName.OutReply,
			data.MockGetName.OutError)
	}

	serviceBundle := services.Bundle{
		JokesSvc: jokesMock,
		NamesSvc: namesMock,
	}

	return &serviceBundle
}
