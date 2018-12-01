package liff

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type LiffIDResponse struct {
	LiffID string `json:"liffId"`
}

func (c *Client) Add(option *AppOption) (*LiffIDResponse, error) {
	form, err := json.Marshal(option)
	//fmt.Println(string(form))
	resp, err := c.request("POST", baseURL, bytes.NewBuffer(form))
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer resp.Body.Close()
	if code := resp.StatusCode; code == 400 {
		return nil, fmt.Errorf("Error: %s", resp.Status)
	}
	if code := resp.StatusCode; code == 401 {
		return nil, fmt.Errorf("Error: %s authorization failed", resp.Status)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var liffIDResponse LiffIDResponse
	err = json.Unmarshal(body, &liffIDResponse)
	if err != nil {
		return nil, err
	}
	return &liffIDResponse, nil
}
