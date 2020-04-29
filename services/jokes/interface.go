package jokes

import (
	"github.com/bmsandoval/apple-coding-challenge/library/appcontext"
)

type Helper struct {
	AppCtx appcontext.Context
}
type Helpable struct{}

func(h Helpable) NewHelper(appCtx appcontext.Context) (interface{}, error) {
	return &Helper{
		AppCtx: appCtx,
	}, nil
}

func (h Helpable) ServiceName() string {
	return "JokesSvc"
}

type Service interface {
	Get() (*string, error)
}
