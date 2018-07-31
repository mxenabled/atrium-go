package models

type AccountNumber struct {
	AccountGuid   string `json:"account_guid,omitempty"`
	AccountNumber string `json:"account_number,omitempty"`
	Guid          string `json:"guid,omitempty"`
	MemberGuid    string `json:"member_guid,omitempty"`
	RoutingNumber string `json:"routing_number,omitempty"`
	UserGuid      string `json:"user_guid,omitempty"`
}

type AccountNumbersResponse struct {
	AccountNumbers []*AccountNumber `json:"account_numbers"`
}

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
