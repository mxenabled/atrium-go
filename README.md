atrium-go
=========

This is a wrapper for the [MX Atrium API](https://atrium.mx.com/).

#### Installation

The only requirement is the Go Programming Language. This wrapper only uses the standard library.

```
$ go get -u github.com/mxenabled/atrium-go
```

#### Example

Check out the examples directory of this repo for more working examples.

```go
package main

import (
    "github.com/mxenabled/atrium-go/client"
    "github.com/mxenabled/atrium-go/models"
    "fmt"
)

func main() {
    client := &client.Client{
		ApiKey:   "YOUR_API_KEY",
		ClientId: "YOUR_CLIENT_ID",
		ApiURL:   "https://vestibule.mx.com",
	}

    newUser := &models.User{
        Identifier: "my_unique_id",
        Metadata:   "info to store on the user",
    }

    user, _ := client.CreateUser(newUser)
    fmt.Println("Created a user with guid:", user.Guid)

    err := client.DeleteUser(user)
    if err != nil {
        fmt.Println("Error deleting user:", err)
    } else {
        fmt.Println("User", user.Guid, "was deleted")
    }
}
```

#### License

MIT License.
