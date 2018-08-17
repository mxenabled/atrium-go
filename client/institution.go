package client

import (
	"bytes"
	"encoding/json"
	"github.com/mxenabled/atrium-go/models"
	"net/http"
)

func parseInstitutionResponse(response *http.Response) (*models.Institution, error) {
	if err := parseResponseErrors(response.StatusCode); err != nil {
		return nil, err
	}

	buffer := new(bytes.Buffer)
	buffer.ReadFrom(response.Body)
	bufferStr := buffer.String()
	response.Body.Close()

	if response.StatusCode == 200 {
		institutionResponse := &models.InstitutionResponse{}
		json.Unmarshal([]byte(bufferStr), institutionResponse)
		return institutionResponse.Institution, nil
	}

	return nil, makeGenericError(response.StatusCode, bufferStr)
}

func (c *Client) ListInstitutions(name string) ([]*models.Institution, error) {
	apiEndpointUrl := c.ApiURL + "/institutions" + buildparams(name, "", "")
	response, err := Get(apiEndpointUrl, c.defaultHeaders(), c.LocalIP)
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
	response.Body.Close()

	if response.StatusCode == 200 {
		institutionResponse := &models.InstitutionsResponse{}
		json.Unmarshal([]byte(bufferStr), institutionResponse)
		return institutionResponse.Institutions, nil
	}

	return nil, makeGenericError(response.StatusCode, bufferStr)
}

func (c *Client) GetInstitution(code string) (*models.Institution, error) {
	if code == "" {
		return nil, MissingGuid
	}

	apiEndpointUrl := c.ApiURL + "/institutions/" + code
	response, err := Get(apiEndpointUrl, c.defaultHeaders(), c.LocalIP)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	return parseInstitutionResponse(response)
}
