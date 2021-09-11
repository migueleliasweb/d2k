package d2k

import (
	"net/http"
	"testing"

	"github.com/docker/docker/client"
	"github.com/stretchr/testify/assert"
)

type RoundTripperFunc func(*http.Request) (*http.Response, error)

func (f RoundTripperFunc) RoundTrip(r *http.Request) (*http.Response, error) {
	return f(r)
}

func TestHandlersContainerCreate(t *testing.T) {
	c := &http.Client{
		Transport: RoundTripperFunc(func(r *http.Request) (*http.Response, error) {
			return &http.Response{}, nil
		}),
	}

	dockerClient, dockerClientErr := client.NewClientWithOpts(
		client.WithHTTPClient(c),
	)

	assert.Nil(t, dockerClientErr)
	assert.NotNil(t, dockerClient)
}
