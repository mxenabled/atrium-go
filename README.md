# Go API client for atrium

The MX Atrium API supports over 48,000 data connections to thousands of financial institutions. It provides secure access to your users' accounts and transactions with industry-leading cleansing, categorization, and classification.  Atrium is designed according to resource-oriented REST architecture and responds with JSON bodies and HTTP response codes.  Use Atrium's development environment, vestibule.mx.com, to quickly get up and running. The development environment limits are 100 users, 25 members per user, and access to the top 15 institutions. Contact MX to purchase production access. 

## Installation
Put the package under your project folder and add the following in import:
```golang
import "github.com/mxenabled/atrium-go"
```

## Documentation for API Endpoints

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*AccountsApi* | [**ListAccountTransactions**](docs/AccountsApi.md#listaccounttransactions) | **Get** /users/{user_guid}/accounts/{account_guid}/transactions | List account transactions
*AccountsApi* | [**ListUserAccounts**](docs/AccountsApi.md#listuseraccounts) | **Get** /users/{user_guid}/accounts | List accounts for a user
*AccountsApi* | [**ReadAccount**](docs/AccountsApi.md#readaccount) | **Get** /users/{user_guid}/accounts/{account_guid} | Read an account
*AccountsApi* | [**ReadAccountByMemberGUID**](docs/AccountsApi.md#readaccountbymemberguid) | **Get** /users/{user_guid}/members/{member_guid}/accounts/{account_guid} | Read an account
*ConnectWidgetApi* | [**GetConnectWidget**](docs/ConnectWidgetApi.md#getconnectwidget) | **Post** /users/{user_guid}/connect_widget_url | Embedding in a website
*IdentityApi* | [**IdentifyMember**](docs/IdentityApi.md#identifymember) | **Post** /users/{user_guid}/members/{member_guid}/identify | Identify
*IdentityApi* | [**ListAccountOwners**](docs/IdentityApi.md#listaccountowners) | **Get** /users/{user_guid}/members/{member_guid}/account_owners | List member account owners
*InstitutionsApi* | [**ListInstitutions**](docs/InstitutionsApi.md#listinstitutions) | **Get** /institutions | List institutions
*InstitutionsApi* | [**ReadInstitution**](docs/InstitutionsApi.md#readinstitution) | **Get** /institutions/{institution_code} | Read institution
*InstitutionsApi* | [**ReadInstitutionCredentials**](docs/InstitutionsApi.md#readinstitutioncredentials) | **Get** /institutions/{institution_code}/credentials | Read institution credentials
*MembersApi* | [**AggregateMember**](docs/MembersApi.md#aggregatemember) | **Post** /users/{user_guid}/members/{member_guid}/aggregate | Aggregate member
*MembersApi* | [**CreateMember**](docs/MembersApi.md#createmember) | **Post** /users/{user_guid}/members | Create member
*MembersApi* | [**DeleteMember**](docs/MembersApi.md#deletemember) | **Delete** /users/{user_guid}/members/{member_guid} | Delete member
*MembersApi* | [**ListMemberAccounts**](docs/MembersApi.md#listmemberaccounts) | **Get** /users/{user_guid}/members/{member_guid}/accounts | List member accounts
*MembersApi* | [**ListMemberCredentials**](docs/MembersApi.md#listmembercredentials) | **Get** /users/{user_guid}/members/{member_guid}/credentials | List member credentials
*MembersApi* | [**ListMemberMFAChallenges**](docs/MembersApi.md#listmembermfachallenges) | **Get** /users/{user_guid}/members/{member_guid}/challenges | List member MFA challenges
*MembersApi* | [**ListMemberTransactions**](docs/MembersApi.md#listmembertransactions) | **Get** /users/{user_guid}/members/{member_guid}/transactions | List member transactions
*MembersApi* | [**ListMembers**](docs/MembersApi.md#listmembers) | **Get** /users/{user_guid}/members | List members
*MembersApi* | [**ReadMember**](docs/MembersApi.md#readmember) | **Get** /users/{user_guid}/members/{member_guid} | Read member
*MembersApi* | [**ReadMemberStatus**](docs/MembersApi.md#readmemberstatus) | **Get** /users/{user_guid}/members/{member_guid}/status | Read member connection status
*MembersApi* | [**ResumeMember**](docs/MembersApi.md#resumemember) | **Put** /users/{user_guid}/members/{member_guid}/resume | Resume aggregation from MFA
*MembersApi* | [**UpdateMember**](docs/MembersApi.md#updatemember) | **Put** /users/{user_guid}/members/{member_guid} | Update member
*TransactionsApi* | [**CleanseAndCategorizeTransactions**](docs/TransactionsApi.md#cleanseandcategorizetransactions) | **Post** /cleanse_and_categorize | Categorize transactions
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
 - [AccountAttributes](docs/AccountAttributes.md)
 - [AccountNumberAttributes](docs/AccountNumberAttributes.md)
 - [AccountNumbers](docs/AccountNumbers.md)
 - [AccountOwnerAttributes](docs/AccountOwnerAttributes.md)
 - [AccountOwners](docs/AccountOwners.md)
 - [Accounts](docs/Accounts.md)
 - [ChallengeAttributes](docs/ChallengeAttributes.md)
 - [ChallengeAttributesOptions](docs/ChallengeAttributesOptions.md)
 - [Challenges](docs/Challenges.md)
 - [ConnectWidget](docs/ConnectWidget.md)
 - [ConnectWidgetAttributes](docs/ConnectWidgetAttributes.md)
 - [ConnectWidgetRequestBody](docs/ConnectWidgetRequestBody.md)
 - [CredentialAttributes](docs/CredentialAttributes.md)
 - [CredentialOptionAttributes](docs/CredentialOptionAttributes.md)
 - [CredentialResponseAttributes](docs/CredentialResponseAttributes.md)
 - [Credentials](docs/Credentials.md)
 - [Institution](docs/Institution.md)
 - [InstitutionAttributes](docs/InstitutionAttributes.md)
 - [Institutions](docs/Institutions.md)
 - [Member](docs/Member.md)
 - [MemberAttributes](docs/MemberAttributes.md)
 - [MemberConnectionStatus](docs/MemberConnectionStatus.md)
 - [MemberConnectionStatusAttributes](docs/MemberConnectionStatusAttributes.md)
 - [MemberCreateRequestBody](docs/MemberCreateRequestBody.md)
 - [MemberCreateRequestBodyAttributes](docs/MemberCreateRequestBodyAttributes.md)
 - [MemberResumeRequestBody](docs/MemberResumeRequestBody.md)
 - [MemberResumeRequestBodyAttributes](docs/MemberResumeRequestBodyAttributes.md)
 - [MemberUpdateRequestBody](docs/MemberUpdateRequestBody.md)
 - [MemberUpdateRequestBodyAttributes](docs/MemberUpdateRequestBodyAttributes.md)
 - [Members](docs/Members.md)
 - [Pagination](docs/Pagination.md)
 - [Transaction](docs/Transaction.md)
 - [TransactionAttributes](docs/TransactionAttributes.md)
 - [Transactions](docs/Transactions.md)
 - [TransactionsCleanseAndCategorize](docs/TransactionsCleanseAndCategorize.md)
 - [TransactionsCleanseAndCategorizeAttributes](docs/TransactionsCleanseAndCategorizeAttributes.md)
 - [TransactionsCleanseAndCategorizeRequestBody](docs/TransactionsCleanseAndCategorizeRequestBody.md)
 - [TransactionsCleanseAndCategorizeRequestBodyAttributes](docs/TransactionsCleanseAndCategorizeRequestBodyAttributes.md)
 - [User](docs/User.md)
 - [UserAttributes](docs/UserAttributes.md)
 - [UserCreateRequestBody](docs/UserCreateRequestBody.md)
 - [UserUpdateRequestBody](docs/UserUpdateRequestBody.md)
 - [Users](docs/Users.md)

