#!/usr/bin/env bash

echo "Waiting for https://github.com/deepmap/oapi-codegen/pull/436 to be merged."

exit 1

OAPI_CODEGEN_VERSION="v1.8.3" # hopefully
DOCKER_CONVERTED_OPENAPI3_VERSION="v1.41.yaml"

go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@${OAPI_CODEGEN_VERSION}

for GEN in types server; do
    oapi-codegen -package openapigen -generate ${GEN} ./build/${DOCKER_CONVERTED_OPENAPI3_VERSION} > ../src/openapigen/${GEN}.go
done 