package k8s

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/go-openapi/runtime/middleware"
	"github.com/migueleliasweb/d2k/src/openapi/gen/models"
	"github.com/migueleliasweb/d2k/src/openapi/gen/restapi/operations"
	"github.com/migueleliasweb/d2k/src/openapi/gen/restapi/operations/container"
	"k8s.io/client-go/kubernetes"
)

func ApiConfigurator(
	clientset kubernetes.Interface,
	api *operations.DockerEngineAPIAPI,
) {

	api.AddMiddlewareFor(
		http.MethodPost,
		"/containers/create", // can be auto generated via openapi definition
		func(h http.Handler) http.Handler {
			return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
				unmarshalledParams := container.ContainerCreateParams{}

				bodyBytes, err := io.ReadAll(r.Body)

				if err != nil {
					http.Error(
						rw,
						"error reading body",
						http.StatusInternalServerError,
					)
				}

				json.Unmarshal(bodyBytes, &unmarshalledParams)

				r.Body.Close() //  must close now
				r.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

				fmt.Println("params", unmarshalledParams)
			})
		},
	)

	api.ContainerContainerCreateHandler = container.ContainerCreateHandlerFunc(
		containerCreateHandler(clientset),
	)

	api.ContainerContainerListHandler = container.ContainerListHandlerFunc(
		func(clp container.ContainerListParams) middleware.Responder {
			return container.NewContainerListOK().WithPayload([]*models.ContainerSummaryItems0{
				{
					ID:    "123",
					Names: []string{"container1"},
					Image: "alpine",
				},
			})
		},
	)
}
