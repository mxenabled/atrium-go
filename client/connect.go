package client

import (
	"bytes"
	"encoding/json"
	"github.com/mxenabled/atrium-go/models"
	"net/http"
)

func parseConnectResponse(response *http.Response) (*models.Connect, error) {
	if err := parseResponseErrors(response.StatusCode); err != nil {
		return nil, err
	}

	buffer := new(bytes.Buffer)
	buffer.ReadFrom(response.Body)
	bufferStr := buffer.String()
	response.Body.Close()

	if response.StatusCode == 200 {
		connectResponse := &models.ConnectResponse{}
		if err := json.Unmarshal([]byte(bufferStr), connectResponse); err != nil {
			return nil, err
		}
		return connectResponse.Connect, nil
	}

	return nil, makeGenericError(response.StatusCode, bufferStr)
}

func (c *Client) GetWidget(userGuid string) (*models.Connect, error) {
	return c.GetWidgetWithConnectParams(userGuid, models.ConnectParams{})
}

func (c *Client) GetWidgetWithConnectParams(userGuid string, params models.ConnectParams) (*models.Connect, error) {
	if userGuid == "" {
		return nil, MissingGuid
	}

	paramsJSON, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	apiEndpointUrl := c.ApiURL + "/users/" + userGuid + "/connect_widget_url"
	response, err := Post(apiEndpointUrl, string(paramsJSON), c.defaultHeaders())
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	return parseConnectResponse(response)
}
