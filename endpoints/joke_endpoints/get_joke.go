package joke_endpoints

import (
	"context"
	"github.com/bmsandoval/apple-coding-challenge/api/transport_http/codecs/joke_codecs"
	"github.com/bmsandoval/apple-coding-challenge/api/transport_http/codecs/name_codecs"
	"github.com/go-kit/kit/endpoint"
	"github.com/bmsandoval/apple-coding-challenge/api/transport_http/codecs/requests/hello_requests"
	"github.com/bmsandoval/apple-coding-challenge/api/transport_http/codecs/responses/hello_responses"
	"github.com/bmsandoval/apple-coding-challenge/library/appcontext"
	"github.com/bmsandoval/apple-coding-challenge/services"
	"github.com/bmsandoval/apple-coding-challenge/services/grpc_service"
	"log"
	"time"
)

// This function gets the joke, do you???
func MakeGetHelloEndpoint(appCtx appcontext.Context, services services.Bundle) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		_ = request.(joke_codecs.GetNameRequest)

		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		response, err := services.GrpcSvc.GreeterClient.GetHello(ctx, &grpc_service.GetHelloRequest{})
		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}
		log.Printf("Greeting: %s", response)

		return joke_codecs.GetHelloResponse{
			Response: response.Greetings,
		}, nil
	}
}