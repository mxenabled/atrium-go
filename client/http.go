package client

import (
	"io"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type Headers map[string]string

func request(method, url string, body io.Reader, headers *Headers, localIP *string) (*http.Response, error) {
	request, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	for key, value := range *headers {
		request.Header.Add(key, value)
	}

	client := &http.Client{}

	if localIP != nil {
		localAddr, err := net.ResolveIPAddr("ip", *localIP)
		if err != nil {
			return nil, err
		}

		dialer := &net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
			LocalAddr: &net.TCPAddr{
				IP: localAddr.IP,
			},
		}

		client = &http.Client{
			Transport: &http.Transport{
				Proxy:               http.ProxyFromEnvironment,
				TLSHandshakeTimeout: 10 * time.Second,
				Dial:                dialer.Dial,
			},
		}
	}

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func Get(url string, headers *Headers, localIP *string) (*http.Response, error) {
	return request("GET", url, nil, headers, localIP)
}

func Post(url, body string, headers *Headers, localIP *string) (*http.Response, error) {
	return request("POST", url, strings.NewReader(body), headers, localIP)
}

func Put(url, body string, headers *Headers, localIP *string) (*http.Response, error) {
	return request("PUT", url, strings.NewReader(body), headers, localIP)
}

func Delete(url string, headers *Headers, localIP *string) (*http.Response, error) {
	return request("DELETE", url, nil, headers, localIP)
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
