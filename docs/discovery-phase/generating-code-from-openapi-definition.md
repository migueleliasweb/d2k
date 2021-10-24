# Generating code from OpenAPI definitition

## From https://github.com/moby/moby/blob/master/hack/generate-swagger-api.sh

```go
swagger generate model -f api/swagger.yaml \
	-t api -m types --skip-validator -C api/swagger-gen.yaml \
	-n ErrorResponse \
	-n GraphDriverData \
	-n IdResponse \
	-n ImageDeleteResponseItem \
	-n ImageSummary \
	-n Plugin -n PluginDevice -n PluginMount -n PluginEnv -n PluginInterfaceType \
	-n Port \
	-n ServiceUpdateResponse \
	-n Volume
```