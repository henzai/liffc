package api

import (
	"github.com/henzai/liffc/api/liff"
	"github.com/henzai/liffc/api/things"
)

type Client struct {
	LIFF   *liff.LIFFClient
	Things *things.ThingsClient
}

func NewClient(token string) *Client {
	return &Client{
		LIFF:   liff.NewLIFFClient(token),
		Things: things.NewThingsClient(token),
	}
}
