# \InstitutionsApi

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListInstitutions**](InstitutionsApi.md#ListInstitutions) | **Get** /institutions | List institutions
[**ReadInstitution**](InstitutionsApi.md#ReadInstitution) | **Get** /institutions/{institution_code} | Read institution
[**ReadInstitutionCredentials**](InstitutionsApi.md#ReadInstitutionCredentials) | **Get** /institutions/{institution_code}/credentials | Read institution credentials


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
  "github.com/mxenabled/atrium-go/v2"
  "github.com/antihax/optional"
)

func main() {
  client := atrium.AtriumClient("YOUR_API_KEY", "YOUR_CLIENT_ID")
  ctx := context.Background()
  
  opts := &atrium.ListInstitutionsOpts{ 
    Name: optional.NewString(name_example), // string | This will list only institutions in which the appended string appears.
    Page: optional.NewInt32(1), // int32 | Specify current page.
    RecordsPerPage: optional.NewInt32(12), // int32 | Specify records per page.
    SupportsAccountIdentification: optional.NewBool(true), // bool | Filter only institutions which support account identification.
    SupportsAccountStatement: optional.NewBool(true), // bool | Filter only institutions which support account statements.
    SupportsAccountVerification: optional.NewBool(true), // bool | Filter only institutions which support account verification.
    SupportsTransactionHistory: optional.NewBool(true), // bool | Filter only institutions which support extended transaction history.
  }

  response, _, err := client.Institutions.ListInstitutions(ctx, opts)
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
 **supportsAccountIdentification** | **optional.Bool**| Filter only institutions which support account identification. | 
 **supportsAccountStatement** | **optional.Bool**| Filter only institutions which support account statements. | 
 **supportsAccountVerification** | **optional.Bool**| Filter only institutions which support account verification. | 
 **supportsTransactionHistory** | **optional.Bool**| Filter only institutions which support extended transaction history. | 

### Return type

[**InstitutionsResponseBody**](InstitutionsResponseBody.md)

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
  "github.com/mxenabled/atrium-go/v2"
)

func main() {
  client := atrium.AtriumClient("YOUR_API_KEY", "YOUR_CLIENT_ID")
  ctx := context.Background()
  
  institutionCode := "example_institution_code" // string | The institution_code of the institution.

  response, _, err := client.Institutions.ReadInstitution(ctx, institutionCode)
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
  "github.com/mxenabled/atrium-go/v2"
)

func main() {
  client := atrium.AtriumClient("YOUR_API_KEY", "YOUR_CLIENT_ID")
  ctx := context.Background()
  
  institutionCode := "example_institution_code" // string | The institution_code of the institution.

  response, _, err := client.Institutions.ReadInstitutionCredentials(ctx, institutionCode)
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

