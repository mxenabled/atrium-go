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

	response, _, err := client.Accounts.ListAccountTransactions(ctx, accountGUID, userGUID, , opts)
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
*AccountsApi* | [**ListAccountTransactions**](docs/AccountsApi.md#listaccounttransactions) | **Get** /users/{user_guid}/accounts/{account_guid}/transactions | List account transactions
*AccountsApi* | [**ListUserAccounts**](docs/AccountsApi.md#listuseraccounts) | **Get** /users/{user_guid}/accounts | List accounts for a user
*AccountsApi* | [**ReadAccount**](docs/AccountsApi.md#readaccount) | **Get** /users/{user_guid}/accounts/{account_guid} | Read an account
*AccountsApi* | [**ReadAccountByMemberGUID**](docs/AccountsApi.md#readaccountbymemberguid) | **Get** /users/{user_guid}/members/{member_guid}/accounts/{account_guid} | Read an account
*ConnectWidgetApi* | [**GetConnectWidget**](docs/ConnectWidgetApi.md#getconnectwidget) | **Post** /users/{user_guid}/connect_widget_url | Embedding in a website
*HoldingsApi* | [**ListHoldings**](docs/HoldingsApi.md#listholdings) | **Get** /users/{user_guid}/holdings | List holdings
*HoldingsApi* | [**ListHoldingsByAccount**](docs/HoldingsApi.md#listholdingsbyaccount) | **Get** /users/{user_guid}/accounts/{account_guid}/holdings | List holdings by account
*HoldingsApi* | [**ListHoldingsByMember**](docs/HoldingsApi.md#listholdingsbymember) | **Get** /users/{user_guid}/members/{member_guid}/holdings | List holdings by member
*HoldingsApi* | [**ReadHolding**](docs/HoldingsApi.md#readholding) | **Get** /users/{user_guid}/holdings/{holding_guid} | Read holding
*IdentityApi* | [**IdentifyMember**](docs/IdentityApi.md#identifymember) | **Post** /users/{user_guid}/members/{member_guid}/identify | Identify
*IdentityApi* | [**ListAccountOwners**](docs/IdentityApi.md#listaccountowners) | **Get** /users/{user_guid}/members/{member_guid}/account_owners | List member account owners
*InstitutionsApi* | [**ListInstitutions**](docs/InstitutionsApi.md#listinstitutions) | **Get** /institutions | List institutions
*InstitutionsApi* | [**ReadInstitution**](docs/InstitutionsApi.md#readinstitution) | **Get** /institutions/{institution_code} | Read institution
*InstitutionsApi* | [**ReadInstitutionCredentials**](docs/InstitutionsApi.md#readinstitutioncredentials) | **Get** /institutions/{institution_code}/credentials | Read institution credentials
*MembersApi* | [**AggregateMember**](docs/MembersApi.md#aggregatemember) | **Post** /users/{user_guid}/members/{member_guid}/aggregate | Aggregate member
*MembersApi* | [**AggregateMemberBalances**](docs/MembersApi.md#aggregatememberbalances) | **Post** /users/{user_guid}/members/{member_guid}/balance | Aggregate member account balances
*MembersApi* | [**CreateMember**](docs/MembersApi.md#createmember) | **Post** /users/{user_guid}/members | Create member
*MembersApi* | [**DeleteMember**](docs/MembersApi.md#deletemember) | **Delete** /users/{user_guid}/members/{member_guid} | Delete member
*MembersApi* | [**ExtendHistory**](docs/MembersApi.md#extendhistory) | **Post** /users/{user_guid}/members/{member_guid}/extend_history | Extend history
*MembersApi* | [**ListMemberAccounts**](docs/MembersApi.md#listmemberaccounts) | **Get** /users/{user_guid}/members/{member_guid}/accounts | List member accounts
*MembersApi* | [**ListMemberCredentials**](docs/MembersApi.md#listmembercredentials) | **Get** /users/{user_guid}/members/{member_guid}/credentials | List member credentials
*MembersApi* | [**ListMemberMFAChallenges**](docs/MembersApi.md#listmembermfachallenges) | **Get** /users/{user_guid}/members/{member_guid}/challenges | List member MFA challenges
*MembersApi* | [**ListMemberTransactions**](docs/MembersApi.md#listmembertransactions) | **Get** /users/{user_guid}/members/{member_guid}/transactions | List member transactions
*MembersApi* | [**ListMembers**](docs/MembersApi.md#listmembers) | **Get** /users/{user_guid}/members | List members
*MembersApi* | [**ReadMember**](docs/MembersApi.md#readmember) | **Get** /users/{user_guid}/members/{member_guid} | Read member
*MembersApi* | [**ReadMemberStatus**](docs/MembersApi.md#readmemberstatus) | **Get** /users/{user_guid}/members/{member_guid}/status | Read member connection status
*MembersApi* | [**ResumeMember**](docs/MembersApi.md#resumemember) | **Put** /users/{user_guid}/members/{member_guid}/resume | Resume aggregation from MFA
*MembersApi* | [**UpdateMember**](docs/MembersApi.md#updatemember) | **Put** /users/{user_guid}/members/{member_guid} | Update member
*MerchantsApi* | [**ReadMerchant**](docs/MerchantsApi.md#readmerchant) | **Get** /merchants/{merchant_guid} | Read merchant
*StatementsApi* | [**DownloadStatementPdf**](docs/StatementsApi.md#downloadstatementpdf) | **Get** /users/{user_guid}/members/{member_guid}/statements/{statement_guid}.pdf | Download statement PDF
*StatementsApi* | [**FetchStatements**](docs/StatementsApi.md#fetchstatements) | **Post** /users/{user_guid}/members/{member_guid}/fetch_statements | Fetch statements
*StatementsApi* | [**ListMemberStatements**](docs/StatementsApi.md#listmemberstatements) | **Get** /users/{user_guid}/members/{member_guid}/statements | List member statements
*StatementsApi* | [**ReadMemberStatement**](docs/StatementsApi.md#readmemberstatement) | **Get** /users/{user_guid}/members/{member_guid}/statements/{statement_guid} | Read statement JSON
*TransactionsApi* | [**CleanseAndCategorizeTransactions**](docs/TransactionsApi.md#cleanseandcategorizetransactions) | **Post** /transactions/cleanse_and_categorize | Categorize transactions
*TransactionsApi* | [**ListUserTransactions**](docs/TransactionsApi.md#listusertransactions) | **Get** /users/{user_guid}/transactions | List transactions for a user
*TransactionsApi* | [**ReadTransaction**](docs/TransactionsApi.md#readtransaction) | **Get** /users/{user_guid}/transactions/{transaction_guid} | Read a transaction
*UsersApi* | [**CreateUser**](docs/UsersApi.md#createuser) | **Post** /users | Create user
*UsersApi* | [**DeleteUser**](docs/UsersApi.md#deleteuser) | **Delete** /users/{user_guid} | Delete user
*UsersApi* | [**ListUsers**](docs/UsersApi.md#listusers) | **Get** /users | List users
*UsersApi* | [**ReadUser**](docs/UsersApi.md#readuser) | **Get** /users/{user_guid} | Read user
*UsersApi* | [**UpdateUser**](docs/UsersApi.md#updateuser) | **Put** /users/{user_guid} | Update user
*VerificationApi* | [**ListAccountNumbers**](docs/VerificationApi.md#listaccountnumbers) | **Get** /users/{user_guid}/members/{member_guid}/account_numbers | Read account numbers
*VerificationApi* | [**ListAccountNumbersByAccount**](docs/VerificationApi.md#listaccountnumbersbyaccount) | **Get** /users/{user_guid}/accounts/{account_guid}/account_numbers | Read account numbers by account GUID
*VerificationApi* | [**VerifyMember**](docs/VerificationApi.md#verifymember) | **Post** /users/{user_guid}/members/{member_guid}/verify | Verify


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
 - [Holding](docs/Holding.md)
 - [HoldingResponseBody](docs/HoldingResponseBody.md)
 - [HoldingsResponseBody](docs/HoldingsResponseBody.md)
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
 - [Merchant](docs/Merchant.md)
 - [MerchantResponseBody](docs/MerchantResponseBody.md)
 - [Pagination](docs/Pagination.md)
 - [Statement](docs/Statement.md)
 - [StatementResponseBody](docs/StatementResponseBody.md)
 - [StatementsResponseBody](docs/StatementsResponseBody.md)
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

