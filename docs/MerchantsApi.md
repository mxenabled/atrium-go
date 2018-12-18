# \MerchantsApi

Method | HTTP request | Description
------------- | ------------- | -------------
[**ReadMerchant**](MerchantsApi.md#ReadMerchant) | **Get** /merchants/{merchant_guid} | Read merchant


# **ReadMerchant**
> MerchantResponseBody ReadMerchant(ctx, merchantGUID)
Read merchant

Returns information about a particular merchant, such as a logo, name, and website.

### Example
```go
package main

import (
  "context"
  "fmt"
  "github.com/mxenabled/atrium-go"
)

func main() {
  client := atrium.AtriumClient("YOUR_API_KEY", "YOUR_CLIENT_ID")
  ctx := context.Background()
  
  merchantGUID := "MCH-123" // string | The unique identifier for a `merchant`.

  response, _, err := client.Merchants.ReadMerchant(ctx, merchantGUID)
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
  **merchantGUID** | **string**| The unique identifier for a &#x60;merchant&#x60;. | 

### Return type

[**MerchantResponseBody**](MerchantResponseBody.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

