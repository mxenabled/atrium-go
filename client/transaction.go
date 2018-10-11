package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/mxenabled/atrium-go/models"
)

func parseTransactionResponse(response *http.Response) (*models.Transaction, error) {
	if err := parseResponseErrors(response.StatusCode); err != nil {
		return nil, err
	}

	buffer := new(bytes.Buffer)
	buffer.ReadFrom(response.Body)
	bufferStr := buffer.String()
	response.Body.Close()

	if response.StatusCode == 200 {
		transactionResponse := &models.TransactionResponse{}
		if err := json.Unmarshal([]byte(bufferStr), transactionResponse); err != nil {
			return nil, err
		}
		return transactionResponse.Transaction, nil
	}

	return nil, makeGenericError(response.StatusCode, bufferStr)
}

func (c *Client) CategorizeAndDescribeTransactions(transactionsToCategorize []models.Transaction) ([]*models.Transaction, error) {
	apiEndpointURL := c.ApiURL + "/categorize_and_describe"

	var transactionsRequest = struct {
		Transaction []models.Transaction `json:"transactions"`
	}{transactionsToCategorize}

	transactionsJSON, err := json.Marshal(transactionsRequest)

	if err != nil {
		return nil, err
	}

	body := string(transactionsJSON)

	fmt.Println(body)
	response, err := c.Post(apiEndpointURL, body, c.defaultHeaders())

	defer response.Body.Close()

	if err != nil {
		return nil, err
	}

	if err = parseResponseErrors(response.StatusCode); err != nil {
		return nil, err
	}

	buffer := new(bytes.Buffer)
	buffer.ReadFrom(response.Body)
	bufferStr := buffer.String()

	fmt.Printf("Response: %+v\n", bufferStr)

	transactionsResponse := &models.TransactionsResponse{}
	if err = json.Unmarshal([]byte(bufferStr), transactionsResponse); err != nil {
		return nil, err
	}
	return transactionsResponse.Transactions, nil
}

func (c *Client) ListTransactions(userGuid string) ([]*models.Transaction, error) {
	return c.ListTransactionsWithDateRange(userGuid, "", "")
}

func (c *Client) ListTransactionsWithDateRange(userGuid, fromDate, toDate string) ([]*models.Transaction, error) {
	if userGuid == "" {
		return nil, MissingGuid
	}

	apiEndpointUrl := c.ApiURL + "/users/" + userGuid + "/transactions" + buildparams("", fromDate, toDate)
	response, err := c.Get(apiEndpointUrl, c.defaultHeaders())
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

func (c *Client) GetTransaction(userGuid, transactionGuid string) (*models.Transaction, error) {
	if userGuid == "" || transactionGuid == "" {
		return nil, MissingGuid
	}

	apiEndpointUrl := c.ApiURL + "/users/" + userGuid + "/transactions/" + transactionGuid
	response, err := c.Get(apiEndpointUrl, c.defaultHeaders())
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	return parseTransactionResponse(response)
}
