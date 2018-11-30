package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Option struct {
	View     View    `json:"view"`
	Features Feature `json:"features"`
}

type View struct {
	Description string `json:"description"`
	LiffType    string `json:"type"`
	Url         string `json:"url"`
}

type Feature struct {
	Ble bool `json:"ble"`
}

func (c *Client) Add(option *Option) error {
	form, err := json.Marshal(option)
	//fmt.Println(string(form))
	resp, err := c.request("POST", baseURL, bytes.NewBuffer(form))
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	fmt.Println(string(body))
	return nil
}

func NewAddOption(description, liffType, liffURL string, ble bool) *Option {
	addOption := &Option{
		View{
			description,
			liffType,
			liffURL,
		},
		Feature{
			ble,
		},
	}
	return addOption
}
