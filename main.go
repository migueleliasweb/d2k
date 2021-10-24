package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"

	"github.com/migueleliasweb/d2k/src/openapi/gen/models"
	"github.com/migueleliasweb/d2k/src/openapi/gen/restapi"
	"github.com/migueleliasweb/d2k/src/openapi/gen/restapi/operations"
	"github.com/migueleliasweb/d2k/src/openapi/gen/restapi/operations/container"
	"github.com/migueleliasweb/d2k/src/openapi/gen/restapi/operations/system"
)

func main() {

	swaggerSpec, err := loads.Embedded(restapi.SwaggerJSON, restapi.FlatSwaggerJSON)
	if err != nil {
		log.Fatalln(err)
	}

	api := operations.NewDockerEngineAPIAPI(swaggerSpec)

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

	api.SystemSystemVersionHandler = system.SystemVersionHandlerFunc(
		func(svp system.SystemVersionParams) middleware.Responder {
			return system.NewSystemVersionOK().WithPayload(&models.SystemVersion{
				APIVersion:    "v1.4.1",
				MinAPIVersion: "v1.4.1",
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

	http.ListenAndServe("127.0.0.1:8080", http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		fmt.Println(r.URL.Path, r.Method)

		splittedPath := strings.Split(r.URL.Path, "/")

		if splittedPath[1] == "_ping" || splittedPath[1] == "ping" || (len(splittedPath) >= 3 && strings.HasPrefix(splittedPath[1], "v1.") && splittedPath[1] != "v1.41") {
			splittedPath[1] = "v1.41"

			fmt.Println("changed to:", strings.Join(splittedPath, "/"))
			r.URL.Path = strings.Join(splittedPath, "/")
		}

		rw.Header().Add("Api-Version", "v1.41")
		fmt.Println("-----------------------------------")

		api.Serve(nil).ServeHTTP(rw, r)
	}))
}
