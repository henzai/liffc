package liff

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
)

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
