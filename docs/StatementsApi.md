# \StatementsApi

Method | HTTP request | Description
------------- | ------------- | -------------
[**DownloadStatementPdf**](StatementsApi.md#DownloadStatementPdf) | **Get** /users/{user_guid}/members/{member_guid}/statements/{statement_guid}.pdf | Download statement PDF
[**FetchStatements**](StatementsApi.md#FetchStatements) | **Post** /users/{user_guid}/members/{member_guid}/fetch_statements | Fetch statements
[**ListMemberStatements**](StatementsApi.md#ListMemberStatements) | **Get** /users/{user_guid}/members/{member_guid}/statements | List member statements
[**ReadMemberStatement**](StatementsApi.md#ReadMemberStatement) | **Get** /users/{user_guid}/members/{member_guid}/statements/{statement_guid} | Read statement JSON


# **DownloadStatementPdf**
> *os.File DownloadStatementPdf(ctx, memberGUID, userGUID, statementGUID)
Download statement PDF

Use this endpoint to download a specified statement. The endpoint URL is the same as the URI given in each `statement` object. 

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
  
  memberGUID := "MBR-123" // string | The unique identifier for a `member`.
  userGUID := "USR-123" // string | The unique identifier for a `user`.
  statementGUID := "STA-123" // string | The unique identifier for an `statement`.

  response, _, err := client.Statements.DownloadStatementPdf(ctx, memberGUID, userGUID, statementGUID)
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
  **statementGUID** | **string**| The unique identifier for an &#x60;statement&#x60;. | 

### Return type

[***os.File**](*os.File.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FetchStatements**
> MemberResponseBody FetchStatements(ctx, memberGUID, userGUID)
Fetch statements

The fetch statements endpoint begins fetching statements for a member.

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
  
  memberGUID := "MBR-123" // string | The unique identifier for a `member`.
  userGUID := "USR-123" // string | The unique identifier for a `user`.

  response, _, err := client.Statements.FetchStatements(ctx, memberGUID, userGUID, )
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

# **ListMemberStatements**
> StatementsResponseBody ListMemberStatements(ctx, memberGUID, userGUID, optional)
List member statements

Certain institutions in Atrium allow developers to access account statements associated with a particular `member`. Use this endpoint to get an array of available statements.  Before this endpoint can be used, `fetch_statements` should be performed on the relevant `member`. 

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
  
  memberGUID := "MBR-123" // string | The unique identifier for a `member`.
  userGUID := "USR-123" // string | The unique identifier for a `user`.
  opts := &atrium.ListMemberStatementsOpts{ 
    Page: optional.NewInt32(1), // int32 | Specify current page.
    RecordsPerPage: optional.NewInt32(12), // int32 | Specify records per page.
  }

  response, _, err := client.Statements.ListMemberStatements(ctx, memberGUID, userGUID, , opts)
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
 **optional** | ***ListMemberStatementsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListMemberStatementsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **optional.Int32**| Specify current page. | 
 **recordsPerPage** | **optional.Int32**| Specify records per page. | 

### Return type

[**StatementsResponseBody**](StatementsResponseBody.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadMemberStatement**
> StatementResponseBody ReadMemberStatement(ctx, memberGUID, userGUID, statementGUID)
Read statement JSON

Use this endpoint to download a specified statement. The endpoint URL is the same as the URI given in each `statement` object. 

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
  
  memberGUID := "MBR-123" // string | The unique identifier for a `member`.
  userGUID := "USR-123" // string | The unique identifier for a `user`.
  statementGUID := "STA-123" // string | The unique identifier for an `statement`.

  response, _, err := client.Statements.ReadMemberStatement(ctx, memberGUID, userGUID, statementGUID)
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
  **statementGUID** | **string**| The unique identifier for an &#x60;statement&#x60;. | 

### Return type

[**StatementResponseBody**](StatementResponseBody.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

