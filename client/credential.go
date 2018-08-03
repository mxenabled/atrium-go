package client

import (
	"bytes"
	"encoding/json"
	"github.com/mxenabled/atrium-go/models"
)

func (c *Client) ListCredentials(institutionCode string) ([]*models.Credential, error) {
	if institutionCode == "" {
		return nil, MissingGuid
	}

	apiEndpointUrl := c.ApiURL + "/institutions/" + institutionCode + "/credentials"
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
		credentialResponse := &models.CredentialsResponse{}
		if err := json.Unmarshal([]byte(bufferStr), credentialResponse); err != nil {
			return nil, err
		}
		return credentialResponse.Credentials, nil
	}

	return nil, makeGenericError(response.StatusCode, bufferStr)
}
