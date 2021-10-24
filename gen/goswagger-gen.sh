#!/usr/bin/env bash

# Source definition from: https://docs.docker.com/engine/api/v1.41.yaml
# Example generation from: https://github.com/moby/moby/blob/v20.10.8/hack/generate-swagger-api.sh

docker run -it --rm \
    -u $UID \
    -v ${PWD}:/opt/d2k \
    --workdir /opt/d2k \
    --entrypoint swagger \
    quay.io/goswagger/swagger:v0.27.0 \
    generate server -f ./gen/v1.41.yaml --target=./src/openapi/gen/

docker run -it --rm \
    -u $UID \
    -v ${PWD}:/opt/d2k \
    --workdir /opt/d2k \
    --entrypoint sh \
    quay.io/goswagger/swagger:v0.27.0