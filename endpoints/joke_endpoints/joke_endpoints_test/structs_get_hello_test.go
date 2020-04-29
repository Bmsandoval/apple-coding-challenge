package joke_endpoints_test

type GetJokeTestData struct {
	BaseTestData
	MockGetJoke *MockGetJoke
	MockGetName *MockGetName
	ResponseError error
}

type MockGetJoke struct {
	OutReply *string
	OutError   error
}

type MockGetName struct {
	OutReply *string
	OutError   error
}
