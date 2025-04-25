# \HealthAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Healthz**](HealthAPI.md#Healthz) | **Get** /healthz | Get health information about the service.



## Healthz

> Health Healthz(ctx).Execute()

Get health information about the service.



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
	resp, r, err := apiClient.HealthAPI.Healthz(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HealthAPI.Healthz``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Healthz`: Health
	fmt.Fprintf(os.Stdout, "Response from `HealthAPI.Healthz`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiHealthzRequest struct via the builder pattern


### Return type

[**Health**](Health.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

