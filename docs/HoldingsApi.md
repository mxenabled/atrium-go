# \HoldingsApi

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListHoldings**](HoldingsApi.md#ListHoldings) | **Get** /users/{user_guid}/holdings | List holdings
[**ListHoldingsByAccount**](HoldingsApi.md#ListHoldingsByAccount) | **Get** /users/{user_guid}/accounts/{account_guid}/holdings | List holdings by account
[**ListHoldingsByMember**](HoldingsApi.md#ListHoldingsByMember) | **Get** /users/{user_guid}/members/{member_guid}/holdings | List holdings by member
[**ReadHolding**](HoldingsApi.md#ReadHolding) | **Get** /users/{user_guid}/holdings/{holding_guid} | Read holding


# **ListHoldings**
> HoldingsResponseBody ListHoldings(ctx, userGUID, optional)
List holdings

Use this endpoint to read all holdings associated with a specific user.

### Example
```go
package main

import (
  "context"
  "fmt"
  "github.com/mxenabled/atrium-go/v2"
  "github.com/antihax/optional"
)

func main() {
  client := atrium.AtriumClient("YOUR_API_KEY", "YOUR_CLIENT_ID")
  ctx := context.Background()
  
  userGUID := "USR-123" // string | The unique identifier for a `user`.
  opts := &atrium.ListHoldingsOpts{ 
    Page: optional.NewInt32(1), // int32 | Specify current page.
    RecordsPerPage: optional.NewInt32(12), // int32 | Specify records per page.
  }

  response, _, err := client.Holdings.ListHoldings(ctx, userGUID, , opts)
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
 **optional** | ***ListHoldingsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListHoldingsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Int32**| Specify current page. | 
 **recordsPerPage** | **optional.Int32**| Specify records per page. | 

### Return type

[**HoldingsResponseBody**](HoldingsResponseBody.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListHoldingsByAccount**
> HoldingsResponseBody ListHoldingsByAccount(ctx, accountGUID, userGUID, optional)
List holdings by account

Use this endpoint to read all holdings associated with a specific account.

### Example
```go
package main

import (
  "context"
  "fmt"
  "github.com/mxenabled/atrium-go/v2"
  "github.com/antihax/optional"
)

func main() {
  client := atrium.AtriumClient("YOUR_API_KEY", "YOUR_CLIENT_ID")
  ctx := context.Background()
  
  accountGUID := "ACT-123" // string | The unique identifier for an `account`.
  userGUID := "USR-123" // string | The unique identifier for a `user`.
  opts := &atrium.ListHoldingsByAccountOpts{ 
    Page: optional.NewInt32(1), // int32 | Specify current page.
    RecordsPerPage: optional.NewInt32(12), // int32 | Specify records per page.
  }

  response, _, err := client.Holdings.ListHoldingsByAccount(ctx, accountGUID, userGUID, , opts)
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
  **accountGUID** | **string**| The unique identifier for an &#x60;account&#x60;. | 
  **userGUID** | **string**| The unique identifier for a &#x60;user&#x60;. | 
 **optional** | ***ListHoldingsByAccountOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListHoldingsByAccountOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **optional.Int32**| Specify current page. | 
 **recordsPerPage** | **optional.Int32**| Specify records per page. | 

### Return type

[**HoldingsResponseBody**](HoldingsResponseBody.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListHoldingsByMember**
> HoldingsResponseBody ListHoldingsByMember(ctx, memberGUID, userGUID, optional)
List holdings by member

Use this endpoint to read all holdings associated with a specific member.

### Example
```go
package main

import (
  "context"
  "fmt"
  "github.com/mxenabled/atrium-go/v2"
  "github.com/antihax/optional"
)

func main() {
  client := atrium.AtriumClient("YOUR_API_KEY", "YOUR_CLIENT_ID")
  ctx := context.Background()
  
  memberGUID := "MBR-123" // string | The unique identifier for a `member`.
  userGUID := "USR-123" // string | The unique identifier for a `user`.
  opts := &atrium.ListHoldingsByMemberOpts{ 
    Page: optional.NewInt32(1), // int32 | Specify current page.
    RecordsPerPage: optional.NewInt32(12), // int32 | Specify records per page.
  }

  response, _, err := client.Holdings.ListHoldingsByMember(ctx, memberGUID, userGUID, , opts)
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
  **memberGUID** | **string**| The unique identifier for a &#x60;member&#x60;. | 
  **userGUID** | **string**| The unique identifier for a &#x60;user&#x60;. | 
 **optional** | ***ListHoldingsByMemberOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListHoldingsByMemberOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **optional.Int32**| Specify current page. | 
 **recordsPerPage** | **optional.Int32**| Specify records per page. | 

### Return type

[**HoldingsResponseBody**](HoldingsResponseBody.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadHolding**
> HoldingResponseBody ReadHolding(ctx, holdingGUID, userGUID)
Read holding

Use this endpoint to read the attributes of a specific holding.

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
  
  holdingGUID := "HOL-123" // string | The unique identifier for a `holding`.
  userGUID := "USR-123" // string | The unique identifier for a `user`.

  response, _, err := client.Holdings.ReadHolding(ctx, holdingGUID, userGUID, )
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
  **holdingGUID** | **string**| The unique identifier for a &#x60;holding&#x60;. | 
  **userGUID** | **string**| The unique identifier for a &#x60;user&#x60;. | 

### Return type

[**HoldingResponseBody**](HoldingResponseBody.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

