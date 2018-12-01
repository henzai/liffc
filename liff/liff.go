package liff

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
)

const (
	baseURL = "https://api.line.me/liff/v1/apps"
	pushURL = "https://api.line.me/v2/bot/message/push"
)

type Client struct {
	lineAccessToken string
}

func NewClient(token string) *Client {
	return &Client{token}
}

func (liffc *Client) request(method, url string, body io.Reader) (*http.Response, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %v", liffc.lineAccessToken))
	req.Header.Add("Content-Type", "application/json")

	client := http.Client{}
	return client.Do(req)
}

type Apps struct {
	Apps []App
}

func (apps *Apps) StringArray() [][]string {
	s := make([][]string, len(apps.Apps))
	for _, v := range apps.Apps {
		i := []string{
			v.LiffID,
			v.AppOption.View.Description,
			v.AppOption.View.Type.String(),
			v.AppOption.View.URL,
			strconv.FormatBool(v.Features.Ble),
		}
		s = append(s, i)
	}
	return s
}

type App struct {
	LiffID string `json:"liffId"`
	AppOption
}

type AppOption struct {
	View     View    `json:"view"`
	Features Feature `json:"features"`
}

func NewAppOption(description, liffType, liffURL string, ble bool) (*AppOption, error) {
	viewType, err := NewViewType(liffType)
	if err != nil {
		return nil, err
	}
	addOption := &AppOption{
		View{
			description,
			viewType,
			liffURL,
		},
		Feature{
			ble,
		},
	}
	return addOption, nil
}

type View struct {
	Description string   `json:"description"`
	Type        ViewType `json:"type"`
	URL         string   `json:"url"`
}

type Feature struct {
	Ble bool `json:"ble"`
}

type PushMessage struct {
	To       string    `json:"to"`
	Messages []Message `json:"messages"`
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

type Message struct {
	Type string `json:"type"`
	Text string `json:"text"`
}
