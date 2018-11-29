# \VerificationApi

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListAccountNumbers**](VerificationApi.md#ListAccountNumbers) | **Get** /users/{user_guid}/members/{member_guid}/account_numbers | Read account numbers
[**ListAccountNumbersByAccount**](VerificationApi.md#ListAccountNumbersByAccount) | **Get** /users/{user_guid}/accounts/{account_guid}/account_numbers | Read account numbers by account GUID
[**VerifyMember**](VerificationApi.md#VerifyMember) | **Post** /users/{user_guid}/members/{member_guid}/verify | Verify


# **ListAccountNumbers**
> AccountNumbers ListAccountNumbers(ctx, memberGuid, userGuid)
Read account numbers

Use this endpoint to check whether account and routing numbers are available for accounts associated with a particular member. It returns the account_numbers object, which contains account and routing number data for each account associated with the member.

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

  response, _, err := client.VerificationApi.ListAccountNumbers(nil, memberGuid, userGuid)
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

[**AccountNumbers**](AccountNumbers.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListAccountNumbersByAccount**
> AccountNumbers ListAccountNumbersByAccount(ctx, accountGuid, userGuid)
Read account numbers by account GUID

Use this endpoint to check whether account and routing numbers are available for a specific account. It returns the account_numbers object, which contains account and routing number data.

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
  
  accountGuid := "accountGuid_example" // string | The unique identifier for an `account`.
  userGuid := "userGuid_example" // string | The unique identifier for a `user`.

  response, _, err := client.VerificationApi.ListAccountNumbersByAccount(nil, accountGuid, userGuid)
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
  **accountGuid** | **string**| The unique identifier for an &#x60;account&#x60;. | 
  **userGuid** | **string**| The unique identifier for a &#x60;user&#x60;. | 

### Return type

[**AccountNumbers**](AccountNumbers.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VerifyMember**
> Member VerifyMember(ctx, memberGuid, userGuid)
Verify

The verify endpoint begins a verification process for a member.

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

  response, _, err := client.VerificationApi.VerifyMember(nil, memberGuid, userGuid)
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

