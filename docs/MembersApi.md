# \MembersApi

Method | HTTP request | Description
------------- | ------------- | -------------
[**AggregateMember**](MembersApi.md#AggregateMember) | **Post** /users/{user_guid}/members/{member_guid}/aggregate | Aggregate member
[**CreateMember**](MembersApi.md#CreateMember) | **Post** /users/{user_guid}/members | Create member
[**DeleteMember**](MembersApi.md#DeleteMember) | **Delete** /users/{user_guid}/members/{member_guid} | Delete member
[**ListMemberAccounts**](MembersApi.md#ListMemberAccounts) | **Get** /users/{user_guid}/members/{member_guid}/accounts | List member accounts
[**ListMemberCredentials**](MembersApi.md#ListMemberCredentials) | **Get** /users/{user_guid}/members/{member_guid}/credentials | List member credentials
[**ListMemberMFAChallenges**](MembersApi.md#ListMemberMFAChallenges) | **Get** /users/{user_guid}/members/{member_guid}/challenges | List member MFA challenges
[**ListMemberTransactions**](MembersApi.md#ListMemberTransactions) | **Get** /users/{user_guid}/members/{member_guid}/transactions | List member transactions
[**ListMembers**](MembersApi.md#ListMembers) | **Get** /users/{user_guid}/members | List members
[**ReadMember**](MembersApi.md#ReadMember) | **Get** /users/{user_guid}/members/{member_guid} | Read member
[**ReadMemberStatus**](MembersApi.md#ReadMemberStatus) | **Get** /users/{user_guid}/members/{member_guid}/status | Read member connection status
[**ResumeMember**](MembersApi.md#ResumeMember) | **Put** /users/{user_guid}/members/{member_guid}/resume | Resume aggregation from MFA
[**UpdateMember**](MembersApi.md#UpdateMember) | **Put** /users/{user_guid}/members/{member_guid} | Update member


# **AggregateMember**
> Member AggregateMember(ctx, memberGuid, userGuid)
Aggregate member

Calling this endpoint initiates an aggregation event for the member. This brings in the latest account and transaction data from the connected institution. If this data has recently been updated, MX may not initiate an aggregation event. 

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

  response, _, err := client.MembersApi.AggregateMember(nil, memberGuid, userGuid)
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

# **CreateMember**
> Member CreateMember(ctx, userGuid, body)
Create member

This endpoint allows you to create a new member. Members are created with the required parameters credentials and institution_code, and the optional parameters identifier and metadata.<br> When creating a member, you'll need to include the correct type of credential required by the financial institution and provided by the user. You can find out which credential type is required with the /institutions/{institution_code}/credentials endpoint.<br> If successful, Atrium will respond with the newly-created member object.<br> Once you successfully create a member, MX will immediately validate the provided credentials and attempt to aggregate data for accounts and transactions. 

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
  body := atrium-go.MemberCreateRequestBody{} // MemberCreateRequestBody | Member object to be created with optional parameters (identifier and metadata) and required parameters (credentials and institution_code)

  response, _, err := client.MembersApi.CreateMember(nil, userGuidbody)
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
  **body** | [**MemberCreateRequestBody**](MemberCreateRequestBody.md)| Member object to be created with optional parameters (identifier and metadata) and required parameters (credentials and institution_code) | 

### Return type

[**Member**](Member.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteMember**
> DeleteMember(ctx, memberGuid, userGuid)
Delete member

Accessing this endpoint will permanently delete a member.

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

  response, _, err := client.MembersApi.DeleteMember(nil, memberGuid, userGuid)
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

 (empty response body)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListMemberAccounts**
> Accounts ListMemberAccounts(ctx, memberGuid, userGuid, optional)
List member accounts

This endpoint returns an array with information about every account associated with a particular member.

### Example
```go
package main

import (
  "fmt"
  "github.com/mxenabled/atrium-go"
  "github.com/antihax/optional"
)

func main() {
  config := mx.NewConfiguration()
  config.DefaultHeader["MX-Client-ID"] = "YOUR MX-Client-ID"
  config.DefaultHeader["MX-API-Key"] = "YOUR MX-API-Key"

  client := mx.NewAPIClient(config)
  
  memberGuid := "memberGuid_example" // string | The unique identifier for a `member`.
  userGuid := "userGuid_example" // string | The unique identifier for a `user`.
  opts := &mx.ListMemberAccountsOpts{ 
    Page: optional.NewInt32(12), // int32 | Specify current page.
    RecordsPerPage: optional.NewInt32(12), // int32 | Specify records per page.
  }

  response, _, err := client.MembersApi.ListMemberAccounts(nil, memberGuid, userGuid, opts)
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
 **optional** | ***ListMemberAccountsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListMemberAccountsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **optional.Int32**| Specify current page. | 
 **recordsPerPage** | **optional.Int32**| Specify records per page. | 

### Return type

[**Accounts**](Accounts.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListMemberCredentials**
> Credentials ListMemberCredentials(ctx, memberGuid, userGuid)
List member credentials

This endpoint returns an array which contains information on every non-MFA credential associated with a specific member.

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

  response, _, err := client.MembersApi.ListMemberCredentials(nil, memberGuid, userGuid)
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

[**Credentials**](Credentials.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListMemberMFAChallenges**
> Challenges ListMemberMFAChallenges(ctx, memberGuid, userGuid)
List member MFA challenges

Use this endpoint for information on what multi-factor authentication challenges need to be answered in order to aggregate a member.<br> If the aggregation is not challenged, i.e., the member does not have a connection status of CHALLENGED, then code 204 No Content will be returned.<br> If the aggregation has been challenged, i.e., the member does have a connection status of CHALLENGED, then code 200 OK will be returned — along with the corresponding credentials. 

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

  response, _, err := client.MembersApi.ListMemberMFAChallenges(nil, memberGuid, userGuid)
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

[**Challenges**](Challenges.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListMemberTransactions**
> Transactions ListMemberTransactions(ctx, memberGuid, userGuid, optional)
List member transactions

Use this endpoint to get all transactions from all accounts associated with a specific member.<br> This endpoint accepts optional URL query parameters — from_date and to_date — which are used to filter transactions according to the date they were posted. If no values are given for the query parameters, from_date will default to 90 days prior to the request and to_date will default to 5 days from the time of the request. 

### Example
```go
package main

import (
  "fmt"
  "github.com/mxenabled/atrium-go"
  "github.com/antihax/optional"
)

func main() {
  config := mx.NewConfiguration()
  config.DefaultHeader["MX-Client-ID"] = "YOUR MX-Client-ID"
  config.DefaultHeader["MX-API-Key"] = "YOUR MX-API-Key"

  client := mx.NewAPIClient(config)
  
  memberGuid := "memberGuid_example" // string | The unique identifier for a `member`.
  userGuid := "userGuid_example" // string | The unique identifier for a `user`.
  opts := &mx.ListMemberTransactionsOpts{ 
    FromDate: optional.NewString("fromDate_example"), // string | Filter transactions from this date.
    ToDate: optional.NewString("toDate_example"), // string | Filter transactions to this date.
    Page: optional.NewInt32(12), // int32 | Specify current page.
    RecordsPerPage: optional.NewInt32(12), // int32 | Specify records per page.
  }

  response, _, err := client.MembersApi.ListMemberTransactions(nil, memberGuid, userGuid, opts)
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
 **optional** | ***ListMemberTransactionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListMemberTransactionsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fromDate** | **optional.String**| Filter transactions from this date. | 
 **toDate** | **optional.String**| Filter transactions to this date. | 
 **page** | **optional.Int32**| Specify current page. | 
 **recordsPerPage** | **optional.Int32**| Specify records per page. | 

### Return type

[**Transactions**](Transactions.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListMembers**
> Members ListMembers(ctx, userGuid, optional)
List members

This endpoint returns an array which contains information on every member associated with a specific user.

### Example
```go
package main

import (
  "fmt"
  "github.com/mxenabled/atrium-go"
  "github.com/antihax/optional"
)

func main() {
  config := mx.NewConfiguration()
  config.DefaultHeader["MX-Client-ID"] = "YOUR MX-Client-ID"
  config.DefaultHeader["MX-API-Key"] = "YOUR MX-API-Key"

  client := mx.NewAPIClient(config)
  
  userGuid := "userGuid_example" // string | The unique identifier for a `user`.
  opts := &mx.ListMembersOpts{ 
    Page: optional.NewInt32(12), // int32 | Specify current page.
    RecordsPerPage: optional.NewInt32(12), // int32 | Specify records per page.
  }

  response, _, err := client.MembersApi.ListMembers(nil, userGuid, opts)
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
 **optional** | ***ListMembersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListMembersOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Int32**| Specify current page. | 
 **recordsPerPage** | **optional.Int32**| Specify records per page. | 

### Return type

[**Members**](Members.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadMember**
> Member ReadMember(ctx, memberGuid, userGuid)
Read member

Use this endpoint to read the attributes of a specific member.

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

  response, _, err := client.MembersApi.ReadMember(nil, memberGuid, userGuid)
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

# **ReadMemberStatus**
> MemberConnectionStatus ReadMemberStatus(ctx, memberGuid, userGuid)
Read member connection status

This endpoint provides the status of the member's most recent aggregation event. This is an important step in the aggregation process, and the results returned by this endpoint should determine what you do next in order to successfully aggregate a member.<br> MX has introduced new, more detailed information on the current status of a member's connection to a financial institution and the state of its aggregation: the connection_status field. These are intended to replace and expand upon the information provided in the status field, which will soon be deprecated; support for the status field remains for the time being. 

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

  response, _, err := client.MembersApi.ReadMemberStatus(nil, memberGuid, userGuid)
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

[**MemberConnectionStatus**](MemberConnectionStatus.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ResumeMember**
> Member ResumeMember(ctx, memberGuid, userGuid, body)
Resume aggregation from MFA

This endpoint answers the challenges needed when a member has been challenged by multi-factor authentication.

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
  body := atrium-go.MemberResumeRequestBody{} // MemberResumeRequestBody | Member object with MFA challenge answers

  response, _, err := client.MembersApi.ResumeMember(nil, memberGuid, userGuidbody)
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
  **body** | [**MemberResumeRequestBody**](MemberResumeRequestBody.md)| Member object with MFA challenge answers | 

### Return type

[**Member**](Member.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateMember**
> Member UpdateMember(ctx, memberGuid, userGuid, optional)
Update member

Use this endpoint to update a member's attributes. Only the credentials, identifier, and metadata parameters can be updated. To get a list of the required credentials for the member, use the list member credentials endpoint. 

### Example
```go
package main

import (
  "fmt"
  "github.com/mxenabled/atrium-go"
  "github.com/antihax/optional"
)

func main() {
  config := mx.NewConfiguration()
  config.DefaultHeader["MX-Client-ID"] = "YOUR MX-Client-ID"
  config.DefaultHeader["MX-API-Key"] = "YOUR MX-API-Key"

  client := mx.NewAPIClient(config)
  
  memberGuid := "memberGuid_example" // string | The unique identifier for a `member`.
  userGuid := "userGuid_example" // string | The unique identifier for a `user`.
  opts := &mx.UpdateMemberOpts{ 
    Body: optional.NewInterface(atrium-go.MemberUpdateRequestBody{}), // MemberUpdateRequestBody | Member object to be updated with optional parameters (credentials, identifier, metadata)
  }

  response, _, err := client.MembersApi.UpdateMember(nil, memberGuid, userGuid, opts)
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
 **optional** | ***UpdateMemberOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UpdateMemberOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of MemberUpdateRequestBody**](MemberUpdateRequestBody.md)| Member object to be updated with optional parameters (credentials, identifier, metadata) | 

### Return type

[**Member**](Member.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

