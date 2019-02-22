# \TransactionsApi

Method | HTTP request | Description
------------- | ------------- | -------------
[**CleanseAndCategorizeTransactions**](TransactionsApi.md#CleanseAndCategorizeTransactions) | **Post** /cleanse_and_categorize | Categorize transactions
[**ListUserTransactions**](TransactionsApi.md#ListUserTransactions) | **Get** /users/{user_guid}/transactions | List transactions for a user
[**ReadTransaction**](TransactionsApi.md#ReadTransaction) | **Get** /users/{user_guid}/transactions/{transaction_guid} | Read a transaction


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
  client := atrium.AtriumClient("YOUR_API_KEY", "YOUR_CLIENT_ID")
  ctx := context.Background()
  
  body := atrium.TransactionsCleanseAndCategorizeRequestBody{} // TransactionsCleanseAndCategorizeRequestBody | User object to be created with optional parameters (amount, type) amd required parameters (description, identifier)

  response, _, err := client.Transactions.CleanseAndCategorizeTransactions(ctx, body)
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
  client := atrium.AtriumClient("YOUR_API_KEY", "YOUR_CLIENT_ID")
  ctx := context.Background()
  
  userGUID := "USR-123" // string | The unique identifier for a `user`.
  opts := &atrium.ListUserTransactionsOpts{ 
    Page: optional.NewInt32(1), // int32 | Specify current page.
    FromDate: optional.NewString("2016-09-20"), // string | Filter transactions from this date.
    RecordsPerPage: optional.NewInt32(12), // int32 | Specify records per page.
    ToDate: optional.NewString("2016-10-20"), // string | Filter transactions to this date.
  }

  response, _, err := client.Transactions.ListUserTransactions(ctx, userGUID, , opts)
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
  client := atrium.AtriumClient("YOUR_API_KEY", "YOUR_CLIENT_ID")
  ctx := context.Background()
  
  transactionGUID := "TRN-123" // string | The unique identifier for a `transaction`.
  userGUID := "USR-123" // string | The unique identifier for a `user`.

  response, _, err := client.Transactions.ReadTransaction(ctx, transactionGUID, userGUID, )
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

