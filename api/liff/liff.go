package liff

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/henzai/liffc/api/util"
)

const (
	baseURL = "https://api.line.me/liff/v1/apps"
	pushURL = "https://api.line.me/v2/bot/message/push"
)

type LIFFClient struct {
	lineAccessToken string
}

func NewLIFFClient(token string) *LIFFClient {
	return &LIFFClient{token}
}

func (liff *LIFFClient) GetLineAccessToken() string {
	return liff.lineAccessToken
}

func (liff *LIFFClient) Add(option *AppOption) (*LIFFIDResponse, error) {
	form, err := json.Marshal(option)
	//fmt.Println(string(form))
	resp, err := util.Request(liff, "POST", baseURL, bytes.NewBuffer(form))
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
	var liffIDResponse LIFFIDResponse
	err = json.Unmarshal(body, &liffIDResponse)
	if err != nil {
		return nil, err
	}
	return &liffIDResponse, nil
}

func (liff *LIFFClient) Delete(liffID string) error {

	resp, err := util.Request(liff, "DELETE", fmt.Sprintf("%v/%v", baseURL, liffID), nil)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer resp.Body.Close()
	return nil
}

func (liff *LIFFClient) List() (*Apps, error) {

	resp, err := util.Request(liff, "GET", baseURL, nil)
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

func (liff *LIFFClient) Send(pushMessage *PushMessage) error {
	form, err := json.Marshal(pushMessage)

	resp, err := util.Request(liff, "POST", pushURL, bytes.NewBuffer(form))
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

func (liff *LIFFClient) Update(liffID string, option *AppOption) error {
	form, err := json.Marshal(option)
	resp, err := util.Request(liff, "PUT", fmt.Sprintf("%v/%v", baseURL, liffID), bytes.NewBuffer(form))
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
