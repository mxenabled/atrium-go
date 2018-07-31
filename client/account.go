package client

import (
	"bytes"
	"encoding/json"
	"github.com/mxenabled/atrium-go/models"
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

func (c *Client) ListAccountNumbers(userGuid, accountGuid string) ([]*models.AccountNumber, error) {
	if userGuid == "" || accountGuid == "" {
		return nil, MissingGuid
	}

	apiEndpointUrl := c.ApiURL + "/users/" + userGuid + "/accounts/" + accountGuid + "/account_numbers"
	response, err := Get(apiEndpointUrl, c.defaultHeaders())
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	return parseAccountNumbersResponse(response)
}

func (c *Client) ListAccountTransactions(userGuid, accountGuid string) ([]*models.Transaction, error) {
	return c.ListAccountTransactionsWithDateRange(userGuid, accountGuid, "", "")
}

func (c *Client) ListAccountTransactionsWithDateRange(userGuid, accountGuid, fromDate, toDate string) ([]*models.Transaction, error) {
	if userGuid == "" || accountGuid == "" {
		return nil, MissingGuid
	}

	apiEndpointUrl := c.ApiURL + "/users/" + userGuid + "/accounts/" + accountGuid + "/transactions" + buildparams("", fromDate, toDate)
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
