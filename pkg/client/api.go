package client

import "fmt"

const baseURL string = "http://%s:8082/v3/clusters/"

type Client struct {
	ClusterId string
	Host      string
}

func NewClient(clusterid string, host string) *Client {
	return &Client{
		ClusterId: clusterid,
		Host:      fmt.Sprintf(baseURL, host),
	}
}
