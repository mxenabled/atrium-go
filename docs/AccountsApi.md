# \AccountsApi

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListAccountTransactions**](AccountsApi.md#ListAccountTransactions) | **Get** /users/{user_guid}/accounts/{account_guid}/transactions | List account transactions
[**ListUserAccounts**](AccountsApi.md#ListUserAccounts) | **Get** /users/{user_guid}/accounts | List accounts for a user
[**ReadAccount**](AccountsApi.md#ReadAccount) | **Get** /users/{user_guid}/accounts/{account_guid} | Read an account
[**ReadAccountByMemberGUID**](AccountsApi.md#ReadAccountByMemberGUID) | **Get** /users/{user_guid}/members/{member_guid}/accounts/{account_guid} | Read an account


# **ListAccountTransactions**
> Transactions ListAccountTransactions(ctx, accountGuid, userGuid, optional)
List account transactions

This endpoint allows you to see every transaction that belongs to a specific account. The default from_date is 90 days prior to the request, and the default to_date is 5 days from the time of the request.<br> The from_date and to_date parameters can optionally be appended to the request. 

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
  
  accountGuid := "accountGuid_example" // string | The unique identifier for an `account`.
  userGuid := "userGuid_example" // string | The unique identifier for a `user`.
  opts := &mx.ListAccountTransactionsOpts{ 
    FromDate: optional.NewString("fromDate_example"), // string | Filter transactions from this date.
    ToDate: optional.NewString("toDate_example"), // string | Filter transactions to this date.
    Page: optional.NewInt32(12), // int32 | Specify current page.
    RecordsPerPage: optional.NewInt32(12), // int32 | Specify records per page.
  }

  response, _, err := client.AccountsApi.ListAccountTransactions(nil, accountGuid, userGuid, opts)
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

[**Transactions**](Transactions.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListUserAccounts**
> Accounts ListUserAccounts(ctx, userGuid, optional)
List accounts for a user

Use this endpoint to view information about every account that belongs to a user. You'll need the user's GUID to access this list. The information will include the account type — e.g., CHECKING, MONEY_MARKET, or PROPERTY — the account balance, the date the account was started, etc.

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
  opts := &mx.ListUserAccountsOpts{ 
    Page: optional.NewInt32(12), // int32 | Specify current page.
    RecordsPerPage: optional.NewInt32(12), // int32 | Specify records per page.
  }

  response, _, err := client.AccountsApi.ListUserAccounts(nil, userGuid, opts)
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
 **optional** | ***ListUserAccountsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListUserAccountsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Int32**| Specify current page. | 
 **recordsPerPage** | **optional.Int32**| Specify records per page. | 

### Return type

[**Accounts**](Accounts.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadAccount**
> Account ReadAccount(ctx, accountGuid, userGuid)
Read an account

Reading an account allows you to get information about a specific account that belongs to a user. That includes the account type — e.g., CHECKING, MONEY_MARKET, or PROPERTY — the balance, the date the account was started, and much more.<br> There are two endpoints for reading an account. Both will return the same information.<br> It's important to remember that balance and available_balance will normally be positive numbers — for all account types. But this should be interpreted differently for debt accounts and asset accounts.<br> An asset account, e.g., CHECKING, SAVINGS, or INVESTMENT, will have a positive balance unless it is in an overdraft condition, in which case the balance will be negative.<br> On the other hand, a debt account, e.g., CREDIT CARD, LOAN, MORTGAGE, would have a positivebalance when the user owes money on the account. It would have a negative balance if the account has been overpaid. 

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

  response, _, err := client.AccountsApi.ReadAccount(nil, accountGuid, userGuid)
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

[**Account**](Account.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadAccountByMemberGUID**
> Account ReadAccountByMemberGUID(ctx, accountGuid, memberGuid, userGuid)
Read an account

Reading an account allows you to get information about a specific account that belongs to a user. That includes the account type — e.g., CHECKING, MONEY_MARKET, or PROPERTY — the balance, the date the account was started, and much more.<br> There are two endpoints for reading an account. Both will return the same information.<br> It's important to remember that balance and available_balance will normally be positive numbers — for all account types. But this should be interpreted differently for debt accounts and asset accounts.<br> An asset account, e.g., CHECKING, SAVINGS, or INVESTMENT, will have a positive balance unless it is in an overdraft condition, in which case the balance will be negative.<br> On the other hand, a debt account, e.g., CREDIT CARD, LOAN, MORTGAGE, would have a positivebalance when the user owes money on the account. It would have a negative balance if the account has been overpaid. 

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
  memberGuid := "memberGuid_example" // string | The unique identifier for a `member`.
  userGuid := "userGuid_example" // string | The unique identifier for a `user`.

  response, _, err := client.AccountsApi.ReadAccountByMemberGUID(nil, accountGuid, memberGuid, userGuid)
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
  **memberGuid** | **string**| The unique identifier for a &#x60;member&#x60;. | 
  **userGuid** | **string**| The unique identifier for a &#x60;user&#x60;. | 

### Return type

[**Account**](Account.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

