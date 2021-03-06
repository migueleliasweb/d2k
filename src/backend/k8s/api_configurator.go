package k8s

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/migueleliasweb/d2k/src/openapi/gen/models"
	"github.com/migueleliasweb/d2k/src/openapi/gen/restapi/operations"
	"github.com/migueleliasweb/d2k/src/openapi/gen/restapi/operations/container"
)

func ApiConfigurator(api *operations.DockerEngineAPIAPI) {
	api.ContainerContainerCreateHandler = container.ContainerCreateHandlerFunc(
		func(params container.ContainerCreateParams) middleware.Responder {
			return container.NewContainerCreateCreated().WithPayload(&container.ContainerCreateCreatedBody{
				ID: "123",
			})
		},
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
