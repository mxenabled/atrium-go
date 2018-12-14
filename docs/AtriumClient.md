# \AtriumClient

Method | HTTP request | Description
------------- | ------------- | -------------
[**AggregateMember**](AtriumClient.md#AggregateMember) | **Post** /users/{user_guid}/members/{member_guid}/aggregate | Aggregate member
[**CleanseAndCategorizeTransactions**](AtriumClient.md#CleanseAndCategorizeTransactions) | **Post** /cleanse_and_categorize | Categorize transactions
[**CreateMember**](AtriumClient.md#CreateMember) | **Post** /users/{user_guid}/members | Create member
[**CreateUser**](AtriumClient.md#CreateUser) | **Post** /users | Create user
[**DeleteMember**](AtriumClient.md#DeleteMember) | **Delete** /users/{user_guid}/members/{member_guid} | Delete member
[**DeleteUser**](AtriumClient.md#DeleteUser) | **Delete** /users/{user_guid} | Delete user
[**GetConnectWidget**](AtriumClient.md#GetConnectWidget) | **Post** /users/{user_guid}/connect_widget_url | Embedding in a website
[**IdentifyMember**](AtriumClient.md#IdentifyMember) | **Post** /users/{user_guid}/members/{member_guid}/identify | Identify
[**ListAccountNumbers**](AtriumClient.md#ListAccountNumbers) | **Get** /users/{user_guid}/members/{member_guid}/account_numbers | Read account numbers
[**ListAccountNumbersByAccount**](AtriumClient.md#ListAccountNumbersByAccount) | **Get** /users/{user_guid}/accounts/{account_guid}/account_numbers | Read account numbers by account GUID
[**ListAccountOwners**](AtriumClient.md#ListAccountOwners) | **Get** /users/{user_guid}/members/{member_guid}/account_owners | List member account owners
[**ListAccountTransactions**](AtriumClient.md#ListAccountTransactions) | **Get** /users/{user_guid}/accounts/{account_guid}/transactions | List account transactions
[**ListInstitutions**](AtriumClient.md#ListInstitutions) | **Get** /institutions | List institutions
[**ListMemberAccounts**](AtriumClient.md#ListMemberAccounts) | **Get** /users/{user_guid}/members/{member_guid}/accounts | List member accounts
[**ListMemberCredentials**](AtriumClient.md#ListMemberCredentials) | **Get** /users/{user_guid}/members/{member_guid}/credentials | List member credentials
[**ListMemberMFAChallenges**](AtriumClient.md#ListMemberMFAChallenges) | **Get** /users/{user_guid}/members/{member_guid}/challenges | List member MFA challenges
[**ListMemberTransactions**](AtriumClient.md#ListMemberTransactions) | **Get** /users/{user_guid}/members/{member_guid}/transactions | List member transactions
[**ListMembers**](AtriumClient.md#ListMembers) | **Get** /users/{user_guid}/members | List members
[**ListUserAccounts**](AtriumClient.md#ListUserAccounts) | **Get** /users/{user_guid}/accounts | List accounts for a user
[**ListUserTransactions**](AtriumClient.md#ListUserTransactions) | **Get** /users/{user_guid}/transactions | List transactions for a user
[**ListUsers**](AtriumClient.md#ListUsers) | **Get** /users | List users
[**ReadAccount**](AtriumClient.md#ReadAccount) | **Get** /users/{user_guid}/accounts/{account_guid} | Read an account
[**ReadAccountByMemberGUID**](AtriumClient.md#ReadAccountByMemberGUID) | **Get** /users/{user_guid}/members/{member_guid}/accounts/{account_guid} | Read an account
[**ReadInstitution**](AtriumClient.md#ReadInstitution) | **Get** /institutions/{institution_code} | Read institution
[**ReadInstitutionCredentials**](AtriumClient.md#ReadInstitutionCredentials) | **Get** /institutions/{institution_code}/credentials | Read institution credentials
[**ReadMember**](AtriumClient.md#ReadMember) | **Get** /users/{user_guid}/members/{member_guid} | Read member
[**ReadMemberStatus**](AtriumClient.md#ReadMemberStatus) | **Get** /users/{user_guid}/members/{member_guid}/status | Read member connection status
[**ReadTransaction**](AtriumClient.md#ReadTransaction) | **Get** /users/{user_guid}/transactions/{transaction_guid} | Read a transaction
[**ReadUser**](AtriumClient.md#ReadUser) | **Get** /users/{user_guid} | Read user
[**ResumeMember**](AtriumClient.md#ResumeMember) | **Put** /users/{user_guid}/members/{member_guid}/resume | Resume aggregation from MFA
[**UpdateMember**](AtriumClient.md#UpdateMember) | **Put** /users/{user_guid}/members/{member_guid} | Update member
[**UpdateUser**](AtriumClient.md#UpdateUser) | **Put** /users/{user_guid} | Update user
[**VerifyMember**](AtriumClient.md#VerifyMember) | **Post** /users/{user_guid}/members/{member_guid}/verify | Verify


# **AggregateMember**
> MemberResponseBody AggregateMember(ctx, memberGUID, userGUID)
Aggregate member

Calling this endpoint initiates an aggregation event for the member. This brings in the latest account and transaction data from the connected institution. If this data has recently been updated, MX may not initiate an aggregation event. 

### Example
```go
package main

import (
  "context"
  "fmt"
  "github.com/mxenabled/atrium-go"
)

func main() {
  client := atrium.NewAtriumClient("YOUR_API_KEY", "YOUR_CLIENT_ID")
  ctx := context.Background()
  
  memberGUID := "memberGUID_example" // string | The unique identifier for a `member`.
  userGUID := "userGUID_example" // string | The unique identifier for a `user`.

  response, _, err := client.AggregateMember(ctx, memberGUID, userGUID)
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

[**MemberResponseBody**](MemberResponseBody.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CleanseAndCategorizeTransactions**
> TransactionsCleanseAndCategorizeResponseBody CleanseAndCategorizeTransactions(ctx, body)
Categorize transactions

Use this endpoint to categorize, cleanse, and classify transactions. These transactions are not persisted or stored on the MX platform.

### Example
```go
package main

import (
  "context"
  "fmt"
  "github.com/mxenabled/atrium-go"
)

func main() {
  client := atrium.NewAtriumClient("YOUR_API_KEY", "YOUR_CLIENT_ID")
  ctx := context.Background()
  
  body := atrium.TransactionsCleanseAndCategorizeRequestBody{} // TransactionsCleanseAndCategorizeRequestBody | User object to be created with optional parameters (amount, type) amd required parameters (description, identifier)

  response, _, err := client.CleanseAndCategorizeTransactions(ctx, body)
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
  **body** | [**TransactionsCleanseAndCategorizeRequestBody**](TransactionsCleanseAndCategorizeRequestBody.md)| User object to be created with optional parameters (amount, type) amd required parameters (description, identifier) | 

### Return type

[**TransactionsCleanseAndCategorizeResponseBody**](TransactionsCleanseAndCategorizeResponseBody.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateMember**
> MemberResponseBody CreateMember(ctx, userGUID, body)
Create member

This endpoint allows you to create a new member. Members are created with the required parameters credentials and institution_code, and the optional parameters identifier and metadata.<br> When creating a member, you'll need to include the correct type of credential required by the financial institution and provided by the user. You can find out which credential type is required with the /institutions/{institution_code}/credentials endpoint.<br> If successful, Atrium will respond with the newly-created member object.<br> Once you successfully create a member, MX will immediately validate the provided credentials and attempt to aggregate data for accounts and transactions. 

### Example
```go
package main

import (
  "context"
  "fmt"
  "github.com/mxenabled/atrium-go"
)

func main() {
  client := atrium.NewAtriumClient("YOUR_API_KEY", "YOUR_CLIENT_ID")
  ctx := context.Background()
  
  userGUID := "userGUID_example" // string | The unique identifier for a `user`.
  body := atrium.MemberCreateRequestBody{} // MemberCreateRequestBody | Member object to be created with optional parameters (identifier and metadata) and required parameters (credentials and institution_code)

  response, _, err := client.CreateMember(ctx, userGUIDbody)
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
  **body** | [**MemberCreateRequestBody**](MemberCreateRequestBody.md)| Member object to be created with optional parameters (identifier and metadata) and required parameters (credentials and institution_code) | 

### Return type

[**MemberResponseBody**](MemberResponseBody.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateUser**
> UserResponseBody CreateUser(ctx, body)
Create user

Call this endpoint to create a new user. Atrium will respond with the newly-created user object if successful. This endpoint accepts several parameters: identifier, metadata, and is_disabled.<br> Disabling a user means that accounts and transactions associated with it will not be updated in the background by MX. It will also restrict access to that user's data until they are no longer disabled. Users who are disabled for the entirety of an Atrium billing period will not be factored into that month's bill. 

### Example
```go
package main

import (
  "context"
  "fmt"
  "github.com/mxenabled/atrium-go"
)

func main() {
  client := atrium.NewAtriumClient("YOUR_API_KEY", "YOUR_CLIENT_ID")
  ctx := context.Background()
  
  body := atrium.UserCreateRequestBody{} // UserCreateRequestBody | User object to be created with optional parameters (identifier, is_disabled, metadata)

  response, _, err := client.CreateUser(ctx, body)
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
  **body** | [**UserCreateRequestBody**](UserCreateRequestBody.md)| User object to be created with optional parameters (identifier, is_disabled, metadata) | 

### Return type

[**UserResponseBody**](UserResponseBody.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteMember**
> DeleteMember(ctx, memberGUID, userGUID)
Delete member

Accessing this endpoint will permanently delete a member.

### Example
```go
package main

import (
  "context"
  "fmt"
  "github.com/mxenabled/atrium-go"
)

func main() {
  client := atrium.NewAtriumClient("YOUR_API_KEY", "YOUR_CLIENT_ID")
  ctx := context.Background()
  
  memberGUID := "memberGUID_example" // string | The unique identifier for a `member`.
  userGUID := "userGUID_example" // string | The unique identifier for a `user`.

  response, _, err := client.DeleteMember(ctx, memberGUID, userGUID)
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

 (empty response body)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteUser**
> DeleteUser(ctx, userGUID)
Delete user

Calling this endpoint will permanently delete a user from Atrium. If successful, the API will respond with Status: 204 No Content. 

### Example
```go
package main

import (
  "context"
  "fmt"
  "github.com/mxenabled/atrium-go"
)

func main() {
  client := atrium.NewAtriumClient("YOUR_API_KEY", "YOUR_CLIENT_ID")
  ctx := context.Background()
  
  userGUID := "userGUID_example" // string | The unique identifier for a `user`.

  response, _, err := client.DeleteUser(ctx, userGUID)
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

### Return type

 (empty response body)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetConnectWidget**
> ConnectWidgetResponseBody GetConnectWidget(ctx, userGUID, body)
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
  client := atrium.NewAtriumClient("YOUR_API_KEY", "YOUR_CLIENT_ID")
  ctx := context.Background()
  
  userGUID := "userGUID_example" // string | The unique identifier for a `user`.
  body := atrium.ConnectWidgetRequestBody{} // ConnectWidgetRequestBody | Optional config options for WebView (is_mobile_webview, current_institution_code, current_member_guid, update_credentials)

  response, _, err := client.GetConnectWidget(ctx, userGUIDbody)
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

[**ConnectWidgetResponseBody**](ConnectWidgetResponseBody.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **IdentifyMember**
> MemberResponseBody IdentifyMember(ctx, memberGUID, userGUID)
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
  client := atrium.NewAtriumClient("YOUR_API_KEY", "YOUR_CLIENT_ID")
  ctx := context.Background()
  
  memberGUID := "memberGUID_example" // string | The unique identifier for a `member`.
  userGUID := "userGUID_example" // string | The unique identifier for a `user`.

  response, _, err := client.IdentifyMember(ctx, memberGUID, userGUID)
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

[**MemberResponseBody**](MemberResponseBody.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListAccountNumbers**
> AccountNumbersResponseBody ListAccountNumbers(ctx, memberGUID, userGUID)
Read account numbers

Use this endpoint to check whether account and routing numbers are available for accounts associated with a particular member. It returns the account_numbers object, which contains account and routing number data for each account associated with the member.

### Example
```go
package main

import (
  "context"
  "fmt"
  "github.com/mxenabled/atrium-go"
)

func main() {
  client := atrium.NewAtriumClient("YOUR_API_KEY", "YOUR_CLIENT_ID")
  ctx := context.Background()
  
  memberGUID := "memberGUID_example" // string | The unique identifier for a `member`.
  userGUID := "userGUID_example" // string | The unique identifier for a `user`.

  response, _, err := client.ListAccountNumbers(ctx, memberGUID, userGUID)
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

[**AccountNumbersResponseBody**](AccountNumbersResponseBody.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListAccountNumbersByAccount**
> AccountNumbersResponseBody ListAccountNumbersByAccount(ctx, accountGUID, userGUID)
Read account numbers by account GUID

Use this endpoint to check whether account and routing numbers are available for a specific account. It returns the account_numbers object, which contains account and routing number data.

### Example
```go
package main

import (
  "context"
  "fmt"
  "github.com/mxenabled/atrium-go"
)

func main() {
  client := atrium.NewAtriumClient("YOUR_API_KEY", "YOUR_CLIENT_ID")
  ctx := context.Background()
  
  accountGUID := "accountGUID_example" // string | The unique identifier for an `account`.
  userGUID := "userGUID_example" // string | The unique identifier for a `user`.

  response, _, err := client.ListAccountNumbersByAccount(ctx, accountGUID, userGUID)
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

### Return type

[**AccountNumbersResponseBody**](AccountNumbersResponseBody.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListAccountOwners**
> AccountOwnersResponseBody ListAccountOwners(ctx, memberGUID, userGUID)
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
  client := atrium.NewAtriumClient("YOUR_API_KEY", "YOUR_CLIENT_ID")
  ctx := context.Background()
  
  memberGUID := "memberGUID_example" // string | The unique identifier for a `member`.
  userGUID := "userGUID_example" // string | The unique identifier for a `user`.

  response, _, err := client.ListAccountOwners(ctx, memberGUID, userGUID)
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

[**AccountOwnersResponseBody**](AccountOwnersResponseBody.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListAccountTransactions**
> TransactionsResponseBody ListAccountTransactions(ctx, accountGUID, userGUID, optional)
List account transactions

This endpoint allows you to see every transaction that belongs to a specific account. The default from_date is 90 days prior to the request, and the default to_date is 5 days from the time of the request.<br> The from_date and to_date parameters can optionally be appended to the request. 

### Example
```go
package main

import (
  "context"
  "fmt"
  "github.com/mxenabled/atrium-go"
  "github.com/antihax/optional"
)

func main() {
  client := atrium.NewAtriumClient("YOUR_API_KEY", "YOUR_CLIENT_ID")
  ctx := context.Background()
  
  accountGUID := "accountGUID_example" // string | The unique identifier for an `account`.
  userGUID := "userGUID_example" // string | The unique identifier for a `user`.
  opts := &atrium.ListAccountTransactionsOpts{ 
    FromDate: optional.NewString("fromDate_example"), // string | Filter transactions from this date.
    ToDate: optional.NewString("toDate_example"), // string | Filter transactions to this date.
    Page: optional.NewInt32(12), // int32 | Specify current page.
    RecordsPerPage: optional.NewInt32(12), // int32 | Specify records per page.
  }

  response, _, err := client.ListAccountTransactions(ctx, accountGUID, userGUID, opts)
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
 **optional** | ***ListAccountTransactionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListAccountTransactionsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fromDate** | **optional.String**| Filter transactions from this date. | 
 **toDate** | **optional.String**| Filter transactions to this date. | 
 **page** | **optional.Int32**| Specify current page. | 
 **recordsPerPage** | **optional.Int32**| Specify records per page. | 

### Return type

[**TransactionsResponseBody**](TransactionsResponseBody.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListInstitutions**
> InstitutionsResponseBody ListInstitutions(ctx, optional)
List institutions

This endpoint allows you to see what institutions are available for connection. Information returned will include the institution_code assigned to a particular institution, URLs for the financial institution's logo, and the URL for its website.<br> This endpoint takes an optional query string, name={string}. This will list only institutions in which the appended string appears. 

### Example
```go
package main

import (
  "context"
  "fmt"
  "github.com/mxenabled/atrium-go"
  "github.com/antihax/optional"
)

func main() {
  client := atrium.NewAtriumClient("YOUR_API_KEY", "YOUR_CLIENT_ID")
  ctx := context.Background()
  
  opts := &atrium.ListInstitutionsOpts{ 
    Name: optional.NewString("name_example"), // string | This will list only institutions in which the appended string appears.
    Page: optional.NewInt32(12), // int32 | Specify current page.
    RecordsPerPage: optional.NewInt32(12), // int32 | Specify records per page.
  }

  response, _, err := client.ListInstitutions(ctx, opts)
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
 **optional** | ***ListInstitutionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListInstitutionsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **optional.String**| This will list only institutions in which the appended string appears. | 
 **page** | **optional.Int32**| Specify current page. | 
 **recordsPerPage** | **optional.Int32**| Specify records per page. | 

### Return type

[**InstitutionsResponseBody**](InstitutionsResponseBody.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListMemberAccounts**
> AccountsResponseBody ListMemberAccounts(ctx, memberGUID, userGUID, optional)
List member accounts

This endpoint returns an array with information about every account associated with a particular member.

### Example
```go
package main

import (
  "context"
  "fmt"
  "github.com/mxenabled/atrium-go"
  "github.com/antihax/optional"
)

func main() {
  client := atrium.NewAtriumClient("YOUR_API_KEY", "YOUR_CLIENT_ID")
  ctx := context.Background()
  
  memberGUID := "memberGUID_example" // string | The unique identifier for a `member`.
  userGUID := "userGUID_example" // string | The unique identifier for a `user`.
  opts := &atrium.ListMemberAccountsOpts{ 
    Page: optional.NewInt32(12), // int32 | Specify current page.
    RecordsPerPage: optional.NewInt32(12), // int32 | Specify records per page.
  }

  response, _, err := client.ListMemberAccounts(ctx, memberGUID, userGUID, opts)
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
 **optional** | ***ListMemberAccountsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListMemberAccountsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **optional.Int32**| Specify current page. | 
 **recordsPerPage** | **optional.Int32**| Specify records per page. | 

### Return type

[**AccountsResponseBody**](AccountsResponseBody.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListMemberCredentials**
> CredentialsResponseBody ListMemberCredentials(ctx, memberGUID, userGUID)
List member credentials

This endpoint returns an array which contains information on every non-MFA credential associated with a specific member.

### Example
```go
package main

import (
  "context"
  "fmt"
  "github.com/mxenabled/atrium-go"
)

func main() {
  client := atrium.NewAtriumClient("YOUR_API_KEY", "YOUR_CLIENT_ID")
  ctx := context.Background()
  
  memberGUID := "memberGUID_example" // string | The unique identifier for a `member`.
  userGUID := "userGUID_example" // string | The unique identifier for a `user`.

  response, _, err := client.ListMemberCredentials(ctx, memberGUID, userGUID)
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

[**CredentialsResponseBody**](CredentialsResponseBody.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListMemberMFAChallenges**
> ChallengesResponseBody ListMemberMFAChallenges(ctx, memberGUID, userGUID)
List member MFA challenges

Use this endpoint for information on what multi-factor authentication challenges need to be answered in order to aggregate a member.<br> If the aggregation is not challenged, i.e., the member does not have a connection status of CHALLENGED, then code 204 No Content will be returned.<br> If the aggregation has been challenged, i.e., the member does have a connection status of CHALLENGED, then code 200 OK will be returned — along with the corresponding credentials. 

### Example
```go
package main

import (
  "context"
  "fmt"
  "github.com/mxenabled/atrium-go"
)

func main() {
  client := atrium.NewAtriumClient("YOUR_API_KEY", "YOUR_CLIENT_ID")
  ctx := context.Background()
  
  memberGUID := "memberGUID_example" // string | The unique identifier for a `member`.
  userGUID := "userGUID_example" // string | The unique identifier for a `user`.

  response, _, err := client.ListMemberMFAChallenges(ctx, memberGUID, userGUID)
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

[**ChallengesResponseBody**](ChallengesResponseBody.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListMemberTransactions**
> TransactionsResponseBody ListMemberTransactions(ctx, memberGUID, userGUID, optional)
List member transactions

Use this endpoint to get all transactions from all accounts associated with a specific member.<br> This endpoint accepts optional URL query parameters — from_date and to_date — which are used to filter transactions according to the date they were posted. If no values are given for the query parameters, from_date will default to 90 days prior to the request and to_date will default to 5 days from the time of the request. 

### Example
```go
package main

import (
  "context"
  "fmt"
  "github.com/mxenabled/atrium-go"
  "github.com/antihax/optional"
)

func main() {
  client := atrium.NewAtriumClient("YOUR_API_KEY", "YOUR_CLIENT_ID")
  ctx := context.Background()
  
  memberGUID := "memberGUID_example" // string | The unique identifier for a `member`.
  userGUID := "userGUID_example" // string | The unique identifier for a `user`.
  opts := &atrium.ListMemberTransactionsOpts{ 
    FromDate: optional.NewString("fromDate_example"), // string | Filter transactions from this date.
    ToDate: optional.NewString("toDate_example"), // string | Filter transactions to this date.
    Page: optional.NewInt32(12), // int32 | Specify current page.
    RecordsPerPage: optional.NewInt32(12), // int32 | Specify records per page.
  }

  response, _, err := client.ListMemberTransactions(ctx, memberGUID, userGUID, opts)
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

[**TransactionsResponseBody**](TransactionsResponseBody.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListMembers**
> MembersResponseBody ListMembers(ctx, userGUID, optional)
List members

This endpoint returns an array which contains information on every member associated with a specific user.

### Example
```go
package main

import (
  "context"
  "fmt"
  "github.com/mxenabled/atrium-go"
  "github.com/antihax/optional"
)

func main() {
  client := atrium.NewAtriumClient("YOUR_API_KEY", "YOUR_CLIENT_ID")
  ctx := context.Background()
  
  userGUID := "userGUID_example" // string | The unique identifier for a `user`.
  opts := &atrium.ListMembersOpts{ 
    Page: optional.NewInt32(12), // int32 | Specify current page.
    RecordsPerPage: optional.NewInt32(12), // int32 | Specify records per page.
  }

  response, _, err := client.ListMembers(ctx, userGUID, opts)
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
 **optional** | ***ListMembersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListMembersOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Int32**| Specify current page. | 
 **recordsPerPage** | **optional.Int32**| Specify records per page. | 

### Return type

[**MembersResponseBody**](MembersResponseBody.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListUserAccounts**
> AccountsResponseBody ListUserAccounts(ctx, userGUID, optional)
List accounts for a user

Use this endpoint to view information about every account that belongs to a user. You'll need the user's GUID to access this list. The information will include the account type — e.g., CHECKING, MONEY_MARKET, or PROPERTY — the account balance, the date the account was started, etc.

### Example
```go
package main

import (
  "context"
  "fmt"
  "github.com/mxenabled/atrium-go"
  "github.com/antihax/optional"
)

func main() {
  client := atrium.NewAtriumClient("YOUR_API_KEY", "YOUR_CLIENT_ID")
  ctx := context.Background()
  
  userGUID := "userGUID_example" // string | The unique identifier for a `user`.
  opts := &atrium.ListUserAccountsOpts{ 
    Page: optional.NewInt32(12), // int32 | Specify current page.
    RecordsPerPage: optional.NewInt32(12), // int32 | Specify records per page.
  }

  response, _, err := client.ListUserAccounts(ctx, userGUID, opts)
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
 **optional** | ***ListUserAccountsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListUserAccountsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Int32**| Specify current page. | 
 **recordsPerPage** | **optional.Int32**| Specify records per page. | 

### Return type

[**AccountsResponseBody**](AccountsResponseBody.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListUserTransactions**
> TransactionsResponseBody ListUserTransactions(ctx, userGUID, optional)
List transactions for a user

Use this endpoint to get all transactions that belong to a specific user, across all the user's members and accounts.<br> This endpoint accepts optional query parameters, from_date and to_date, which filter transactions according to the date they were posted. If no values are given, from_date will default to 90 days prior to the request, and to_date will default to 5 days from the time of the request. 

### Example
```go
package main

import (
  "context"
  "fmt"
  "github.com/mxenabled/atrium-go"
  "github.com/antihax/optional"
)

func main() {
  client := atrium.NewAtriumClient("YOUR_API_KEY", "YOUR_CLIENT_ID")
  ctx := context.Background()
  
  userGUID := "userGUID_example" // string | The unique identifier for a `user`.
  opts := &atrium.ListUserTransactionsOpts{ 
    Page: optional.NewInt32(12), // int32 | Specify current page.
    FromDate: optional.NewString("fromDate_example"), // string | Filter transactions from this date.
    RecordsPerPage: optional.NewInt32(12), // int32 | Specify records per page.
    ToDate: optional.NewString("toDate_example"), // string | Filter transactions to this date.
  }

  response, _, err := client.ListUserTransactions(ctx, userGUID, opts)
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
 **optional** | ***ListUserTransactionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListUserTransactionsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Int32**| Specify current page. | 
 **fromDate** | **optional.String**| Filter transactions from this date. | 
 **recordsPerPage** | **optional.Int32**| Specify records per page. | 
 **toDate** | **optional.String**| Filter transactions to this date. | 

### Return type

[**TransactionsResponseBody**](TransactionsResponseBody.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListUsers**
> UsersResponseBody ListUsers(ctx, optional)
List users

Use this endpoint to list every user you've created in Atrium.

### Example
```go
package main

import (
  "context"
  "fmt"
  "github.com/mxenabled/atrium-go"
  "github.com/antihax/optional"
)

func main() {
  client := atrium.NewAtriumClient("YOUR_API_KEY", "YOUR_CLIENT_ID")
  ctx := context.Background()
  
  opts := &atrium.ListUsersOpts{ 
    Page: optional.NewInt32(12), // int32 | Specify current page.
    RecordsPerPage: optional.NewInt32(12), // int32 | Specify records per page.
  }

  response, _, err := client.ListUsers(ctx, opts)
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
 **optional** | ***ListUsersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListUsersOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int32**| Specify current page. | 
 **recordsPerPage** | **optional.Int32**| Specify records per page. | 

### Return type

[**UsersResponseBody**](UsersResponseBody.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadAccount**
> AccountResponseBody ReadAccount(ctx, accountGUID, userGUID)
Read an account

Reading an account allows you to get information about a specific account that belongs to a user. That includes the account type — e.g., CHECKING, MONEY_MARKET, or PROPERTY — the balance, the date the account was started, and much more.<br> There are two endpoints for reading an account. Both will return the same information.<br> It's important to remember that balance and available_balance will normally be positive numbers — for all account types. But this should be interpreted differently for debt accounts and asset accounts.<br> An asset account, e.g., CHECKING, SAVINGS, or INVESTMENT, will have a positive balance unless it is in an overdraft condition, in which case the balance will be negative.<br> On the other hand, a debt account, e.g., CREDIT CARD, LOAN, MORTGAGE, would have a positivebalance when the user owes money on the account. It would have a negative balance if the account has been overpaid. 

### Example
```go
package main

import (
  "context"
  "fmt"
  "github.com/mxenabled/atrium-go"
)

func main() {
  client := atrium.NewAtriumClient("YOUR_API_KEY", "YOUR_CLIENT_ID")
  ctx := context.Background()
  
  accountGUID := "accountGUID_example" // string | The unique identifier for an `account`.
  userGUID := "userGUID_example" // string | The unique identifier for a `user`.

  response, _, err := client.ReadAccount(ctx, accountGUID, userGUID)
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

### Return type

[**AccountResponseBody**](AccountResponseBody.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadAccountByMemberGUID**
> AccountResponseBody ReadAccountByMemberGUID(ctx, accountGUID, memberGUID, userGUID)
Read an account

Reading an account allows you to get information about a specific account that belongs to a user. That includes the account type — e.g., CHECKING, MONEY_MARKET, or PROPERTY — the balance, the date the account was started, and much more.<br> There are two endpoints for reading an account. Both will return the same information.<br> It's important to remember that balance and available_balance will normally be positive numbers — for all account types. But this should be interpreted differently for debt accounts and asset accounts.<br> An asset account, e.g., CHECKING, SAVINGS, or INVESTMENT, will have a positive balance unless it is in an overdraft condition, in which case the balance will be negative.<br> On the other hand, a debt account, e.g., CREDIT CARD, LOAN, MORTGAGE, would have a positivebalance when the user owes money on the account. It would have a negative balance if the account has been overpaid. 

### Example
```go
package main

import (
  "context"
  "fmt"
  "github.com/mxenabled/atrium-go"
)

func main() {
  client := atrium.NewAtriumClient("YOUR_API_KEY", "YOUR_CLIENT_ID")
  ctx := context.Background()
  
  accountGUID := "accountGUID_example" // string | The unique identifier for an `account`.
  memberGUID := "memberGUID_example" // string | The unique identifier for a `member`.
  userGUID := "userGUID_example" // string | The unique identifier for a `user`.

  response, _, err := client.ReadAccountByMemberGUID(ctx, accountGUID, memberGUID, userGUID)
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
  **memberGUID** | **string**| The unique identifier for a &#x60;member&#x60;. | 
  **userGUID** | **string**| The unique identifier for a &#x60;user&#x60;. | 

### Return type

[**AccountResponseBody**](AccountResponseBody.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadInstitution**
> InstitutionResponseBody ReadInstitution(ctx, institutionCode)
Read institution

This endpoint allows you to see information for a specific institution.

### Example
```go
package main

import (
  "context"
  "fmt"
  "github.com/mxenabled/atrium-go"
)

func main() {
  client := atrium.NewAtriumClient("YOUR_API_KEY", "YOUR_CLIENT_ID")
  ctx := context.Background()
  
  institutionCode := "institutionCode_example" // string | The institution_code of the institution.

  response, _, err := client.ReadInstitution(ctx, institutionCode)
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
  **institutionCode** | **string**| The institution_code of the institution. | 

### Return type

[**InstitutionResponseBody**](InstitutionResponseBody.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadInstitutionCredentials**
> CredentialsResponseBody ReadInstitutionCredentials(ctx, institutionCode)
Read institution credentials

Use this endpoint to see which credentials will be needed to create a member for a specific institution.

### Example
```go
package main

import (
  "context"
  "fmt"
  "github.com/mxenabled/atrium-go"
)

func main() {
  client := atrium.NewAtriumClient("YOUR_API_KEY", "YOUR_CLIENT_ID")
  ctx := context.Background()
  
  institutionCode := "institutionCode_example" // string | The institution_code of the institution.

  response, _, err := client.ReadInstitutionCredentials(ctx, institutionCode)
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
  **institutionCode** | **string**| The institution_code of the institution. | 

### Return type

[**CredentialsResponseBody**](CredentialsResponseBody.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadMember**
> MemberResponseBody ReadMember(ctx, memberGUID, userGUID)
Read member

Use this endpoint to read the attributes of a specific member.

### Example
```go
package main

import (
  "context"
  "fmt"
  "github.com/mxenabled/atrium-go"
)

func main() {
  client := atrium.NewAtriumClient("YOUR_API_KEY", "YOUR_CLIENT_ID")
  ctx := context.Background()
  
  memberGUID := "memberGUID_example" // string | The unique identifier for a `member`.
  userGUID := "userGUID_example" // string | The unique identifier for a `user`.

  response, _, err := client.ReadMember(ctx, memberGUID, userGUID)
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

[**MemberResponseBody**](MemberResponseBody.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadMemberStatus**
> MemberConnectionStatusResponseBody ReadMemberStatus(ctx, memberGUID, userGUID)
Read member connection status

This endpoint provides the status of the member's most recent aggregation event. This is an important step in the aggregation process, and the results returned by this endpoint should determine what you do next in order to successfully aggregate a member.<br> MX has introduced new, more detailed information on the current status of a member's connection to a financial institution and the state of its aggregation: the connection_status field. These are intended to replace and expand upon the information provided in the status field, which will soon be deprecated; support for the status field remains for the time being. 

### Example
```go
package main

import (
  "context"
  "fmt"
  "github.com/mxenabled/atrium-go"
)

func main() {
  client := atrium.NewAtriumClient("YOUR_API_KEY", "YOUR_CLIENT_ID")
  ctx := context.Background()
  
  memberGUID := "memberGUID_example" // string | The unique identifier for a `member`.
  userGUID := "userGUID_example" // string | The unique identifier for a `user`.

  response, _, err := client.ReadMemberStatus(ctx, memberGUID, userGUID)
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

[**MemberConnectionStatusResponseBody**](MemberConnectionStatusResponseBody.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadTransaction**
> TransactionResponseBody ReadTransaction(ctx, transactionGUID, userGUID)
Read a transaction

This endpoint allows you to view information about a specific transaction that belongs to a user.<br>

### Example
```go
package main

import (
  "context"
  "fmt"
  "github.com/mxenabled/atrium-go"
)

func main() {
  client := atrium.NewAtriumClient("YOUR_API_KEY", "YOUR_CLIENT_ID")
  ctx := context.Background()
  
  transactionGUID := "transactionGUID_example" // string | The unique identifier for a `transaction`.
  userGUID := "userGUID_example" // string | The unique identifier for a `user`.

  response, _, err := client.ReadTransaction(ctx, transactionGUID, userGUID)
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
  **transactionGUID** | **string**| The unique identifier for a &#x60;transaction&#x60;. | 
  **userGUID** | **string**| The unique identifier for a &#x60;user&#x60;. | 

### Return type

[**TransactionResponseBody**](TransactionResponseBody.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadUser**
> UserResponseBody ReadUser(ctx, userGUID)
Read user

Use this endpoint to read the attributes of a specific user.

### Example
```go
package main

import (
  "context"
  "fmt"
  "github.com/mxenabled/atrium-go"
)

func main() {
  client := atrium.NewAtriumClient("YOUR_API_KEY", "YOUR_CLIENT_ID")
  ctx := context.Background()
  
  userGUID := "userGUID_example" // string | The unique identifier for a `user`.

  response, _, err := client.ReadUser(ctx, userGUID)
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

### Return type

[**UserResponseBody**](UserResponseBody.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ResumeMember**
> MemberResponseBody ResumeMember(ctx, memberGUID, userGUID, body)
Resume aggregation from MFA

This endpoint answers the challenges needed when a member has been challenged by multi-factor authentication.

### Example
```go
package main

import (
  "context"
  "fmt"
  "github.com/mxenabled/atrium-go"
)

func main() {
  client := atrium.NewAtriumClient("YOUR_API_KEY", "YOUR_CLIENT_ID")
  ctx := context.Background()
  
  memberGUID := "memberGUID_example" // string | The unique identifier for a `member`.
  userGUID := "userGUID_example" // string | The unique identifier for a `user`.
  body := atrium.MemberResumeRequestBody{} // MemberResumeRequestBody | Member object with MFA challenge answers

  response, _, err := client.ResumeMember(ctx, memberGUID, userGUIDbody)
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
  **body** | [**MemberResumeRequestBody**](MemberResumeRequestBody.md)| Member object with MFA challenge answers | 

### Return type

[**MemberResponseBody**](MemberResponseBody.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateMember**
> MemberResponseBody UpdateMember(ctx, memberGUID, userGUID, optional)
Update member

Use this endpoint to update a member's attributes. Only the credentials, identifier, and metadata parameters can be updated. To get a list of the required credentials for the member, use the list member credentials endpoint. 

### Example
```go
package main

import (
  "context"
  "fmt"
  "github.com/mxenabled/atrium-go"
  "github.com/antihax/optional"
)

func main() {
  client := atrium.NewAtriumClient("YOUR_API_KEY", "YOUR_CLIENT_ID")
  ctx := context.Background()
  
  memberGUID := "memberGUID_example" // string | The unique identifier for a `member`.
  userGUID := "userGUID_example" // string | The unique identifier for a `user`.
  opts := &atrium.UpdateMemberOpts{ 
    Body: optional.NewInterface(atrium.MemberUpdateRequestBody{}), // MemberUpdateRequestBody | Member object to be updated with optional parameters (credentials, identifier, metadata)
  }

  response, _, err := client.UpdateMember(ctx, memberGUID, userGUID, opts)
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
 **optional** | ***UpdateMemberOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UpdateMemberOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of MemberUpdateRequestBody**](MemberUpdateRequestBody.md)| Member object to be updated with optional parameters (credentials, identifier, metadata) | 

### Return type

[**MemberResponseBody**](MemberResponseBody.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateUser**
> UserResponseBody UpdateUser(ctx, userGUID, optional)
Update user

Use this endpoint to update the attributes of a specific user. Atrium will respond with the updated user object.<br> Disabling a user means that accounts and transactions associated with it will not be updated in the background by MX. It will also restrict access to that user's data until they are no longer disabled. Users who are disabled for the entirety of an Atrium billing period will not be factored into that month's bill.<br> To disable a user, update it and set the is_disabled parameter to true. Set it to false to re-enable the user. 

### Example
```go
package main

import (
  "context"
  "fmt"
  "github.com/mxenabled/atrium-go"
  "github.com/antihax/optional"
)

func main() {
  client := atrium.NewAtriumClient("YOUR_API_KEY", "YOUR_CLIENT_ID")
  ctx := context.Background()
  
  userGUID := "userGUID_example" // string | The unique identifier for a `user`.
  opts := &atrium.UpdateUserOpts{ 
    Body: optional.NewInterface(atrium.UserUpdateRequestBody{}), // UserUpdateRequestBody | User object to be updated with optional parameters (identifier, is_disabled, metadata)
  }

  response, _, err := client.UpdateUser(ctx, userGUID, opts)
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
 **optional** | ***UpdateUserOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UpdateUserOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of UserUpdateRequestBody**](UserUpdateRequestBody.md)| User object to be updated with optional parameters (identifier, is_disabled, metadata) | 

### Return type

[**UserResponseBody**](UserResponseBody.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VerifyMember**
> MemberResponseBody VerifyMember(ctx, memberGUID, userGUID)
Verify

The verify endpoint begins a verification process for a member.

### Example
```go
package main

import (
  "context"
  "fmt"
  "github.com/mxenabled/atrium-go"
)

func main() {
  client := atrium.NewAtriumClient("YOUR_API_KEY", "YOUR_CLIENT_ID")
  ctx := context.Background()
  
  memberGUID := "memberGUID_example" // string | The unique identifier for a `member`.
  userGUID := "userGUID_example" // string | The unique identifier for a `user`.

  response, _, err := client.VerifyMember(ctx, memberGUID, userGUID)
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

[**MemberResponseBody**](MemberResponseBody.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

