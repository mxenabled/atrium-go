# \ConnectWidgetAPI

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetConnectWidget**](ConnectWidgetAPI.md#GetConnectWidget) | **Post** /users/{user_guid}/connect_widget_url | Embedding in a website


# **GetConnectWidget**
> ConnectWidget GetConnectWidget(ctx, userGUID, body)
Embedding in a website

This endpoint will return a URL for an embeddable version of MX Connect.

### Example
```go
package main

import (
  "context"
  "fmt"
  "github.com/mxenabled/atrium-go"
)

func main() {
  client := atrium.NewAPIClient("YOUR_API_KEY", "YOUR_CLIENT_ID")
  ctx := context.Background()
  
  userGUID := "userGUID_example" // string | The unique identifier for a `user`.
  body := atrium.ConnectWidgetRequestBody{} // ConnectWidgetRequestBody | Optional config options for WebView (is_mobile_webview, current_institution_code, current_member_guid, update_credentials)

  response, _, err := client.ConnectWidgetAPI.GetConnectWidget(ctx, userGUIDbody)
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
  **userGUID** | **string**| The unique identifier for a &#x60;user&#x60;. | 
  **body** | [**ConnectWidgetRequestBody**](ConnectWidgetRequestBody.md)| Optional config options for WebView (is_mobile_webview, current_institution_code, current_member_guid, update_credentials) | 

### Return type

[**ConnectWidget**](ConnectWidget.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

