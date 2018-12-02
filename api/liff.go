package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
)

const (
	pushURL = "https://api.line.me/v2/bot/message/push"
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

func (c *Client) Delete(liffID string) error {

	resp, err := c.request("DELETE", fmt.Sprintf("%v/%v", baseURL, liffID), nil)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer resp.Body.Close()
	return nil
}

func (c *Client) List() (*Apps, error) {

	resp, err := c.request("GET", baseURL, nil)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var apps Apps
	err = json.Unmarshal(body, &apps)
	if err != nil {
		return nil, err
	}
	return &apps, nil
}

func (c *Client) Send(pushMessage *PushMessage) error {
	form, err := json.Marshal(pushMessage)

	resp, err := c.request("POST", pushURL, bytes.NewBuffer(form))
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		var i interface{}
		body, _ := ioutil.ReadAll(resp.Body)
		err = json.Unmarshal(body, &i)
		if err != nil {
			return err
		}
		return fmt.Errorf("Error %s %v", resp.Status, i)
	}
	return nil
}

func (c *Client) Update(liffID string, option *AppOption) error {
	form, err := json.Marshal(option)
	resp, err := c.request("PUT", fmt.Sprintf("%v/%v", baseURL, liffID), bytes.NewBuffer(form))
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer resp.Body.Close()
	if code := resp.StatusCode; code == 401 {
		return fmt.Errorf("Error: %s authorization failed", resp.Status)
	}
	if code := resp.StatusCode; code == 404 {
		return fmt.Errorf("Error: %s %v", resp.Status, liffID)
	}
	return nil
}
