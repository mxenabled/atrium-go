package client

import (
	"io"
	"net/http"
	"net/url"
	"strings"
)

type Headers map[string]string

func (c *Client) request(method, url string, body io.Reader, headers Headers) (*http.Response, error) {
	request, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	for key, value := range headers {
		request.Header.Add(key, value)
	}

	// NOTE: This could be safer by adding a lock, but this will eventually memoize.
	if c.HttpClient == nil {
		c.HttpClient = &http.Client{}
	}
	response, err := c.HttpClient.Do(request)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) Get(url string, headers Headers) (*http.Response, error) {
	return c.request("GET", url, nil, headers)
}

func (c *Client) Post(url, body string, headers Headers) (*http.Response, error) {
	return c.request("POST", url, strings.NewReader(body), headers)
}

func (c *Client) Put(url, body string, headers Headers) (*http.Response, error) {
	return c.request("PUT", url, strings.NewReader(body), headers)
}

func (c *Client) Delete(url string, headers Headers) (*http.Response, error) {
	return c.request("DELETE", url, nil, headers)
}

func buildparams(name, fromDate, toDate string) string {
	params := url.Values{}
	if len(name) > 0 {
		params.Add("name", name)
	}
	if len(fromDate) > 0 {
		params.Add("from_date", fromDate)
	}
	if len(toDate) > 0 {
		params.Add("to_date", toDate)
	}
	paramsStr := "?" + params.Encode()
	return paramsStr
}
