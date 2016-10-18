package client

import (
	"github.com/mxenabled/atrium-go/models"
	"net/http"
	"encoding/json"
	"bytes"
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
		json.Unmarshal([]byte(bufferStr), membersResponse)
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
		json.Unmarshal([]byte(bufferStr), memberResponse)
		return memberResponse.Member, nil
	}

	return nil, makeGenericError(response.StatusCode, bufferStr)
}

func (c *Client) GetMember(userGuid, memberGuid string) (*models.Member, error) {
	if userGuid == "" {
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
	if userGuid == "" {
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

func (c *Client) CreateMember(userGuid string, member *models.Member, credentials []*models.Credential) (*models.Member, error) {
	if userGuid == "" {
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

func (c *Client) ResumeMember(userGuid string, member *models.Member, challenges []*models.Challenge) (*models.Member, error) {
	if userGuid == "" {
		return nil, MissingGuid
	}

	return nil, nil
}
