package joke_endpoints

import (
	"context"
	"github.com/bmsandoval/apple-coding-challenge/api/transport_http/codecs/joke_codecs"
	"github.com/bmsandoval/apple-coding-challenge/library/appcontext"
	"github.com/bmsandoval/apple-coding-challenge/services"
	"github.com/go-kit/kit/endpoint"
	"regexp"
)

// This function gets the joke, do you???
func MakeGetHelloEndpoint(appCtx appcontext.Context, services services.Bundle) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		// request is empty, don't really need to do this
		_ = request.(joke_codecs.GetJokeRequest)

		joke, err := services.JokesSvc.Get()
		if err != nil {
			return nil, err
		}
		name, err := services.NamesSvc.Get()
		if err != nil {
			return nil, err
		}

		var reg = regexp.MustCompile(`John Doe`)
		updatedJoke := reg.ReplaceAllString(*joke, *name)

		return joke_codecs.GetJokeResponse(updatedJoke), nil
	}
}