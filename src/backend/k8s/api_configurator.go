package k8s

import (
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

	api.ContainerContainerStartHandler = container.ContainerStartHandlerFunc(
		func(csp container.ContainerStartParams) middleware.Responder {
			return container.NewContainerStartNoContent()
		},
	)
}
