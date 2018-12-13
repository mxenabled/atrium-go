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
*AccountsAPI* | [**ListAccountTransactions**](docs/AccountsAPI.md#listaccounttransactions) | **Get** /users/{user_guid}/accounts/{account_guid}/transactions | List account transactions
*AccountsAPI* | [**ListUserAccounts**](docs/AccountsAPI.md#listuseraccounts) | **Get** /users/{user_guid}/accounts | List accounts for a user
*AccountsAPI* | [**ReadAccount**](docs/AccountsAPI.md#readaccount) | **Get** /users/{user_guid}/accounts/{account_guid} | Read an account
*AccountsAPI* | [**ReadAccountByMemberGUID**](docs/AccountsAPI.md#readaccountbymemberguid) | **Get** /users/{user_guid}/members/{member_guid}/accounts/{account_guid} | Read an account
*ConnectWidgetAPI* | [**GetConnectWidget**](docs/ConnectWidgetAPI.md#getconnectwidget) | **Post** /users/{user_guid}/connect_widget_url | Embedding in a website
*IdentityAPI* | [**IdentifyMember**](docs/IdentityAPI.md#identifymember) | **Post** /users/{user_guid}/members/{member_guid}/identify | Identify
*IdentityAPI* | [**ListAccountOwners**](docs/IdentityAPI.md#listaccountowners) | **Get** /users/{user_guid}/members/{member_guid}/account_owners | List member account owners
*InstitutionsAPI* | [**ListInstitutions**](docs/InstitutionsAPI.md#listinstitutions) | **Get** /institutions | List institutions
*InstitutionsAPI* | [**ReadInstitution**](docs/InstitutionsAPI.md#readinstitution) | **Get** /institutions/{institution_code} | Read institution
*InstitutionsAPI* | [**ReadInstitutionCredentials**](docs/InstitutionsAPI.md#readinstitutioncredentials) | **Get** /institutions/{institution_code}/credentials | Read institution credentials
*MembersAPI* | [**AggregateMember**](docs/MembersAPI.md#aggregatemember) | **Post** /users/{user_guid}/members/{member_guid}/aggregate | Aggregate member
*MembersAPI* | [**CreateMember**](docs/MembersAPI.md#createmember) | **Post** /users/{user_guid}/members | Create member
*MembersAPI* | [**DeleteMember**](docs/MembersAPI.md#deletemember) | **Delete** /users/{user_guid}/members/{member_guid} | Delete member
*MembersAPI* | [**ListMemberAccounts**](docs/MembersAPI.md#listmemberaccounts) | **Get** /users/{user_guid}/members/{member_guid}/accounts | List member accounts
*MembersAPI* | [**ListMemberCredentials**](docs/MembersAPI.md#listmembercredentials) | **Get** /users/{user_guid}/members/{member_guid}/credentials | List member credentials
*MembersAPI* | [**ListMemberMFAChallenges**](docs/MembersAPI.md#listmembermfachallenges) | **Get** /users/{user_guid}/members/{member_guid}/challenges | List member MFA challenges
*MembersAPI* | [**ListMemberTransactions**](docs/MembersAPI.md#listmembertransactions) | **Get** /users/{user_guid}/members/{member_guid}/transactions | List member transactions
*MembersAPI* | [**ListMembers**](docs/MembersAPI.md#listmembers) | **Get** /users/{user_guid}/members | List members
*MembersAPI* | [**ReadMember**](docs/MembersAPI.md#readmember) | **Get** /users/{user_guid}/members/{member_guid} | Read member
*MembersAPI* | [**ReadMemberStatus**](docs/MembersAPI.md#readmemberstatus) | **Get** /users/{user_guid}/members/{member_guid}/status | Read member connection status
*MembersAPI* | [**ResumeMember**](docs/MembersAPI.md#resumemember) | **Put** /users/{user_guid}/members/{member_guid}/resume | Resume aggregation from MFA
*MembersAPI* | [**UpdateMember**](docs/MembersAPI.md#updatemember) | **Put** /users/{user_guid}/members/{member_guid} | Update member
*TransactionsAPI* | [**CleanseAndCategorizeTransactions**](docs/TransactionsAPI.md#cleanseandcategorizetransactions) | **Post** /cleanse_and_categorize | Categorize transactions
*TransactionsAPI* | [**ListUserTransactions**](docs/TransactionsAPI.md#listusertransactions) | **Get** /users/{user_guid}/transactions | List transactions for a user
*TransactionsAPI* | [**ReadTransaction**](docs/TransactionsAPI.md#readtransaction) | **Get** /users/{user_guid}/transactions/{transaction_guid} | Read a transaction
*UsersAPI* | [**CreateUser**](docs/UsersAPI.md#createuser) | **Post** /users | Create user
*UsersAPI* | [**DeleteUser**](docs/UsersAPI.md#deleteuser) | **Delete** /users/{user_guid} | Delete user
*UsersAPI* | [**ListUsers**](docs/UsersAPI.md#listusers) | **Get** /users | List users
*UsersAPI* | [**ReadUser**](docs/UsersAPI.md#readuser) | **Get** /users/{user_guid} | Read user
*UsersAPI* | [**UpdateUser**](docs/UsersAPI.md#updateuser) | **Put** /users/{user_guid} | Update user
*VerificationAPI* | [**ListAccountNumbers**](docs/VerificationAPI.md#listaccountnumbers) | **Get** /users/{user_guid}/members/{member_guid}/account_numbers | Read account numbers
*VerificationAPI* | [**ListAccountNumbersByAccount**](docs/VerificationAPI.md#listaccountnumbersbyaccount) | **Get** /users/{user_guid}/accounts/{account_guid}/account_numbers | Read account numbers by account GUID
*VerificationAPI* | [**VerifyMember**](docs/VerificationAPI.md#verifymember) | **Post** /users/{user_guid}/members/{member_guid}/verify | Verify


## Documentation For Models

 - [Account](docs/Account.md)
 - [AccountAttributes](docs/AccountAttributes.md)
 - [AccountNumberAttributes](docs/AccountNumberAttributes.md)
 - [AccountNumbers](docs/AccountNumbers.md)
 - [AccountOwnerAttributes](docs/AccountOwnerAttributes.md)
 - [AccountOwners](docs/AccountOwners.md)
 - [Accounts](docs/Accounts.md)
 - [ChallengeAttributes](docs/ChallengeAttributes.md)
 - [ChallengeOptionAttributes](docs/ChallengeOptionAttributes.md)
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

