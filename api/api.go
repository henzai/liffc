package api

import (
	"fmt"
	"io"
	"net/http"
)

const (
	baseURL = "https://api.line.me/liff/v1/apps"
)

type Client struct {
	lineAccessToken string
}

func NewClient(token string) *Client {
	return &Client{token}
}

func (liffc *Client) request(method, url string, body io.Reader) (*http.Response, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %v", liffc.lineAccessToken))
	req.Header.Add("Content-Type", "application/json")

	client := http.Client{}
	return client.Do(req)
}
