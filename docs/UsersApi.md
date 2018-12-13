# \UsersAPI

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateUser**](UsersAPI.md#CreateUser) | **Post** /users | Create user
[**DeleteUser**](UsersAPI.md#DeleteUser) | **Delete** /users/{user_guid} | Delete user
[**ListUsers**](UsersAPI.md#ListUsers) | **Get** /users | List users
[**ReadUser**](UsersAPI.md#ReadUser) | **Get** /users/{user_guid} | Read user
[**UpdateUser**](UsersAPI.md#UpdateUser) | **Put** /users/{user_guid} | Update user


# **CreateUser**
> User CreateUser(ctx, body)
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
  client := atrium.NewAPIClient("YOUR_API_KEY", "YOUR_CLIENT_ID")
  ctx := context.Background()
  
  body := atrium.UserCreateRequestBody{} // UserCreateRequestBody | User object to be created with optional parameters (identifier, is_disabled, metadata)

  response, _, err := client.UsersAPI.CreateUser(ctx, body)
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

[**User**](User.md)

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
  client := atrium.NewAPIClient("YOUR_API_KEY", "YOUR_CLIENT_ID")
  ctx := context.Background()
  
  userGUID := "userGUID_example" // string | The unique identifier for a `user`.

  response, _, err := client.UsersAPI.DeleteUser(ctx, userGUID)
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

# **ListUsers**
> Users ListUsers(ctx, optional)
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
  client := atrium.NewAPIClient("YOUR_API_KEY", "YOUR_CLIENT_ID")
  ctx := context.Background()
  
  opts := &atrium.ListUsersOpts{ 
    Page: optional.NewInt32(12), // int32 | Specify current page.
    RecordsPerPage: optional.NewInt32(12), // int32 | Specify records per page.
  }

  response, _, err := client.UsersAPI.ListUsers(ctx, opts)
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

[**Users**](Users.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadUser**
> User ReadUser(ctx, userGUID)
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
  client := atrium.NewAPIClient("YOUR_API_KEY", "YOUR_CLIENT_ID")
  ctx := context.Background()
  
  userGUID := "userGUID_example" // string | The unique identifier for a `user`.

  response, _, err := client.UsersAPI.ReadUser(ctx, userGUID)
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

[**User**](User.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateUser**
> User UpdateUser(ctx, userGUID, optional)
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
  client := atrium.NewAPIClient("YOUR_API_KEY", "YOUR_CLIENT_ID")
  ctx := context.Background()
  
  userGUID := "userGUID_example" // string | The unique identifier for a `user`.
  opts := &atrium.UpdateUserOpts{ 
    Body: optional.NewInterface(atrium.UserUpdateRequestBody{}), // UserUpdateRequestBody | User object to be updated with optional parameters (identifier, is_disabled, metadata)
  }

  response, _, err := client.UsersAPI.UpdateUser(ctx, userGUID, opts)
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

[**User**](User.md)

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

