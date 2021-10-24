package main

import (
	"net/http"

	"github.com/migueleliasweb/d2k/src/backend/baseapi"
	"github.com/migueleliasweb/d2k/src/backend/docker141"
	"github.com/migueleliasweb/d2k/src/backend/k8s"
)

func main() {

	api := baseapi.ConfigureBaseApi()

	// sets up handlers
	k8s.ApiConfigurator(api)

	http.ListenAndServe("127.0.0.1:8080", http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		docker141.Docker141Handler(
			rw,
			r,
		)

		api.Serve(nil).ServeHTTP(rw, r)
	}))
}
