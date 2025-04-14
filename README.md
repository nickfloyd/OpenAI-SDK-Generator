# OpenAI-SDK-Generator
SDK generator that uses the OpenAI OpenAPI descriptions to build out SDKs using Microsoft's Kiota


## Development requirements
* .NET 7 or greater - this is for the Kiota toolset
* Go latest - this is for the generative tooling
* [Kiota](https://learn.microsoft.com/en-us/openapi/kiota/install?tabs=bash) `dotnet tool install --global Microsoft.OpenApi.Kiot


## Usage

`./scripts/generate-go.sh` (this will install the tools as well)