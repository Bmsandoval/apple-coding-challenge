package joke_endpoints_test

func stringPointer(sample string) *string {
	return &sample
}

type BaseTestData struct {
	Description string
	Request interface{}
	Response interface{}
}