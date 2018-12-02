package util

import (
	"fmt"
	"io"
	"net/http"
)

type LineClient interface {
	GetLineAccessToken() string
}

func Request(c LineClient, method, url string, body io.Reader) (*http.Response, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %v", c.GetLineAccessToken()))
	req.Header.Add("Content-Type", "application/json")

	client := http.Client{}
	return client.Do(req)
}
