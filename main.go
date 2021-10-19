package main

import (
	"log"
	"os"

	"github.com/go-openapi/loads"
	flags "github.com/jessevdk/go-flags"

	"github.com/go-openapi/runtime/middleware"
	"github.com/migueleliasweb/d2k/src/openapi/gen/restapi"
	"github.com/migueleliasweb/d2k/src/openapi/gen/restapi/operations"
	"github.com/migueleliasweb/d2k/src/openapi/gen/restapi/operations/container"
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

	server := restapi.NewServer(api)
	defer server.Shutdown()

	parser := flags.NewParser(server, flags.Default)
	parser.ShortDescription = "Docker Engine API"
	parser.LongDescription = "The Engine API is an HTTP API served by Docker Engine. It is the API the\nDocker client uses to communicate with the Engine, so everything the Docker\nclient can do can be done with the API.\n\nMost of the client's commands map directly to API endpoints (e.g. `docker ps`\nis `GET /containers/json`). The notable exception is running containers,\nwhich consists of several API calls.\n\n# Errors\n\nThe API uses standard HTTP status codes to indicate the success or failure\nof the API call. The body of the response will be JSON in the following\nformat:\n\n```\n{\n  \"message\": \"page not found\"\n}\n```\n\n# Versioning\n\nThe API is usually changed in each release, so API calls are versioned to\nensure that clients don't break. To lock to a specific version of the API,\nyou prefix the URL with its version, for example, call `/v1.30/info` to use\nthe v1.30 version of the `/info` endpoint. If the API version specified in\nthe URL is not supported by the daemon, a HTTP `400 Bad Request` error message\nis returned.\n\nIf you omit the version-prefix, the current version of the API (v1.41) is used.\nFor example, calling `/info` is the same as calling `/v1.41/info`. Using the\nAPI without a version-prefix is deprecated and will be removed in a future release.\n\nEngine releases in the near future should support this version of the API,\nso your client will continue to work even if it is talking to a newer Engine.\n\nThe API uses an open schema model, which means server may add extra properties\nto responses. Likewise, the server will ignore any extra query parameters and\nrequest body properties. When you write clients, you need to ignore additional\nproperties in responses to ensure they do not break when talking to newer\ndaemons.\n\n\n# Authentication\n\nAuthentication for registries is handled client side. The client has to send\nauthentication details to various endpoints that need to communicate with\nregistries, such as `POST /images/(name)/push`. These are sent as\n`X-Registry-Auth` header as a [base64url encoded](https://tools.ietf.org/html/rfc4648#section-5)\n(JSON) string with the following structure:\n\n```\n{\n  \"username\": \"string\",\n  \"password\": \"string\",\n  \"email\": \"string\",\n  \"serveraddress\": \"string\"\n}\n```\n\nThe `serveraddress` is a domain/IP without a protocol. Throughout this\nstructure, double quotes are required.\n\nIf you have already got an identity token from the [`/auth` endpoint](#operation/SystemAuth),\nyou can just pass this instead of credentials:\n\n```\n{\n  \"identitytoken\": \"9cbaf023786cd7...\"\n}\n```\n"
	server.ConfigureFlags()
	for _, optsGroup := range api.CommandLineOptionsGroups {
		_, err := parser.AddGroup(optsGroup.ShortDescription, optsGroup.LongDescription, optsGroup.Options)
		if err != nil {
			log.Fatalln(err)
		}
	}

	if _, err := parser.Parse(); err != nil {
		code := 1
		if fe, ok := err.(*flags.Error); ok {
			if fe.Type == flags.ErrHelp {
				code = 0
			}
		}
		os.Exit(code)
	}

	server.ConfigureAPI()

	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}

}
