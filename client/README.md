# Go API client for client

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: unstable
- Package version: 1.0.0
- Generator version: 7.12.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```sh
go get github.com/stretchr/testify/assert
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```go
import client "github.com/klearwave/service-info/client"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```go
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `client.ContextServerIndex` of type `int`.

```go
ctx := context.WithValue(context.Background(), client.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `client.ContextServerVariables` of type `map[string]string`.

```go
ctx := context.WithValue(context.Background(), client.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `client.ContextOperationServerIndices` and `client.ContextOperationServerVariables` context maps.

```go
ctx := context.WithValue(context.Background(), client.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), client.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *http://localhost*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*AboutAPI* | [**GetAbout**](docs/AboutAPI.md#getabout) | **Get** /api/about | Get overall information about the service.
*ContainerImagesAPI* | [**CreateContainerImage**](docs/ContainerImagesAPI.md#createcontainerimage) | **Post** /api/v0/container_images | Create a new container image.
*ContainerImagesAPI* | [**DeleteContainerImage**](docs/ContainerImagesAPI.md#deletecontainerimage) | **Delete** /api/v0/container_images/{id} | Delete a specific container image.
*ContainerImagesAPI* | [**GetContainerImage**](docs/ContainerImagesAPI.md#getcontainerimage) | **Get** /api/v0/container_images/{id} | Get specific container image information.
*ContainerImagesAPI* | [**ListContainerImageVersions**](docs/ContainerImagesAPI.md#listcontainerimageversions) | **Get** /api/v0/container_images/{id}/versions | List versions for a specific container image.
*ContainerImagesAPI* | [**ListContainerImages**](docs/ContainerImagesAPI.md#listcontainerimages) | **Get** /api/v0/container_images | List all container image information.
*HealthAPI* | [**Healthz**](docs/HealthAPI.md#healthz) | **Get** /healthz | Get health information about the service.
*VersionsAPI* | [**CreateVersion**](docs/VersionsAPI.md#createversion) | **Post** /api/v0/versions | Create a new Version.
*VersionsAPI* | [**DeleteVersion**](docs/VersionsAPI.md#deleteversion) | **Delete** /api/v0/versions/{id} | Delete a specific version.
*VersionsAPI* | [**GetVersion**](docs/VersionsAPI.md#getversion) | **Get** /api/v0/versions/{id} | Get specific version information.
*VersionsAPI* | [**ListVersionContainerImages**](docs/VersionsAPI.md#listversioncontainerimages) | **Get** /api/v0/versions/{id}/container_images | List container images for a specific version.
*VersionsAPI* | [**ListVersions**](docs/VersionsAPI.md#listversions) | **Get** /api/v0/versions | List all version information.


## Documentation For Models

 - [About](docs/About.md)
 - [ContainerImage](docs/ContainerImage.md)
 - [ContainerImageBody](docs/ContainerImageBody.md)
 - [ContainerImageResponseBody](docs/ContainerImageResponseBody.md)
 - [ErrorDetail](docs/ErrorDetail.md)
 - [ErrorModel](docs/ErrorModel.md)
 - [Health](docs/Health.md)
 - [Version](docs/Version.md)
 - [VersionBody](docs/VersionBody.md)
 - [VersionResponseBody](docs/VersionResponseBody.md)


## Documentation For Authorization

Endpoints do not require authorization.


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author



