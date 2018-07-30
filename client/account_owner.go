package client

import (
	"bytes"
	"encoding/json"
	"github.com/mxenabled/atrium-go/models"
	"net/http"
)

func parseAccountOwnersResponse(response *http.Response) ([]*models.AccountOwner, error) {
	if err := parseResponseErrors(response.StatusCode); err != nil {
		return nil, err
	}

	buffer := new(bytes.Buffer)
	buffer.ReadFrom(response.Body)
	bufferStr := buffer.String()
	response.Body.Close()

	if response.StatusCode == 200 {
		accountOwnersResponse := &models.AccountOwnersResponse{}
		json.Unmarshal([]byte(bufferStr), accountOwnersResponse)
		return accountOwnersResponse.AccountOwners, nil
	}

	return nil, makeGenericError(response.StatusCode, bufferStr)
}

func (c *Client) ListAccountOwners(userGuid, memberGuid string) ([]*models.AccountOwner, error) {
	if userGuid == "" || memberGuid == "" {
		return nil, MissingGuid
	}

	apiEndpointUrl := c.ApiURL + "/users/" + userGuid + "/members/" + memberGuid + "/account_owners"
	response, err := Get(apiEndpointUrl, c.defaultHeaders())
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	return parseAccountOwnersResponse(response)
}
