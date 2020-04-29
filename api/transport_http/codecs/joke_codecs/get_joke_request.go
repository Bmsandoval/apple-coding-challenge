package joke_codecs

import (
	"context"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/bmsandoval/apple-coding-challenge/library/appcontext"
	"net/http"
)

type GetNameRequest struct {
}

//Decode the Push notification Req.
func MakeGetNameRequestDecoder(appCtx appcontext.Context) (kithttp.DecodeRequestFunc, error) {
	return func(ctx context.Context, r *http.Request) (interface{}, error) {
		var Req GetNameRequest
		//err := decode.Decode(&Req)
		//if err != nil {
		//	return nil, errors.New("inconsistent mapping between route and handler")
		//}

		// GET requests don't have bodies
		return Req, nil
	}, nil
}
