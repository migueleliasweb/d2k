package k8s

import (
	"encoding/base64"

	"github.com/go-openapi/runtime/middleware"
	"github.com/migueleliasweb/d2k/src/openapi/gen/models"
	"github.com/migueleliasweb/d2k/src/openapi/gen/restapi/operations"
	"github.com/migueleliasweb/d2k/src/openapi/gen/restapi/operations/container"
	"github.com/migueleliasweb/d2k/src/openapi/gen/restapi/operations/system"
	"k8s.io/client-go/kubernetes"
)

func ApiConfigurator(
	clientset kubernetes.Interface,
	api *operations.DockerEngineAPIAPI,
) {
	api.ContainerContainerCreateHandler = container.ContainerCreateHandlerFunc(
		containerCreateHandler(clientset),
	)

	api.ContainerContainerListHandler = container.ContainerListHandlerFunc(
		func(params container.ContainerListParams) middleware.Responder {
			return container.NewContainerListOK().WithPayload([]*models.ContainerSummaryItems0{
				{
					ID:    "123",
					Names: []string{"container1"},
					Image: "ubuntu",
				},
			})
		},
	)

	api.ContainerContainerStartHandler = container.ContainerStartHandlerFunc(
		func(csp container.ContainerStartParams) middleware.Responder {
			return container.NewContainerStartNoContent()
		},
	)

	// https://docs.docker.com/engine/api/v1.41/#operation/SystemEvents
	// Type: "container"
	// Action: "create"
	// Actor:
	//   ID: "ede54ee1afda366ab42f824e8a5ffd195155d853ceaec74a927f249ea270c743"
	//   Attributes:
	// 	com.example.some-label: "some-label-value"
	// 	image: "alpine"
	// 	name: "my-container"
	// time: 1461943101
	//
	// "create" and "start"
	api.SystemSystemEventsHandler = system.SystemEventsHandlerFunc(
		func(params system.SystemEventsParams) middleware.Responder {
			return system.NewSystemEventsOK().WithPayload(
				&system.SystemEventsOKBody{
					Action: "create",
					Type:   "container",
					Actor: &system.SystemEventsOKBodyActor{
						ID: base64.StdEncoding.EncodeToString([]byte("migueleliasweb/d2k")),
						Attributes: addCommonLabels(map[string]string{
							"foo":   "bar",
							"image": "ubuntu",
							"name":  "container1",
						}),
					},
				},
			)
		},
	)
}
