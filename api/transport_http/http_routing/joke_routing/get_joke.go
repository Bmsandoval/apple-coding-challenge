package joke_routing

import (
	"github.com/bmsandoval/apple-coding-challenge/api/transport_http/codecs/joke_codecs"
	"github.com/bmsandoval/apple-coding-challenge/endpoints/joke_endpoints"
	httpTransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"github.com/bmsandoval/apple-coding-challenge/api/transport_http/codecs/requests/hello_requests"
	"github.com/bmsandoval/apple-coding-challenge/api/transport_http/codecs/responses/hello_responses"
	"github.com/bmsandoval/apple-coding-challenge/api/transport_http/http_routing"
	"github.com/bmsandoval/apple-coding-challenge/endpoints/greeting_endpoints"
	"github.com/bmsandoval/apple-coding-challenge/library/appcontext"
	"github.com/bmsandoval/apple-coding-challenge/services"
)

func init() {
	http_routing.Bundle(MakeGetJokeHttpHandler())
}

func MakeGetJokeHttpHandler() http_routing.Bundlable {
	return func(appCtx appcontext.Context, router *mux.Router, services services.Bundle) {
		api := router.PathPrefix("/api").Subrouter()

		endpoint := joke_endpoints.MakeGetHelloEndpoint(appCtx, services)
		decoder, _ := joke_codecs.MakeGetNameRequestDecoder(appCtx)
		encoder, _ := joke_codecs.MakeGetNameResponseEncoder(appCtx)

		//POST /Find Campaing ID
		api.Methods("GET").Path("/get").Handler(httpTransport.NewServer(
			endpoint,
			decoder,
			encoder,
		))
	}
}
