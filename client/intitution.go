package client

import (
	"github.com/mxenabled/atrium-go/models"
	"encoding/json"
	"bytes"
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

	if err := parseResponseErrors(response.StatusCode); err != nil {
		return nil, err
	}

	if response.StatusCode == 200 {
		institutionResponse := &models.InstitutionsResponse{}
		json.Unmarshal([]byte(bufferStr), institutionResponse)
		return institutionResponse.Institutions, nil
	}

	return nil, makeGenericError(response.StatusCode, bufferStr)
}
