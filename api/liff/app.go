package liff

import (
	"strconv"
)

type Apps struct {
	Apps []App
}

type App struct {
	LiffID string `json:"liffId"`
	AppOption
}

type AppOption struct {
	View     View    `json:"view"`
	Features Feature `json:"features"`
}

type View struct {
	Description string   `json:"description"`
	Type        ViewType `json:"type"`
	URL         string   `json:"url"`
}

type Feature struct {
	Ble bool `json:"ble"`
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

func (_ LIFFClient) NewAppOption(description, liffType, liffURL string, ble bool) (*AppOption, error) {
	viewType, err := NewViewType(liffType)
	if err != nil {
		return nil, err
	}
	addOption := &AppOption{
		View: View{
			Description: description,
			Type:        viewType,
			URL:         liffURL,
		},
		Features: Feature{
			Ble: ble,
		},
	}
	return addOption, nil
}
