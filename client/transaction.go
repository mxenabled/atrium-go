package client

import (
	"bytes"
	"encoding/json"
	"github.com/mxenabled/atrium-go/models"
	"net/http"
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
		json.Unmarshal([]byte(bufferStr), transactionResponse)
		return transactionResponse.Transaction, nil
	}

	return nil, makeGenericError(response.StatusCode, bufferStr)
}

func (c *Client) ListTransactions(userGuid, fromDate, toDate string) ([]*models.Transaction, error) {
	if userGuid == "" {
		return nil, MissingGuid
	}

	var params = "?"
	if fromDate != "" {
		params += "from_date=" + fromDate + "&"
	}
	if toDate != "" {
		params += "to_date=" + toDate + "&";
	}
	params = params[:len(params)-1]

	apiEndpointUrl := c.ApiURL + "/users/" + userGuid + "/transactions" + params
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

func (c *Client) GetTransaction(userGuid, transactionGuid string) (*models.Transaction, error) {
	if userGuid == "" || transactionGuid == "" {
		return nil, MissingGuid
	}

	apiEndpointUrl := c.ApiURL + "/users/" + userGuid + "/transactions/" + transactionGuid
	response, err := Get(apiEndpointUrl, c.defaultHeaders())
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	return parseTransactionResponse(response)
}
