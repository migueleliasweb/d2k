package docker141

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/docker/docker/client"
	"github.com/migueleliasweb/d2k/src/backend/docker141"
	"github.com/stretchr/testify/assert"
)

func TestIntegrationServerVersion(t *testing.T) {
	fakeServer := httptest.NewServer(
		http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
			docker141.Docker141Handler(rw, r)

			api := ConfigureBaseApi()

			api.Serve(nil).ServeHTTP(rw, r)
		}),
	)

	defer fakeServer.Close()

	ctx := context.Background()

	cli, err := client.NewClientWithOpts(
		client.WithAPIVersionNegotiation(),
		client.WithHost(fakeServer.URL),
	)

	if err != nil {
		panic(err)
	}

	version, err := cli.ServerVersion(ctx)

	if err != nil {
		panic(err)
	}

	assert.Equal(
		t,
		"v1.41", // TODO: get this from build flag?
		version.APIVersion,
	)
}
