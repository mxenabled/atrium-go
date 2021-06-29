# \MerchantsApi

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListMerchantLocations**](MerchantsApi.md#ListMerchantLocations) | **Get** /merchants/{merchant_guid}/merchant_locations | List merchant locations
[**ListMerchants**](MerchantsApi.md#ListMerchants) | **Get** /merchants | List merchants
[**ReadMerchant**](MerchantsApi.md#ReadMerchant) | **Get** /merchants/{merchant_guid} | Read merchant
[**ReadMerchantLocation**](MerchantsApi.md#ReadMerchantLocation) | **Get** /merchants/{merchant_guid}/merchant_locations/{merchant_location_guid} | Read merchant location


# **ListMerchantLocations**
> MerchantLocationsResponseBody ListMerchantLocations(ctx, merchantGUID)
List merchant locations

Returns a list of all the merchant locations associated with a merchant, including physical location, latitude, longitude, etc.

### Example
```go
package main

import (
  "context"
  "fmt"
  "github.com/mxenabled/atrium-go/v2"
)

func main() {
  client := atrium.AtriumClient("YOUR_API_KEY", "YOUR_CLIENT_ID")
  ctx := context.Background()
  
  merchantGUID := "MCH-123" // string | The unique identifier for a `merchant`.

  response, _, err := client.Merchants.ListMerchantLocations(ctx, merchantGUID, )
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

[**MerchantLocationsResponseBody**](MerchantLocationsResponseBody.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListMerchants**
> MerchantsResponseBody ListMerchants(ctx, )
List merchants

Returns a list of merchnants.

### Example
```go
package main

import (
  "context"
  "fmt"
  "github.com/mxenabled/atrium-go/v2"
)

func main() {
  client := atrium.AtriumClient("YOUR_API_KEY", "YOUR_CLIENT_ID")
  ctx := context.Background()

  response, _, err := client.Merchants.ListMerchants
  if err != nil {
    fmt.Printf("Error: %v\n", err)
  } else {
    fmt.Printf("Response: %s\n", response)
  }
}
```

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**MerchantsResponseBody**](MerchantsResponseBody.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

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
  "github.com/mxenabled/atrium-go/v2"
)

func main() {
  client := atrium.AtriumClient("YOUR_API_KEY", "YOUR_CLIENT_ID")
  ctx := context.Background()
  
  merchantGUID := "MCH-123" // string | The unique identifier for a `merchant`.

  response, _, err := client.Merchants.ReadMerchant(ctx, merchantGUID, )
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

# **ReadMerchantLocation**
> MerchantLocationResponseBody ReadMerchantLocation(ctx, merchantGUID, merchantLocationGUID)
Read merchant location

Retuns a specific location associated with a merchant, including physical location, latitude, longitude, etc.

### Example
```go
package main

import (
  "context"
  "fmt"
  "github.com/mxenabled/atrium-go/v2"
)

func main() {
  client := atrium.AtriumClient("YOUR_API_KEY", "YOUR_CLIENT_ID")
  ctx := context.Background()
  
  merchantGUID := "MCH-123" // string | The unique identifier for a `merchant`.
  merchantLocationGUID := "MCL-123" // string | The unique identifier for a `merchant_location`.

  response, _, err := client.Merchants.ReadMerchantLocation(ctx, merchantGUID, merchantLocationGUID)
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
  **merchantLocationGUID** | **string**| The unique identifier for a &#x60;merchant_location&#x60;. | 

### Return type

[**MerchantLocationResponseBody**](MerchantLocationResponseBody.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

