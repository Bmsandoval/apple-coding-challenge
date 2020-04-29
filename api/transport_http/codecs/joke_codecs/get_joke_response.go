package joke_codecs

import (
	"context"
	"encoding/json"
	"errors"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/bmsandoval/apple-coding-challenge/library/appcontext"
	"net/http"
)

// Errorer is implemented by all concrete response types that may contain
// errors. It allows us to change the HTTP response code without needing to
// trigger an endpoint (transport-level) error.
type Errorer interface {
	error() error
}

type GetJokeResponse string

func MakeGetJokeResponseEncoder(appCtx appcontext.Context) (kithttp.EncodeResponseFunc, error) {
	return func(ctx context.Context, httpResponse http.ResponseWriter, endpointResponse interface{}) error {
		response, ok := endpointResponse.(GetJokeResponse)
		if ! ok {
			return errors.New("response was not a string")
		}
		_, err := httpResponse.Write([]byte(response))
		return err
	}, nil
}

// EncodeError - Provide those as HTTP errors
func EncodeError(_ context.Context, err error, w http.ResponseWriter) {
	if err == nil {
		panic("EncodeError with nil error")
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}
