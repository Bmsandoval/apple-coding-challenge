package joke_endpoints

import (
	"context"
	"github.com/bmsandoval/apple-coding-challenge/api/transport_http/codecs/joke_codecs"
	"github.com/bmsandoval/apple-coding-challenge/library/appcontext"
	"github.com/bmsandoval/apple-coding-challenge/services"
	"github.com/go-kit/kit/endpoint"
)

// This function gets the joke, do you???
func MakeGetHelloEndpoint(appCtx appcontext.Context, services services.Bundle) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		_ = request.(joke_codecs.GetJokeRequest)

		return joke_codecs.GetJokeResponse{
			Response: "hello world",
		}, nil
	}
}