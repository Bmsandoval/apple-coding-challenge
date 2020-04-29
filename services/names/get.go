package names

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

type apiResponse struct {
	person
}

type person struct {
	Name string `json:"name"`
	Surname string `json:"surname"`
	Gender string `json:"gender"`
	Region string `json:"region"`
}

func (h *Helper) Get() (*string, error) {
	response, err := http.Get("http://uinames.com/api/")
	if err != nil {
		return nil, err
	}
	if response.StatusCode != 200 {
		return nil, errors.New("names api not accessible")
	}
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var person apiResponse
	if err := json.Unmarshal(data, &person); err != nil {
		return nil, err
	}

	return &person.Name, nil
}
