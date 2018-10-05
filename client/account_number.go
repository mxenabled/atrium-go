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

func (c *Client) ListAccountAccountNumbers(userGuid, accountGuid string) ([]*models.AccountNumber, error) {
	if userGuid == "" || accountGuid == "" {
		return nil, MissingGuid
	}

	apiEndpointUrl := c.ApiURL + "/users/" + userGuid + "/accounts/" + accountGuid + "/account_numbers"

	response, err := c.Get(apiEndpointUrl, c.defaultHeaders())
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	return parseAccountNumbersResponse(response)
}

func (c *Client) ListMemberAccountNumbers(userGuid, memberGuid string) ([]*models.AccountNumber, error) {
	if userGuid == "" || memberGuid == "" {
		return nil, MissingGuid
	}

	apiEndpointUrl := c.ApiURL + "/users/" + userGuid + "/members/" + memberGuid + "/account_numbers"

	response, err := c.Get(apiEndpointUrl, c.defaultHeaders())
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	return parseAccountNumbersResponse(response)
}
