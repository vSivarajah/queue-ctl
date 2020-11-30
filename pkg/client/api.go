package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const baseURL string = "http://localhost:8082/v3/clusters/aCyyjt7nR1C38M0pvE54Yw"

type Client struct {
	Client *http.Client
}

func NewClient() (*Client, error) {
	client := Client{
		Client: &http.Client{},
	}
	return &client, nil
}

func (c *Client) Get(url string, responseType interface{}) error {
	return c.CallAPI("GET", url, nil, responseType)
}

func (c *Client) Post(url string, reqBody, responseType interface{}) error {
	return c.CallAPI("POST", url, reqBody, responseType)
}

func (c *Client) CallAPI(method, path string, reqBody, responseType interface{}) error {
	return c.CallAPIWithContext(context.Background(), method, path, reqBody, responseType)
}

func (c *Client) CallAPIWithContext(ctx context.Context, method, path string, reqBody, responseType interface{}) error {
	req, err := c.NewRequest(method, path, reqBody)
	if err != nil {
		return err
	}

	req = req.WithContext(ctx)
	response, err := c.Do(req)
	if err != nil {
		return err
	}
	return c.UnmarshalJson(response, responseType)
}

func (c *Client) UnmarshalJson(response *http.Response, responseType interface{}) error {
	// Read all the response body
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}
	// API error (200-300 code)
	if response.StatusCode < http.StatusOK || response.StatusCode >= http.StatusMultipleChoices {
		apiError := &RestError{
			StatusCode: response.StatusCode,
		}
		if err = json.Unmarshal(body, apiError); err != nil {
			apiError.Message = string(body)
		}
		return apiError
	}

	//No data sent in
	if len(body) == 0 || responseType == nil {
		return nil
	}

	data := json.NewDecoder(bytes.NewReader(body))
	return data.Decode(&responseType)
}

func (c *Client) Do(req *http.Request) (*http.Response, error) {
	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *Client) NewRequest(method, path string, reqBody interface{}) (*http.Request, error) {
	var body []byte
	var err error

	if reqBody != nil {
		body, err = json.Marshal(reqBody)
		if err != nil {
			return nil, err
		}
	}
	target := fmt.Sprintf("%s%s", baseURL, path)

	req, err := http.NewRequest(method, target, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}

	return req, nil
}
