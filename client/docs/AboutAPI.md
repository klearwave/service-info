# \AboutAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAbout**](AboutAPI.md#GetAbout) | **Get** /api/about | Get overall information about the service.



## GetAbout

> About GetAbout(ctx).Execute()

Get overall information about the service.



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
	resp, r, err := apiClient.AboutAPI.GetAbout(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AboutAPI.GetAbout``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAbout`: About
	fmt.Fprintf(os.Stdout, "Response from `AboutAPI.GetAbout`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAboutRequest struct via the builder pattern


### Return type

[**About**](About.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

