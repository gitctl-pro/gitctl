package promethus

import (
	"golang.org/x/net/http2"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

type Client struct {
	httpClient *http.Client
	endpoint   string
}

func NewPrometheusClient(endpoint string) *Client {
	client := &http.Client{
		Timeout: time.Duration(10) * time.Second,
	}
	client.Transport = &http2.Transport{}
	return &Client{
		httpClient: client,
		endpoint:   endpoint,
	}
}

func (c *Client) QueryRange(data url.Values) ([]byte, error) {
	url := c.endpoint + "/api/v1/query_range?" + data.Encode()
	response, err := c.httpClient.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	return body, err
}

func (c *Client) Series(data url.Values) ([]byte, error) {
	url := c.endpoint + "/api/v1/series?" + data.Encode()
	response, err := c.httpClient.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	return body, err
}

func (c *Client) Query(data url.Values) ([]byte, error) {
	url := c.endpoint + "/api/v1/query"
	response, err := c.httpClient.PostForm(url, data)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	respBody, err := ioutil.ReadAll(response.Body)
	return respBody, err
}
