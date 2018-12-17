# \AccountsApi

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListAccountTransactions**](AccountsApi.md#ListAccountTransactions) | **Get** /users/{user_guid}/accounts/{account_guid}/transactions | List account transactions
[**ListUserAccounts**](AccountsApi.md#ListUserAccounts) | **Get** /users/{user_guid}/accounts | List accounts for a user
[**ReadAccount**](AccountsApi.md#ReadAccount) | **Get** /users/{user_guid}/accounts/{account_guid} | Read an account
[**ReadAccountByMemberGUID**](AccountsApi.md#ReadAccountByMemberGUID) | **Get** /users/{user_guid}/members/{member_guid}/accounts/{account_guid} | Read an account


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
  client := atrium.AtriumClient("YOUR_API_KEY", "YOUR_CLIENT_ID")
  ctx := context.Background()
  
  accountGUID := "ACT-123" // string | The unique identifier for an `account`.
  userGUID := "USR-123" // string | The unique identifier for a `user`.
  opts := &atrium.ListAccountTransactionsOpts{ 
    FromDate: optional.NewString("2016-09-20"), // string | Filter transactions from this date.
    ToDate: optional.NewString("2016-10-20"), // string | Filter transactions to this date.
    Page: optional.NewInt32(1), // int32 | Specify current page.
    RecordsPerPage: optional.NewInt32(12), // int32 | Specify records per page.
  }

  response, _, err := client.Accounts.ListAccountTransactions(ctx, accountGUID, userGUID, opts)
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
  client := atrium.AtriumClient("YOUR_API_KEY", "YOUR_CLIENT_ID")
  ctx := context.Background()
  
  userGUID := "USR-123" // string | The unique identifier for a `user`.
  opts := &atrium.ListUserAccountsOpts{ 
    Page: optional.NewInt32(1), // int32 | Specify current page.
    RecordsPerPage: optional.NewInt32(12), // int32 | Specify records per page.
  }

  response, _, err := client.Accounts.ListUserAccounts(ctx, userGUID, opts)
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
  client := atrium.AtriumClient("YOUR_API_KEY", "YOUR_CLIENT_ID")
  ctx := context.Background()
  
  accountGUID := "ACT-123" // string | The unique identifier for an `account`.
  userGUID := "USR-123" // string | The unique identifier for a `user`.

  response, _, err := client.Accounts.ReadAccount(ctx, accountGUID, userGUID)
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
  client := atrium.AtriumClient("YOUR_API_KEY", "YOUR_CLIENT_ID")
  ctx := context.Background()
  
  accountGUID := "ACT-123" // string | The unique identifier for an `account`.
  memberGUID := "MBR-123" // string | The unique identifier for a `member`.
  userGUID := "USR-123" // string | The unique identifier for a `user`.

  response, _, err := client.Accounts.ReadAccountByMemberGUID(ctx, accountGUID, memberGUID, userGUID)
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

