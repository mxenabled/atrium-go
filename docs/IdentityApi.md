# \IdentityApi

Method | HTTP request | Description
------------- | ------------- | -------------
[**IdentifyMember**](IdentityApi.md#IdentifyMember) | **Post** /users/{user_guid}/members/{member_guid}/identify | Identify
[**ListAccountOwners**](IdentityApi.md#ListAccountOwners) | **Get** /users/{user_guid}/members/{member_guid}/account_owners | List member account owners


# **IdentifyMember**
> Member IdentifyMember(ctx, memberGuid, userGuid)
Identify

The identify endpoint begins an identification process for an already-existing member.

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
  
  memberGuid := "memberGuid_example" // string | The unique identifier for a `member`.
  userGuid := "userGuid_example" // string | The unique identifier for a `user`.

  response, _, err := client.IdentityApi.IdentifyMember(nil, memberGuid, userGuid)
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
  **memberGuid** | **string**| The unique identifier for a &#x60;member&#x60;. | 
  **userGuid** | **string**| The unique identifier for a &#x60;user&#x60;. | 

### Return type

[**Member**](Member.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListAccountOwners**
> AccountOwners ListAccountOwners(ctx, memberGuid, userGuid)
List member account owners

This endpoint returns an array with information about every account associated with a particular member.

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
  
  memberGuid := "memberGuid_example" // string | The unique identifier for a `member`.
  userGuid := "userGuid_example" // string | The unique identifier for a `user`.

  response, _, err := client.IdentityApi.ListAccountOwners(nil, memberGuid, userGuid)
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
  **memberGuid** | **string**| The unique identifier for a &#x60;member&#x60;. | 
  **userGuid** | **string**| The unique identifier for a &#x60;user&#x60;. | 

### Return type

[**AccountOwners**](AccountOwners.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

