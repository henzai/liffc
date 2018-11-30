package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func (c *Client) Update(liffID string, option *Option) (*http.Response, error) {
	form, err := json.Marshal(option)

	resp, err := c.request("PUT", fmt.Sprintf("%v/%v/view", baseURL, liffID), bytes.NewBuffer(form))
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer resp.Body.Close()
	return resp, nil
}
