package docker141

import (
	"log"

	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"
	"github.com/migueleliasweb/d2k/src/openapi/gen/models"
	"github.com/migueleliasweb/d2k/src/openapi/gen/restapi"
	"github.com/migueleliasweb/d2k/src/openapi/gen/restapi/operations"
	"github.com/migueleliasweb/d2k/src/openapi/gen/restapi/operations/system"
)

func ConfigureBaseApi() *operations.DockerEngineAPIAPI {
	swaggerSpec, err := loads.Embedded(restapi.SwaggerJSON, restapi.FlatSwaggerJSON)
	if err != nil {
		log.Fatalln(err)
	}

	api := operations.NewDockerEngineAPIAPI(swaggerSpec)

	api.SystemSystemVersionHandler = system.SystemVersionHandlerFunc(
		func(svp system.SystemVersionParams) middleware.Responder {
			return system.NewSystemVersionOK().WithPayload(&models.SystemVersion{
				APIVersion:    "v1.41",
				MinAPIVersion: "v1.41",
				Version:       "v0.0.1-alpha", // d2k version
				Components: []*models.ComponentVersion{
					{
						Name:    swag.String("D2K"),
						Version: "v0.0.1-alpha",
					},
					{
						Name:    swag.String("Kubernetes"),
						Version: "v1.22.2",
					},
				},
			})
		},
	)

	api.SystemSystemPingHeadHandler = system.SystemPingHeadHandlerFunc(
		func(sphp system.SystemPingHeadParams) middleware.Responder {
			return system.NewSystemPingOK()
		},
	)

	api.SystemSystemPingHandler = system.SystemPingHandlerFunc(
		func(spp system.SystemPingParams) middleware.Responder {
			return system.NewSystemPingOK()
		},
	)

	return api
}
