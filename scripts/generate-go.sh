#!/bin/sh

set -ex

# TODO: Add options for version 

SCHEMA_FILE="schemas/openai-openapi.yaml"
NAMESPACE="go-sdk"

./scripts/tools.sh
VERSION="v1.0.0" #Currently noop

go run schemas/main.go --version=$VERSION
kiota generate -l go --ll trace -o $(pwd)/stage/go/$NAMESPACE/pkg/openai -n github.com/openai/$NAMESPACE/pkg/openai -d $SCHEMA_FILE --ebc
cd $(pwd)/stage/go/$NAMESPACE
go build ./...