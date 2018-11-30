package api

import (
	"fmt"
)

func (c *Client) Delete(liffID string) error {

	resp, err := c.request("DELETE", fmt.Sprintf("%v/%v", baseURL, liffID), nil)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer resp.Body.Close()
	return nil
}
