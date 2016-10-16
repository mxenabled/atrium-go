package client

import (
	"github.com/mxenabled/atrium-go/models"
	"bytes"
	"strconv"
	"encoding/json"
	"net/http"
	"errors"
)

func parseUserResponse(response *http.Response) (*models.User, error) {
	buffer := new(bytes.Buffer)
	buffer.ReadFrom(response.Body)
	bufferStr := buffer.String()
	response.Body.Close()

	switch response.StatusCode {
	case 200:
		userResponse := &models.UserResponse{}
		json.Unmarshal([]byte(bufferStr), userResponse)
		return userResponse.User, nil
	case 422:
		return nil, UserLimitExceeded
	case 509:
		return nil, RateLimitExceeded
	}

	return nil, errors.New("Received response " + strconv.Itoa(response.StatusCode) + " - " + bufferStr)
}

func (c *Client) CreateUser(user *models.User) (*models.User, error) {
	userRequest := models.NewUserRequest(user)
	jsonStr, err := json.Marshal(userRequest)
	if err != nil {
		return nil, err
	}

	apiEndpointUrl := c.ApiURL + "/users"
	response, err := Post(apiEndpointUrl, string(jsonStr), c.defaultHeaders())
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	return parseUserResponse(response)
}

func (c *Client) UpdateUser(user *models.User) (*models.User, error) {
	if user.Guid == "" {
		return nil, MissingGuid
	}

	userRequest := models.NewUserRequest(user)
	jsonStr, err := json.Marshal(userRequest)
	if err != nil {
		return nil, err
	}

	apiEndpointUrl := c.ApiURL + "/users/" + user.Guid
	response, err := Put(apiEndpointUrl, string(jsonStr), c.defaultHeaders())
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	return parseUserResponse(response)
}


func (c *Client) GetUser(userGuid string) (*models.User, error) {
	if userGuid == "" {
		return nil, MissingGuid
	}

	apiEndpointUrl := c.ApiURL + "/users/" + userGuid
	response, err := Get(apiEndpointUrl, c.defaultHeaders())
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	return parseUserResponse(response)
}

func (c *Client) DeleteUser(userGuid string) error {
	if userGuid == "" {
		return MissingGuid
	}

	apiEndpointUrl := c.ApiURL + "/users/" + userGuid
	response, err := Get(apiEndpointUrl, c.defaultHeaders())
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if response.StatusCode == 204 {
		return nil
	}
	_, err = parseUserResponse(response)
	return err
}
