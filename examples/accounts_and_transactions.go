package main

import (
	"fmt"
  "time"
	"github.com/mxenabled/atrium-go/client"
	"github.com/mxenabled/atrium-go/models"
)

func main() {
	// Create a new Client
	client := &client.Client {
    ApiKey:   "YOUR_MX_API_KEY",
    ClientId: "YOUR_MX_CLIENT_ID",
    ApiURL:   "https://vestibule.mx.com",
  }

  fmt.Println("\n* Creating test user and member *")
  newUser := &models.User {
    Metadata: "some metadata",
    IsDisabled: false,
  }
  user, err := client.CreateUser(newUser)
  if err != nil {
    fmt.Println("Error creating user:", err)
    return
  }
	fmt.Println("Created user:", user.Guid)

  loginCredentials := []*models.Credential{}
  username := &models.Credential{Guid: "CRD-9f61fb4c-912c-bd1e-b175-ccc7f0275cc1", Value: "test_atrium"}
  loginCredentials = append(loginCredentials, username)
  password := &models.Credential{Guid: "CRD-e3d7ea81-aac7-05e9-fbdd-4b493c6e474d", Value: "password"}
  loginCredentials = append(loginCredentials, password)

  newMember := &models.Member{
    InstitutionCode: "mxbank",
  }
  member, memberError := client.CreateMember(user.Guid, newMember, loginCredentials)
  if err != nil {
    fmt.Println("Error creating member:", memberError)
    return
  }
	fmt.Println("Created member:", member.Guid)
  time.Sleep(time.Second)


  fmt.Println("\n* Aggregating member *")
  _, aggError := client.AggregateMember(user.Guid, member.Guid)
  if err != nil {
    fmt.Println("Error aggregating member:", aggError)
    return
  }
  time.Sleep(4 * time.Second)


  fmt.Println("\n* Listing all member accounts and transactions *")
  accounts, err := client.ListMemberAccounts(user.Guid, member.Guid)
  if err != nil {
    fmt.Println("Error listing member accounts:", err)
    return
  }
  for _, account := range accounts {
    fmt.Println("Type:", account.Type, "\tName:", account.Name, "\tAvailable Balance:", account.AvailableBalance, "\tAvailable Credit:", account.AvailableCredit)
    fmt.Println("Transactions")
    transactions, accountError := client.ListAccountTransactions(user.Guid, account.Guid)
    if err != nil {
      fmt.Println("Error listing account transactions:", accountError)
      return
    }
    for _, transaction := range transactions {
      fmt.Println("\tDate:", transaction.PostedAt, "\tDescription:", transaction.Description, "\tAmount:", transaction.Amount)
    }
  }


	fmt.Println("\n* Deleting test user *")
  deleteError := client.DeleteUser(user.Guid)
  if err != nil {
    fmt.Println("Error deleting user:", deleteError)
    return
  }
	fmt.Println("Deleted user:", user.Guid)
}
