#!/usr/bin/env bash

go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.59.1

# Install Kiota
# dotnet tool uninstall Microsoft.OpenApi.Kiota --global
dotnet tool install --global Microsoft.OpenApi.Kiota --version 1.19.0