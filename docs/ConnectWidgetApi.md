# \ConnectWidgetApi

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetConnectWidget**](ConnectWidgetApi.md#GetConnectWidget) | **Post** /users/{user_guid}/connect_widget_url | Embedding in a website


# **GetConnectWidget**
> ConnectWidget GetConnectWidget(ctx, userGuid, body)
Embedding in a website

This endpoint will return a URL for an embeddable version of MX Connect.

### Example
```go
package main

import (
  "fmt"
  "github.com/mxenabled/atrium-go"
)

func main() {
  config := mx.NewConfiguration()
  config.DefaultHeader["MX-Client-ID"] = "YOUR MX-Client-ID"
  config.DefaultHeader["MX-API-Key"] = "YOUR MX-API-Key"

  client := mx.NewAPIClient(config)
  
  userGuid := "userGuid_example" // string | The unique identifier for a `user`.
  body := atrium-go.ConnectWidgetRequestBody{} // ConnectWidgetRequestBody | Optional config options for WebView (is_mobile_webview, current_institution_code, current_member_guid, update_credentials)

  response, _, err := client.ConnectWidgetApi.GetConnectWidget(nil, userGuidbody)
  if err != nil {
    fmt.Printf("Error: %v\n", err)
  } else {
    fmt.Printf("Response: %s\n", response)
  }
}
```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userGuid** | **string**| The unique identifier for a &#x60;user&#x60;. | 
  **body** | [**ConnectWidgetRequestBody**](ConnectWidgetRequestBody.md)| Optional config options for WebView (is_mobile_webview, current_institution_code, current_member_guid, update_credentials) | 

### Return type

[**ConnectWidget**](ConnectWidget.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

