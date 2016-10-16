package client

import (
	"errors"
)

var (
	UserLimitExceeded = errors.New("You have exceeded the maximum amount of users for your account.")
	RateLimitExceeded = errors.New("The server is under load, please try again later.")
	MissingGuid = errors.New("A guid is missing")
)

type Client struct {
	ApiKey string
	ClientId string
	ApiURL string
}

func (c *Client) defaultHeaders() *Headers {
	headers := Headers{}
	headers["Content-Type"] = "application/json"
	headers["MX-API-KEY"] = c.ApiKey
	headers["MX-CLIENT-ID"] = c.ClientId
	return &headers
}
