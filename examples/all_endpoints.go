package main

import (
	"fmt"
  "time"
	"github.com/mxenabled/atrium-go/client"
	"github.com/mxenabled/atrium-go/models"
)

func main() {
	// Create a new Client
	client := &client.Client{
		ApiKey:   "YOUR_MX_API_KEY",
		ClientId: "YOUR_MX_CLIENT_ID",
		ApiURL:   "https://vestibule.mx.com",
	}

  fmt.Println("\n************************** Create User **************************")
  newUser := &models.User {
    Identifier: "unique_id",
    Metadata: "{\"first_name\": \"Steven\"}",
  }
  user, err := client.CreateUser(newUser)
  if err != nil {
    fmt.Println("Error creating user:", err)
    return
  }
  fmt.Printf("%+v\n", user)
  userGUID := user.Guid

  fmt.Println("\n************************** Read User **************************")
  user, err = client.GetUser(userGUID)
  if err != nil {
    fmt.Println("Error reading user:", err)
    return
  }
  fmt.Printf("%+v\n", user)

  fmt.Println("\n************************** Update User **************************")
  updatedUser := &models.User {
    Guid: userGUID,
    Metadata: "{\"first_name\": \"Steven\", \"last_name\": \"Universe\"}",
  }
  user, err = client.UpdateUser(updatedUser)
  if err != nil {
    fmt.Println("Error updating user:", err)
    return
  }
  fmt.Printf("%+v\n", user)

  fmt.Println("\n************************** List Users **************************")
  users, err := client.ListUsers()
  if err != nil {
    fmt.Println("Error listing users:", err)
    return
  }
  for _, user := range users {
    fmt.Printf("%+v\n", user)
  }

  fmt.Println("\n************************** List Institutions **************************")
  institutions, err := client.ListInstitutions("bank")
  if err != nil {
    fmt.Println("Error listing institutions:", err)
    return
  }
  for _, institution := range institutions {
    fmt.Printf("%+v\n", institution)
  }

  fmt.Println("\n************************** Read Institution **************************")
  institution, err := client.GetInstitution("mxbank")
  if err != nil {
    fmt.Println("Error reading institution:", err)
    return
  }
  fmt.Printf("%+v\n", institution)


  fmt.Println("\n************************** Read Institution Credentials ************************** ")
  credentials, err := client.ListCredentials("mxbank")
  if err != nil {
    fmt.Println("Error reading institution credentials:", err)
    return
  }
  for _, credential := range credentials {
    fmt.Printf("%+v\n", credential)
  }

  fmt.Println("\n************************** Create Member **************************")
  loginCredentials := []*models.Credential{}
  username := &models.Credential{Guid: "CRD-9f61fb4c-912c-bd1e-b175-ccc7f0275cc1", Value: "test_atrium1"}
  loginCredentials = append(loginCredentials, username)
  password := &models.Credential{Guid: "CRD-e3d7ea81-aac7-05e9-fbdd-4b493c6e474d", Value: "challenge1"}
  loginCredentials = append(loginCredentials, password)

  newMember := &models.Member{
    InstitutionCode: "mxbank",
  }
  member, err := client.CreateMember(userGUID, newMember, loginCredentials)
  if err != nil {
    fmt.Println("Error creating member:", err)
    return
  }
  fmt.Printf("%+v\n", member)
  memberGUID := member.Guid

  fmt.Println("\n************************** Read Member **************************")
  member, err = client.GetMember(userGUID, memberGUID)
  if err != nil {
    fmt.Println("Error reading member:", err)
    return
  }
  fmt.Printf("%+v\n", member)

  fmt.Println("\n************************** Update Member **************************")
  updatedCredentials := []*models.Credential{}
  username = &models.Credential{Guid: "CRD-9f61fb4c-912c-bd1e-b175-ccc7f0275cc1", Value: "test_atrium"}
  updatedCredentials = append(updatedCredentials, username)
  password = &models.Credential{Guid: "CRD-e3d7ea81-aac7-05e9-fbdd-4b493c6e474d", Value: "challenge"}
  updatedCredentials = append(updatedCredentials, password)

  updatedMember := &models.Member{
    Guid: memberGUID,
    Metadata: "{\"credentials_last_refreshed_at\": \"2015-10-16\"}",
  }

  member, err = client.UpdateMember(userGUID, updatedMember, updatedCredentials)
  if err != nil {
    fmt.Println("Error updating member:", err)
    return
  }
  fmt.Printf("%+v\n", member)

  fmt.Println("\n************************** List Members **************************")
  members, err := client.ListMembers(userGUID)
  if err != nil {
    fmt.Println("Error listing members:", err)
    return
  }
  for _, member := range members {
    fmt.Printf("%+v\n", member)
  }

  fmt.Println("\n************************** Aggregate Member **************************")
  member, err = client.AggregateMember(userGUID, memberGUID)
  if err != nil {
    fmt.Println("Error aggregating member:", err)
    return
  }
  fmt.Printf("%+v\n", member)

  fmt.Println("\n************************** Read Member Status **************************")
  time.Sleep(5 * time.Second)
  member, err = client.GetMemberStatus(userGUID, memberGUID)
  if err != nil {
    fmt.Println("Error getting member status:", err)
    return
  }
  fmt.Printf("%+v\n", member)

  fmt.Println("\n************************** List Member MFA Challenges **************************")
  challenges, err := client.GetMemberChallenges(userGUID, memberGUID)
  if err != nil {
    fmt.Println("Error listing member MFA challenges:", err)
    return
  }
  for _, challenge := range challenges {
    fmt.Printf("%+v\n", challenge)
  }

  fmt.Println("\n************************** Resume Aggregation **************************")
  challenges1 := []*models.Challenge{}
  challengeOne := &models.Challenge{Guid: challenges[0].Guid, Value: "correct"}
  challenges1 = append(challenges1, challengeOne)

  member, err = client.ResumeMember(userGUID, memberGUID, challenges1)
  if err != nil {
    fmt.Println("Error resuming member:", err)
    return
  }
  fmt.Printf("%+v\n", member)

  fmt.Println("\n************************** List Member Credentials **************************")
  credentials, err = client.ListMemberCredentials(userGUID, memberGUID)
  if err != nil {
    fmt.Println("Error listing member credentials:", err)
    return
  }
  for _, credential := range credentials {
    fmt.Printf("%+v\n", credential)
  }

  fmt.Println("\n************************** List Member Accounts **************************")
  time.Sleep(5 * time.Second)
  accounts, err := client.ListMemberAccounts(userGUID, memberGUID)
  if err != nil {
    fmt.Println("Error listing member accounts:", err)
    return
  }
  for _, account := range accounts {
    fmt.Printf("%+v\n", account)
  }
  accountGUID := accounts[0].Guid

  fmt.Println("\n************************** List Member Transactions **************************")
  transactions, err := client.ListMemberTransactions(userGUID, memberGUID)
  if err != nil {
    fmt.Println("Error listing member transactions:", err)
    return
  }
  for _, transaction := range transactions {
    fmt.Printf("%+v\n", transaction)
  }

  fmt.Println("\n************************** Read Account **************************")
  account, err := client.GetAccount(userGUID, accountGUID)
  if err != nil {
    fmt.Println("Error reading account:", err)
    return
  }
  fmt.Printf("%+v\n", account)


  fmt.Println("\n************************** List Accounts for User **************************")
  accounts, err = client.ListAccounts(userGUID)
  if err != nil {
    fmt.Println("Error listing accounts:", err)
    return
  }
  for _, account := range accounts {
    fmt.Printf("%+v\n", account)
  }

  fmt.Println("\n************************** List Account Transactions **************************")
  transactions, err = client.ListAccountTransactions(userGUID, accountGUID)
  if err != nil {
    fmt.Println("Error listing account transactions:", err)
    return
  }
  for _, transaction := range transactions {
    fmt.Printf("%+v\n", transaction)
  }
  transactionGUID := transactions[0].Guid

  fmt.Println("\n************************** Read a Transaction **************************")
  transaction, err := client.GetTransaction(userGUID, transactionGUID)
  if err != nil {
    fmt.Println("Error reading Transaction:", err)
    return
  }
  fmt.Printf("%+v\n", transaction)

  fmt.Println("\n************************** List Transactions **************************")
  transactions, err = client.ListTransactions(userGUID)
  if err != nil {
    fmt.Println("Error listing transactions:", err)
    return
  }
  for _, transaction := range transactions {
    fmt.Printf("%+v\n", transaction)
  }

  fmt.Println("\n************************** Connect Widget **************************")
  connect, err := client.GetWidget(userGUID)
  if err != nil {
    fmt.Println("Error getting connect widget:", err)
    return
  }
  fmt.Printf("%+v\n", connect)

  fmt.Println("\n************************** Delete Member **************************")
  err = client.DeleteMember(userGUID, memberGUID)
  if err != nil {
    fmt.Println("Error deleting member:", err)
    return
  }

  fmt.Println("\n************************** Delete User **************************")
  err = client.DeleteUser(userGUID)
  if err != nil {
    fmt.Println("Error deleting user:", err)
    return
  }
}
