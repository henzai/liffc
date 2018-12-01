package liff

import (
	"bytes"
	"encoding/json"
	"fmt"
)

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
