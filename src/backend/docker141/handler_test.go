package docker141

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDocker141Handler(t *testing.T) {
	tests := []struct {
		name string
		rw   http.ResponseWriter
		r    *http.Request
	}{
		{
			name: "_ping",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fakeServer := httptest.NewServer(
				http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
					Docker141Handler(rw, r)

					apiVersionHeaderValue := rw.Header().Values("Api-Version")

					assert.Equal(
						t,
						"v1.41",
						apiVersionHeaderValue[0],
					)
				}),
			)

			fakeServer.Client().Get(
				tt.name,
			)
		})
	}
}
