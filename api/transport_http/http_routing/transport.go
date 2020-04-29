package http_routing

import (
	"github.com/gorilla/mux"
	"github.com/bmsandoval/apple-coding-challenge/library/appcontext"
	"github.com/bmsandoval/apple-coding-challenge/services"
)

type ServerContext struct {
	AppCtx appcontext.Context
	Bundle services.Bundle
}

type Bundlable func(appCtx appcontext.Context, router *mux.Router, services services.Bundle)

var Bundlables []Bundlable

func Bundle(bundlable Bundlable) {
	Bundlables = append(Bundlables, bundlable)
}

func BundleAll(appCtx appcontext.Context, router *mux.Router, services services.Bundle) {
	for _, bundlable := range Bundlables {
		bundlable(appCtx, router, services)
	}
}
