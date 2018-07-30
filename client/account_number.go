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
		json.Unmarshal([]byte(bufferStr), accountNumbersResponse)
		return accountNumbersResponse.AccountNumbers, nil
	}

	return nil, makeGenericError(response.StatusCode, bufferStr)
}

func (c *Client) ListAccountNumbers(userGuid, memberGuid string) ([]*models.AccountNumber, error) {
	if userGuid == "" || memberGuid == "" {
		return nil, MissingGuid
	}

	apiEndpointUrl := c.ApiURL + "/users/" + userGuid + "/members/" + memberGuid + "/account_numbers"
	response, err := Get(apiEndpointUrl, c.defaultHeaders())
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	return parseAccountNumbersResponse(response)
}
