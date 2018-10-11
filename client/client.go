package client

import (
	"errors"
	"net/http"
	"strconv"
)

var (
	UserLimitExceeded = errors.New("You have exceeded the maximum amount of users for your account.")
	RateLimitExceeded = errors.New("The server is under load, please try again later.")
	MissingGuid       = errors.New("A guid is missing.")
	NotFound          = errors.New("Resource was not found.")
)

type Client struct {
	ApiKey   string
	ClientId string
	ApiURL   string

	HttpClient *http.Client
}

func (c *Client) defaultHeaders() Headers {
	headers := Headers{}
	headers["Content-Type"] = "application/json"
	headers["MX-API-KEY"] = c.ApiKey
	headers["MX-CLIENT-ID"] = c.ClientId
	return headers
}

func parseResponseErrors(statusCode int) error {
	switch statusCode {
	case 404:
		return NotFound
	case 422:
		return UserLimitExceeded
	case 509:
		return RateLimitExceeded
	case 200: // Make OK explicitly nil
		return nil
	}

	return nil
}

func makeGenericError(statusCode int, message string) error {
	return errors.New("Received response " + strconv.Itoa(statusCode) + " - " + message)
}
