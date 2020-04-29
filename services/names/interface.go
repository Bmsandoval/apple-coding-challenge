package names

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
	return "NamesSvc"
}

type Service interface {
	Get() (error)
}
