# \VersionsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateVersion**](VersionsAPI.md#CreateVersion) | **Post** /api/v0/versions | Create a new Version.
[**DeleteVersion**](VersionsAPI.md#DeleteVersion) | **Delete** /api/v0/versions/{id} | Delete a specific version.
[**GetVersion**](VersionsAPI.md#GetVersion) | **Get** /api/v0/versions/{id} | Get specific version information.
[**ListVersionContainerImages**](VersionsAPI.md#ListVersionContainerImages) | **Get** /api/v0/versions/{id}/container_images | List container images for a specific version.
[**ListVersions**](VersionsAPI.md#ListVersions) | **Get** /api/v0/versions | List all version information.



## CreateVersion

> VersionResponseBody CreateVersion(ctx).VersionBody(versionBody).Authorization(authorization).Execute()

Create a new Version.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/klearwave/service-info/client"
)

func main() {
	versionBody := *openapiclient.NewVersionBody() // VersionBody | 
	authorization := "authorization_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VersionsAPI.CreateVersion(context.Background()).VersionBody(versionBody).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VersionsAPI.CreateVersion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateVersion`: VersionResponseBody
	fmt.Fprintf(os.Stdout, "Response from `VersionsAPI.CreateVersion`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **versionBody** | [**VersionBody**](VersionBody.md) |  | 
 **authorization** | **string** |  | 

### Return type

[**VersionResponseBody**](VersionResponseBody.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteVersion

> VersionResponseBody DeleteVersion(ctx, id).Authorization(authorization).Execute()

Delete a specific version.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/klearwave/service-info/client"
)

func main() {
	id := "id_example" // string | Database ID of the stored object.
	authorization := "authorization_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VersionsAPI.DeleteVersion(context.Background(), id).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VersionsAPI.DeleteVersion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteVersion`: VersionResponseBody
	fmt.Fprintf(os.Stdout, "Response from `VersionsAPI.DeleteVersion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Database ID of the stored object. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authorization** | **string** |  | 

### Return type

[**VersionResponseBody**](VersionResponseBody.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVersion

> VersionResponseBody GetVersion(ctx, id).Execute()

Get specific version information.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/klearwave/service-info/client"
)

func main() {
	id := "id_example" // string | Database ID of the stored object.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VersionsAPI.GetVersion(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VersionsAPI.GetVersion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVersion`: VersionResponseBody
	fmt.Fprintf(os.Stdout, "Response from `VersionsAPI.GetVersion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Database ID of the stored object. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**VersionResponseBody**](VersionResponseBody.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListVersionContainerImages

> ContainerImageResponseBody ListVersionContainerImages(ctx, id).Execute()

List container images for a specific version.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/klearwave/service-info/client"
)

func main() {
	id := "id_example" // string | Database ID of the stored object.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VersionsAPI.ListVersionContainerImages(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VersionsAPI.ListVersionContainerImages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListVersionContainerImages`: ContainerImageResponseBody
	fmt.Fprintf(os.Stdout, "Response from `VersionsAPI.ListVersionContainerImages`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Database ID of the stored object. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListVersionContainerImagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ContainerImageResponseBody**](ContainerImageResponseBody.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListVersions

> VersionResponseBody ListVersions(ctx).Execute()

List all version information.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/klearwave/service-info/client"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VersionsAPI.ListVersions(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VersionsAPI.ListVersions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListVersions`: VersionResponseBody
	fmt.Fprintf(os.Stdout, "Response from `VersionsAPI.ListVersions`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListVersionsRequest struct via the builder pattern


### Return type

[**VersionResponseBody**](VersionResponseBody.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

