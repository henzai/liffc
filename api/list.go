package api

import (
	"fmt"
	"io/ioutil"
)

func (c *Client) List() error {

	resp, err := c.request("GET", baseURL, nil)
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
