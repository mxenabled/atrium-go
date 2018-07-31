package client

import (
	"bytes"
	"encoding/json"
	"github.com/mxenabled/atrium-go/models"
	"net/http"
)

func parseAccountNumbersResponse(response *http.Response) ([]*models.AccountNumber, error) {
	if err := parseResponseErrors(response.StatusCode); err != nil {
		return nil, err
	}

	buffer := new(bytes.Buffer)
	buffer.ReadFrom(response.Body)
	bufferStr := buffer.String()
	response.Body.Close()

	if response.StatusCode == 200 {
		accountNumbersResponse := &models.AccountNumbersResponse{}
		if err := json.Unmarshal([]byte(bufferStr), accountNumbersResponse); err != nil {
			return nil, err
		}
		return accountNumbersResponse.AccountNumbers, nil
	}

	return nil, makeGenericError(response.StatusCode, bufferStr)
}

func (c *Client) ListAccountNumbers(userGuid, accountOrMemberGuid string) ([]*models.AccountNumber, error) {
	if userGuid == "" || accountOrMemberGuid == "" {
		return nil, MissingGuid
	}

	if accountOrMemberGuid[0:3] == "ACT" {
		apiEndpointUrl := c.ApiURL + "/users/" + userGuid + "/accounts/" + accountOrMemberGuid + "/account_numbers"
	} else {
		apiEndpointUrl := c.ApiURL + "/users/" + userGuid + "/members/" + accountOrMemberGuid + "/account_numbers"
	}

	response, err := Get(apiEndpointUrl, c.defaultHeaders())
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	return parseAccountNumbersResponse(response)
}
