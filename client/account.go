package client

import (
	"bytes"
	"encoding/json"
	"github.com/mxenabled/atrium-go/models"
)

func (c *Client) ListAccounts(userGuid string) ([]*models.Account, error) {
	if userGuid == "" {
		return nil, MissingGuid
	}

	apiEndpointUrl := c.ApiURL + "/users/" + userGuid + "/accounts"
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
		json.Unmarshal([]byte(bufferStr), accountsResponse)
		return accountsResponse.Accounts, nil
	}

	return nil, makeGenericError(response.StatusCode, bufferStr)
}

func (c *Client) GetAccount(userGuid, accountGuid string) (*models.Account, error) {
	if userGuid == "" || accountGuid == "" {
		return nil, MissingGuid
	}

	apiEndpointUrl := c.ApiURL + "/users/" + userGuid + "/accounts/" + accountGuid
	response, err := Get(apiEndpointUrl, c.defaultHeaders())
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	buffer := new(bytes.Buffer)
	buffer.ReadFrom(response.Body)
	bufferStr := buffer.String()

	if response.StatusCode == 200 {
		accountResponse := &models.AccountResponse{}
		json.Unmarshal([]byte(bufferStr), accountResponse)
		return accountResponse.Account, nil
	}

	return nil, makeGenericError(response.StatusCode, bufferStr)
}

func (c *Client) ListAccountTransations(userGuid, accountGuid string) ([]*models.Transaction, error) {
	if userGuid == "" || accountGuid == "" {
		return nil, MissingGuid
	}

	apiEndpointUrl := c.ApiURL + "/users/" + userGuid + "/accounts/" + accountGuid + "/transactions"
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
		json.Unmarshal([]byte(bufferStr), transactionsResponse)
		return transactionsResponse.Transactions, nil
	}

	return nil, makeGenericError(response.StatusCode, bufferStr)
}
