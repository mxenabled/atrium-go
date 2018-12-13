# \IdentityAPI

Method | HTTP request | Description
------------- | ------------- | -------------
[**IdentifyMember**](IdentityAPI.md#IdentifyMember) | **Post** /users/{user_guid}/members/{member_guid}/identify | Identify
[**ListAccountOwners**](IdentityAPI.md#ListAccountOwners) | **Get** /users/{user_guid}/members/{member_guid}/account_owners | List member account owners


# **IdentifyMember**
> Member IdentifyMember(ctx, memberGUID, userGUID)
Identify

The identify endpoint begins an identification process for an already-existing member.

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
  
  memberGUID := "memberGUID_example" // string | The unique identifier for a `member`.
  userGUID := "userGUID_example" // string | The unique identifier for a `user`.

  response, _, err := client.IdentityAPI.IdentifyMember(ctx, memberGUID, userGUID)
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

### Return type

[**Member**](Member.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListAccountOwners**
> AccountOwners ListAccountOwners(ctx, memberGUID, userGUID)
List member account owners

This endpoint returns an array with information about every account associated with a particular member.

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
  
  memberGUID := "memberGUID_example" // string | The unique identifier for a `member`.
  userGUID := "userGUID_example" // string | The unique identifier for a `user`.

  response, _, err := client.IdentityAPI.ListAccountOwners(ctx, memberGUID, userGUID)
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

### Return type

[**AccountOwners**](AccountOwners.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

