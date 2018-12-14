# Go API client for atrium

The MX Atrium API supports over 48,000 data connections to thousands of financial institutions. It provides secure access to your users' accounts and transactions with industry-leading cleansing, categorization, and classification.  Atrium is designed according to resource-oriented REST architecture and responds with JSON bodies and HTTP response codes.  Use Atrium's development environment, vestibule.mx.com, to quickly get up and running. The development environment limits are 100 users, 25 members per user, and access to the top 15 institutions. Contact MX to purchase production access. 

## Installation
Put the package under your project folder and add the following in import:
```golang
import "github.com/mxenabled/atrium-go"
```

## Example Usage

Please see `docs` directory for additional endpoint examples

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

## Documentation for API Endpoints

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*AtriumClient* | [**AggregateMember**](docs/AtriumClient.md#aggregatemember) | **Post** /users/{user_guid}/members/{member_guid}/aggregate | Aggregate member
*AtriumClient* | [**CleanseAndCategorizeTransactions**](docs/AtriumClient.md#cleanseandcategorizetransactions) | **Post** /cleanse_and_categorize | Categorize transactions
*AtriumClient* | [**CreateMember**](docs/AtriumClient.md#createmember) | **Post** /users/{user_guid}/members | Create member
*AtriumClient* | [**CreateUser**](docs/AtriumClient.md#createuser) | **Post** /users | Create user
*AtriumClient* | [**DeleteMember**](docs/AtriumClient.md#deletemember) | **Delete** /users/{user_guid}/members/{member_guid} | Delete member
*AtriumClient* | [**DeleteUser**](docs/AtriumClient.md#deleteuser) | **Delete** /users/{user_guid} | Delete user
*AtriumClient* | [**GetConnectWidget**](docs/AtriumClient.md#getconnectwidget) | **Post** /users/{user_guid}/connect_widget_url | Embedding in a website
*AtriumClient* | [**IdentifyMember**](docs/AtriumClient.md#identifymember) | **Post** /users/{user_guid}/members/{member_guid}/identify | Identify
*AtriumClient* | [**ListAccountNumbers**](docs/AtriumClient.md#listaccountnumbers) | **Get** /users/{user_guid}/members/{member_guid}/account_numbers | Read account numbers
*AtriumClient* | [**ListAccountNumbersByAccount**](docs/AtriumClient.md#listaccountnumbersbyaccount) | **Get** /users/{user_guid}/accounts/{account_guid}/account_numbers | Read account numbers by account GUID
*AtriumClient* | [**ListAccountOwners**](docs/AtriumClient.md#listaccountowners) | **Get** /users/{user_guid}/members/{member_guid}/account_owners | List member account owners
*AtriumClient* | [**ListAccountTransactions**](docs/AtriumClient.md#listaccounttransactions) | **Get** /users/{user_guid}/accounts/{account_guid}/transactions | List account transactions
*AtriumClient* | [**ListInstitutions**](docs/AtriumClient.md#listinstitutions) | **Get** /institutions | List institutions
*AtriumClient* | [**ListMemberAccounts**](docs/AtriumClient.md#listmemberaccounts) | **Get** /users/{user_guid}/members/{member_guid}/accounts | List member accounts
*AtriumClient* | [**ListMemberCredentials**](docs/AtriumClient.md#listmembercredentials) | **Get** /users/{user_guid}/members/{member_guid}/credentials | List member credentials
*AtriumClient* | [**ListMemberMFAChallenges**](docs/AtriumClient.md#listmembermfachallenges) | **Get** /users/{user_guid}/members/{member_guid}/challenges | List member MFA challenges
*AtriumClient* | [**ListMemberTransactions**](docs/AtriumClient.md#listmembertransactions) | **Get** /users/{user_guid}/members/{member_guid}/transactions | List member transactions
*AtriumClient* | [**ListMembers**](docs/AtriumClient.md#listmembers) | **Get** /users/{user_guid}/members | List members
*AtriumClient* | [**ListUserAccounts**](docs/AtriumClient.md#listuseraccounts) | **Get** /users/{user_guid}/accounts | List accounts for a user
*AtriumClient* | [**ListUserTransactions**](docs/AtriumClient.md#listusertransactions) | **Get** /users/{user_guid}/transactions | List transactions for a user
*AtriumClient* | [**ListUsers**](docs/AtriumClient.md#listusers) | **Get** /users | List users
*AtriumClient* | [**ReadAccount**](docs/AtriumClient.md#readaccount) | **Get** /users/{user_guid}/accounts/{account_guid} | Read an account
*AtriumClient* | [**ReadAccountByMemberGUID**](docs/AtriumClient.md#readaccountbymemberguid) | **Get** /users/{user_guid}/members/{member_guid}/accounts/{account_guid} | Read an account
*AtriumClient* | [**ReadInstitution**](docs/AtriumClient.md#readinstitution) | **Get** /institutions/{institution_code} | Read institution
*AtriumClient* | [**ReadInstitutionCredentials**](docs/AtriumClient.md#readinstitutioncredentials) | **Get** /institutions/{institution_code}/credentials | Read institution credentials
*AtriumClient* | [**ReadMember**](docs/AtriumClient.md#readmember) | **Get** /users/{user_guid}/members/{member_guid} | Read member
*AtriumClient* | [**ReadMemberStatus**](docs/AtriumClient.md#readmemberstatus) | **Get** /users/{user_guid}/members/{member_guid}/status | Read member connection status
*AtriumClient* | [**ReadTransaction**](docs/AtriumClient.md#readtransaction) | **Get** /users/{user_guid}/transactions/{transaction_guid} | Read a transaction
*AtriumClient* | [**ReadUser**](docs/AtriumClient.md#readuser) | **Get** /users/{user_guid} | Read user
*AtriumClient* | [**ResumeMember**](docs/AtriumClient.md#resumemember) | **Put** /users/{user_guid}/members/{member_guid}/resume | Resume aggregation from MFA
*AtriumClient* | [**UpdateMember**](docs/AtriumClient.md#updatemember) | **Put** /users/{user_guid}/members/{member_guid} | Update member
*AtriumClient* | [**UpdateUser**](docs/AtriumClient.md#updateuser) | **Put** /users/{user_guid} | Update user
*AtriumClient* | [**VerifyMember**](docs/AtriumClient.md#verifymember) | **Post** /users/{user_guid}/members/{member_guid}/verify | Verify


## Documentation For Models

 - [Account](docs/Account.md)
 - [AccountNumber](docs/AccountNumber.md)
 - [AccountNumbersResponseBody](docs/AccountNumbersResponseBody.md)
 - [AccountOwner](docs/AccountOwner.md)
 - [AccountOwnersResponseBody](docs/AccountOwnersResponseBody.md)
 - [AccountResponseBody](docs/AccountResponseBody.md)
 - [AccountsResponseBody](docs/AccountsResponseBody.md)
 - [Challenge](docs/Challenge.md)
 - [ChallengeOption](docs/ChallengeOption.md)
 - [ChallengesResponseBody](docs/ChallengesResponseBody.md)
 - [ConnectWidget](docs/ConnectWidget.md)
 - [ConnectWidgetRequestBody](docs/ConnectWidgetRequestBody.md)
 - [ConnectWidgetResponseBody](docs/ConnectWidgetResponseBody.md)
 - [CredentialOption](docs/CredentialOption.md)
 - [CredentialRequest](docs/CredentialRequest.md)
 - [CredentialResponse](docs/CredentialResponse.md)
 - [CredentialsResponseBody](docs/CredentialsResponseBody.md)
 - [Institution](docs/Institution.md)
 - [InstitutionResponseBody](docs/InstitutionResponseBody.md)
 - [InstitutionsResponseBody](docs/InstitutionsResponseBody.md)
 - [Member](docs/Member.md)
 - [MemberConnectionStatus](docs/MemberConnectionStatus.md)
 - [MemberConnectionStatusResponseBody](docs/MemberConnectionStatusResponseBody.md)
 - [MemberCreateRequest](docs/MemberCreateRequest.md)
 - [MemberCreateRequestBody](docs/MemberCreateRequestBody.md)
 - [MemberResponseBody](docs/MemberResponseBody.md)
 - [MemberResumeRequest](docs/MemberResumeRequest.md)
 - [MemberResumeRequestBody](docs/MemberResumeRequestBody.md)
 - [MemberUpdateRequest](docs/MemberUpdateRequest.md)
 - [MemberUpdateRequestBody](docs/MemberUpdateRequestBody.md)
 - [MembersResponseBody](docs/MembersResponseBody.md)
 - [Pagination](docs/Pagination.md)
 - [Transaction](docs/Transaction.md)
 - [TransactionCleanseAndCategorizeRequest](docs/TransactionCleanseAndCategorizeRequest.md)
 - [TransactionCleanseAndCategorizeResponse](docs/TransactionCleanseAndCategorizeResponse.md)
 - [TransactionResponseBody](docs/TransactionResponseBody.md)
 - [TransactionsCleanseAndCategorizeRequestBody](docs/TransactionsCleanseAndCategorizeRequestBody.md)
 - [TransactionsCleanseAndCategorizeResponseBody](docs/TransactionsCleanseAndCategorizeResponseBody.md)
 - [TransactionsResponseBody](docs/TransactionsResponseBody.md)
 - [User](docs/User.md)
 - [UserCreateRequestBody](docs/UserCreateRequestBody.md)
 - [UserResponseBody](docs/UserResponseBody.md)
 - [UserUpdateRequestBody](docs/UserUpdateRequestBody.md)
 - [UsersResponseBody](docs/UsersResponseBody.md)

