package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	pushURL = "https://api.line.me/v2/bot/message/push"
)

type PushMessage struct {
	To       string
	Messages []Message
}

type Message struct {
	Type string
	Text string
}

func NewPushMessage(liffID, userID string) *PushMessage {
	return &PushMessage{
		userID,
		[]Message{
			Message{"text", fmt.Sprintf("line://app/%v", liffID)},
		},
	}
}

func (c *Client) Send(pushMessage *PushMessage) (*http.Response, error) {
	form, err := json.Marshal(pushMessage)

	resp, err := c.request("POST", pushURL, bytes.NewBuffer(form))
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer resp.Body.Close()
	return resp, nil
}
