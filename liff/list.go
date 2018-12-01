package liff

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

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
