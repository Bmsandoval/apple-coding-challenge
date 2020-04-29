package jokes

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

type apiResponse struct {
	Type string `json:"type"`
	Value joke `json:"value"`
}

type joke struct {
	Id int `json:"id"`
	Joke string `json:"joke"`
	Categories []string `json:"categories"`
}

func (h *Helper) Get() (*string, error) {
	response, err := http.Get("http://api.icndb.com/jokes/random?firstName=John&lastName=Doe&limitTo=\\[nerdy\\]")
	if err != nil {
		return nil, err
	}
	if response.StatusCode != 200 {
		return nil, errors.New("jokes api not accessible")

	}
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var joke apiResponse
	if err := json.Unmarshal(data, &joke); err != nil {
		return nil, err
	}

	return &joke.Value.Joke, nil
}
