# \ContainerImagesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateContainerImage**](ContainerImagesAPI.md#CreateContainerImage) | **Post** /api/v0/container_images | Create a new container image.
[**DeleteContainerImage**](ContainerImagesAPI.md#DeleteContainerImage) | **Delete** /api/v0/container_images/{id} | Delete a specific container image.
[**GetContainerImage**](ContainerImagesAPI.md#GetContainerImage) | **Get** /api/v0/container_images/{id} | Get specific container image information.
[**ListContainerImageVersions**](ContainerImagesAPI.md#ListContainerImageVersions) | **Get** /api/v0/container_images/{id}/versions | List versions for a specific container image.
[**ListContainerImages**](ContainerImagesAPI.md#ListContainerImages) | **Get** /api/v0/container_images | List all container image information.



## CreateContainerImage

> ContainerImageResponseBody CreateContainerImage(ctx).ContainerImageBody(containerImageBody).Authorization(authorization).Execute()

Create a new container image.



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
	containerImageBody := *openapiclient.NewContainerImageBody() // ContainerImageBody | 
	authorization := "authorization_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContainerImagesAPI.CreateContainerImage(context.Background()).ContainerImageBody(containerImageBody).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContainerImagesAPI.CreateContainerImage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateContainerImage`: ContainerImageResponseBody
	fmt.Fprintf(os.Stdout, "Response from `ContainerImagesAPI.CreateContainerImage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateContainerImageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **containerImageBody** | [**ContainerImageBody**](ContainerImageBody.md) |  | 
 **authorization** | **string** |  | 

### Return type

[**ContainerImageResponseBody**](ContainerImageResponseBody.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteContainerImage

> ContainerImageResponseBody DeleteContainerImage(ctx, id).Authorization(authorization).Execute()

Delete a specific container image.



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
	id := int64(1) // int64 | Database ID of the stored object.
	authorization := "authorization_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContainerImagesAPI.DeleteContainerImage(context.Background(), id).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContainerImagesAPI.DeleteContainerImage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteContainerImage`: ContainerImageResponseBody
	fmt.Fprintf(os.Stdout, "Response from `ContainerImagesAPI.DeleteContainerImage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Database ID of the stored object. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteContainerImageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authorization** | **string** |  | 

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


## GetContainerImage

> ContainerImageResponseBody GetContainerImage(ctx, id).Execute()

Get specific container image information.



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
	id := int64(1) // int64 | Database ID of the stored object.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContainerImagesAPI.GetContainerImage(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContainerImagesAPI.GetContainerImage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetContainerImage`: ContainerImageResponseBody
	fmt.Fprintf(os.Stdout, "Response from `ContainerImagesAPI.GetContainerImage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Database ID of the stored object. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetContainerImageRequest struct via the builder pattern


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


## ListContainerImageVersions

> VersionResponseBody ListContainerImageVersions(ctx, id).Execute()

List versions for a specific container image.



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
	id := int64(1) // int64 | Database ID of the stored object.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContainerImagesAPI.ListContainerImageVersions(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContainerImagesAPI.ListContainerImageVersions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListContainerImageVersions`: VersionResponseBody
	fmt.Fprintf(os.Stdout, "Response from `ContainerImagesAPI.ListContainerImageVersions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Database ID of the stored object. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListContainerImageVersionsRequest struct via the builder pattern


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


## ListContainerImages

> ContainerImageResponseBody ListContainerImages(ctx).Execute()

List all container image information.



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
	resp, r, err := apiClient.ContainerImagesAPI.ListContainerImages(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContainerImagesAPI.ListContainerImages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListContainerImages`: ContainerImageResponseBody
	fmt.Fprintf(os.Stdout, "Response from `ContainerImagesAPI.ListContainerImages`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListContainerImagesRequest struct via the builder pattern


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

