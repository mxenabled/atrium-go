package client

import (
	"io"
	"net/http"
	"net/url"
	"strings"
)

type Headers map[string]string

func request(method, url string, body io.Reader, headers *Headers) (*http.Response, error) {
	request, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	for key, value := range *headers {
		request.Header.Add(key, value)
	}

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func Get(url string, headers *Headers) (*http.Response, error) {
	return request("GET", url, nil, headers)
}

func Post(url, body string, headers *Headers) (*http.Response, error) {
	return request("POST", url, strings.NewReader(body), headers)
}

func Put(url, body string, headers *Headers) (*http.Response, error) {
	return request("PUT", url, strings.NewReader(body), headers)
}

func Delete(url string, headers *Headers) (*http.Response, error) {
	return request("DELETE", url, nil, headers)
}

func buildparams(name, fromDate, toDate string) (string) {
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
