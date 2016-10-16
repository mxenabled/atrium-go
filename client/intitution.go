package client

import (
	"github.com/mxenabled/atrium-go/models"
	"encoding/json"
	"bytes"
	"strconv"
	"errors"
)

func (c *Client) ListInstitutions() ([]*models.Institution, error) {
	apiEndpointUrl := c.ApiURL + "/institutions"
	response, err := Get(apiEndpointUrl, c.defaultHeaders())
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	buffer := new(bytes.Buffer)
	buffer.ReadFrom(response.Body)
	bufferStr := buffer.String()
	response.Body.Close()

	switch response.StatusCode {
	case 200:
		institutionResponse := &models.InstitutionsResponse{}
		json.Unmarshal([]byte(bufferStr), institutionResponse)
		return institutionResponse.Institutions, nil
	case 422:
		return nil, UserLimitExceeded
	case 509:
		return nil, RateLimitExceeded
	}

	return nil, errors.New("Received response " + strconv.Itoa(response.StatusCode) + " - " + bufferStr)
}
