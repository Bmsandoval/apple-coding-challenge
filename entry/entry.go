package entry

import (
	"github.com/bmsandoval/apple-coding-challenge/api/transport_http/http_routing"
	"github.com/bmsandoval/apple-coding-challenge/library/appcontext"
	"github.com/bmsandoval/apple-coding-challenge/services"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func Entry() {
	// Build Context
	ctx := appcontext.Context{
		// Config: *config,
	}

	// Bundle Services
	serviceBundle, err := services.NewBundle(ctx)
	if err != nil {
		panic(err) }

	router := mux.NewRouter()
	http_routing.BundleAll(ctx, router, *serviceBundle)

	svr := http.Server{
		Addr: ":8080",
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
			router.ServeHTTP(w, r)
		}),
	}

	log.Println("Starting Server...")
	err = svr.ListenAndServe()
	if err != nil {
		log.Fatalf("error listening: %v", err.Error())
	}
}
