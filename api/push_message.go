package api

import "fmt"

type PushMessage struct {
	To       string    `json:"to"`
	Messages []Message `json:"messages"`
}

type Message struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

func NewPushMessage(liffID, userID string) *PushMessage {
	return &PushMessage{
		To: userID,
		Messages: []Message{
			Message{
				Type: "text",
				Text: fmt.Sprintf("line://app/%v", liffID),
			},
		},
	}
}
