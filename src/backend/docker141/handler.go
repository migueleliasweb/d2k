package docker141

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/migueleliasweb/d2k/src/openapi/gen/restapi/operations/container"
)

type EndpointMatch struct {
	HttpMethod string // e.g. http.Method*
	Endpoint   string
}

var ContainerCreateEndpointMatch = EndpointMatch{
	HttpMethod: http.MethodPost,
	Endpoint:   "/v1.41/containers/create",
}

func Docker141Handler(rw http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/_ping" {
		r.URL.Path = "/v1.41" + r.URL.Path
	}

	// docker cli uses this header to figure out the version of the daemon
	// instead of using `/version` *sigh*
	rw.Header().Add("Api-Version", "v1.41")
	fmt.Println("-----------------------------------")
}

// FixOptionalParamsContainerCreate fixes some weird problems with inconsistencies with the
// 	docker api regarding the ContainerCreate endpoint.
//
// The method call `AddMiddlewareFor` can't be used to mutate the request for some reason
// as it messes up with the request pointer inside go-swagger.
func FixOptionalParamsContainerCreate(rw http.ResponseWriter, r *http.Request) {
	if r.Method == ContainerCreateEndpointMatch.HttpMethod &&
		r.URL.Path == ContainerCreateEndpointMatch.Endpoint {

		containerCreateBody := &container.ContainerCreateBody{}

		bodyBytes, err := io.ReadAll(r.Body)
		r.Body.Close() //  must close now

		if err != nil {
			fmt.Println("err:", err)

			http.Error(
				rw,
				"error reading body",
				http.StatusInternalServerError,
			)

			return
		}

		json.Unmarshal(bodyBytes, &containerCreateBody)

		// The docker cli instead of sending a `zero`, it sends a `-1`... sigh
		if *containerCreateBody.HostConfig.MemorySwappiness < 0 {
			zero := int64(0)
			containerCreateBody.HostConfig.MemorySwappiness = &zero
		}

		// The docker cli sends a "no" but the v1.41 reference only mentions
		// `empty string`, `on-failure`, `unless-stopped` or `always` as valid values fos this field.
		// So we need to fix `no` to `empty string`.
		if strings.ToLower(containerCreateBody.HostConfig.RestartPolicy.Name) == "no" {
			containerCreateBody.HostConfig.RestartPolicy.Name = ""
		}

		containerCreateBodyBytes, err := json.Marshal(containerCreateBody)

		if err != nil {
			fmt.Println("err:", err)

			http.Error(
				rw,
				"error reconstructing body",
				http.StatusInternalServerError,
			)

			return
		}

		r.Body = io.NopCloser(bytes.NewBuffer(containerCreateBodyBytes))
	}
}
