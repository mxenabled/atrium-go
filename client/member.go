package client

import (
	"bytes"
	"encoding/json"
	"github.com/mxenabled/atrium-go/models"
	"net/http"
)

func parseMembersResponse(response *http.Response) ([]*models.Member, error) {
	if err := parseResponseErrors(response.StatusCode); err != nil {
		return nil, err
	}

	buffer := new(bytes.Buffer)
	buffer.ReadFrom(response.Body)
	bufferStr := buffer.String()
	response.Body.Close()

	if response.StatusCode == 200 {
		membersResponse := &models.MembersResponse{}
		if err := json.Unmarshal([]byte(bufferStr), membersResponse); err != nil {
			return nil, err
		}
		return membersResponse.Members, nil
	}

	return nil, makeGenericError(response.StatusCode, bufferStr)
}

func (c *Client) ListMembers(userGuid string) ([]*models.Member, error) {
	if userGuid == "" {
		return nil, MissingGuid
	}

	apiEndpointUrl := c.ApiURL + "/users/" + userGuid + "/members"
	response, err := Get(apiEndpointUrl, c.defaultHeaders())
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	return parseMembersResponse(response)
}

func parseMemberResponse(response *http.Response) (*models.Member, error) {
	if err := parseResponseErrors(response.StatusCode); err != nil {
		return nil, err
	}

	buffer := new(bytes.Buffer)
	buffer.ReadFrom(response.Body)
	bufferStr := buffer.String()
	response.Body.Close()

	if response.StatusCode == 200 || response.StatusCode == 202 {
		memberResponse := &models.MemberResponse{}
		if err := json.Unmarshal([]byte(bufferStr), memberResponse); err != nil {
			return nil, err
		}
		return memberResponse.Member, nil
	}

	return nil, makeGenericError(response.StatusCode, bufferStr)
}

func (c *Client) GetMember(userGuid, memberGuid string) (*models.Member, error) {
	if userGuid == "" || memberGuid == "" {
		return nil, MissingGuid
	}

	apiEndpointUrl := c.ApiURL + "/users/" + userGuid + "/members/" + memberGuid
	response, err := Get(apiEndpointUrl, c.defaultHeaders())
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	return parseMemberResponse(response)
}

func (c *Client) GetMemberStatus(userGuid, memberGuid string) (*models.Member, error) {
	if userGuid == "" || memberGuid == "" {
		return nil, MissingGuid
	}

	apiEndpointUrl := c.ApiURL + "/users/" + userGuid + "/members/" + memberGuid + "/status"
	response, err := Get(apiEndpointUrl, c.defaultHeaders())
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	return parseMemberResponse(response)
}

func (c *Client) GetMemberChallenges(userGuid, memberGuid string) ([]*models.Challenge, error) {
	if userGuid == "" || memberGuid == "" {
		return nil, MissingGuid
	}

	apiEndpointUrl := c.ApiURL + "/users/" + userGuid + "/members/" + memberGuid + "/challenges"
	response, err := Get(apiEndpointUrl, c.defaultHeaders())
	if err != nil {
		return nil, err
	}

	buffer := new(bytes.Buffer)
	buffer.ReadFrom(response.Body)
	bufferStr := buffer.String()
	response.Body.Close()

	if response.StatusCode == 200 {
		challengesResponse := &models.ChallengesResponse{}
		if err := json.Unmarshal([]byte(bufferStr), challengesResponse); err != nil {
			return nil, err
		}
		return challengesResponse.Challenges, nil
	}

	return nil, makeGenericError(response.StatusCode, bufferStr)
}

func (c *Client) CreateMember(userGuid string, member *models.Member, credentials []*models.Credential) (*models.Member, error) {
	if userGuid == "" || len(credentials) == 0 {
		return nil, MissingGuid
	}

	memberCreateRequest := models.NewMemberCreateRequest(member, credentials)
	jsonStr, err := json.Marshal(memberCreateRequest)
	if err != nil {
		return nil, err
	}

	apiEndpointUrl := c.ApiURL + "/users/" + userGuid + "/members"
	response, err := Post(apiEndpointUrl, string(jsonStr), c.defaultHeaders())
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	return parseMemberResponse(response)
}

func (c *Client) UpdateMember(userGuid string, member *models.Member, credentials []*models.Credential) (*models.Member, error) {
	if userGuid == "" || member == nil || member.Guid == "" {
		return nil, MissingGuid
	}

	memberUpdateRequest := models.NewMemberUpdateRequest(member, credentials)
	jsonStr, err := json.Marshal(memberUpdateRequest)
	if err != nil {
		return nil, err
	}

	apiEndpointUrl := c.ApiURL + "/users/" + userGuid + "/members/" + member.Guid
	response, err := Put(apiEndpointUrl, string(jsonStr), c.defaultHeaders())
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	return parseMemberResponse(response)
}

func (c *Client) DeleteMember(userGuid string, memberGuid string) error {
	if userGuid == "" || memberGuid == "" {
		return MissingGuid
	}

	apiEndpointUrl := c.ApiURL + "/users/" + userGuid + "/members/" + memberGuid
	response, err := Delete(apiEndpointUrl, c.defaultHeaders())
	defer response.Body.Close()
	if err != nil {
		return err
	}

	if err = parseResponseErrors(response.StatusCode); err != nil {
		return err
	}

	if response.StatusCode == 204 {
		return nil
	}

	buffer := new(bytes.Buffer)
	buffer.ReadFrom(response.Body)
	bufferStr := buffer.String()
	return makeGenericError(response.StatusCode, bufferStr)
}

func (c *Client) AggregateMember(userGuid string, memberGuid string) (*models.Member, error) {
	if userGuid == "" || memberGuid == "" {
		return nil, MissingGuid
	}

	apiEndpointUrl := c.ApiURL + "/users/" + userGuid + "/members/" + memberGuid + "/aggregate"
	response, err := Post(apiEndpointUrl, "", c.defaultHeaders())
	defer response.Body.Close()
	if err != nil {
		return nil, err
	}

	if err = parseResponseErrors(response.StatusCode); err != nil {
		return nil, err
	}

	defer response.Body.Close()

	return parseMemberResponse(response)
}

func (c *Client) ResumeMember(userGuid, memberGuid string, challenges []*models.Challenge) (*models.Member, error) {
	if userGuid == "" || memberGuid == "" {
		return nil, MissingGuid
	}

	memberResumeRequest := models.NewMemberResumeRequest(challenges)
	jsonStr, err := json.Marshal(memberResumeRequest)
	if err != nil {
		return nil, err
	}

	apiEndpointUrl := c.ApiURL + "/users/" + userGuid + "/members/" + memberGuid + "/resume"
	response, err := Put(apiEndpointUrl, string(jsonStr), c.defaultHeaders())
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	return parseMemberResponse(response)
}

func (c *Client) ListMemberAccounts(userGuid, memberGuid string) ([]*models.Account, error) {
	if userGuid == "" || memberGuid == "" {
		return nil, MissingGuid
	}

	apiEndpointUrl := c.ApiURL + "/users/" + userGuid + "/members/" + memberGuid + "/accounts"
	response, err := Get(apiEndpointUrl, c.defaultHeaders())
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if err := parseResponseErrors(response.StatusCode); err != nil {
		return nil, err
	}

	buffer := new(bytes.Buffer)
	buffer.ReadFrom(response.Body)
	bufferStr := buffer.String()

	if response.StatusCode == 200 {
		accountsResponse := &models.AccountsResponse{}
		if err := json.Unmarshal([]byte(bufferStr), accountsResponse); err != nil {
			return nil, err
		}
		return accountsResponse.Accounts, nil
	}

	return nil, makeGenericError(response.StatusCode, bufferStr)
}

func (c *Client) ListMemberTransactions(userGuid, memberGuid string) ([]*models.Transaction, error) {
	return c.ListMemberTransactionsWithDateRange(userGuid, memberGuid, "", "")
}

func (c *Client) ListMemberTransactionsWithDateRange(userGuid, memberGuid, fromDate, toDate string) ([]*models.Transaction, error) {
	if userGuid == "" || memberGuid == "" {
		return nil, MissingGuid
	}

	apiEndpointUrl := c.ApiURL + "/users/" + userGuid + "/members/" + memberGuid + "/transactions" + buildparams("", fromDate, toDate)
	response, err := Get(apiEndpointUrl, c.defaultHeaders())
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if err := parseResponseErrors(response.StatusCode); err != nil {
		return nil, err
	}

	buffer := new(bytes.Buffer)
	buffer.ReadFrom(response.Body)
	bufferStr := buffer.String()

	if response.StatusCode == 200 {
		transactionsResponse := &models.TransactionsResponse{}
		if err := json.Unmarshal([]byte(bufferStr), transactionsResponse); err != nil {
			return nil, err
		}
		return transactionsResponse.Transactions, nil
	}

	return nil, makeGenericError(response.StatusCode, bufferStr)
}

func (c *Client) ListMemberCredentials(userGuid, memberGuid string) ([]*models.Credential, error) {
	if userGuid == "" || memberGuid == "" {
		return nil, MissingGuid
	}

	apiEndpointUrl := c.ApiURL + "/users/" + userGuid + "/members/" + memberGuid + "/credentials"
	response, err := Get(apiEndpointUrl, c.defaultHeaders())
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if err := parseResponseErrors(response.StatusCode); err != nil {
		return nil, err
	}

	buffer := new(bytes.Buffer)
	buffer.ReadFrom(response.Body)
	bufferStr := buffer.String()
	response.Body.Close()

	if response.StatusCode == 200 {
		credentialResponse := &models.CredentialsResponse{}
		if err := json.Unmarshal([]byte(bufferStr), credentialResponse); err != nil {
			return nil, err
		}
		return credentialResponse.Credentials, nil
	}

	return nil, makeGenericError(response.StatusCode, bufferStr)
}

func (c *Client) VerifyMember(userGuid, memberGuid string) (*models.Member, error) {
	if userGuid == "" || memberGuid == "" {
		return nil, MissingGuid
	}

	apiEndpointUrl := c.ApiURL + "/users/" + userGuid + "/members/" + memberGuid + "/verify"
	response, err := Post(apiEndpointUrl, c.defaultHeaders())
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	return parseMemberResponse(response)
}

func (c *Client) IdentifyMember(userGuid, memberGuid string) (*models.Member, error) {
	if userGuid == "" || memberGuid == "" {
		return nil, MissingGuid
	}

	apiEndpointUrl := c.ApiURL + "/users/" + userGuid + "/members/" + memberGuid + "/identify"
	response, err := Post(apiEndpointUrl, c.defaultHeaders())
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	return parseMemberResponse(response)
}
