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

func (c *Client) ListInstitutions(Name string) ([]*models.Institution, error) {
	var param = "?"
	if Name != "" {
		param += "name=" + Name + "&"
	}
	param = param[:len(param)-1]

	apiEndpointUrl := c.ApiURL + "/institutions" + param
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
	response.Body.Close()

	if response.StatusCode == 200 {
		institutionResponse := &models.InstitutionsResponse{}
		json.Unmarshal([]byte(bufferStr), institutionResponse)
		return institutionResponse.Institutions, nil
	}

	return nil, makeGenericError(response.StatusCode, bufferStr)
}

func (c *Client) GetInstitution(Code string) (*models.Institution, error) {
	if Code == "" {
		return nil, MissingGuid
	}

	apiEndpointUrl := c.ApiURL + "/institutions/" + Code
	response, err := Get(apiEndpointUrl, c.defaultHeaders())
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	return parseInstitutionResponse(response)
}
